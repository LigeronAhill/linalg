package matrix

import (
	"testing"

	"github.com/LigeronAhill/linalg/internal/rational"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)
	m := New(3, 3)
	assert.Equal(3, len(*m))
	assert.Equal(3, len((*m)[0]))
}

func SetRow(t *testing.T) {
	assert := assert.New(t)
	m := New(3, 3)
	m.
		SetRow(0, []*rational.Rational{rational.ParseInt(1), rational.ParseInt(2), rational.ParseInt(3)}).
		SetRow(1, []*rational.Rational{rational.ParseInt(4), rational.ParseInt(5), rational.ParseInt(6)}).
		SetRow(2, []*rational.Rational{rational.ParseInt(7), rational.ParseInt(8), rational.ParseInt(9)})
	assert.Equal(rational.ParseInt(5), (*m)[1][1])
}

func TestProduct(t *testing.T) {
	assert := assert.New(t)
	r0 := []*rational.Rational{rational.ParseInt(0), rational.ParseInt(0)}
	r1 := []*rational.Rational{rational.ParseInt(1), rational.ParseInt(0)}
	r2 := []*rational.Rational{rational.ParseInt(0), rational.ParseInt(1)}
	A := New(2, 2).SetRow(0, r2).SetRow(1, r0)    // 01 00
	B := New(2, 2).SetRow(0, r0).SetRow(1, r1)    // 00 10
	want := New(2, 2).SetRow(0, r1).SetRow(1, r0) // 10 00
	C, err := A.Product(B)
	assert.NoError(err)
	D, err := B.Product(A)
	assert.NoError(err)
	for i, row := range *C {
		for j := range row {
			got, err := C.Get(i, j)
			assert.NoError(err)
			want, err := want.Get(i, j)
			assert.NoError(err)
			assert.Equal(got, want)
		}
	}
	assert.Equal(want, C)
	assert.NotEqual(want, D)
	E := New(3, 1)
	_, err = A.Product(E)
	assert.Error(err)
}

func TestAdd(t *testing.T) {
	assert := assert.New(t)
	r0 := []*rational.Rational{rational.ParseInt(0), rational.ParseInt(0)}
	r1 := []*rational.Rational{rational.ParseInt(1), rational.ParseInt(0)}
	r2 := []*rational.Rational{rational.ParseInt(0), rational.ParseInt(1)}
	A := New(2, 2).SetRow(0, r2).SetRow(1, r0)
	B := New(2, 2).SetRow(0, r0).SetRow(1, r1)
	want := New(2, 2).SetRow(0, r2).SetRow(1, r1)
	C, err := A.Add(B)
	assert.NoError(err)
	assert.Equal(want, C)
	D, err := B.Add(A)
	assert.NoError(err)
	assert.Equal(want, D)
	E := New(3, 1)
	_, err = A.Add(E)
	assert.Error(err)
}

func TestScalar(t *testing.T) {
	assert := assert.New(t)
	m := New(3, 3)
	m.
		SetRow(0, []*rational.Rational{rational.ParseInt(1), rational.ParseInt(2), rational.ParseInt(3)}).
		SetRow(1, []*rational.Rational{rational.ParseInt(4), rational.ParseInt(5), rational.ParseInt(6)}).
		SetRow(2, []*rational.Rational{rational.ParseInt(7), rational.ParseInt(8), rational.ParseInt(9)})
	want := New(3, 3)
	want.
		SetRow(0, []*rational.Rational{rational.ParseInt(3), rational.ParseInt(6), rational.ParseInt(9)}).
		SetRow(1, []*rational.Rational{rational.ParseInt(12), rational.ParseInt(15), rational.ParseInt(18)}).
		SetRow(2, []*rational.Rational{rational.ParseInt(21), rational.ParseInt(24), rational.ParseInt(27)})
	assert.Equal(want, m.Scalar(rational.ParseInt(3)))
}

func makeTestMatrix() *Matrix {
	a := New(5, 5).
		SetRow(0, []*rational.Rational{rational.ParseInt(77), rational.ParseInt(88), rational.ParseInt(99), rational.ParseInt(12), rational.ParseInt(42)}).
		SetRow(1, []*rational.Rational{rational.ParseInt(61), rational.ParseInt(47), rational.ParseInt(8), rational.ParseInt(19), rational.ParseInt(41)}).
		SetRow(2, []*rational.Rational{rational.ParseInt(1), rational.ParseInt(22), rational.ParseInt(13), rational.ParseInt(74), rational.ParseInt(55)}).
		SetRow(3, []*rational.Rational{rational.ParseInt(3), rational.ParseInt(17), rational.ParseInt(58), rational.ParseInt(3), rational.ParseInt(32)}).
		SetRow(4, []*rational.Rational{rational.ParseInt(91), rational.ParseInt(27), rational.ParseInt(49), rational.ParseInt(4), rational.ParseInt(65)})
	return a
}

func TestDeterminantClassic(t *testing.T) {
	assert := assert.New(t)
	a := makeTestMatrix()
	want := rational.ParseInt(-546499540)
	got, err := a.DeterminantClassic()
	assert.NoError(err)
	assert.Equal(want, got)
	b := New(2, 2).
		SetRow(0, []*rational.Rational{rational.ParseInt(4), rational.ParseInt(7)}).
		SetRow(1, []*rational.Rational{rational.ParseInt(3), rational.ParseInt(2)})
	got, err = b.DeterminantClassic()
	assert.NoError(err)
	want = rational.ParseInt(4*2 - 3*7)
	assert.Equal(want, got)
}

func TestTranspose(t *testing.T) {
	assert := assert.New(t)
	a := New(5, 3).
		SetRow(0, []*rational.Rational{rational.ParseInt(77), rational.ParseInt(88), rational.ParseInt(99)}).
		SetRow(1, []*rational.Rational{rational.ParseInt(61), rational.ParseInt(47), rational.ParseInt(8)}).
		SetRow(2, []*rational.Rational{rational.ParseInt(1), rational.ParseInt(22), rational.ParseInt(13)}).
		SetRow(3, []*rational.Rational{rational.ParseInt(3), rational.ParseInt(17), rational.ParseInt(58)}).
		SetRow(4, []*rational.Rational{rational.ParseInt(91), rational.ParseInt(27), rational.ParseInt(49)})
	want := New(3, 5).
		SetRow(0, []*rational.Rational{rational.ParseInt(77), rational.ParseInt(61), rational.ParseInt(1), rational.ParseInt(3), rational.ParseInt(91)}).
		SetRow(1, []*rational.Rational{rational.ParseInt(88), rational.ParseInt(47), rational.ParseInt(22), rational.ParseInt(17), rational.ParseInt(27)}).
		SetRow(2, []*rational.Rational{rational.ParseInt(99), rational.ParseInt(8), rational.ParseInt(13), rational.ParseInt(58), rational.ParseInt(49)})
	got := a.Transpose()
	assert.Equal(want, got)
}

func TestDeterminant(t *testing.T) {
	assert := assert.New(t)
	a := makeTestMatrix()
	got, err := a.Determinant()
	assert.NoError(err)
	want := rational.ParseInt(-546499540)
	assert.NoError(err)
	assert.Equal(want, got)
}

func BenchmarkDeterminant(b *testing.B) {
	a := makeTestMatrix()
	for i := 0; i < b.N; i++ {
		if _, err := a.Determinant(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDeterminantClassic(b *testing.B) {
	a := makeTestMatrix()
	for i := 0; i < b.N; i++ {
		if _, err := a.DeterminantClassic(); err != nil {
			b.Fatal(err)
		}
	}
}
