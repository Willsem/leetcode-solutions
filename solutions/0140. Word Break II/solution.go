func wordBreak(s string, wordDict []string) []string {
	t := NewTrie()
	for i := 0; i < len(wordDict); i++ {
		t.AddWord(wordDict[i])
	}

	return solve(0, s, "", []string{}, t)
}

func solve(ind int, s, curr string, answer []string, t *Trie) []string {
	if ind == len(s) {
		answer = append(answer, curr[1:])
		return answer
	}

	for i := ind + 1; i <= len(s); i++ {
		substr := s[ind:i]
		if t.search(substr) {
			answer = solve(i, s, curr+" "+substr, answer, t)
		}
	}

	return answer
}

type Trie struct {
	Nodes []*Trie
	End   bool
}

func NewTrie() *Trie {
	return &Trie{
		Nodes: make([]*Trie, 26),
		End:   false,
	}
}

func (t *Trie) AddWord(word string) {
	ptr := t
	for i := 0; i < len(word); i++ {
		if ptr.Nodes[word[i]-'a'] == nil {
			ptr.Nodes[word[i]-'a'] = NewTrie()
		}
		ptr = ptr.Nodes[word[i]-'a']
	}

	ptr.End = true
}

func (t *Trie) search(word string) bool {
	ptr := t
	for i := 0; i < len(word); i++ {
		if ptr.Nodes[word[i]-'a'] == nil {
			return false
		}

		ptr = ptr.Nodes[word[i]-'a']
	}

	return ptr.End
}
