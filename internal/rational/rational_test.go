package rational

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	intTest := 15
	floatTest := 4.75
	t1 := ParseInt(intTest)
	t2, err := ParseFloat(floatTest)
	if err != nil {
		t.Fatal(err)
	}
	want1 := &Rational{n: 15, d: 1}
	want2 := &Rational{
		n: 19,
		d: 4,
	}
	assert.Equal(t, want1, t1)
	assert.Equal(t, want2, t2)
}

func TestNew(t *testing.T) {
	got := New(16, 8)
	want := &Rational{
		n: 2,
		d: 1,
	}
	assert.Equal(t, want, got)
}

func TestParseInt(t *testing.T) {
	assert := assert.New(t)
	i := 42
	r := ParseInt(i)
	want := New(42, 1)
	assert.Equal(want, r)
}

func TestParseFloat(t *testing.T) {
	assert := assert.New(t)
	f := 4.63
	want := New(463, 100)
	got, err := ParseFloat(f)
	assert.NoError(err)
	assert.Equal(want, got)
}

func TestAdd(t *testing.T) {
	assert := assert.New(t)
	tt := []struct {
		a    *Rational
		b    *Rational
		want *Rational
	}{
		{
			a:    New(46, 8),
			b:    New(2, 1),
			want: New(62, 8),
		},
		{
			a:    New(7*43+15, 43),
			b:    New(4*98+65, 98),
			want: New(12*4214+51, 4214),
		},
		{
			a:    New(69, 96),
			b:    New(5*87+55, 87),
			want: New(6*2784+977, 2784),
		},
	}
	for _, test := range tt {
		assert.Equal(test.want, test.a.Add(test.b))
	}
}

func TestSub(t *testing.T) {
	assert := assert.New(t)
	tt := []struct {
		a    *Rational
		b    *Rational
		want *Rational
	}{
		{
			a:    New(5*8+6, 8),
			b:    New(2, 1),
			want: New(3*8+6, 8),
		},
		{
			a:    New(7*43+15, 43),
			b:    New(4*98+65, 98),
			want: New(2*4214+2889, 4214),
		},
		{
			a:    New(10*96+69, 96),
			b:    New(5*87+55, 87),
			want: New(5*2784+241, 2784),
		},
	}
	for _, test := range tt {
		assert.Equal(test.want, test.a.Sub(test.b))
	}
}

func TestMultiply(t *testing.T) {
	assert := assert.New(t)
	tt := []struct {
		a    *Rational
		b    *Rational
		want *Rational
	}{
		{
			a:    New(5*8+6, 8),
			b:    New(-2, 1),
			want: New(-11*2-1, 2),
		},
		{
			a:    New(7*43+15, 43),
			b:    New(4*98+65, 98),
			want: New(34*2107+568, 2107),
		},
		{
			a:    New(10*96+69, 96),
			b:    New(5*87+55, 87),
			want: New(60*1392+515, 1392),
		},
	}
	for _, test := range tt {
		assert.Equal(test.want, test.a.Multiply(test.b))
	}
}

func TestDivide(t *testing.T) {
	assert := assert.New(t)
	tt := []struct {
		a    *Rational
		b    *Rational
		want *Rational
	}{
		{
			a:    New(5*8+6, 8),
			b:    New(-2, 1),
			want: New(-2*8-7, 8),
		},
		{
			a:    New(7*43+15, 43),
			b:    New(4*98+65, 98),
			want: New(19651+11317, 19651),
		},
		{
			a:    New(10*96+69, 96),
			b:    New(5*87+55, 87),
			want: New(320+289, 320),
		},
	}
	for _, test := range tt {
		assert.Equal(test.want, test.a.Divide(test.b))
	}
}
