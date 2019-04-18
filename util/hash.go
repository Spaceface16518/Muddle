package util

import (
	"fmt"
	"hash/fnv"
)

// HashTranslations takes a pointer to a list of translation names and outputs the hash of the result as a hexadecimal string
func HashTranslations(used *[]string) string {
	hasher := fnv.New32()
	for _, lang := range *used {
		hasher.Write([]byte(lang))
	}
	sum := hasher.Sum32()
	return fmt.Sprintf("%x", sum)
}
