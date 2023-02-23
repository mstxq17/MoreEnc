package main

import (
	"MoreEnc/internal/core"
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	encode = flag.Bool("e", false, "Encode input")
	decode = flag.Bool("d", false, "Decode input")
	help   = flag.Bool("help", false, "Show usage")
)

func main() {
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	input := readInput()

	var output string
	if *encode {
		output = input
	} else if *decode {
		output = input
	} else {
		_, _ = fmt.Fprintln(os.Stderr, "Error: encode or decode flag must be set")
		os.Exit(1)
	}
	fmt.Println(output)
}

// Read input from stdin or specified string
func readInput() string {
	var input string
	if stat, _ := os.Stdin.Stat(); (stat.Mode() & os.ModeCharDevice) == 0 {
		// Input is piped in
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if input != "" {
				input += "\n"
			}
			if *encode {
				line = core.UrlEncode(line)
			} else if *decode {
				line = core.UrlDecode(line)
			}
			input += line
		}
		if err := scanner.Err(); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
	} else if flag.NArg() > 0 {
		// Input is specified as a command line argument
		input = flag.Arg(0)
	} else {
		// Prompt for input
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter input: ")
		input, _ = reader.ReadString('\n')
	}
	return input
}
