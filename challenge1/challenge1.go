package challenge1

// Run calculates the nth value of the Fibonacci Sequence
func Run(n int) uint64 {
  if n == 0 || n == 1 {
    return uint64(n)
  }
	return fib(n-2, 1, 0)
}

func fib(n int, n1, n2 uint64) uint64 {
  if n == 0 {
    return n1 + n2
  }
  return fib(n-1, n1+n2, n1)
}

