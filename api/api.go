package api

type Action struct {
	Score       int
	Category    string
	Description string
}

type Expansion interface {
	Name() string
	GetRequiredThrows() int
	Run(day int, throws ...int) ([]Action, error)
	SetVerbose(v bool)
}

type ExpansionStruct struct {
	Name              func() string
	GetRequiredThrows func() int
	Run               func(int, []int) ([]string, error)
	SetVerbose        func(v bool)
}
