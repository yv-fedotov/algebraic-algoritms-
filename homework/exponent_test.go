package homework

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"
)

func Test_Exp(t *testing.T) {
	type set struct {
		name   string
		inputA float64
		inputN int
		answer float64
	}

	dirName := "../test_Power"
	dir, err := os.ReadDir(dirName)
	if err != nil {
		t.Fatal(err)
	}

	data := make([]set, 10)
	for _, file := range dir {
		if strings.HasSuffix(file.Name(), ".txt") {
			continue
		}
		name := file.Name()

		nList := strings.Split(name, ".")
		idx, _ := strconv.Atoi(nList[1])

		switch {
		case strings.HasSuffix(file.Name(), ".in"):

			input, _ := os.ReadFile(path.Join(dirName, name))

			listInput := strings.Split(string(input), "\r\n")
			aNumber, _ := strconv.ParseFloat(listInput[0], 64)
			exp, _ := strconv.ParseInt(listInput[1], 10, 64)

			data[idx] = set{
				name:   name,
				inputA: aNumber,
				inputN: int(exp),
				answer: float64(exp),
			}
		case strings.HasSuffix(file.Name(), ".out"):

			answer, _ := os.ReadFile(path.Join(dirName, name))
			anw, _ := strconv.ParseFloat(string(answer), 64)
			data[idx].answer = anw
		}
	}

	t.Run("iter", func(t *testing.T) {
		e := &Exponent{}

		for _, d := range data {
			log.Printf("file: %s \ta %f\tn %d ans %f\n ", d.name, d.inputA, d.inputN, d.answer)
			answ := e.Exp(d.inputA, d.inputN)
			assert.Equal(t, d.answer, answ)
		}
	})
	t.Run("mul", func(t *testing.T) {
		e := &Exponent{}
		for _, d := range data {
			log.Printf("file: %s \ta %f\tn %d ans %f\n ", d.name, d.inputA, d.inputN, d.answer)
			answ := e.Mul(d.inputA, d.inputN)
			assert.Equal(t, d.answer, answ)
		}
	})
}
