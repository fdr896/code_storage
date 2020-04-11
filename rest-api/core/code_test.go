package core

import "testing"

func TestNewCodeWithValidCode(t *testing.T) {
	code := &Code{
		Source:      "test code",
		Language:    "test lang",
		Description: "test descr",
	}

	if err := code.Validate(); err != nil {
		t.Errorf("NewCode failed on proper request with this error: %v", err)
	}
}

func TestNewCodeWithEmptyAndOnlyBlankCharsFields(t *testing.T) {
	code := &Code{}

	if err := code.Validate(); err == nil {
		t.Errorf("NewCode didn't handle empty fields properly")
	}

	code = &Code{
		Source:      "   ",
		Language:    "test lang",
		Description: "test descr",
	}

	if err := code.Validate(); err != ErrEmptySource {
		t.Errorf("NewCode didn't recognise Source field which contains only blank characters")
	}

	code = &Code{
		Source:      "test code",
		Language:    "     ",
		Description: "test descr",
	}

	if err := code.Validate(); err != ErrEmptyLanguage {
		t.Errorf("NewCode didn't recognise empty Language field which contains only blank characters")
	}

	code = &Code{
		Source:      "test code",
		Language:    "test lang",
		Description: "\t\t\n",
	}

	if err := code.Validate(); err != ErrEmptyDescription {
		t.Errorf("NewCode didn't recignise empty Description field which contains only blank characters")
	}
}
