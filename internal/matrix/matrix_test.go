package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)
	m := New[int](3, 3)
	assert.Equal(3, len(*m))
	assert.Equal(3, len((*m)[0]))
}

func SetRow(t *testing.T) {
	assert := assert.New(t)
	m := New[int](3, 3)
	m.
		SetRow(0, []int{1, 2, 3}).
		SetRow(1, []int{4, 5, 6}).
		SetRow(2, []int{7, 8, 9})
	assert.Equal(5, (*m)[1][1])
}

func TestProduct(t *testing.T) {
	assert := assert.New(t)
	A := New[int](2, 2).SetRow(0, []int{0, 1}).SetRow(1, []int{0, 0})
	B := New[int](2, 2).SetRow(0, []int{0, 0}).SetRow(1, []int{1, 0})
	want := New[int](2, 2).SetRow(0, []int{1, 0}).SetRow(1, []int{0, 0})
	C, err := A.Product(B)
	assert.NoError(err)
	D, err := B.Product(A)
	assert.NoError(err)
	assert.Equal(want, C)
	assert.NotEqual(want, D)
	E := New[int](3, 1)
	_, err = A.Product(E)
	assert.Error(err)
}

func TestAdd(t *testing.T) {
	assert := assert.New(t)
	A := New[int](2, 2).SetRow(0, []int{0, 1}).SetRow(1, []int{0, 0})
	B := New[int](2, 2).SetRow(0, []int{0, 0}).SetRow(1, []int{1, 0})
	want := New[int](2, 2).SetRow(0, []int{0, 1}).SetRow(1, []int{1, 0})
	C, err := A.Add(B)
	assert.NoError(err)
	assert.Equal(want, C)
	D, err := B.Add(A)
	assert.NoError(err)
	assert.Equal(want, D)
	E := New[int](3, 1)
	_, err = A.Add(E)
	assert.Error(err)
}

func TestScalar(t *testing.T) {
	assert := assert.New(t)
	m := New[int](3, 3)
	m.
		SetRow(0, []int{1, 2, 3}).
		SetRow(1, []int{4, 5, 6}).
		SetRow(2, []int{7, 8, 9})
	want := New[int](3, 3)
	want.
		SetRow(0, []int{3, 6, 9}).
		SetRow(1, []int{12, 15, 18}).
		SetRow(2, []int{21, 24, 27})
	assert.Equal(want, m.Scalar(3))
}

func TestDeterminantClassic(t *testing.T) {
	assert := assert.New(t)
	a := New[int](5, 5).
		SetRow(0, []int{77, 88, 99, 12, 42}).
		SetRow(1, []int{61, 47, 8, 19, 41}).
		SetRow(2, []int{1, 22, 13, 74, 55}).
		SetRow(3, []int{3, 17, 58, 3, 32}).
		SetRow(4, []int{91, 27, 49, 4, 65})
	want := -546499540
	got, err := a.DeterminantClassic()
	assert.NoError(err)
	assert.Equal(want, got)
	b := New[int](2, 2).
		SetRow(0, []int{4, 7}).
		SetRow(1, []int{3, 2})
	got, err = b.DeterminantClassic()
	assert.NoError(err)
	want = 4*2 - 3*7
	assert.Equal(want, got)
}

func TestTranspose(t *testing.T) {
	assert := assert.New(t)
	a := New[int](5, 3).
		SetRow(0, []int{77, 88, 99}).
		SetRow(1, []int{61, 47, 8}).
		SetRow(2, []int{1, 22, 13}).
		SetRow(3, []int{3, 17, 58}).
		SetRow(4, []int{91, 27, 49})
	want := New[int](3, 5).
		SetRow(0, []int{77, 61, 1, 3, 91}).
		SetRow(1, []int{88, 47, 22, 17, 27}).
		SetRow(2, []int{99, 8, 13, 58, 49})
	got := a.Transpose()
	assert.Equal(want, got)
}

func TestDeterminant(t *testing.T) {
	assert := assert.New(t)
	a := New[int](5, 5).
		SetRow(0, []int{77, 88, 99, 12, 42}).
		SetRow(1, []int{61, 47, 8, 19, 41}).
		SetRow(2, []int{1, 22, 13, 74, 55}).
		SetRow(3, []int{3, 17, 58, 3, 32}).
		SetRow(4, []int{91, 27, 49, 4, 65})
	got, err := a.Determinant()
	assert.NoError(err)
	want, err := a.DeterminantClassic()
	assert.NoError(err)
	assert.Equal(want, got)
}

func BenchmarkDeterminant(b *testing.B) {
	a := New[int](5, 5).
		SetRow(0, []int{77, 88, 99, 12, 42}).
		SetRow(1, []int{61, 47, 8, 19, 41}).
		SetRow(2, []int{1, 22, 13, 74, 55}).
		SetRow(3, []int{3, 17, 58, 3, 32}).
		SetRow(4, []int{91, 27, 49, 4, 65})
	for i := 0; i < b.N; i++ {
		if _, err := a.Determinant(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDeterminantClassic(b *testing.B) {
	a := New[int](5, 5).
		SetRow(0, []int{77, 88, 99, 12, 42}).
		SetRow(1, []int{61, 47, 8, 19, 41}).
		SetRow(2, []int{1, 22, 13, 74, 55}).
		SetRow(3, []int{3, 17, 58, 3, 32}).
		SetRow(4, []int{91, 27, 49, 4, 65})
	for i := 0; i < b.N; i++ {
		if _, err := a.DeterminantClassic(); err != nil {
			b.Fatal(err)
		}
	}
}
