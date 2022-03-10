package sliceconv

func AnyUint64(r []interface{}) []uint64 {

	var nr []uint64

	for _, d := range r {
		v, er := d.(uint64)

		if er == true {
			return []uint64{}
		}

		nr = append(nr, v)
	}

	return nr

}