## Which are in?

### Detail

Given two arrays of strings `a1` and `a2` return a sorted array `r` in lexicographical order of the strings of `a1`
which are substrings of strings of `a2`.

### Examples

###### Example 1

```golang
a1 := {"arp", "live", "strong"}
a2 := {"lively", "alive", "harp", "sharp", "armstrong"}

r := {"arp", "live", "strong"}
```

###### Example 2

```golang
a1 := {"tarp", "mice", "bull"}
a2 := {"lively", "alive", "harp", "sharp", "armstrong"}

r := {}
```

### Notes

- Arrays are written in "general" notation. See "Your Test Cases" for examples in your language.
- In Shell bash `a1` and `a2` are strings. The return is a string where words are separated by commas.
- Beware: `r` must be without duplicates.
- Don't mutate the inputs.

### Link

https://www.codewars.com/kata/which-are-in