package helpers_test

import (
	"testing"
	s "shortify/internal/helpers"
)

func TestBase62Encode(t *testing.T)  {

	type testStruct struct {
		input		uint64
		expected	string
	}

	testData := []testStruct{
		{1, "1"},
		{2, "2"},
		{3, "3"},
		{61, "Z"},
		{62, "10"},
		{123456789, "8m0Kx"},
		// Проверка уникальности
		{123456789, "8m0Kx"},
	}

	for _, v := range testData {
		test := s.Base62Encode(v.input)
		if test != v.expected {
			t.Errorf("Несовпадение полученного %s и ожидаемого %s", test, v.expected)
		} else {
			t.Logf("%s = %s", test, v.expected)
		}
	}
}