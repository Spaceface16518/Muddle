package util

import (
	"io"
	"log"
	"os"
)

// Dump prints the result to standard output
func Dump(result []string) {
	bytesWritten, err := DumpTo(result, os.Stdout)
	if err != nil {
		log.Fatalf("Dumping failed. %v bytes written\n", bytesWritten)
	}
}

// DumpTo prints the result to a specified output source that implements io.Writer
func DumpTo(result []string, file io.Writer) (int, error) {
	interspersedResult := intersperse(&result, "\n")
	return writeAll(&interspersedResult, file)
}
