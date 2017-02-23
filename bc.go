package bc

import "strings"

var codes = map[string]string{
	"a": "AAAAA", "b": "AAAAB", "c": "AAABA", "d": "AAABB", "e": "AABAA", "f": "AABAB", "g": "AABBA", "h": "AABBB", "i": "ABAAA",
	"j": "ABAAB", "k": "ABABA", "l": "ABABB", "m": "ABBAA", "n": "ABBAB", "o": "ABBBA", "p": "ABBBB", "q": "BAAAA", "r": "BAAAB",
	"s": "BAABA", "t": "BAABB", "u": "BABAA", "v": "BABAB", "w": "BABBA", "x": "BABBB", "y": "BBAAA", "z": "BBAAB", " ": "BBBAA",

	"AAAAA": "a", "AAAAB": "b", "AAABA": "c", "AAABB": "d", "AABAA": "e", "AABAB": "f", "AABBA": "g", "AABBB": "h", "ABAAA": "i",
	"ABAAB": "j", "ABABA": "k", "ABABB": "l", "ABBAA": "m", "ABBAB": "n", "ABBBA": "o", "ABBBB": "p", "BAAAA": "q", "BAAAB": "r",
	"BAABA": "s", "BAABB": "t", "BABAA": "u", "BABAB": "v", "BABBA": "w", "BABBB": "x", "BBAAA": "y", "BBAAB": "z", "BBBAA": " ",
}

func Encrypt(message, text string) string {
	var et []byte
	var em []byte

	message = strings.ToLower(message)
	text = strings.ToLower(text)

	for _, char := range message {
		if char >= 97 || char <= 122 {
			et = append(et, []byte(codes[string(char)])...)
		} else {
			et = append(et, []byte(codes[" "])...)
		}
	}

	var count int
	for _, char := range text {
		if char >= 97 && char <= 122 {
			if et[count] == 65 {
				em = append(em, byte(char))
			} else {
				em = append(em, byte(char-32))
			}
			count++
			if count == len(et) {
				break
			}
		} else {
			em = append(em, byte(char))
		}
	}

	return string(em)
}

func Decrypt(message string) string {
	var dt []byte
	var dm []byte

	for _, char := range message {
		if char >= 97 && char <= 122 {
			dt = append(dt, []byte("A")...)
		} else if char >= 65 && char <= 90 {
			dt = append(dt, []byte("B")...)
		}
	}

	var quintet []byte
	for c := 0; c <= len(dt); c += 5 {
		quintet = dt[c : c+5]
		dm = append(dm, codes[string(quintet)]...)
	}

	return string(dm)
}
