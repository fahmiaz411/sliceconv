package sliceconv

func AnyUint64(r []interface{}) []uint64 {

	var nr []uint64

	for _, d := range r {
		v, s := d.(uint64)

		if s == false {
			return []uint64{}
		}

		nr = append(nr, v)
	}

	return nr

}