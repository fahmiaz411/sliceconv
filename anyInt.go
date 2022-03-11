package sliceconv

import "errors"

func AnyInt8(r []interface{}) ([]int8, error) {

	var nr []int8

	for _, d := range r {
		v, s := d.(int8)

		if s == false {
			return []int8{}, errors.New("Convert failed to int8")
		}

		nr = append(nr, v)
	}

	return nr, nil

}

func AnyInt16(r []interface{}) ([]int16, error) {

	var nr []int16

	for _, d := range r {
		v, s := d.(int16)

		if s == false {
			return []int16{}, errors.New("Convert failed to int16")
		}

		nr = append(nr, v)
	}

	return nr, nil

}

func AnyInt(r []interface{}) ([]int, error) {

	var nr []int

	for _, d := range r {
		v, s := d.(int)

		if s == false {
			return []int{}, errors.New("Convert failed to int")
		}

		nr = append(nr, v)
	}

	return nr, nil

}

func AnyInt64(r []interface{}) ([]int64, error) {

	var nr []int64

	for _, d := range r {
		v, s := d.(int64)

		if s == false {
			return []int64{}, errors.New("Convert failed to int64")
		}

		nr = append(nr, v)
	}

	return nr, nil

}