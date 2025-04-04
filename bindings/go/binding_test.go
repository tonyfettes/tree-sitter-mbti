package tree_sitter_mbti_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_mbti "github.com/moonbitlang/tree-sitter-mbti/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_mbti.Language())
	if language == nil {
		t.Errorf("Error loading MoonBit Interface grammar")
	}
}
