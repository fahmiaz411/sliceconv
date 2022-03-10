package sliceconv

func AnyInt16(r []interface{}) []int16 {

	var nr []int16

	for _, d := range r {
		v, er := d.(int16)

		if er == true {
			return []int16{}
		}

		nr = append(nr, v)
	}

	return nr

}