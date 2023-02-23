package main

import (
	"MoreEnc/internal/core"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// Read input from stdin or specified string
func readInput() string {
	var input string
	if stat, _ := os.Stdin.Stat(); (stat.Mode() & os.ModeCharDevice) == 0 {
		// Input is piped in
		bytes, _ := io.ReadAll(os.Stdin)
		input = string(bytes)
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

func main() {
	encode := flag.Bool("e", false, "URL encode input")
	decode := flag.Bool("d", false, "URL encode input")
	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [INPUT]\n\n", os.Args[0])
		_, _ = fmt.Fprintln(os.Stderr, "OPTIONS:")
		flag.PrintDefaults()
		_, _ = fmt.Fprintln(os.Stderr, "\nINPUT:")
		_, _ = fmt.Fprintln(os.Stderr, "  The string to encode or decode. If not specified, the program will prompt for input.")
	}
	// parse the setting
	// 初始化设置
	flag.Parse()
	var output string
	var input string = readInput()
	if *encode {
		output = core.UrlEncode(input)
		fmt.Println(output)
	}
	if *decode {
		output = core.UrlDecode(input)
		fmt.Println(output)
	}
	if !*encode && !*decode {
		_, _ = fmt.Fprintln(os.Stderr, "Must specify either -encode or -decode flag")
		os.Exit(1)
	}
}
