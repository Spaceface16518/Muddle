package util

import (
	"bufio"
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
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	rollingSum := 0
	for _, line := range result {
		bytesWritten, err := writer.WriteString(line)
		rollingSum += bytesWritten
		if err != nil {
			return rollingSum, err
		}

		bytesWrittenNewline, err := writer.WriteRune('\n')
		rollingSum += bytesWrittenNewline
		if err != nil {
			return rollingSum, err
		}
	}
	return rollingSum, nil
}
