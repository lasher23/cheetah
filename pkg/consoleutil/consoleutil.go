package consoleutil

import (
	"bufio"
	"fmt"
	"github.com/lasher23/cheetah/pkg/stringutil"
	"os"
)

func DisplayQuestion(answerOptions []string, displayText string) string {
	var input string
	var e error
	reader := bufio.NewReader(os.Stdin)
	for !stringutil.ContainsIgnoreCase(answerOptions, stringutil.TrimReturn(input)) {
		fmt.Print(displayText)
		input, e = reader.ReadString('\n')
		input = stringutil.TrimReturn(input)
		if e != nil {
			fmt.Println("unexpected error")
			os.Exit(1)
		}
	}
	return input
}
