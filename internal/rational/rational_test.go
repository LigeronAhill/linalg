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
	want1 := &Rational{i: 15}
	want2 := &Rational{
		i: 4,
		n: 3,
		d: 4,
	}
	assert.Equal(t, want1, t1)
	assert.Equal(t, want2, t2)
}

func TestNew(t *testing.T) {
	got := New(0, 16, 8)
	want := &Rational{
		i: 2,
		n: 0,
		d: 0,
	}
	assert.Equal(t, want, got)
}

func TestParseInt(t *testing.T) {
	assert := assert.New(t)
	i := 42
	r := ParseInt(i)
	want := New(i, 0, 0)
	assert.Equal(want, r)
}

func TestParseFloat(t *testing.T) {
	assert := assert.New(t)
	f := 4.63
	want := New(4, 63, 100)
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
			a:    New(5, 6, 8),
			b:    New(2, 0, 0),
			want: New(7, 6, 8),
		},
		{
			a:    New(7, 15, 43),
			b:    New(4, 65, 98),
			want: New(12, 51, 4214),
		},
		{
			a:    New(0, 69, 96),
			b:    New(5, 55, 87),
			want: New(6, 977, 2784),
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
			a:    New(5, 6, 8),
			b:    New(2, 0, 0),
			want: New(3, 6, 8),
		},
		{
			a:    New(7, 15, 43),
			b:    New(4, 65, 98),
			want: New(2, 2889, 4214),
		},
		{
			a:    New(10, 69, 96),
			b:    New(5, 55, 87),
			want: New(5, 241, 2784),
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
			a:    New(5, 6, 8),
			b:    New(2, 0, 0),
			want: New(11, 1, 2),
		},
		{
			a:    New(7, 15, 43),
			b:    New(4, 65, 98),
			want: New(34, 568, 2107),
		},
		{
			a:    New(10, 69, 96),
			b:    New(5, 55, 87),
			want: New(60, 515, 1392),
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
			a:    New(5, 6, 8),
			b:    New(2, 0, 0),
			want: New(2, 7, 8),
		},
		{
			a:    New(7, 15, 43),
			b:    New(4, 65, 98),
			want: New(1, 11317, 19651),
		},
		{
			a:    New(10, 69, 96),
			b:    New(5, 55, 87),
			want: New(1, 289, 320),
		},
	}
	for _, test := range tt {
		assert.Equal(test.want, test.a.Divide(test.b))
	}
}
