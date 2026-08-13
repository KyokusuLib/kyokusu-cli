package models

type Definition struct {
	Name     string
	Short    string
	Terminal bool
	Handler  Handler
	Children map[string]*Definition
	Options  map[string]*Definition
}