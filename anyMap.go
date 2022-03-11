package sliceconv

import "errors"

func AnyMapStrStr(r []interface{}) ([]map[string]string, error) {

	var nr []map[string]string

	for _, sl := range r {

		m := map[string]string{}
		
		n_sl, s := sl.(map[string]string)

		if s == false {
			return nil, errors.New("Convert failed to map[string]string")
		}

		for k, v := range n_sl{
			m[k] = v
		}

		nr = append(nr, m)
	}

	return nr, nil

}


func AnyMapStrInt(r []interface{}) ([]map[string]int, error) {

	var nr []map[string]int

	for _, sl := range r {

		m := map[string]int{}
		
		n_sl, s := sl.(map[string]int)

		if s == false {
			return nil, errors.New("Convert failed to map[string]int")
		}

		for k, v := range n_sl{
			m[k] = v
		}

		nr = append(nr, m)
	}

	return nr, nil

}

func AnyMapStrInt64(r []interface{}) ([]map[string]int64, error) {

	var nr []map[string]int64

	for _, sl := range r {

		m := map[string]int64{}
		
		n_sl, s := sl.(map[string]int64)

		if s == false {
			return nil, errors.New("Convert failed to map[string]int")
		}

		for k, v := range n_sl{
			m[k] = v
		}

		nr = append(nr, m)
	}

	return nr, nil

}
