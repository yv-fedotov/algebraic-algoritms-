package homework

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"
)

func TestFib_Fibonacci(t *testing.T) {
	type set struct {
		name   string
		input  int
		answer uint64
	}

	dirName := "../test_Fibo"
	dir, err := os.ReadDir(dirName)
	if err != nil {
		t.Fatal(err)
	}

	data := make([]set, len(dir))
	for _, file := range dir {
		if strings.HasSuffix(file.Name(), ".txt") {
			continue
		}
		name := file.Name()

		nList := strings.Split(name, ".")
		idx, _ := strconv.Atoi(nList[1])

		switch {
		case strings.HasSuffix(file.Name(), ".in"):

			line, _ := os.ReadFile(path.Join(dirName, name))
			listInput := strings.Split(string(line), "\n")

			input, _ := strconv.ParseInt(listInput[0], 10, 64)

			data[idx] = set{
				name:  name,
				input: int(input),
			}
		case strings.HasSuffix(file.Name(), ".out"):
			line, _ := os.ReadFile(path.Join(dirName, name))
			listInput := strings.Split(string(line), "\n")
			anw, _ := strconv.ParseUint(listInput[0], 10, 64)
			data[idx].answer = anw
		}
	}

	t.Run("Gold", func(t *testing.T) {
		f := &Fib{}

		for _, d := range data {
			res := f.FibonacciGold(d.input)
			assert.Equal(t, d.answer, res, fmt.Sprintf("%s %d %d\t[%d]\n", d.name, d.input, d.answer, res))
		}

	})

	t.Run("Iter", func(t *testing.T) {
		f := &Fib{}

		for _, d := range data {
			res := f.FibonacciIter(d.input)
			assert.Equal(t, d.answer, res, fmt.Sprintf("%s %d %d\t[%d]\n", d.name, d.input, d.answer, res))
		}
	})

}
