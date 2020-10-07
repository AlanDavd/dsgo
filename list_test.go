/* Linked list tests. */
package dsgo

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

var (
	listUnxOpt = cmp.AllowUnexported(List{})
	nodeUnxOpt = cmp.AllowUnexported(Node{})
)

func compareLists(t *testing.T, expected, got *List, producer string) {
	diff := cmp.Diff(expected, got, listUnxOpt, nodeUnxOpt)
	if diff != "" {
		t.Fatalf("%s produced unwanted list: %v\nwant %v\ndiff want -> got\n%s",
			producer, got, expected, diff)
	}
}

func TestSinglyList_Constructor(t *testing.T) {
	expected := List{Head: nil, Length: 0}
	got := New()
	compareLists(t, &expected, &got, "CreateList")
}

func TestSinglyList_Get(t *testing.T) {
	tests := []struct {
		name		string
		init		List
		index, want	int
	} {
		{
			name: "Get -1 by negative index in empty list",
			init: List{Head: nil, Length: 0},
			want: -1,
			index: 1,
		},
		{
			name: "Get -1 by index zero in empty list",
			init: List{Head: nil, Length: 0},
			want: -1,
			index: 0,
		},
		{
			name: "Get value by index zero in list with one element",
			init: List{Head: &Node{Value: 4}, Length: 1},
			want: 4,
			index: 0,
		},
		{
			name: "Get out of range response by index in list with elements",
			init: List{Head: &Node{Value: 4, Next: &Node{Value: 1, Next: &Node{Value: 60}}}, Length: 3},
			want: -1,
			index: 4,
		},
		{
			name: "Get value by index in list with elements",
			init: List{Head: &Node{Value: 4, Next: &Node{Value: 1, Next: &Node{Value: 60, Next: &Node{Value: 5}}}}, Length: 4},
			want: 60,
			index: 2,
		},
		{
			name: "Get last value by index in list with elements",
			init: List{Head: &Node{Value: 4, Next: &Node{Value: 1, Next: &Node{Value: 60, Next: &Node{Value: 5}}}}, Length: 4},
			want: 5,
			index: 3,
		},
		{
			name: "Get value by index in list with elements when index is equal to list's length",
			init: List{Head: &Node{Value: 4, Next: &Node{Value: 1, Next: &Node{Value: 60, Next: &Node{Value: 5}}}}, Length: 4},
			want: -1,
			index: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.init.Get(test.index)
		})
	}
}

func TestSinglyList_AddAtHead(t *testing.T) {
	tests := []struct {
		name 		  string
		init, want 	  List
		insertedData  []int
	} {
		{
			name: "Add at head in empty list",
			init: List{Head: nil, Length: 0},
			want: List{Head: &Node{Value: 1}, Length: 1},
			insertedData: []int{1},
		},
		{
			name: "Add at head with one element",
			init: List{Head: nil, Length: 0},
			want: List{Head: &Node{Value: 2, Next: &Node{Value: 1}}, Length: 2},
			insertedData: []int{1, 2},
		},
		{
			name: "Add at head in list with elements",
			init: List{Head: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3}}}, Length: 3},
			want: List{Head: &Node{Value: -5, Next: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3}}}}, Length: 4},
			insertedData: []int{-5},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, i := range test.insertedData {
				test.init.AddAtHead(i)
			}
		})
		compareLists(t, &test.want, &test.init, "AddAtHead")
	}
}

func TestSinglyList_AddAtTail(t *testing.T) {
	tests := []struct {
		name 		  string
		init, want 	  List
		insertedData  []int
	} {
		{
			name: "Add at tail in empty list",
			init: List{Head: nil, Length: 0},
			want: List{Head: &Node{Value: 1}, Length: 1},
			insertedData: []int{1},
		},
		{
			name: "Add at tail with one element",
			init: List{Head: nil, Length: 0},
			want: List{Head: &Node{Value: 1, Next: &Node{Value: 2}}, Length: 2},
			insertedData: []int{1, 2},
		},
		{
			name: "Add at tail in list with elements",
			init: List{Head: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3}}}, Length: 3},
			want: List{Head: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3, Next: &Node{Value: -5}}}}, Length: 4},
			insertedData: []int{-5},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, i := range test.insertedData {
				test.init.AddAtTail(i)
			}
		})
		compareLists(t, &test.want, &test.init, "AddAtTail")
	}
}

func TestSinglyList_AddAtIndex(t *testing.T) {
	tests := []struct {
		name 		  		string
		init, want 	  		List
		index, insertedData int
	} {
		{
			name: "Add at negative index in empty list",
			init: List{Head: nil, Length: 0},
			want: List{Head: nil, Length: 0},
			insertedData: 1,
			index: -1,
		},
		{
			name: "Add at out of range index in list with elements",
			init: List{Head: &Node{Value: 1, Next: &Node{Value: 2}}, Length: 2},
			want: List{Head: &Node{Value: 1, Next: &Node{Value: 2}}, Length: 2},
			insertedData: 60,
			index: 8,
		},
		{
			name: "Add at head by index in list with elements",
			init: List{Head: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3}}}, Length: 3},
			want: List{Head: &Node{Value: -5, Next: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3}}}}, Length: 4},
			insertedData: -5,
			index: 0,
		},
		{
			name: "Add at tail by index in list with elements",
			init: List{Head: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3}}}, Length: 3},
			want: List{Head: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3, Next: &Node{Value: -5}}}}, Length: 4},
			insertedData: -5,
			index: 3,
		},
		{
			name: "Add somewhere in list with elements by index",
			init: List{Head: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3}}}, Length: 3},
			want: List{Head: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: -5, Next: &Node{Value: 3}}}}, Length: 4},
			insertedData: -5,
			index: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.init.AddAtIndex(test.index, test.insertedData)
		})
		compareLists(t, &test.want, &test.init, "AddAtIndex")
	}
}

func TestSinglyList_DeleteAtIndex(t *testing.T) {
	tests := []struct {
		name		string
		init, want	List
		index		int
	} {
		{
			name: "Delete at negative index in empty list",
			init: List{Head: nil, Length: 0},
			want: List{Head: nil, Length: 0},
			index: -1,
		},
		{
			name: "Delete at out of range index in list with elements",
			init: List{Head: &Node{Value: 1, Next: &Node{Value: 2}}, Length: 2},
			want: List{Head: &Node{Value: 1, Next: &Node{Value: 2}}, Length: 2},
			index: 8,
		},
		{
			name: "Delete at head by index in list with elements",
			init: List{Head: &Node{Value: -5, Next: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3}}}}, Length: 4},
			want: List{Head: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3}}}, Length: 3},
			index: 0,
		},
		{
			name: "Delete at tail by index in list with elements",
			init: List{Head: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3, Next: &Node{Value: -5}}}}, Length: 4},
			want: List{Head: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3}}}, Length: 3},
			index: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.init.DeleteAtIndex(test.index)
		})
		compareLists(t, &test.want, &test.init, "DeleteAtIndex")
	}
}
