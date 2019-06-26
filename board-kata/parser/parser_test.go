package parser_test

import (
	"testing"

	"github.com/aleixcam/gopher/board-kata/parser"
)

func TestParseUrl(t *testing.T) {
	expected := `<a href="https://friendsofgo.tech">https://friendsofgo.tech</a>`
	got := parser.Parse("https://friendsofgo.tech")
	if got != expected {
		t.Errorf("Wrong parse - expected %s, got: %s", expected, got)
	}
}

func TestParseTag(t *testing.T) {
	expected := `<a href="https://fogo-parser.dev/FriendsOfGoTech">@FriendsOfGoTech</a>`
	got := parser.Parse("@FriendsOfGoTech")
	if got != expected {
		t.Errorf("Wrong parse - expected %s, got: %s", expected, got)
	}
}

func TestParseHash(t *testing.T) {
	expected := `<a href="https://fogo-parser.dev/hash/friendsofgo">#friendsofgo</a>`
	got := parser.Parse("#friendsofgo")
	if got != expected {
		t.Errorf("Wrong parse - expected %s, got: %s", expected, got)
	}
}
