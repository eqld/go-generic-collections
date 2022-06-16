package collections

// AccumulatorNumberSum returns generic function to calculate sum of all elemnts of a generic collection (e.g. `SetReduce()`).
func AccumulatorNumberSum[T Number]() func(partialResult, element T) T {
	return func(partialResult, element T) T { return partialResult + element }
}

// AccumulatorNumberMul returns generic function to calculate product of all elemnts of a generic collection (e.g. `SetReduce()`).
func AccumulatorNumberMul[T Number]() func(partialResult, element T) T {
	return func(partialResult, element T) T { return partialResult * element }
}
