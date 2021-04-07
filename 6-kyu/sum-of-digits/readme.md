## Sum of Digits / Digital Root

### Detail

A digital root is the recursive sum of all the digits in a number. Given `n`, take the sum of the digits of `n`. If that
value has two digits, continue reducing in this way until a single-digit number is produced. This is only applicable to
the natural numbers.

### Examples

###### Example 1

```golang
DigitalRoot(16) //should return 7 (1 + 6)
```

###### Example 2

```golang
Solution(942) //should return 6 (9 + 4 + 2 => 15 => 1 + 5)
```

### Link

https://www.codewars.com/kata/sum-of-digits-slash-digital-root