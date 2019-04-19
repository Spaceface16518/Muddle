package util

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestWriteAll(t *testing.T) {
	strings := []string{"these", "are", "the", "input", "strings"}
	expectedResult := "thesearetheinputstrings"
	buf := bytes.NewBuffer(nil)
	writeAll(&strings, buf)
	actualResult := buf.String()
	if actualResult != expectedResult {
		t.Errorf("Expected %s, got %s", expectedResult, actualResult)
	}
}

func TestIntersperse(t *testing.T) {
	strings := []string{"these", "are", "the", "input", "strings"}
	delim := "delim"
	expectedResult := []string{"these", "delim", "are", "delim", "the", "delim", "input", "delim", "strings", "delim"}
	actualResult := intersperse(&strings, delim)
	if concat(&actualResult) != concat(&expectedResult) {
		t.Errorf("Expected %v, got %v", expectedResult, actualResult)
	}
}

func BenchmarkWriteAll0(b *testing.B) {
	var strings []string
	nullDevice, err := os.Open(os.DevNull)
	if err != nil {
		b.Error("There was an error opening the null device")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		writeAll(&strings, nullDevice)
	}
}

func BenchmarkWriteAll1(b *testing.B) {
	strings := []string{"a string"}
	nullDevice, err := os.Open(os.DevNull)
	if err != nil {
		b.Error("There was an error opening the null device")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		writeAll(&strings, nullDevice)
	}
}

func BenchmarkWriteAll5(b *testing.B) {
	strings := []string{"these", "are", "the", "input", "strings"}
	nullDevice, err := os.Open(os.DevNull)
	if err != nil {
		b.Error("There was an error opening the null device")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		writeAll(&strings, nullDevice)
	}
}
func BenchmarkWriteAll10(b *testing.B) {
	strings := []string{"these", "are", "the", "input", "strings", "and", "these", "are", "some", "more"}
	nullDevice, err := os.Open(os.DevNull)
	if err != nil {
		b.Error("There was an error opening the null device")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		writeAll(&strings, nullDevice)
	}
}

func BenchmarkWriteAll20(b *testing.B) {
	strings := []string{"these", "are", "the", "input", "strings", "and", "these", "are", "some", "more", "and", "here", "are", "even", "more", "strings", "that", "mean", "absolutely", "nothing"}
	nullDevice, err := os.Open(os.DevNull)
	if err != nil {
		b.Error("There was an error opening the null device")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		writeAll(&strings, nullDevice)
	}
}

func BenchmarkWriteAll1_SmallInput(b *testing.B) {
	strings := []string{" "}
	nullDevice, err := os.Open(os.DevNull)
	if err != nil {
		b.Error("There was an error opening the null device")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		writeAll(&strings, nullDevice)
	}
}

func BenchmarkWriteAll1_LargeInput(b *testing.B) {
	strings := []string{"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."}
	nullDevice, err := os.Open(os.DevNull)
	if err != nil {
		b.Error("There was an error opening the null device")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		writeAll(&strings, nullDevice)
	}
}

func BenchmarkIntersperse0(b *testing.B) {
	var strings []string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersperse(&strings, "delim")
	}
}

func BenchmarkIntersperse1(b *testing.B) {
	strings := []string{"a string"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersperse(&strings, "delim")
	}
}

func BenchmarkIntersperse5(b *testing.B) {
	strings := []string{"these", "are", "the", "input", "strings"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersperse(&strings, "delim")
	}
}

func BenchmarkIntersperse10(b *testing.B) {
	strings := []string{"these", "are", "the", "input", "strings", "and", "these", "are", "some", "more"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersperse(&strings, "delim")
	}
}

func BenchmarkIntersperse20(b *testing.B) {
	strings := []string{"these", "are", "the", "input", "strings", "and", "these", "are", "some", "more", "and", "here", "are", "even", "more", "strings", "that", "mean", "absolutely", "nothing"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersperse(&strings, "delim")
	}
}

func BenchmarkIntersperse5_SmallDelim(b *testing.B) {
	strings := []string{"these", "are", "the", "input", "strings"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersperse(&strings, " ")
	}
}

func BenchmarkIntersperse5_LargeDelim(b *testing.B) {
	strings := []string{"these", "are", "the", "input", "strings"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersperse(&strings, "this is a very very very very large delimeter except not really its not super big in the grand scheme of things in fact we are all not very big in the grand scheme of things we are all just insignificant gnats making but a ripple in the grand ocean that is space-time")
	}
}

func concat(stringSlice *[]string) string {
	builder := strings.Builder{}
	for _, s := range *stringSlice {
		builder.WriteString(s)
	}
	return builder.String()
}
