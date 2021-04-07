## Factorial decomposition

### Detail

The aim of the kata is to decompose `n!` (factorial `n`) into its prime factors.

Prime numbers should be in increasing order. When the exponent of a prime is 1 don't put the exponent.

### Examples

###### Example 1

```golang
n = 12
Decomp(12) //should return "2^10 * 3^5 * 5^2 * 7 * 11"
```

###### Example 2

```golang
n = 22
Decomp(12) //should return "2^19 * 3^9 * 5^4 * 7^3 * 11^2 * 13 * 17 * 19"
```

###### Example 3

```golang
n = 25
Decomp(12) //should return "2^22 * 3^10 * 5^6 * 7^3 * 11^2 * 13 * 17 * 19 * 23"
```

### Notes

- The function is `Decomp(n)` and should return the decomposition of `n!` into its prime factors in increasing order of
  the primes, as a string.
- Factorial can be a very big number (`4000!` has 12674 digits, `n` will go from 300 to 4000).

### Link

https://www.codewars.com/kata/factorial-decomposition