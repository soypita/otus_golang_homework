package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasic(t *testing.T) {
	t.Run("should return error when process non go source file", func(t *testing.T) {
		wrongTypeFile := "test/dump_file.yml"
		err := processFile(wrongTypeFile)
		require.NotNil(t, err)
	})

	t.Run("should succesfully generate validation source file", func(t *testing.T) {
		sourceFile := "test/good_model.go"
		validationSourceFile := "test/good_model_validation.go"
		err := processFile(sourceFile)
		require.Nil(t, err)

		_, err = os.Stat(validationSourceFile)
		require.Nil(t, err)
		os.Remove(validationSourceFile)
	})

	t.Run("should return error when process wrong field type-tag connection", func(t *testing.T) {
		sourceFile := "test/wrong_type_model.go"
		validationSourceFile := "test/wrong_type_model_validation.go"
		err := processFile(sourceFile)
		require.NotNil(t, err)
		_, err = os.Stat(validationSourceFile)
		require.NotNil(t, err)
	})

	t.Run("should return error when process unknown tag", func(t *testing.T) {
		sourceFile := "test/unsupported_tag_model.go"
		validationSourceFile := "test/unsupported_tag_model_validation.go"
		err := processFile(sourceFile)
		require.NotNil(t, err)
		_, err = os.Stat(validationSourceFile)
		require.NotNil(t, err)
	})

	t.Run("should succesfully process model with user defined acceptable type", func(t *testing.T) {
		sourceFile := "test/model_with_user_type.go"
		validationSourceFile := "test/model_with_user_type_validation.go"
		err := processFile(sourceFile)
		require.Nil(t, err)
		os.Remove(validationSourceFile)
	})

	t.Run("should return error when process model with unsupported user type", func(t *testing.T) {
		sourceFile := "test/model_with_user_wrong_type.go"
		validationSourceFile := "test/model_with_user_wrong_type_validation.go"
		err := processFile(sourceFile)
		require.NotNil(t, err)
		_, err = os.Stat(validationSourceFile)
		require.NotNil(t, err)
	})

	t.Run("should return error when process model with multi-dimension array", func(t *testing.T) {
		sourceFile := "test/nested_arrays_model.go"
		validationSourceFile := "test/nested_arrays_model_validation.go"
		err := processFile(sourceFile)
		require.NotNil(t, err)
		_, err = os.Stat(validationSourceFile)
		require.NotNil(t, err)
	})

	t.Run("should return error when process model with multi-dimension user type array", func(t *testing.T) {
		sourceFile := "test/nested_array_user_type.go"
		validationSourceFile := "test/nested_array_user_type_validation.go"
		err := processFile(sourceFile)
		require.NotNil(t, err)
		_, err = os.Stat(validationSourceFile)
		require.NotNil(t, err)
	})
}
