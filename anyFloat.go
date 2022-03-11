package sliceconv

import "errors"

func AnyFloat32(r []interface{}) ([]float32, error) {

	var nr []float32

	for _, d := range r {
		v, s := d.(float32)

		if s == false {
			return []float32{}, errors.New("Convert failed to float32")
		}

		nr = append(nr, v)
	}

	return nr, nil

}

func AnyFloat64(r []interface{}) ([]float64, error) {

	var nr []float64

	for _, d := range r {
		v, s := d.(float64)

		if s == false {
			return []float64{}, errors.New("Convert failed to float64")
		}

		nr = append(nr, v)
	}

	return nr, nil

}