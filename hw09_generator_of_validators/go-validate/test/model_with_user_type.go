package test

type UserType int

type Model struct {
	Field1 string   `validate:"len:5"`
	Field2 UserType `validate:"min:5"`
}
