package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_unpackSequence(t *testing.T) {
	testTable := []struct {
		name        string
		str         string
		expectedStr string
		expectedErr error
	}{
		{
			name:        "№1",
			str:         "a4bc2d5e",
			expectedStr: "aaaabccddddde",
			expectedErr: nil,
		},
		{
			name:        "№2",
			str:         "abcd",
			expectedStr: "abcd",
			expectedErr: nil,
		},
		{
			name:        "№3",
			str:         "45",
			expectedStr: "",
			expectedErr: errors.New("invalid string"),
		},
		{
			name:        "№4",
			str:         "",
			expectedStr: "",
			expectedErr: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			var s Unpacker
			s = Sequence{str: testCase.str}
			str, err := s.unpack()
			assert.Equal(t, testCase.expectedStr, str)
			assert.Equal(t, testCase.expectedErr, err)

		})
	}
}

func Test_unpackEscapeSequence(t *testing.T) {
	testTable := []struct {
		name        string
		str         string
		expectedStr string
		expectedErr error
	}{
		{
			name:        "№1",
			str:         "qwe\\4\\5",
			expectedStr: "qwe45",
			expectedErr: nil,
		},
		{
			name:        "№2",
			str:         "qwe\\45",
			expectedStr: "qwe44444",
			expectedErr: nil,
		},
		{
			name:        "№3",
			str:         "qwe\\\\5",
			expectedStr: "qwe\\\\\\\\\\",
			expectedErr: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			var s Unpacker

			s = EscapeSequence{str: testCase.str}

			str, err := s.unpack()
			assert.Equal(t, testCase.expectedStr, str)
			assert.Equal(t, testCase.expectedErr, err)

		})
	}
}
