package funktion

import (
	"reflect"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", true},
		{"whitespace only", "   ", true},
		{"tabs and spaces", "\t  \n", true},
		{"non-empty string", "hello", false},
		{"string with content", "  hello  ", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEmpty(tt.input)
			if result != tt.expected {
				t.Errorf("IsEmpty(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsBlank(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", true},
		{"whitespace only", "   ", true},
		{"non-empty string", "hello", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsBlank(tt.input)
			if result != tt.expected {
				t.Errorf("IsBlank(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsStruct(t *testing.T) {
	type TestStruct struct {
		Name string
	}

	tests := []struct {
		name     string
		input    interface{}
		expected bool
	}{
		{"struct", TestStruct{}, true},
		{"string", "hello", false},
		{"int", 42, false},
		{"slice", []string{}, false},
		{"map", make(map[string]string), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsStruct(tt.input)
			if result != tt.expected {
				t.Errorf("IsStruct(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected bool
	}{
		{"string slice", []string{}, true},
		{"int slice", []int{}, true},
		{"string", "hello", false},
		{"int", 42, false},
		{"map", make(map[string]string), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsSlice(tt.input)
			if result != tt.expected {
				t.Errorf("IsSlice(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected []interface{}
	}{
		{"string slice", []string{"a", "b", "c"}, []interface{}{"a", "b", "c"}},
		{"int slice", []int{1, 2, 3}, []interface{}{1, 2, 3}},
		{"empty slice", []string{}, []interface{}{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToSlice(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ToSlice(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestArrayContains(t *testing.T) {
	tests := []struct {
		name     string
		list     []string
		contains string
		expected bool
	}{
		{"contains item", []string{"apple", "banana", "orange"}, "banana", true},
		{"contains item case insensitive", []string{"Apple", "Banana", "Orange"}, "banana", true},
		{"does not contain item", []string{"apple", "banana", "orange"}, "grape", false},
		{"empty list", []string{}, "apple", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ArrayContains(tt.list, tt.contains)
			if result != tt.expected {
				t.Errorf("ArrayContains(%v, %q) = %v, want %v", tt.list, tt.contains, result, tt.expected)
			}
		})
	}
}

func TestSliceContains(t *testing.T) {
	tests := []struct {
		name     string
		in       []string
		this     []string
		expected bool
	}{
		{"has overlap", []string{"a", "b", "c"}, []string{"b", "d"}, true},
		{"no overlap", []string{"a", "b", "c"}, []string{"d", "e"}, false},
		{"case insensitive", []string{"Apple", "Banana"}, []string{"apple"}, true},
		{"empty slices", []string{}, []string{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SliceContains(tt.in, tt.this)
			if result != tt.expected {
				t.Errorf("SliceContains(%v, %v) = %v, want %v", tt.in, tt.this, result, tt.expected)
			}
		})
	}
}

func TestSplitLines(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{"single line", "hello world", []string{"hello    world"}},
		{"multiple lines", "line1\nline2\nline3", []string{"line1", "line2", "line3"}},
		{"empty string", "", nil},
		{"with tabs", "hello\tworld", []string{"hello    world"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SplitLines(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SplitLines(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestTruncatePrint(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		length   int
		expected string
	}{
		{"shorter than limit", "hello", 10, "hello"},
		{"longer than limit", "hello world", 5, "hello..."},
		{"exact length", "hello", 5, "hello..."},
		{"empty string", "", 5, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TruncatePrint(tt.input, tt.length)
			if result != tt.expected {
				t.Errorf("TruncatePrint(%q, %d) = %q, want %q", tt.input, tt.length, result, tt.expected)
			}
		})
	}
}

func TestIsEmail(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{"valid email", "test@example.com", true},
		{"valid email with subdomain", "user@mail.example.com", true},
		{"invalid email no @", "testexample.com", false},
		{"invalid email no domain", "test@", false},
		{"empty string", "", false},
		{"too short", "a@", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEmail(tt.email)
			if result != tt.expected {
				t.Errorf("IsEmail(%q) = %v, want %v", tt.email, result, tt.expected)
			}
		})
	}
}

func TestIsValidEmail(t *testing.T) {
	domains := []string{"example.com", "test.org"}

	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{"valid domain", "user@example.com", true},
		{"valid domain 2", "user@test.org", true},
		{"invalid domain", "user@invalid.com", false},
		{"bobmcallan exception", "bobmcallan@anywhere.com", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidEmail(tt.email, domains)
			if result != tt.expected {
				t.Errorf("IsValidEmail(%q, %v) = %v, want %v", tt.email, domains, result, tt.expected)
			}
		})
	}
}

func TestIsValidDomain(t *testing.T) {
	tests := []struct {
		name     string
		domain   string
		expected bool
	}{
		{"valid domain 1", "procul.io", true},
		{"valid domain 2", "dashs.com", true},
		{"valid domain 3", "t3b.io", true},
		{"invalid domain", "google.com", false},
		{"empty domain", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidDomain(tt.domain)
			if result != tt.expected {
				t.Errorf("IsValidDomain(%q) = %v, want %v", tt.domain, result, tt.expected)
			}
		})
	}
}

func TestToJson(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		wantErr bool
	}{
		{"simple struct", struct{ Name string }{"test"}, false},
		{"map", map[string]string{"key": "value"}, false},
		{"slice", []string{"a", "b", "c"}, false},
		{"nil", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToJson(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJson(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && result == "" {
				t.Errorf("ToJson(%v) returned empty string", tt.input)
			}
		})
	}
}

func TestToJsonFlat(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		wantErr bool
	}{
		{"simple struct", struct{ Name string }{"test"}, false},
		{"map", map[string]string{"key": "value"}, false},
		{"slice", []string{"a", "b", "c"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToJsonFlat(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJsonFlat(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && result == "" {
				t.Errorf("ToJsonFlat(%v) returned empty string", tt.input)
			}
		})
	}
}
