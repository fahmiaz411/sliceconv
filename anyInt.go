package sliceconv

func AnyInt(r []interface{}) []int {

	var nr []int

	for _, d := range r {
		v, er := d.(int)

		if er == true {
			return []int{}
		}

		nr = append(nr, v)
	}

	return nr

}