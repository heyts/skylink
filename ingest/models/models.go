package models

type DatabaseRecord interface {
	Insert() bool
	Exists(id string) bool
}
