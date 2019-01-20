package status

import (
	"math/rand"
	"testing"

	"github.com/miltfra/tools"
)

func TestStatusPut1000(t *testing.T) {
	arr := tools.RndFloat64Arr(1000)
	st := New(1000)
	for i := 0; i < 1000; i++ {
		st.Put(&Element{make([][]float64, 1000), make([][]bool, 1000), make([]bool, 1000), 0, 1, arr[i]})
		if !st.Check() {
			t.FailNow()
		}
	}
}
func TestStatusPut10000(t *testing.T) {
	arr := tools.RndIntArr(10000)
	st := New(10000)
	for i := 0; i < 10000; i++ {
		st.Put(&Element{make([][]float64, 10000), make([][]bool, 10000), make([]bool, 10000), 0, 1, arr[i]})
	}
	if !st.Check() {
		t.FailNow()
	}
}

func TestStatusPutGet(t *testing.T) {
	arr := tools.RndFloat64Arr(10000)
	st := New(10000)
	count := 0
	for i := 0; i < 10000; i++ {
		if rand.Intn(10) > 6 {
			if st.Get() != nil {
				count--
			}
			if !st.Check() {
				t.FailNow()
			}
			if st.Length != count {
				t.FailNow()
			}
		}
		st.Put(&Element{make([][]float64, 10000), make([][]bool, 10000), make([]bool, 10000), 0, 1, arr[i]})
		count++
		if !st.Check() {
			t.FailNow()
		}
		if st.Length != count {
			t.FailNow()
		}
	}
}
