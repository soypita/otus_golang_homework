package test

type UserSliceType []string

type NestedUserModel struct {
	Field1 []string        `validate:"len:5"`
	Field2 []UserSliceType `validate:"len:5"`
}
