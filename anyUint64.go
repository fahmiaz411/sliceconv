package sliceconv

func AnyInt64(r []interface{}) []int64 {

	var nr []int64

	for _, d := range r {
		v, s := d.(int64)

		if s == false {
			return []int64{}
		}

		nr = append(nr, v)
	}

	return nr

}