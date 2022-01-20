package eject_test

import (
	"strings"
	"testing"

	"github.com/Triment/eject"
)

func getPath(path string) ([]string, int) {
	paths := strings.Split(path, "/")[1:]
	return paths, len(paths)
}
func TestTrieInsert(t *testing.T) {
	trie := eject.Trie{Children: map[string]*eject.Trie{}}
	trie.Insert("/hello/word/*hh")
	trie.Insert("/hello/word")
	paths, length := getPath("/hello/word/123/hekoj")
	form := make(map[string]string)
	t.Log(paths)
	if trie.Search(paths, length, 0, form) == nil {
		t.Error("tree search failed")
	}
	if form["hh"] != "123/hekoj" {
		t.Error("args decode failed")
	}
	paths, length = getPath("/hello/word")
	if trie.Search(paths, length, 0, form).Part != "word" {
		t.Error("decode failed")
	}
}
