package src

import (
	"regexp"
	"strings"
)

type (
	Value struct {
		Str    string
		Number int
	}
	Res []Value
)

func Mix(s1, s2 string) string {
	resAs := getMaxCount(s1, s2)
	res := resAs.ToString()
	return res
}

func getMaxCount(s1, s2 string) Res {
	ar1 := strings.Split(s1, "")
	ar2 := strings.Split(s2, "")
	res := Res{}
	allChars := map[string]int{}
	for _, c := range ar1 {
		if !checkCharInAr(c, allChars) {
			if checkLowerCaseChar(c) {
				allChars[c]++
				v1 := strings.Count(s1, c)
				v2 := strings.Count(s2, c)
				if v1 > 1 || v2 > 1 {
					if v1 == v2 {
						res = res.setValue("=:"+strings.Repeat(c, v1), v1)
					} else if v1 > v2 {
						res = res.setValue("1:"+strings.Repeat(c, v1), v1)
					} else {
						res = res.setValue("2:"+strings.Repeat(c, v2), v2)
					}
				}
			}
		}
	}

	for _, c := range ar2 {
		if !checkCharInAr(c, allChars) {
			if checkLowerCaseChar(c) {
				allChars[c]++
				v := strings.Count(s2, c)
				if v > 1 {
					res = res.setValue("2:"+strings.Repeat(c, v), v)
				}
			}
		}
	}

	return res
}

func (r Res) setValue(s string, n int) Res {
	res := Res{}
	if len(r) == 0 {
		res = append(res, Value{s, n})
		return res
	}

	for key, value := range r {
		if n > value.Number {
			res = append(res, Value{s, n})
			res = append(res, r[key:]...)
			return res
		} else if n == value.Number {
			if strings.Compare(s, value.Str) == -1 {
				res = append(res, Value{s, n})
				res = append(res, r[key:]...)
				return res
			}
			res = append(res, r[key])
		} else {
			res = append(res, r[key])
		}
	}
	res = append(res, Value{s, n})
	return res
}

func (r Res) ToString() string {
	res := ""
	for _, value := range r {
		if res != "" {
			res += "/"
		}
		res += value.Str
	}
	return res
}

func checkCharInAr(c string, a map[string]int) bool {
	_, in := a[c]
	return in
}

func checkLowerCaseChar(c string) bool {
	ok, _ := regexp.MatchString(`[a-z]`, string(c))
	return ok
}
