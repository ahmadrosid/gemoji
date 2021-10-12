package main_test

import (
	"testing"

	gemoji "github.com/ahmadrosid/gemoji"
)

func Test_FormatString(t *testing.T) {
	formatted := gemoji.Format("This is :rocket: and :smile:")
	expected := "This is 🚀 and 😄"
	if formatted != expected {
		t.Fatalf("expected '%s' got '%s'", expected, formatted)
	}
}

func Test_FormatIgnoreInvalidEmoji(t *testing.T) {
	formatted := gemoji.Format("This is :heartbeat: and :invalid:")
	expected := "This is 💓 and :invalid:"
	if formatted != expected {
		t.Fatalf("expected '%s' got '%s'", expected, formatted)
	}
}

func Test_GetEmoji(t *testing.T) {
	formatted := gemoji.Get(":heartbeat:")
	expected := "💓"
	if formatted != expected {
		t.Fatalf("expected '%s' got '%s'", expected, formatted)
	}
}

func Test_GetEmojiNoColon(t *testing.T) {
	formatted := gemoji.Get("rocket")
	expected := "🚀"
	if formatted != expected {
		t.Fatalf("expected '%s' got '%s'", expected, formatted)
	}
}
