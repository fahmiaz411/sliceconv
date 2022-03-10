package sliceconv

func AnyUint8(r []interface{}) []uint8 {

	var nr []uint8

	for _, d := range r {
		v, er := d.(uint8)

		if er == true {
			return []uint8{}
		}

		nr = append(nr, v)
	}

	return nr

}