package natural

import (
	"sort"
	"unicode"
)

func compRight(ra, rb []rune) int {
	bias := 0
	la, lb := len(ra), len(rb)
	var ca, cb rune
	for i := 0; i < la || i < lb; i++ {
		if i < la {
			ca = ra[i]
		} else {
			ca = 0
		}
		if i < lb {
			cb = rb[i]
		} else {
			cb = 0
		}

		da, db := unicode.IsNumber(ca), unicode.IsNumber(cb)
		switch {
		case !da && !db:
			return bias
		case !da:
			return -1
		case !db:
			return 1
		case ca < cb:
			if bias == 0 {
				bias = -1
			}
		case ca > cb:
			if bias == 0 {
				bias = 1
			}
		case ca == 0 && cb == 0:
			return bias
		}
	}

	return 0
}

func compLeft(ra, rb []rune) int {
	la, lb := len(ra), len(rb)
	var ca, cb rune
	i := 0
	for {
		if i < la {
			ca = ra[i]
		} else {
			ca = 0
		}
		if i < lb {
			cb = rb[i]
		} else {
			cb = 0
		}

		da, db := unicode.IsNumber(ca), unicode.IsNumber(cb)
		switch {
		case !da && !db:
			return 0
		case !da:
			return -1
		case !db:
			return 1
		case ca < cb:
			return -1
		case ca > cb:
			return 1
		}
		i++
	}

	return 0
}

func naturalComp(a, b string, foldCase bool) int {
	ra, rb := []rune(a), []rune(b)
	la, lb := len(ra), len(rb)
	ia, ib := 0, 0

	for {
		if ia >= la && ib >= lb {
			return 0
		} else if ia >= la {
			return -1
		} else if ib >= lb {
			return 1
		}
		ca, cb := ra[ia], rb[ib]

		for unicode.IsSpace(ca) {
			ia++
			if ia < la {
				ca = ra[ia]
			} else {
				ca = 0
			}
		}
		for unicode.IsSpace(cb) {
			ib++
			if ib < lb {
				cb = rb[ib]
			} else {
				cb = 0
			}
		}

		if unicode.IsNumber(ca) && unicode.IsNumber(cb) {
			var r int
			if ca == '0' || cb == '0' {
				r = compLeft(ra[ia:], rb[ib:])
				if r != 0 {
					return r
				}
			} else {
				r = compRight(ra[ia:], rb[ib:])
				if r != 0 {
					return r
				}
			}
		}

		if foldCase {
			ca = unicode.ToUpper(ca)
			cb = unicode.ToUpper(cb)
		}

		if ca < cb {
			return -1
		} else if ca > cb {
			return 1
		}

		ia++
		ib++
	}

	return 0
}

// NaturalComp behave like strcmp but natural comparing.
func NaturalComp(a, b string) int {
	return naturalComp(a, b, false)
}

// NaturalCaseComp behave like strcmp but natural comparing with fold case.
func NaturalCaseComp(a, b string) int {
	return naturalComp(a, b, true)
}

// Sort do natural sort.
func Sort(a []string) {
	sort.Sort(StringSlice(a))
}

// SortCase do natural sort with fold case.
func SortCase(a []string) {
	sort.Sort(StringCaseSlice(a))
}

// StringSlice implement sort.Sorter interface by natural comparing.
type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return naturalComp(p[i], p[j], false) < 0 }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p StringSlice) Sort()              { sort.Sort(p) }

// StringCaseSlice implement sort.Sorter interface by natural comparing with fold case.
type StringCaseSlice []string

func (p StringCaseSlice) Len() int           { return len(p) }
func (p StringCaseSlice) Less(i, j int) bool { return naturalComp(p[i], p[j], true) < 0 }
func (p StringCaseSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p StringCaseSlice) Sort()              { sort.Sort(p) }
