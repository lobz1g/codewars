## Mumbling

### Detail

This time no story, no theory. The examples below show you how to write function `Accum`

### Examples

###### Example 1

```golang
Accum("abcd") // should return "A-Bb-Ccc-Dddd"
```

###### Example 2

```golang
Accum("RqaEzty") // should return "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
```

###### Example 3

```golang
Accum("cwAt") // should return "C-Ww-Aaa-Tttt"
```

### Note

The parameter of `Accum` is a string which includes only letters from `a..z` and `A..Z`.

### Link

https://www.codewars.com/kata/mumbling