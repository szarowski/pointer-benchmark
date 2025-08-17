package main

import (
	"math/rand"
	"testing"
)

// LargeValueStruct defines a struct with generically 5 types of arrays
type LargeValueStruct[T [1]int | [10]int | [100]int | [1000]int | [10000]int] struct {
	data T
}

// LargePointerStruct defines a struct with a pointer to generically 5 types of arrays
type LargePointerStruct[T [1]int | [10]int | [100]int | [1000]int | [10000]int] struct {
	data *T
}

// initByValue creates a struct instance and sets values
func initByValue[T [1]int | [10]int | [100]int | [1000]int | [10000]int]() LargeValueStruct[T] {
	s := LargeValueStruct[T]{}
	for i := 0; i < len(s.data); i++ {
		s.data[i] = rand.Intn(len(s.data))
	}
	return s
}

// byValue updates the struct with the array passed by value
func byValue[T [1]int | [10]int | [100]int | [1000]int | [10000]int](s LargeValueStruct[T]) {
	// Do a random write to a random position (not a trivial to a fixed position) to prevent compiler optimization
	s.data[rand.Intn(len(s.data))] = rand.Intn(len(s.data))
}

// initByValue creates a pointer to a struct instance and sets values
func initByPointer[T [1]int | [10]int | [100]int | [1000]int | [10000]int]() *LargePointerStruct[T] {
	s := &LargePointerStruct[T]{}
	s.data = new(T)
	for i := 0; i < len(*s.data); i++ {
		(*s.data)[i] = rand.Intn(len(*s.data))
	}
	return s
}

// byPointer updates the struct with a pointer to the array passed by pointer
func byPointer[T [1]int | [10]int | [100]int | [1000]int | [10000]int](s *LargePointerStruct[T]) {
	// Do a random write to a random position (not a trivial to a fixed position) to prevent compiler optimization
	(*s.data)[rand.Intn(len(*s.data))] = rand.Intn(len(*s.data))
}

// BenchmarkArrayStructs defines a benchmark of updates on structs of arrays or pointers to structs of pointer to arrays
func BenchmarkArrayStructs(b *testing.B) {
	sizes := []int{1, 10, 100, 1000, 10000} // number of ints

	for _, size := range sizes {
		// using array type based on size
		switch size {
		case 1:
			// allocates and initializes the struct of arrays before the benchmark
			sv := initByValue[[1]int]()
			b.Run("Value_Size_1", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					byValue(sv)
				}
			})
			// allocates and initializes the pointer to struct of pointer to arrays before the benchmark
			sp := initByPointer[[1]int]()
			b.Run("Pointer_Size_1", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					byPointer(sp)
				}
			})
		case 10:
			// allocates and initializes the struct of arrays before the benchmark
			sv := initByValue[[10]int]()
			b.Run("Value_Size_10", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					byValue(sv)
				}
			})
			// allocates and initializes the pointer to struct of pointer to arrays before the benchmark
			sp := initByPointer[[10]int]()
			b.Run("Pointer_Size_10", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					byPointer(sp)
				}
			})
		case 100:
			// allocates and initializes the struct of arrays before the benchmark
			sv := initByValue[[100]int]()
			b.Run("Value_Size_100", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					byValue(sv)
				}
			})
			// allocates and initializes the pointer to struct of pointer to arrays before the benchmark
			sp := initByPointer[[100]int]()
			b.Run("Pointer_Size_100", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					byPointer(sp)
				}
			})
		case 1000:
			// allocates and initializes the struct of arrays before the benchmark
			sv := initByValue[[1000]int]()
			b.Run("Value_Size_1000", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					byValue(sv)
				}
			})
			// allocates and initializes the pointer to struct of pointer to arrays before the benchmark
			sp := initByPointer[[1000]int]()
			b.Run("Pointer_Size_1000", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					byPointer(sp)
				}
			})
		case 10000:
			// allocates and initializes the struct of arrays before the benchmark
			sv := initByValue[[10000]int]()
			b.Run("Value_Size_10000", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					byValue(sv)
				}
			})
			// allocates and initializes the pointer to struct of pointer to arrays before the benchmark
			sp := initByPointer[[10000]int]()
			b.Run("Pointer_Size_10000", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					byPointer(sp)
				}
			})
		}
	}
}
