package test

type WrongTypeModel struct {
	Field1 string `validate:"min:5"`
	Field2 int    `validate:"min:5"`
}
