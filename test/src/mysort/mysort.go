package mysort

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

func IsSorted(data Sorter) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

type IntArray []int

func (p IntArray) Len() int           { return len(p) }
func (list IntArray) Less(i, j int) bool { return list[i] < list[j] }
func (list IntArray) Swap(i, j int)      { list[i], list[j] = list[j], list[i] }

type StringArray []string

func (list StringArray) Len() int           { return len(list) }
func (list StringArray) Less(i, j int) bool { return list[i] < list[j] }
func (list StringArray) Swap(i, j int)      { list[i], list[j] = list[j], list[i] }

func SortInts(a []int)       { Sort(IntArray(a)) }
func SortStrings(a []string) { Sort(StringArray(a)) }

func IntsAreSorted(a []int) bool       { return IsSorted(IntArray(a)) }
func StringsAreSorted(a []string) bool { return IsSorted(StringArray(a)) }
