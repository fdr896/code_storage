package core

import "testing"

func TestNewCodeWithValidCode(t *testing.T) {
	code := &Code{
		Source:   "test code",
		Language: "test lang",
	}

	if err := code.NewCode(); err != nil {
		t.Errorf("NewCode failed with this error: %v", err)
	}
}

func TestNewCodeWithEmptyAndOnlyBlankCharsFields(t *testing.T) {
	code := &Code{}

	if err := code.NewCode(); err == nil {
		t.Errorf("NewCode didn't handle empty fields properly")
	}

	code = &Code{
		Source:   "   ",
		Language: "test lang",
	}

	if err := code.NewCode(); err != ErrEmptySource {
		t.Errorf("NewCode didn't recognise Source field which contains only blank characters")
	}

	code = &Code{
		Source:   "test code",
		Language: "     ",
	}

	if err := code.NewCode(); err != ErrEmptyLanguage {
		t.Errorf("NewCode didn't recognise empty Language field which contains only blank characters")
	}
}
