package cachedentries

import (
	aw "github.com/deanishe/awgo"
	"github.com/enniomara/shortify-alfred/internal/api"
)

const (
	alfredEntryCacheName = "entries"
)

type entry struct {
	Name string `json:"name"`
}

type CachedEntries interface {
	GetEntries() ([]entry, error)
	SaveEntries([]api.Entry) error
}

func NewAlfredCacheStorage(cache aw.Cache) CachedEntries {
	return &alfredCacheStorage{
		Cache: cache,
	}
}

type alfredCacheStorage struct {
	Cache aw.Cache
}

func (s *alfredCacheStorage) GetEntries() ([]entry, error) {
	// the non-existence of the cache is not an error, likely it means that the
	// entries must be populated
	if !s.Cache.Exists(alfredEntryCacheName) {
		return []entry{}, nil
	}

	var entries []entry
	err := s.Cache.LoadJSON(alfredEntryCacheName, &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func (s *alfredCacheStorage) SaveEntries(apiEntries []api.Entry) error {
	entries := fromApiEntry(apiEntries)

	err := s.Cache.StoreJSON(alfredEntryCacheName, entries)
	if err != nil {
		return err
	}

	return nil
}

func fromApiEntry(apiEntries []api.Entry) []entry {
	var entries []entry
	for _, e := range apiEntries {
		entries = append(entries, entry{
			Name: e.Name,
		})
	}

	return entries
}
