package parser_test

import (
	"testing"

	"github.com/aleixcam/gopher/board-kata/parser"
)

func TestTable(t *testing.T) {
	tests := map[string]struct{ parsable, expected string }{
		"url":  {"https://friendsofgo.tech", `<a href="https://friendsofgo.tech">https://friendsofgo.tech</a>`},
		"tag":  {"@FriendsOfGoTech", `<a href="https://fogo-parser.dev/FriendsOfGoTech">@FriendsOfGoTech</a>`},
		"hash": {"#friendsofgo", `<a href="https://fogo-parser.dev/hash/friendsofgo">#friendsofgo</a>`},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := parser.Parse(tc.parsable)
			if got != tc.expected {
				t.Errorf("Wrong parse - expected %s, got: %s", tc.expected, got)
			}
		})
	}
}
