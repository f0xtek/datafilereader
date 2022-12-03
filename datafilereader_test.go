package datafilereader

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func createIntsDataFile(name string, numLines int) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}

	defer file.Close()

	for i := 0; i < numLines; i++ {
		s := fmt.Sprintf("%s\n", strconv.Itoa(i))

		if i == numLines-1 {
			s = strings.TrimRight(s, "\n")
		}

		_, err = file.WriteString(s)
		if err != nil {
			return err
		}
	}

	return nil
}

func deleteTestDataFile(filename string) error {
	if err := os.Remove(filename); err != nil {
		return err
	}

	return nil
}

type stringTest struct {
	want string
	got  string
}

type intTest struct {
	want int
	got  int
}

func TestReadLinesFromFile(t *testing.T) {
	t.Parallel()

	const fileName string = "testData.txt"
	numLines := 10000

	err := createIntsDataFile(fileName, numLines)
	if err != nil {
		t.Errorf("error creating test data: %v", err)
	}

	tt := make([]stringTest, 0)

	for i := 0; i < numLines; i++ {
		s := strconv.Itoa(i)
		tt = append(tt, stringTest{want: s, got: ""})
	}

	lines, err := ReadLinesFromFile(fileName)
	if err != nil {
		t.Errorf("error reading lines from file: %v", err)
	}

	for i, line := range lines {
		tt[i].got = line
		if tt[i].want != tt[i].got {
			t.Errorf("got %v, want %v", tt[i].got, tt[i].want)
		}
	}

	err = deleteTestDataFile(fileName)
	if err != nil {
		fmt.Printf("error deleting test data file: %v", err)
	}

}

func BenchmarkReadLinesFromFile(b *testing.B) {
	const fileName string = "testData.txt"
	numLines := 10000

	err := createIntsDataFile(fileName, numLines)
	if err != nil {
		b.Errorf("error creating test data: %v", err)
	}

	for i := 0; i < b.N; i++ {
		ReadLinesFromFile(fileName)
	}

	err = deleteTestDataFile(fileName)
	if err != nil {
		fmt.Printf("error deleting test data file: %v", err)
	}
}

func TestReadInts(t *testing.T) {
	const fileName string = "testData.txt"
	numLines := 10000

	err := createIntsDataFile(fileName, numLines)
	if err != nil {
		t.Errorf("error creating test data: %v", err)
	}

	tt := make([]intTest, numLines)

	for i := 0; i < numLines; i++ {
		tt[i].want = i
	}

	ints, err := ReadInts(fileName)
	if err != nil {
		t.Errorf("error reading ints: %v", err)
	}

	for i, num := range ints {
		tt[i].got = num
		if tt[i].want != tt[i].got {
			t.Errorf("got %v, want %v", tt[i].got, tt[i].want)
		}
	}

	err = deleteTestDataFile(fileName)
	if err != nil {
		fmt.Printf("error deleting test data file: %v", err)
	}
}

func BenchmarkReadInts(b *testing.B) {
	const fileName string = "testData.txt"
	numLines := 10000

	err := createIntsDataFile(fileName, numLines)
	if err != nil {
		b.Errorf("error creating test data: %v", err)
	}

	for i := 0; i < b.N; i++ {
		ReadInts(fileName)
	}

	err = deleteTestDataFile(fileName)
	if err != nil {
		fmt.Printf("error deleting test data file: %v", err)
	}
}
