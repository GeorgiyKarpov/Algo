package WordDictionary

type WordDictionary struct {
	children    map[rune]*WordDictionary
	IsEndOfWord bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		children:    make(map[rune]*WordDictionary),
		IsEndOfWord: false,
	}
}

func (this *WordDictionary) AddWord(word string) {
	currentNode := this
	for _, ch := range word {
		if _, ok := currentNode.children[ch]; !ok {
			newNode := Constructor()
			currentNode.children[ch] = &newNode
		}
		currentNode = currentNode.children[ch]
	}
	currentNode.IsEndOfWord = true
}

func (this *WordDictionary) Search(word string) bool {
	currentNode := this
	for i, ch := range word {
		if _, ok := currentNode.children[ch]; !ok {
			if ch != '.' {
				return false
			}
			for _, node := range currentNode.children {
				if node.Search(word[i+1:]) {
					return true
				}
			}
			return false
		}
		currentNode = currentNode.children[ch]
	}
	return currentNode.IsEndOfWord
}
