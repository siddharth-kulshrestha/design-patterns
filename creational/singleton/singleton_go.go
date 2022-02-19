package main

import (
	"sync"
)

// In single ton pattern only one object for a struct / class can exists at runtime.

// var mutex = &sync.Mutex{}

var once sync.Once

var (
	instantiated *sampleStructWithSingleTon
)

type sampleStructWithSingleTon struct {
	A, B string
}

func GetSampleStructWithSingleTon(a, b string) *sampleStructWithSingleTon {
	// if instantiated == nil {
	// 	mutex.Lock()
	// 	defer mutex.Unlock()
	// 	return &sampleStructWithSingleTon{
	// 		A: a,
	// 		B: b,
	// 	}
	// } else {
	// 	return instantiated
	// }
	once.Do(func() {
		instantiated = &sampleStructWithSingleTon{
			A: a,
			B: b,
		}
	})
	return instantiated
}
