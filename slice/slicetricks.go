package slice

//generic type T (for code readability)
type T interface{}

//delete element at index e
func DeleteAt(s []T, i int) {
	copy(s[i:], s[i+1:])
	s[len(s)-1] = new(T) // or the zero value of T
	s = s[:len(s)-1]
}

//check if element e exists in slice s and return its position
func Search(s []T, e T) (bool,int) {
	for i,t := range s{
		if t == e{
			return true,i
		}
	}
	return false,0
}

//delete element e, return true if the element existed
//func Delete(e T) bool

