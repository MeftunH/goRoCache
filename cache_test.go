package goRoCache

import (
	"errors"
	"reflect"
	"testing"
)

func TestIsAlreadyExists(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAlreadyExists(tt.args.err); got != tt.want {
				t.Errorf("IsAlreadyExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsInvalidKeyType(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInvalidKeyType(tt.args.err); got != tt.want {
				t.Errorf("IsInvalidKeyType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsInvalidMessage(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInvalidMessage(tt.args.err); got != tt.want {
				t.Errorf("IsInvalidMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNilUpdateFunc(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNilUpdateFunc(tt.args.err); got != tt.want {
				t.Errorf("IsNilUpdateFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNonPositivePeriod(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNonPositivePeriod(tt.args.err); got != tt.want {
				t.Errorf("IsNonPositivePeriod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUnexpectedError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUnexpectedError(tt.args.err); got != tt.want {
				t.Errorf("IsUnexpectedError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cacheError_Error(t *testing.T) {
	type fields struct {
		msg         string
		errType     errorType
		nestedError error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ce := cacheError{
				msg:         tt.fields.msg,
				errType:     tt.fields.errType,
				nestedError: tt.fields.nestedError,
			}
			if got := ce.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newError(t *testing.T) {
	type args struct {
		errType errorType
		msg     string
	}
	tests := []struct {
		name string
		args args
		want cacheError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newError(tt.args.errType, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newWrapperError(t *testing.T) {
	type args struct {
		errType     errorType
		msg         string
		nestedError error
	}
	tests := []struct {
		name string
		args args
		want cacheError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newWrapperError(tt.args.errType, tt.args.msg, tt.args.nestedError); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newWrapperError() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestNewErrorWithProvidedValues(t *testing.T) {
	errType := errorType("test")
	msg := "test message"
	result := newError(errType, msg)

	if result.msg != msg {
		t.Errorf("Expected msg to be %s, but got %s", msg, result.msg)
	}

	if result.errType != errType {
		t.Errorf("Expected errType to be %s, but got %s", errType, result.errType)
	}

	if result.nestedError != nil {
		t.Errorf("Expected nestedError to be nil, but got %v", result.nestedError)
	}
}
func TestNewErrorWithNilNestedError(t *testing.T) {
	errType := errorType("test")
	msg := "test message"
	result := newError(errType, msg)

	if result.nestedError != nil {
		t.Errorf("Expected nestedError to be nil, but got %v", result.nestedError)
	}
}
func TestNewErrorWithEmptyMsg(t *testing.T) {
	errType := errorType("test")
	result := newError(errType, "")

	if result.msg != "" {
		t.Errorf("Expected msg to be empty, but got %s", result.msg)
	}
}
func TestNewErrorWithEmptyStringMsg(t *testing.T) {
	errType := errorType("test")
	result := newError(errType, "")

	if result.msg != "" {
		t.Errorf("Expected msg to be empty, but got %s", result.msg)
	}
}
func TestNewWrapperError(t *testing.T) {
	tests := []struct {
		name        string
		errType     errorType
		msg         string
		nestedError error
		want        cacheError
	}{
		{
			name:        "Test with valid inputs",
			errType:     errorTypeUnexpectedError,
			msg:         "Unexpected error occurred",
			nestedError: errors.New("Nested error"),
			want: cacheError{
				msg:         "Unexpected error occurred",
				errType:     errorTypeUnexpectedError,
				nestedError: errors.New("Nested error"),
			},
		},
		{
			name:        "Test with empty message",
			errType:     errorTypeUnexpectedError,
			msg:         "",
			nestedError: errors.New("Nested error"),
			want: cacheError{
				msg:         "",
				errType:     errorTypeUnexpectedError,
				nestedError: errors.New("Nested error"),
			},
		},
		{
			name:        "Test with nil nested error",
			errType:     errorTypeUnexpectedError,
			msg:         "Unexpected error occurred",
			nestedError: nil,
			want: cacheError{
				msg:         "Unexpected error occurred",
				errType:     errorTypeUnexpectedError,
				nestedError: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newWrapperError(tt.errType, tt.msg, tt.nestedError); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newWrapperError() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestErrorTypeFunctions(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want map[string]bool
	}{
		{
			name: "Test with UnexpectedError",
			err:  newError(errorTypeUnexpectedError, "Unexpected error occurred"),
			want: map[string]bool{
				"IsUnexpectedError":   true,
				"IsAlreadyExists":     false,
				"IsNonPositivePeriod": false,
				"IsNilUpdateFunc":     false,
				"IsInvalidKeyType":    false,
				"IsInvalidMessage":    false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUnexpectedError(tt.err); got != tt.want["IsUnexpectedError"] {
				t.Errorf("IsUnexpectedError() = %v, want %v", got, tt.want["IsUnexpectedError"])
			}
			if got := IsAlreadyExists(tt.err); got != tt.want["IsAlreadyExists"] {
				t.Errorf("IsAlreadyExists() = %v, want %v", got, tt.want["IsAlreadyExists"])
			}
			if got := IsNonPositivePeriod(tt.err); got != tt.want["IsNonPositivePeriod"] {
				t.Errorf("IsNonPositivePeriod() = %v, want %v", got, tt.want["IsNonPositivePeriod"])
			}
			if got := IsNilUpdateFunc(tt.err); got != tt.want["IsNilUpdateFunc"] {
				t.Errorf("IsNilUpdateFunc() = %v, want %v", got, tt.want["IsNilUpdateFunc"])
			}
			if got := IsInvalidKeyType(tt.err); got != tt.want["IsInvalidKeyType"] {
				t.Errorf("IsInvalidKeyType() = %v, want %v", got, tt.want["IsInvalidKeyType"])
			}
			if got := IsInvalidMessage(tt.err); got != tt.want["IsInvalidMessage"] {
				t.Errorf("IsInvalidMessage() = %v, want %v", got, tt.want["IsInvalidMessage"])
			}
		})
	}
}
