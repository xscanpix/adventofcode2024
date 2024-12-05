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

func TestSplitFileIntoRows(t *testing.T) {
	var expect = [][]int{
		{
			7, 6, 4, 2, 1,
		},
		{
			1, 2, 7, 8, 9,
		},
		{
			9, 7, 6, 2, 1,
		},
		{
			1, 3, 2, 4, 5,
		},
		{
			8, 6, 4, 4, 1,
		},
		{
			1, 3, 6, 7, 9,
		},
	}

	fs := fstest.MapFS{
		"input.txt": {
			Data: []byte("7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9"),
		},
	}

	file, err := fs.Open("input.txt")

	if err != nil {
		t.Fatal(err)
	}

	result := *SplitFileIntoRowsFromReader(file, " ")

	if !reflect.DeepEqual(result, expect) {
		t.Fatal("Error")
	}
}
