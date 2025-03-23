package rational

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Rational struct {
	n int
	d int
}

func New(n, d int) *Rational {
	if d == 0 {
		d = 1
	}
	if (n < 0 && d < 0) || d < 0 {
		n = -n
		d = -d
	}
	gcd := gcd(n, d)
	n /= gcd
	d /= gcd
	return &Rational{
		n,
		d,
	}
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func ParseInt(i int) *Rational {
	return New(i, 1)
}

func ParseFloat(f float64) (*Rational, error) {
	s := fmt.Sprintf("%f", f)
	sl := strings.Split(s, ".")
	if len(sl) != 2 {
		return nil, errors.New("wrong input")
	}
	iStr := sl[0]
	rStr := sl[1]
	i, err := strconv.Atoi(iStr)
	if err != nil {
		return nil, err
	}
	pow := len(rStr)
	d := int(math.Pow10(pow))
	n64, err := strconv.ParseInt(rStr, 10, 0)
	if err != nil {
		return nil, err
	}
	n := int(n64) + i*d
	if d == 0 {
		d = 1
	}
	return New(n, d), nil
}

func (a *Rational) N() int {
	return a.n
}

func (a *Rational) D() int {
	return a.d
}

func (a *Rational) Add(b *Rational) *Rational {
	commonDenominator := a.d * b.d
	n := a.n*b.d + b.n*a.d
	return New(n, commonDenominator)
}

func (a *Rational) Sub(b *Rational) *Rational {
	commonDenominator := a.d * b.d
	n := a.n*b.d - b.n*a.d
	return New(n, commonDenominator)
}

func (a *Rational) Multiply(b *Rational) *Rational {
	return New(a.n*b.n, a.d*b.d)
}

func (a *Rational) Divide(b *Rational) *Rational {
	return New(a.n*b.d, a.d*b.n)
}

func (r *Rational) String() string {
	var s string
	i := r.n / r.d
	n := r.n % r.d
	if i != 0 {
		s += strconv.FormatInt(int64(i), 10)
	}
	if n != 0 {
		if i != 0 {
			s += " + "
		}
		if r.d < 0 {
			n *= -1
			r.d *= -1
		}
		s += strconv.FormatInt(int64(n), 10) + "/" + strconv.FormatInt(int64(r.d), 10)
	}
	if len(s) == 0 {
		s += "0"
	}
	return s
}
