package challenge1

// Run calculates the nth value of the Fibonacci Sequence
func Run(n int) uint64 {
	return fib(n, 0, 1)
}

func fib(n int, secondToLastNumber, lastNumber uint64) uint64 {
	if n == 0 {
		return secondToLastNumber
	}

	number := secondToLastNumber + lastNumber
	return fib(n-1, lastNumber, number)
}
