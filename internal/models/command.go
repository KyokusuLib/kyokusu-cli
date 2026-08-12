package models

type Command struct {
	Name    string
	Options []Option
	Child   *Command
}

type Option struct {
	Name  string
	Value string
}

type Input struct {
	Command *Command
	Options  []Option
}