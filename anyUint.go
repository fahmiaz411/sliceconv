package sliceconv

func AnyUint(r []interface{}) []uint {

	var nr []uint

	for _, d := range r {
		v, s := d.(uint)

		if s == false {
			return []uint{}
		}

		nr = append(nr, v)
	}

	return nr

}