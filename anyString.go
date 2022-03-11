package sliceconv

func AnyString(r []interface{}) []string {

	var nr []string

	for _, d := range r {
		v, s := d.(string)

		if s == false {
			return []string{}
		}

		nr = append(nr, v)
	}

	return nr

}