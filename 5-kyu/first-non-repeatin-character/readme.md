## First non-repeating character

### Detail

Write a function named `FirstNonRepeating` that takes a string input, and returns the first character that is not
repeated anywhere in the string.

For example, if given the input `"stress"`, the function should return `"t"`, since the letter `t` only occurs once in
the string, and occurs first in the string.

As an added challenge, upper- and lowercase letters are considered the same character, but the function should return
the correct case for the initial letter. For example, the input `"sTreSS"` should return `"T"`.

If a string contains all repeating characters, it should return an empty string `""`.

### Examples

###### Example 1

```golang
FirstNonRepeating("sTreSS") // should return  "T"
```

###### Example 2

```golang
FirstNonRepeating("~><#~><") // should return  "#"
```

### Link

https://www.codewars.com/kata/first-non-repeating-character