package sqrt

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strconv"
)

var (
	cmd    *exec.Cmd
	input  chan string
	output chan float64
	errors chan error
)

func init() {
	errors = make(chan error)
	input = make(chan string)
	output = make(chan float64)
}

func setup() {
	cmd = exec.Command("python", "sqrt.py")

	stdin, err := cmd.StdinPipe()
	defer stdin.Close()

	if err != nil {
		errors <- err
		return
	}

	stdout, err := cmd.StdoutPipe()
	defer stdout.Close()

	if err != nil {
		errors <- err
		return
	}

	if err := cmd.Start(); err != nil {
		errors <- err
		return
	}

	go read(stdout)
	go write(stdin)
	cmd.Wait()
}

func write(w io.Writer) {
	for {
		w.Write([]byte(<-input))
	}
}

func read(r io.Reader) {
	for {
		reader := bufio.NewReader(r)
		line, _, err := reader.ReadLine()

		if err != nil {
			errors <- err
			continue
		}

		result, err := strconv.ParseFloat(string(line), 64)

		if err != nil {
			errors <- err
			continue
		}

		output <- result
	}
}

func pythonSqrt(a float64) (float64, error) {
	if cmd == nil {
		go setup()
	}

	input <- fmt.Sprintf("%f\n", a)

	select {
	case err := <-errors:
		return 0, err
	case result := <-output:
		return result, nil
	}
}
