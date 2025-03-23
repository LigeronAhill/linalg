package rational

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Rational struct {
	i int
	n int
	d int
}

func New(i, n, d int) *Rational {
	if d != 0 {
		gcd := gcd(n, d)
		n /= gcd
		d /= gcd
		for n >= d {
			n -= d
			i += 1
		}
		if n == 0 {
			d = 0
		}
	}
	return &Rational{
		i,
		n,
		d,
	}
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func ParseInt(i int) *Rational {
	return &Rational{
		i: i,
	}
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
	n := int(n64)
	return New(i, n, d), nil
}

func (a *Rational) Add(b *Rational) *Rational {
	commonDenominator := 1
	if a.d != 0 {
		commonDenominator *= a.d
	}
	if b.d != 0 {
		commonDenominator *= b.d
	}
	ai := a.i * commonDenominator
	bi := b.i * commonDenominator
	an := a.n
	if b.d != 0 {
		an *= b.d
	}
	bn := b.n
	if a.d != 0 {
		bn *= a.d
	}
	n := ai + an + bi + bn
	return New(0, n, commonDenominator)
}

func (a *Rational) Sub(b *Rational) *Rational {
	commonDenominator := 1
	if a.d != 0 {
		commonDenominator *= a.d
	}
	if b.d != 0 {
		commonDenominator *= b.d
	}
	ai := a.i * commonDenominator
	bi := b.i * commonDenominator
	an := a.n
	if b.d != 0 {
		an *= b.d
	}
	bn := b.n
	if a.d != 0 {
		bn *= a.d
	}
	n := ai + an - bi - bn
	return New(0, n, commonDenominator)
}

func (a *Rational) Multiply(b *Rational) *Rational {
	an := a.i
	ad := 1
	if a.d != 0 {
		an = a.d*a.i + a.n
		ad *= a.d
	}
	bn := b.i
	bd := 1
	if b.d != 0 {
		bn = b.d*b.i + b.n
		bd *= b.d
	}
	n := an * bn
	d := ad * bd
	return New(0, n, d)
}

func (a *Rational) Divide(b *Rational) *Rational {
	an := a.i
	ad := 1
	if a.d != 0 {
		an = a.d*a.i + a.n
		ad *= a.d
	}
	bn := b.i
	bd := 1
	if b.d != 0 {
		bn = b.d*b.i + b.n
		bd *= b.d
	}
	n := an * bd
	d := ad * bn
	return New(0, n, d)
}

func (r *Rational) String() string {
	var s string
	if r.i != 0 {
		s += strconv.FormatInt(int64(r.i), 10)
	}
	if r.n != 0 {
		if r.i != 0 {
			s += " + "
		}
		if r.d < 0 {
			r.n *= -1
			r.d *= -1
		}
		s += strconv.FormatInt(int64(r.n), 10) + "/" + strconv.FormatInt(int64(r.d), 10)
	}
	if len(s) == 0 {
		s += "0"
	}
	return s
}
