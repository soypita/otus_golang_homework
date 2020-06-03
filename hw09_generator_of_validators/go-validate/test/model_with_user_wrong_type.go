package test

type UserWrongType bool

type WrongModel struct {
	Field1 string   `validate:"len:5"`
	Field2 UserType `validate:"min:5"`
}
