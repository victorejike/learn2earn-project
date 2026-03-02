package main

import (
	"regexp"
	"strconv"
	"strings"
)

func Process(text string) string {
	text = strings.ReplaceAll(text, "\n", " ")

	// Split text but keep modifiers and quotes as single tokens
	tokens := smartTokenize(text)
	tokens = applyModifiers(tokens)
	tokens = fixArticles(tokens)

	result := strings.Join(tokens, " ")
	result = fixPunctuation(result)
	result = fixQuotes(result)

	return strings.TrimSpace(result)
}

// Smart tokenizer: keeps (cap,6), (up,2), (low,3) and quoted text as single tokens
func smartTokenize(text string) []string {
	re := regexp.MustCompile(`\([a-z]+(,\s*\d+)?\)|'[^']*'|[^\s]+`)
	return re.FindAllString(text, -1)
}

func applyModifiers(tokens []string) []string {
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		if token == "(hex)" && i > 0 {
			val, err := strconv.ParseInt(tokens[i-1], 16, 64)
			if err == nil {
				tokens[i-1] = strconv.FormatInt(val, 10)
			}
			tokens = removeToken(tokens, i)
			i--
			continue
		}

		if token == "(bin)" && i > 0 {
			val, err := strconv.ParseInt(tokens[i-1], 2, 64)
			if err == nil {
				tokens[i-1] = strconv.FormatInt(val, 10)
			}
			tokens = removeToken(tokens, i)
			i--
			continue
		}

		if strings.HasPrefix(token, "(") {
			cmd, count := parseModifier(token)
			if cmd == "" {
				continue
			}

			start := i - count
			if start < 0 {
				start = 0
			}

			for j := start; j < i; j++ {
				switch cmd {
				case "up":
					tokens[j] = strings.ToUpper(tokens[j])
				case "low":
					tokens[j] = strings.ToLower(tokens[j])
				case "cap":
					tokens[j] = Capitalize(tokens[j])
				}
			}

			tokens = removeToken(tokens, i)
			i--
		}
	}

	return tokens
}

func parseModifier(token string) (string, int) {
	re := regexp.MustCompile(`^\((up|low|cap)(,\s*(\d+))?\)$`)
	matches := re.FindStringSubmatch(token)
	if len(matches) == 0 {
		return "", 1
	}

	cmd := matches[1]

	if matches[3] != "" {
		n, _ := strconv.Atoi(matches[3])
		return cmd, n
	}

	return cmd, 1
}

func fixArticles(tokens []string) []string {
	for i := 0; i < len(tokens)-1; i++ {
		word := tokens[i]
		next := strings.ToLower(tokens[i+1])
		if strings.ToLower(word) == "a" && startsWithVowelOrH(next) {
			if word == "A" {
				tokens[i] = "An"
			} else {
				tokens[i] = "an"
			}
		}
	}
	return tokens
}

// Fixed punctuation handling
func fixPunctuation(text string) string {
	// 1. Ellipsis and special groups first
	reEllipsis := regexp.MustCompile(`\s*(\.\.\.|!\?|!!|\?\!)\s*`)
	text = reEllipsis.ReplaceAllString(text, "$1")

	// 2. Remove space before punctuation
	reBefore := regexp.MustCompile(`\s+([.,!?;:])`)
	text = reBefore.ReplaceAllString(text, "$1")

	// 3. Add space after punctuation if missing (except for dots in ellipsis)
	reAfter := regexp.MustCompile(`([.,!?;:])([^\s.])`)
	text = reAfter.ReplaceAllString(text, "$1 $2")

	return strings.TrimSpace(text)
}

func fixQuotes(text string) string {
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	return re.ReplaceAllString(text, "'$1'")
}
