package test

type UnsupportedTagModel struct {
	Field1 string `validate:"strange:tag"`
	Field2 int    `validate:"min:5"`
}
