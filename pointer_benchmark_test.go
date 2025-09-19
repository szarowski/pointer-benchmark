package main

import (
	"fmt"
	"math/rand"
	"testing"
)

const (
	maxNestedCount = 16
	maxRecursion   = 3
	maxCallCount   = 16
)

// LargeValueStruct defines a struct with generically 5 types of arrays and a slice of the same nested structs
type LargeValueStruct[D [1]int | [10]int | [100]int | [1000]int | [10000]int] struct {
	data   D
	nested []LargeValueStruct[D]
}

// LargePointerStruct defines a struct with a pointer to generically 5 types of arrays and a pointer to slice of the same nested structs
type LargePointerStruct[D [1]int | [10]int | [100]int | [1000]int | [10000]int] struct {
	data   *D
	nested *[]LargePointerStruct[D]
}

// initByValue creates a struct instance and sets randomized values to the fields
func initByValue[D [1]int | [10]int | [100]int | [1000]int | [10000]int](nestedCount, recursion int) LargeValueStruct[D] {
	s := LargeValueStruct[D]{}
	recursion--
	for j := 0; j < nestedCount; j++ {
		if recursion > 0 {
			// initialize nested structs
			s.nested = append(s.nested, initByValue[D](nestedCount, recursion))
		} else {
			// initialize top-most nested struct
			s.nested = append(s.nested, LargeValueStruct[D]{})
		}
		// initialize the nested struct input data with random values
		for i := 0; i < len(s.data); i++ {
			s.nested[j].data[i] = rand.Intn(len(s.nested[j].data))
		}
	}
	// initialize the input data with random values
	for i := 0; i < len(s.data); i++ {
		s.data[i] = rand.Intn(len(s.data))
	}

	return s
}

// byValue updates the struct with the array passed by value recursively according to number of required calls
func byValue[D [1]int | [10]int | [100]int | [1000]int | [10000]int](lvs LargeValueStruct[D], callCount int) {
	callCount--
	if callCount == 0 {
		return
	}
	if lvs.nested != nil {
		for _, nested := range lvs.nested {
			if nested.nested != nil {
				// update the nested structs by value
				byValue[D](nested, callCount)
			}
			// do a random write to a random position of nested struct data (not a trivial to a fixed position) to prevent compiler optimization
			nested.data[rand.Intn(len(lvs.data))] = rand.Intn(len(nested.data))
		}
	}
	// do a random write to a random position of data (not a trivial to a fixed position) to prevent compiler optimization
	lvs.data[rand.Intn(len(lvs.data))] = rand.Intn(len(lvs.data))
	byValue[D](lvs, callCount)
}

// initByValue creates a pointer to a struct instance and sets randomized values to the fields
func initByPointer[D [1]int | [10]int | [100]int | [1000]int | [10000]int](nestedCount, recursion int) *LargePointerStruct[D] {
	s := &LargePointerStruct[D]{}
	var nested []LargePointerStruct[D]
	recursion--
	for j := 0; j < nestedCount; j++ {
		if recursion > 0 {
			nested = append(nested, *initByPointer[D](nestedCount, recursion))
		} else {
			nested = append(nested, LargePointerStruct[D]{})
		}
		// initialize the nested struct input data with random values
		nested[j].data = new(D)
		for i := 0; i < len(*nested[j].data); i++ {
			(*nested[j].data)[i] = rand.Intn(len(*nested[j].data))
		}
	}
	// initialize the input data with random values
	s.nested = &nested
	s.data = new(D)
	for i := 0; i < len(*s.data); i++ {
		(*s.data)[i] = rand.Intn(len(*s.data))
	}

	return s
}

// byPointer updates the struct with a pointer to the array passed by pointer recursively according to number of required calls
func byPointer[D [1]int | [10]int | [100]int | [1000]int | [10000]int](lps *LargePointerStruct[D], callCount int) {
	callCount--
	if callCount == 0 {
		return
	}
	if lps.nested != nil {
		for _, nested := range *lps.nested {
			if nested.nested != nil {
				// update the nested structs by pointer
				byPointer[D](&nested, callCount)

			}
			// do a random write to a random position of nested struct data (not a trivial to a fixed position) to prevent compiler optimization
			(*nested.data)[rand.Intn(len(*nested.data))] = rand.Intn(len(*nested.data))
		}
	}
	// do a random write to a random position of data (not a trivial to a fixed position) to prevent compiler optimization
	(*lps.data)[rand.Intn(len(*lps.data))] = rand.Intn(len(*lps.data))
	byPointer[D](lps, callCount)
}

