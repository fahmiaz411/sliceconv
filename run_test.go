package sliceconv

import (
	"testing"
)

func TestAnyInt(t *testing.T) {

}

func BenchmarkAnyMapStrStr(b *testing.B) {
	for i:=0; i < b.N; i++{
		a := []interface{}{
			map[string]string{"g":"g", "f":"f",},
			map[string]string{"g":"g", "f":"f",},
			map[string]string{"g":"g", "f":"f",},
		}
	
		AnyMapStrStr(a)
	}
}

func BenchmarkAnyMapStrInt(b *testing.B) {
	for i:=0; i < b.N; i++{
		a := []interface{}{
			map[string]int{"g":1, "f":2,},
			map[string]int{"g":1, "f":2,},
			map[string]int{"g":1, "f":2,},
		}
	
		AnyMapStrInt(a)
	}
}

func BenchmarkAnyMapStrInt64(b *testing.B) {
	for i:=0; i < b.N; i++{
		a := []interface{}{
			map[string]int64{"g":1, "f":2,},
			map[string]int64{"g":1, "f":2,},
			map[string]int64{"g":1, "f":2,},
		}
	
		AnyMapStrInt64(a)
	}
}

func BenchmarkAnyInt(b *testing.B) {
	for i:=0; i < b.N; i++{
		a := []interface{}{
			1,2,3,
		}
		b := []interface{}{
			1,2,3,
		}

		AnyInt(a)
		AnyInt(b)
	}
}