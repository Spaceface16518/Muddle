package util

import (
	"bufio"
	"io"
)

func writeAll(buf *[]string, writer io.Writer) (int, error) {
	bufferedWriter := bufio.NewWriter(writer)
	defer bufferedWriter.Flush()

	rollingSum := 0
	l := len(*buf)

	for i := 0; i < l; i++ {
		bytesWritten, err := bufferedWriter.WriteString((*buf)[i])
		rollingSum += bytesWritten
		if err != nil {
			return rollingSum, err
		}
	}
	return rollingSum, nil
}

func intersperse(orig *[]string, delim string) []string {
	newLen := len(*orig) * 2
	newSlice := make([]string, newLen)

	for i := 0; i < newLen; i++ {
		if i&1 == 0 {
			newSlice[i] = (*orig)[i/2]
		} else {
			newSlice[i] = delim
		}
	}

	return newSlice
}
