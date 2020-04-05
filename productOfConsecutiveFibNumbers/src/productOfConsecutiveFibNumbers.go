package src

func ProductFib(prod uint64) [3]uint64 {
	return fib(1, 2, prod)
}

func fib(a, b, prod uint64) [3]uint64 {
	c := a + b
	tmpProd := c * b
	if tmpProd == prod {
		return [3]uint64{b, c, 1}
	} else if tmpProd < prod {
		return fib(b, c, prod)
	}
	return [3]uint64{b, c, 0}
}
