package sliceconv

import "errors"

func AnyUint8(r []interface{}) ([]uint8, error) {

	var nr []uint8

	for _, d := range r {
		v, s := d.(uint8)

		if s == false {
			return []uint8{}, errors.New("Convert failed to uint8")
		}

		nr = append(nr, v)
	}

	return nr, nil

}

func AnyUint16(r []interface{}) ([]uint16, error) {

	var nr []uint16

	for _, d := range r {
		v, s := d.(uint16)

		if s == false {
			return []uint16{}, errors.New("Convert failed to uint16")
		}

		nr = append(nr, v)
	}

	return nr, nil

}

func AnyUint(r []interface{}) ([]uint, error) {

	var nr []uint

	for _, d := range r {
		v, s := d.(uint)

		if s == false {
			return []uint{}, errors.New("Convert failed to uint")
		}

		nr = append(nr, v)
	}

	return nr, nil

}

func AnyUint64(r []interface{}) ([]uint64, error) {

	var nr []uint64

	for _, d := range r {
		v, s := d.(uint64)

		if s == false {
			return []uint64{}, errors.New("Convert failed to uint64")
		}

		nr = append(nr, v)
	}

	return nr, nil

}