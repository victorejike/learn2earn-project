package main

import (
	"regexp"
	"strconv"
	"strings"
)

func Process(text string) string {
	text = strings.ReplaceAll(text, "\n", " ")
	tokens := tokenize(text)

	tokens = applyModifiers(tokens)
	tokens = fixArticles(tokens)

	result := strings.Join(tokens, " ")
	result = fixPunctuation(result)
	result = fixQuotes(result)

	return strings.TrimSpace(result)
}

func tokenize(text string) []string {
	return strings.Fields(text)
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
		}

		if token == "(bin)" && i > 0 {
			val, err := strconv.ParseInt(tokens[i-1], 2, 64)
			if err == nil {
				tokens[i-1] = strconv.FormatInt(val, 10)
			}
			tokens = removeToken(tokens, i)
			i--
		}

		if strings.HasPrefix(token, "(up") ||
			strings.HasPrefix(token, "(low") ||
			strings.HasPrefix(token, "(cap") {

			cmd, count := parseModifier(token)

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
	re := regexp.MustCompile(`\((up|low|cap)(,\s*(\d+))?\)`)
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

func fixPunctuation(text string) string {
	reBefore := regexp.MustCompile(`\s+([.,!?;:])`)
	text = reBefore.ReplaceAllString(text, "$1")

	reAfter := regexp.MustCompile(`([.,!?;:])([^\s])`)
	text = reAfter.ReplaceAllString(text, "$1 $2")

	reDots := regexp.MustCompile(`\s*(\.\.\.|!\?|!!|\?\!)\s*`)
	text = reDots.ReplaceAllString(text, "$1 ")

	return strings.TrimSpace(text)
}

func fixQuotes(text string) string {
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	return re.ReplaceAllString(text, "'$1'")
}
