package kana

type Trie struct {
	children map[string]*Trie
	letter   string
	values   []string
}

func newTrie() *Trie {
	/*
	   Build a trie for efficient retrieval of entries
	*/
	var root *Trie = &Trie{map[string]*Trie{}, "", []string{}}
	return root
}

func (t *Trie) insert(letters, value string) {
	/*
		Insert a value into the trie
	*/

	lettersRune := []rune(letters)

	// loop through letters in argument word
	for l, letter := range lettersRune {

		letterStr := string(letter)

		// if letter in children
		if t.children[letterStr] != nil {
			t = t.children[letterStr]
		} else {
			// not found, so add letter to children
			t.children[letterStr] = &Trie{map[string]*Trie{}, "", []string{}}
			t = t.children[letterStr]
		}

		if l == len(lettersRune)-1 {
			// last letter, save value and exit
			t.values = append(t.values, value)
			break
		}
	}
}

func (t *Trie) search(srch string) (found []string) {
	/*
		Search for a string in the Trie.

		Returns the corresponding array of strings if found,
		or an empty array otherwise.
	*/
	srchRune := []rune(srch)

	for l, letter := range srchRune {
		letterString := string(letter)
		if t.children[letterString] != nil {
			t = t.children[letterString]
		} else {
			found = []string{""}
			return found
		}
		if l == len(srchRune)-1 {
			found = t.values
		}
	}
	return found
}

func (t *Trie) convert(origin string) (result string) {
	/*
		Convert a given string to the corresponding values
		in the trie. This performed in a greedy fashion,
		replacing the longest valid string it can find at any
		given point.
	*/
	root := t
	originRune := []rune(origin)
	result = ""

	for l := 0; l < len(originRune); l++ {
		t = root
		foundVal := ""
		depth := 0
		for i := 0; i+l < len(originRune); i++ {
			letter := string(originRune[l+i])
			if t.children[letter] == nil {
				// not found
				break
			} else {
				if len(t.children[letter].values) > 0 {
					foundVal = t.children[letter].values[0]
					depth = i
				}
				t = t.children[letter]
			}
		}
		if foundVal != "" {
			result += foundVal
			l += depth
		} else {
			result += string(originRune[l : l+1])
		}
	}
	return result
}
