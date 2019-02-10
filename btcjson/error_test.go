// Copyright (c) 2014 The ohmcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package ohmcjson_test

import (
	"testing"

	"github.com/ohmcsuite/ohmcd/ohmcjson"
)

// TestErrorCodeStringer tests the stringized output for the ErrorCode type.
func TestErrorCodeStringer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   ohmcjson.ErrorCode
		want string
	}{
		{ohmcjson.ErrDuplicateMethod, "ErrDuplicateMethod"},
		{ohmcjson.ErrInvalidUsageFlags, "ErrInvalidUsageFlags"},
		{ohmcjson.ErrInvalidType, "ErrInvalidType"},
		{ohmcjson.ErrEmbeddedType, "ErrEmbeddedType"},
		{ohmcjson.ErrUnexportedField, "ErrUnexportedField"},
		{ohmcjson.ErrUnsupportedFieldType, "ErrUnsupportedFieldType"},
		{ohmcjson.ErrNonOptionalField, "ErrNonOptionalField"},
		{ohmcjson.ErrNonOptionalDefault, "ErrNonOptionalDefault"},
		{ohmcjson.ErrMismatchedDefault, "ErrMismatchedDefault"},
		{ohmcjson.ErrUnregisteredMethod, "ErrUnregisteredMethod"},
		{ohmcjson.ErrNumParams, "ErrNumParams"},
		{ohmcjson.ErrMissingDescription, "ErrMissingDescription"},
		{0xffff, "Unknown ErrorCode (65535)"},
	}

	// Detect additional error codes that don't have the stringer added.
	if len(tests)-1 != int(ohmcjson.TstNumErrorCodes) {
		t.Errorf("It appears an error code was added without adding an " +
			"associated stringer test")
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.String()
		if result != test.want {
			t.Errorf("String #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}

// TestError tests the error output for the Error type.
func TestError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   ohmcjson.Error
		want string
	}{
		{
			ohmcjson.Error{Description: "some error"},
			"some error",
		},
		{
			ohmcjson.Error{Description: "human-readable error"},
			"human-readable error",
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.Error()
		if result != test.want {
			t.Errorf("Error #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}
