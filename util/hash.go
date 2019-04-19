package util

import (
	"fmt"
	"hash/fnv"
)

// HashTranslations takes a pointer to a list of translation names and outputs the hash of the result as a hexadecimal string
func HashTranslations(used *[]string) (string, error) {
	hasher := fnv.New32()
	_, err := writeAll(used, hasher)
	if err != nil {
		return "", err
	}
	sum := hasher.Sum32()
	return hashToString(sum), nil
}

func hashToString(hash uint32) string {
	return fmt.Sprintf("%x", hash)
}
