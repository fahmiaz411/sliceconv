package sliceconv

func AnyUint16(r []interface{}) []uint16 {

	var nr []uint16

	for _, d := range r {
		v, s := d.(uint16)

		if s == false {
			return []uint16{}
		}

		nr = append(nr, v)
	}

	return nr

}