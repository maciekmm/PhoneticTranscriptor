package phonetictranscriptor

import (
	"bufio"
	"os"
	"sort"
)

var (
	languages = make(map[string]Language)
)

// LoadLanguage loads language from languages directory
func LoadLanguage(code string) (ln Language, err error) {
	file, err := os.Open("./languages/" + code + ".lang")
	if err != nil {
		return
	}
	defer file.Close()

	ln.Code = code

	scanner := bufio.NewScanner(file)
	var rawLine string
	for ; scanner.Scan(); rawLine = scanner.Text() {
		if len(rawLine) == 0 {
			continue
		}
		line := []rune(rawLine)
		//Construct new rule
		rule := Rule{}

		hadWhitespace := false

	CharIterator:
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case '#':
				break CharIterator
			case ' ':
				hadWhitespace = true
			default:
				if !hadWhitespace {
					rule.From += string(line[i])
				} else {
					rule.To += string(line[i])
				}
			}
		}
		if len(rule.From) != 0 && len(rule.To) != 0 {
			ln.Rules = append(ln.Rules, rule)
		}
	}
	sort.Sort(sort.Reverse(ln.Rules))
	languages[code] = ln
	return
}

// Language contains rules for specific language
type Language struct {
	Code  string //ISO Code
	Rules rules
}

// Rule defines a rule for transcripting
type Rule struct {
	From string
	To   string
}

type rules []Rule

func (slice rules) Len() int {
	return len(slice)
}

func (slice rules) Less(i, j int) bool {
	return len([]rune(slice[i].From)) < len([]rune(slice[j].From))
}

func (slice rules) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
