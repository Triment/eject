package eject

import (
	"strings"
)

//Trie Tree
type Trie struct {
	Part     string           `json:"path"`
	Children map[string]*Trie `json:"children"`
	IsLeaf   bool             `json:"isleaf"`
	FullPath string           `json:"fullpath"`
}

//Insert Node into Trie
func (t *Trie) Insert(path string) {
	//split path
	paths := strings.Split(path, "/")[1:]
	for _, f := range paths {
		if child, exist := t.Children[f]; exist {
			t = child
		} else if len(f) > 0 {
			isLeaf := f[0] == '*'
			nextTrie := &Trie{Part: f, Children: make(map[string]*Trie)} //if start by '*' it is leaf
			t.Children[f] = nextTrie                                     //move new trie to parent node's children
			t = nextTrie                                                 //pointed it to new trie
			if isLeaf {
				break
			}
		} else {
			break
		}
	}
	t.IsLeaf = true   //set leaf node to true
	t.FullPath = path //set full path
}

//Search Node from Trie
func (t *Trie) Search(paths []string, length int, index int, form map[string]string) *Trie {
	if index == length || len(paths[index])==0 { //the latest item of paths
		return t
	}
	if child, exist := t.Children[paths[index]]; exist { //精确匹配
		return child.Search(paths, length, index+1, form)
	}
	for key, trie := range t.Children { //模糊匹配
		if key[0] == ':' {
			form[key[1:]] = paths[index]
			if length == index+1 {
				return trie
			}
			return trie.Search(paths, length, index+1, form)
		}
		if key[0] == '*' {
			form[key[1:]] = strings.Join(paths[index:], "/")
			return trie
		}
	}
	return nil
}
