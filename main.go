package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cmd := exec.Command("gofmt")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	pipe, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	// strip line within imports
	state := begin
	for scanner.Scan() {
		line := scanner.Text()
		state = state(pipe, line)
	}

	// check for errors and panic
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	if err := pipe.Close(); err != nil {
		panic(err)
	}

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

type State func(io.Writer, string) State

func begin(wr io.Writer, line string) State {
	return findImports(wr, line)
}

func findImports(wr io.Writer, line string) State {
	pwrite(wr, line)
	if line != "import (" {
		return findImports
	} else {
		return cleanImports
	}
}

func cleanImports(wr io.Writer, line string) State {
	switch line {
	case "":
		return cleanImports
	case ")":
		pwrite(wr, line)
		return consumeAll
	default:
		pwrite(wr, line)
		return cleanImports
	}
}

func consumeAll(wr io.Writer, line string) State {
	pwrite(wr, line)
	return consumeAll
}

func pwrite(wr io.Writer, line string) {
	if _, err := fmt.Fprintln(wr, line); err != nil {
		panic(err)
	}
}
