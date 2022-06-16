package collections

func AccumulatorNumberSum[T Number]() func(partialResult, element T) T {
	return func(partialResult, element T) T { return partialResult + element }
}

func AccumulatorNumberMul[T Number]() func(partialResult, element T) T {
	return func(partialResult, element T) T { return partialResult * element }
}
