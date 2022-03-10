package sliceconv

func AnyInt8(r []interface{}) []int8 {

	var nr []int8

	for _, d := range r {
		v, er := d.(int8)

		if er == true {
			return []int8{}
		}

		nr = append(nr, v)
	}

	return nr

}