package util

import (
	"testing"
)

func TestHashTranslations_1(t *testing.T) {
	input := []string{"en", "fr", "de", "es"}
	expectedHash := "f90eb5f9"
	actualHash, err := HashTranslations(&input)
	if err != nil {
		t.Errorf("Hashing resulted in error: %v", err)
	}
	if actualHash != expectedHash {
		t.Errorf("Expected %s, got %s", expectedHash, actualHash)
	}
}

func TestHashTranslations_2(t *testing.T) {
	input := []string{"en", "nb", "en", "de", "en"}
	expectedHash := "f7a6dc19"
	actualHash, err := HashTranslations(&input)
	if err != nil {
		t.Errorf("Hashing resulted in error: %v", err)
	}
	if actualHash != expectedHash {
		t.Errorf("Expected %s, got %s", expectedHash, actualHash)
	}
}

func TestHashTranslations_3(t *testing.T) {
	input := []string{"en", "de", "fr", "de", "it", "de", "en", "pt", "en", "es", "ca", "es", "en", "tr", "en", "hi", "en", "es", "fr", "en", "nb", "en"}
	expectedHash := "e906ca60"
	actualHash, err := HashTranslations(&input)
	if err != nil {
		t.Errorf("Hashing resulted in error: %v", err)
	}
	if actualHash != expectedHash {
		t.Errorf("Expected %s, got %s", expectedHash, actualHash)
	}
}

func BenchmarkHashTranslations20(b *testing.B) {
	input := []string{"en", "de", "fr", "de", "it", "de", "en", "pt", "en", "es", "ca", "es", "en", "tr", "en", "hi", "en", "es", "fr", "en"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HashTranslations(&input)
	}
}

func BenchmarkHashTranslations10(b *testing.B) {
	input := []string{"en", "de", "fr", "de", "it", "de", "en", "pt", "en", "es"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HashTranslations(&input)
	}
}

func BenchmarkHashTranslations5(b *testing.B) {
	input := []string{"en", "de", "fr", "de", "it"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HashTranslations(&input)
	}
}

func BenchmarkHashTranslations1(b *testing.B) {
	input := []string{"en"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HashTranslations(&input)
	}
}

func BenchmarkHashTranslations0(b *testing.B) {
	var input []string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HashTranslations(&input)
	}
}
