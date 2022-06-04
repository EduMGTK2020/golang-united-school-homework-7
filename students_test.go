package coverage

import (
	"os"
	"reflect"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW
func TestPeoples(t *testing.T){
	person0 := Person{"Artur", 	"Dow", time.Now()}
	person1 := Person{"Artur", 	"Down", time.Now()}
	person2 := Person{"John", 	"Down", time.Now().Add(time.Hour)}
	person3 := Person{"Ara", 	"Down", time.Now()}

	peoples := People{}
	len0 := peoples.Len()
	if len0 != 0 {
		t.Errorf("Wrong get len: expected %d, got %d", 0, len0)		
	}

	peoples = append(peoples, person0)
	len1 := peoples.Len()
	if len1 != 1 {
		t.Errorf("Wrong get len: expected %d, got %d", 0, len0)		
	}

	peoples = append(peoples, person1)
	peoples.Swap(0,1)
	if !reflect.DeepEqual(person0, peoples[1]) {
		t.Errorf("Wrong swap: expected %v, got %v", person0, peoples[1])
	}
	if !reflect.DeepEqual(person1, peoples[0]) {
		t.Errorf("Wrong swap: expected %v, got %v", person1, peoples[0])
	} 

	peoples = append(peoples, person2)
	peoples = append(peoples, person3)
	peoples.Swap(0,1)

	testDataLess:= []struct {
		i int
		j int
		expected bool
	}{
		{0,1, true},
		{0,3, false},
		{1,2, false},
		{1,0, false},
	}

	for k, v:= range testDataLess {
		got:= peoples.Less(v.i, v.j)
		if got != v.expected {
			t.Errorf("Step %d - wrong result: expected %v, got %v", k, v.expected, got)
		}
	}
}

func TestMatrix(t *testing.T) {

	var matrix *Matrix
	var err error

	matrix, err = New("")
	if matrix != nil {
		t.Errorf("Wrong error result: expected %v, got %v", nil, err)		
	}
	_, err = New("0\n0 1")
	if err == nil {
		t.Errorf("Wrong error result: expected %v, got %v", nil, err)		
	}
	
	matrix, err = New("0 1 2\n0 1 2")
	if err != nil {
		t.Errorf("Wrong error result: expected %v, got %v", nil, err)		
	}
	
	testSet := []struct{
		r int
		c int
		d int
		set bool
	}{
		{0,0,1,true},
		{3,3,1,false},
		{-3,-3,1,false},
	}

	for k, v := range testSet {
		got := matrix.Set(v.r, v.c, v.d)
		if got != v.set {
			t.Errorf("Step %d - wrong result: expected %v, got %v", k, v.set, got)			
		}  
	}
	
	expM := [][]int {
		{1,1,2},
		{0,1,2},
	}

	if !reflect.DeepEqual(matrix.Rows(), expM) {
		t.Errorf("Wrong result: expected %v, got %v", expM, matrix.Rows()  )		
	}

	expM = [][]int {
		{1,0},
		{1,1},
		{2,2},
	}

	if !reflect.DeepEqual(matrix.Cols(), expM) {
		t.Errorf("Wrong result: expected %v, got %v", expM, matrix.Cols())		
	}
}