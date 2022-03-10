package sliceconv

func AnyUint(r []interface{}) []uint {

	var nr []uint

	for _, d := range r {
		v, er := d.(uint)

		if er == true {
			return []uint{}
		}

		nr = append(nr, v)
	}

	return nr

}