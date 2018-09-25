package main

import (
	"errors"
	"fmt"
	"sort"
)

func main() {

	fixArrayTest()

}

func fixArr(a []int) (error, []int) {
	if len(a) == 0 {
		return errors.New("supplied array cannot be empty"), a
	}

	//grab the largest number in the array by sorting it first
	if s := sort.IntsAreSorted(a); !s {
		sort.Ints(a)
	}

	max := a[len(a)-1:][0]
	min := a[:1][0]
	capacity := (max + 1) - min
	fmt.Println("max: ", max, "min: ", min, "capacity : ", capacity)

	//make a new array from min to max of cap = capacity
	comparator := make([]int, capacity)

	fmt.Println("comparator array : ", comparator)

	// range over the comparator to initialize all elements in icrementing order
	for index := range comparator {
		comparator[index] = min
		min++
	}
	fmt.Println("Comparing : ", comparator, "Against : ", a)

	//start comparing new comparator with “a” and store desimilar values in missingValues array

	missingValues := make([]int, len(comparator)-len(a))
	var i int = 0

	//let’s make a map to reduce the complexity and then delete the “contains() function
	m := make(map[int]int)
	for index, v := range a {
		m[v] = index
	}
	for _, value := range comparator {
		if _, containsValue := m[value]; !containsValue {
			missingValues[i] = value
			i++
		}
	}

	fmt.Println("Missing Values : ", missingValues)

	return nil, missingValues
}

func fixArrayTest() {
	a := []int{9, 3, 5, 7}
	err, _ := fixArr(a)
	if err != nil {
		fmt.Println("function is bad so far")
	}
}
