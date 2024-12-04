package utils

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestSplitFileIntoColumns(t *testing.T) {
	var expect = [][]int{
		{
			3, 4, 2, 1, 3, 3,
		},
		{
			4, 3, 5, 3, 9, 3,
		}}

	fs := fstest.MapFS{
		"input.txt": {
			Data: []byte("3   4\n4   3\n2   5\n1   3\n3   9\n3   3"),
		},
	}

	file, err := fs.Open("input.txt")

	if err != nil {
		t.Fatal(err)
	}

	result := *SplitFileIntoColumnsFromReader(file, "   ", 2)

	if !reflect.DeepEqual(result, expect) {
		t.Fatal("Error")
	}
}
