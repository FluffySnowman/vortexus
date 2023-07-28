package gen

import (
	"bufio"
	"os"
)

// read input from user

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

