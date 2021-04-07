## Simple Fun #122: String Constructing

### Detail

You are given two string `a` an `s`. Starting with an empty string we can do the following two operations:

```
append the given string a to the end of the current string.
erase one symbol of the current string.
```

Your task is to find the least number of operations needed to construct the given string `s`. Assume that all the
letters from `s` appear in `a` at least once.

### Examples

###### Example 1

```golang
StringConstructing("aba", "abbaab") //should return 6 ("" -> "aba" -> "ab" -> "ababa" -> "abba" -> "abbaaba" -> "abbaab")
```

###### Example 2

```golang
StringConstructing("aba", "abaa") //should return 4 ("" -> "aba" -> "abaaba" -> "abaab" -> "abaa")
```

### Link

https://www.codewars.com/kata/simple-fun-number-122-string-constructing