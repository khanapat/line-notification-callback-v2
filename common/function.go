package common

import (
	"encoding/base64"
	"strings"
)

func BToB64(s []byte) string {
	return base64.StdEncoding.EncodeToString(s)
}

func B64ToB(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func FindRuneIndex(s string, r rune) []int {
	var indexes []int
	sub := s
	count := strings.Count(s, string(r))
	for i := 0; i < count; i++ {
		index := strings.IndexRune(sub, r)
		sub = strings.Replace(sub, string(r), "-", 1)
		indexes = append(indexes, index)
	}
	return indexes
}
