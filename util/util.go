package util

import "io"

func PrintParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\t")
		io.WriteString(out, "\n")
	}
}
