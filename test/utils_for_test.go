package dsl_test

// constants used for testing and benchmarking
const (
	num    = 10000
	bigNum = 1000000
)

// data type used for testing purposes only
type vector struct {
	x int
	y int
	z int
}

// helper function to generate different content vectors
func newVector(index int) *vector {
	return &vector{x: index, y: index + 1, z: 2 * index}
}
