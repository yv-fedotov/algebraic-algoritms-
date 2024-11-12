package homework

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"
)

func Test_Prime(t *testing.T) {
	type set struct {
		name   string
		input  int
		answer int
	}

	dirName := "../test_Primes"
	dir, err := os.ReadDir(dirName)
	if err != nil {
		t.Fatal(err)
	}

	data := make([]set, 15)
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
			line = bytes.TrimSpace(line)

			input, _ := strconv.ParseInt(string(line), 10, 64)

			data[idx] = set{
				name:  name,
				input: int(input),
			}
		case strings.HasSuffix(file.Name(), ".out"):
			line, _ := os.ReadFile(path.Join(dirName, name))
			line = bytes.TrimSpace(line)
			anw, _ := strconv.ParseInt(string(line), 10, 32)
			data[idx].answer = int(anw)
		}
	}
	t.Run("iter", func(t *testing.T) {
		p := Prime{}
		for _, d := range data {
			res := p.PrimeIter(d.input)
			assert.Equal(t, d.answer, res, fmt.Sprintf("%s\t%d\t%d\t[%d]\n", d.name, d.input, d.answer, res))
		}
	})
	t.Run("optimal", func(t *testing.T) {
		p := Prime{}
		for _, d := range data {
			res := p.PrimeOptima(d.input)
			assert.Equal(t, d.answer, res, fmt.Sprintf("%s\t%d\t%d\t[%d]\n", d.name, d.input, d.answer, res))
		}
	})
}
