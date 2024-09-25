type TrieNode struct {
	children map[byte]*TrieNode
	count    int
}

type Trie struct {
	root *TrieNode
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[byte]*TrieNode),
	}
}

func NewTrie() *Trie {
	return &Trie{
		root: newTrieNode(),
	}
}

func (trie *Trie) Insert(word string) {
	node := trie.root
	for i := 0; i < len(word); i++ {
		ch := word[i]
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = newTrieNode()
		}
		node = node.children[ch]
		node.count++
	}
}

func (trie *Trie) GetPrefixScoreSum(word string) int {
	node := trie.root
	scoreSum := 0

	for i := 0; i < len(word); i++ {
		ch := word[i]
		node = node.children[ch]
		scoreSum += node.count
	}
	return scoreSum
}

func sumPrefixScores(words []string) []int {
	trie := NewTrie()

	for _, word := range words {
		trie.Insert(word)
	}

	result := make([]int, len(words))

	for i, word := range words {
		result[i] = trie.GetPrefixScoreSum(word)
	}

	return result
}
