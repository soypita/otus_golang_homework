package test

type GoodModel struct {
	Field1 string   `validate:"in:test,test2"`
	Field2 int      `validate:"min:5"`
	Field3 []int    `validate:"min:5|max:10|in:200,300"`
	Field4 []string `validate:"len:10|regexp:^\\w+@\\w+\\.\\w+$"`
}
