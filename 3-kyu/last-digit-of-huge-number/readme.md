## Last digit of a huge number

### Detail

For a given list `{x1, x2, x3, ..., xn}` compute the last (decimal) digit of `x1 ^ (x2 ^ (x3 ^ (... ^ xn)))`.

### Examples

###### Example 1

```golang
LastDigit([]int{3, 4, 2}) // should return 1
```

because 3 ^ (4 ^ 2) = 3 ^ 16 = 43046721

### Notes

- Powers grow incredibly fast. For example, `9 ^ (9 ^ 9)` has more than 369 millions of digits. `LastDigit` has to deal
  with such numbers efficiently.
- `0 ^ 0 = 1`
- `0 ^ (0 ^ 0) = 0 ^ 1 = 0`
- Treat empty list as equals to `1`

### Link

https://www.codewars.com/kata/last-digit-of-a-huge-number