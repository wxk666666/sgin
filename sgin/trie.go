package sgin8

import "strings"

type TrieNode struct {
	pattern  string      // 待匹配路由，例如 /p/:lang
	part     string      // 路由中的一部分，例如 :lang
	children []*TrieNode // 子节点，例如 [doc, tutorial, intro]
	isWild   bool        // 是否精确匹配，part 含有 : 或 * 时为true
}

// matchChild 获得第一个匹配（等于part，或者isWild为true，即模糊匹配）part的child
func (n *TrieNode) matchChild(part string) *TrieNode {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChildes 获得所有匹配（等于part，或者isWild为true，即模糊匹配）part的child们
func (n *TrieNode) matchChildes(part string) []*TrieNode {
	res := make([]*TrieNode, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			res = append(res, child)
		}
	}
	return res
}

func (n *TrieNode) insert(pattern string, parts []string, height int) {
	//字符已经全部插入了，return
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	//只需要找到一个满足当前层的
	child := n.matchChild(part)
	if child == nil {
		//若下一层的孩子含‘:’ or '*'则isWild为true，即不需要part相同,用于搜索
		child = &TrieNode{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	//继续下一层
	child.insert(pattern, parts, height+1)
}

// Search
// For example, p/:lang/doc,you can match p/c/doc,or you can match p/go/doc.
// For example, /static/*filepath, you can match /static/fav.ico, or you can match /static/js/jQuery.js.
// Tips: /static/*filepath/go can not match /static/all/go or /static/all/code/go
func (n *TrieNode) search(parts []string, height int) *TrieNode {
	//字符已经全部搜索入了，或者当前层有‘*’的前缀,但是待匹配的路由为““(也就是说，没完全匹配完路径，而是在半路上)则直接返回
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildes(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil

}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
