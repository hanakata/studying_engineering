package main

import (
	"bufio"
	"flag"
	"strings"

	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func duplicateCheck(wordList []string) bool {

}

func main() {
	word := flag.String("w", "aaa", "word")
	flag.Parse()

	checkWord := *word
	wordList := strings.Split(checkWord, "")
	duplicateCheck(wordList)
	return
}
