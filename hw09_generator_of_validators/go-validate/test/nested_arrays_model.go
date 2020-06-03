package test

type NestedModel struct {
	Field1 []string   `validate:"len:5"`
	Field2 [][]string `validate:"len:5"`
}
