package sliceconv

func AnyUint16(r []interface{}) []uint16 {

	var nr []uint16

	for _, d := range r {
		v, er := d.(uint16)

		if er == true {
			return []uint16{}
		}

		nr = append(nr, v)
	}

	return nr

}