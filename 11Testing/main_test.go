package main

import "testing"

// func TestAdd(t *testing.T){
// 	result, _ := Add(5,6)  //11
// 	if result!=11{
// 		t.Error("incorrect addition: Expected 11")
// 	}
// }

func TestAdd(t *testing.T){
	tables := []struct {
		name string
		n1 float64
		n2 float64
		expected float64
	}{
		{"test1", 4,6,10},
	}

	for _,tt := range tables{
		result, _ := Add(tt.n1, tt.n2)
		if result != tt.expected{
			t.Error("Error expected: ", tt.expected)
		}
	}
}