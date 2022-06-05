package actions

type update struct {
	RunFunc func() error
}

func NewUpdateAction(runFunc func() error) update {
	return update{RunFunc: runFunc}
}

func (u update) Keyword() string     { return "update" }
func (u update) Description() string { return "" }
func (u update) RunText() string     { return "Updating shortify entries..." }
func (u update) Run() error          { return u.RunFunc() }
