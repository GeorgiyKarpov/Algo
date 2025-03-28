package implementTrie

type Trie struct {
	children    map[rune]*Trie
	isEndOfWord bool
}

func Constructor() Trie {
	return Trie{
		children:    make(map[rune]*Trie),
		isEndOfWord: false,
	}
}

func (this *Trie) Insert(word string) {
	currentNode := this
	for _, ch := range word {
		if _, ok := currentNode.children[ch]; !ok {
			newTrieNode := Constructor()
			currentNode.children[ch] = &newTrieNode
		}
		currentNode = currentNode.children[ch]
	}
	currentNode.isEndOfWord = true
}

func (this *Trie) Search(word string) bool {
	currentNode := this
	for _, ch := range word {
		if _, ok := currentNode.children[ch]; !ok {
			return false
		}
		currentNode = currentNode.children[ch]
	}
	return currentNode.isEndOfWord
}

func (this *Trie) StartsWith(prefix string) bool {
	currentNode := this
	for _, ch := range prefix {
		if _, ok := currentNode.children[ch]; !ok {
			return false
		}
		currentNode = currentNode.children[ch]
	}
	return true
}