// BenchmarkArrayStructs defines a benchmark of updates on structs of arrays or pointers to structs of pointer to arrays (of data and of nested structs)
func BenchmarkArrayStructs(b *testing.B) {
	sizes := []int{1, 10, 100, 1000, 10000} // number of ints

	for recursion := 1; recursion <= maxRecursion; recursion++ {
		for nestedCount := 1; nestedCount <= maxNestedCount; nestedCount = nestedCount * 4 {
			for _, size := range sizes {
				for callCount := 1; callCount <= maxCallCount; callCount = callCount * 4 {
					// using array type based on size
					switch size {
					case 1:
						// allocates and initializes the struct of arrays of ints and nested structs before the benchmark
						sv := initByValue[[1]int](nestedCount, recursion)
						b.Run(fmt.Sprintf("ByValue-Recursion_%d_Nested_Struct_Count_%d_Array_Size_%d_CallCount_%d", recursion, nestedCount, size, callCount),
							func(b *testing.B) {
								for i := 0; i < b.N; i++ {
									byValue(sv, callCount)
								}
							})
						// allocates and initializes the pointer to struct of pointer to arrays and nested  structs before the benchmark
						sp := initByPointer[[1]int](nestedCount, recursion)
						b.Run(fmt.Sprintf("ByPointer-Recursion_%d_Nested_Struct_Count_%d_Array_Size_%d_CallCount_%d", recursion, nestedCount, size, callCount),
							func(b *testing.B) {
								for i := 0; i < b.N; i++ {
									byPointer(sp, callCount)
								}
							})
					case 10:
						// allocates and initializes the struct of arrays of ints and nested structs before the benchmark
						sv := initByValue[[10]int](nestedCount, recursion)
						b.Run(fmt.Sprintf("ByValue-Recursion_%d_Nested_Struct_Count_%d_Array_Size_%d_CallCount_%d", recursion, nestedCount, size, callCount),
							func(b *testing.B) {
								for i := 0; i < b.N; i++ {
									byValue(sv, callCount)
								}
							})
						// allocates and initializes the pointer to struct of pointer to arrays and nested  structs before the benchmark
						sp := initByPointer[[10]int](nestedCount, recursion)
						b.Run(fmt.Sprintf("ByPointer-Recursion_%d_Nested_Struct_Count_%d_Array_Size_%d_CallCount_%d", recursion, nestedCount, size, callCount),
							func(b *testing.B) {
								for i := 0; i < b.N; i++ {
									byPointer(sp, callCount)
								}
							})
					case 100:
						// allocates and initializes the struct of arrays of ints and nested structs before the benchmark
						sv := initByValue[[100]int](nestedCount, recursion)
						b.Run(fmt.Sprintf("ByValue-Recursion_%d_Nested_Struct_Count_%d_Array_Size_%d_CallCount_%d", recursion, nestedCount, size, callCount),
							func(b *testing.B) {
								for i := 0; i < b.N; i++ {
									byValue(sv, callCount)
								}
							})
						// allocates and initializes the pointer to struct of pointer to arrays and nested  structs before the benchmark
						sp := initByPointer[[100]int](nestedCount, recursion)
						b.Run(fmt.Sprintf("ByPointer-Recursion_%d_Nested_Struct_Count_%d_Array_Size_%d_CallCount_%d", recursion, nestedCount, size, callCount),
							func(b *testing.B) {
								for i := 0; i < b.N; i++ {
									byPointer(sp, callCount)
								}
							})
					case 1000:
						// allocates and initializes the struct of arrays of ints and nested structs before the benchmark
						sv := initByValue[[1000]int](nestedCount, recursion)
						b.Run(fmt.Sprintf("ByValue-Recursion_%d_Nested_Struct_Count_%d_Array_Size_%d_CallCount_%d", recursion, nestedCount, size, callCount),
							func(b *testing.B) {
								for i := 0; i < b.N; i++ {
									byValue(sv, callCount)
								}
							})
						// allocates and initializes the pointer to struct of pointer to arrays and nested  structs before the benchmark
						sp := initByPointer[[1000]int](nestedCount, recursion)
						b.Run(fmt.Sprintf("ByPointer-Recursion_%d_Nested_Struct_Count_%d_Array_Size_%d_CallCount_%d", recursion, nestedCount, size, callCount),
							func(b *testing.B) {
								for i := 0; i < b.N; i++ {
									byPointer(sp, callCount)
								}
							})
					case 10000:
						// allocates and initializes the struct of arrays of ints and nested structs before the benchmark
						sv := initByValue[[10000]int](nestedCount, recursion)
						b.Run(fmt.Sprintf("ByValue-Recursion_%d_Nested_Struct_Count_%d_Array_Size_%d_CallCount_%d", recursion, nestedCount, size, callCount),
							func(b *testing.B) {
								for i := 0; i < b.N; i++ {
									byValue(sv, callCount)
								}
							})
						// allocates and initializes the pointer to struct of pointer to arrays and nested  structs before the benchmark
						sp := initByPointer[[10000]int](nestedCount, recursion)
						b.Run(fmt.Sprintf("ByPointer-Recursion_%d_Nested_Struct_Count_%d_Array_Size_%d_CallCount_%d", recursion, nestedCount, size, callCount),
							func(b *testing.B) {
								for i := 0; i < b.N; i++ {
									byPointer(sp, callCount)
								}
							})
					}
				}
			}
		}
	}
}
