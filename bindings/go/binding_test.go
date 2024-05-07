package tree_sitter_astro_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-astro"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_astro.Language())
	if language == nil {
		t.Errorf("Error loading Astro grammar")
	}
}
