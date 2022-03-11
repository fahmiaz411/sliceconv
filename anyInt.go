package sliceconv

func AnyInt(r []interface{}) []int {

	var nr []int

	for _, d := range r {
		v, s := d.(int)

		if s == false {
			return []int{}
		}

		nr = append(nr, v)
	}

	return nr

}