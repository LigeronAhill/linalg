package matrix

import (
	"errors"
	"fmt"
	"math"

	"github.com/LigeronAhill/linalg/internal/rational"
)

type Number interface {
	int | float64
}

type Matrix [][]*rational.Rational

func (m *Matrix) Get(row, col int) (*rational.Rational, error) {
	if m.Rows() < row || m.Cols() < col {
		return nil, errors.New("wrong arguments")
	} else {
		return (*m)[row][col], nil
	}
}

func (m *Matrix) String() string {
	var s string
	for _, row := range *m {
		s += fmt.Sprintf("%v\n", row)
	}
	return s
}

func (m *Matrix) GetRow(row int) ([]*rational.Rational, error) {
	if m.Rows() < row {
		return nil, errors.New("wrong arguments")
	} else {
		return (*m)[row], nil
	}
}

func (m *Matrix) Set(row, col int, newValue *rational.Rational) error {
	if m.Rows() < row || m.Cols() < col {
		return errors.New("wrong arguments")
	} else {
		(*m)[row][col] = newValue
		return nil
	}
}

func New(rows, cols int) *Matrix {
	var m Matrix = make([][]*rational.Rational, rows)
	for i := range m {
		m[i] = make([]*rational.Rational, cols)
		for j := range m[i] {
			m[i][j] = rational.ParseInt(0)
		}
	}
	return &m
}

func (m *Matrix) Rows() int {
	return len(*m)
}

func (m *Matrix) Cols() int {
	return len((*m)[0])
}

func (m *Matrix) SetRow(index int, row []*rational.Rational) *Matrix {
	copy((*m)[index], row)
	return m
}

func (a *Matrix) Product(b *Matrix) (*Matrix, error) {
	if a.Rows() != b.Cols() {
		return nil, errors.New("wrong matrices sizes")
	}
	n := b.Rows()
	c := New(a.Rows(), b.Cols())
	for i := 0; i < a.Rows(); i++ {
		for j := 0; j < b.Cols(); j++ {
			for l := 0; l < n; l++ {
				p := (*a)[i][l].Multiply((*b)[l][j])
				(*c)[i][j] = (*c)[i][j].Add(p)
			}
		}
	}
	return c, nil
}

func (a *Matrix) Add(b *Matrix) (*Matrix, error) {
	if a.Rows() != b.Rows() || a.Cols() != b.Cols() {
		return nil, errors.New("wrong matrices sizes")
	}
	c := New(a.Rows(), a.Cols())
	for i, row := range *a {
		for j, element := range row {
			bVal, _ := b.Get(i, j)
			err := c.Set(i, j, element.Add(bVal))
			if err != nil {
				return nil, err
			}
		}
	}
	return c, nil
}

func (a *Matrix) Scalar(p *rational.Rational) *Matrix {
	result := New(a.Rows(), a.Cols())
	for i, row := range *a {
		for j, element := range row {
			result.Set(i, j, element.Multiply(p))
		}
	}
	return result
}

func (m *Matrix) DeterminantClassic() (*rational.Rational, error) {
	if m.Rows() != m.Cols() {
		return nil, errors.New("matrix must be n x n")
	}
	if m.Rows() == 2 {
		a, _ := m.Get(0, 0)
		b, _ := m.Get(1, 1)
		c, _ := m.Get(0, 1)
		d, _ := m.Get(1, 0)
		a = a.Multiply(b)
		b = c.Multiply(d)
		return a.Sub(b), nil
	} else {
		d := rational.ParseInt(0)
		for col, element := range (*m)[0] {
			minorMatrix := m.minor(col)
			signVal := math.Pow(-1.0, float64(col))
			sign := rational.ParseInt(1)
			if signVal < 0 {
				sign = rational.ParseInt(-1)
			}
			minorDet, err := minorMatrix.DeterminantClassic()
			if err != nil {
				return nil, err
			}
			td := element.Multiply(minorDet)
			td = td.Multiply(sign)
			d = d.Add(td)
		}
		return d, nil
	}
}

func (m *Matrix) minor(col int) *Matrix {
	noFirstRowMatrix := New(m.Rows()-1, m.Cols())

	for i, row := range *m {
		if i != 0 {
			noFirstRowMatrix.SetRow(i-1, row)
		}
	}

	result := New(noFirstRowMatrix.Rows(), m.Cols()-1)

	for i, row := range *noFirstRowMatrix {
		var newRow []*rational.Rational
		for j, element := range row {
			if j != col {
				newRow = append(newRow, element)
			}
		}
		result.SetRow(i, newRow)
	}
	return result
}

func (m *Matrix) Transpose() *Matrix {
	result := New(m.Cols(), m.Rows())
	for j, row := range *m {
		for i, element := range row {
			result.Set(i, j, element)
		}
	}
	return result
}

func (m *Matrix) Determinant() (*rational.Rational, error) {
	if m.Rows() != m.Cols() {
		return nil, errors.New("matrix must be n x n")
	}
	result := m
	for j := 0; j < result.Cols(); j++ {
		for i := 1; i < result.Rows(); i++ {
			if j < i {
				element, _ := result.Get(i, j)
				if element == rational.ParseInt(0) {
					continue
				}
				d, _ := result.Get(j, j)
				m := element.Divide(d)
				mr, _ := result.GetRow(j) // 7
				cr, _ := result.GetRow(i) // 5
				for k, el := range cr {
					v := el.Sub(mr[k].Multiply(m))
					if k == j && v.N() != 0 {
						fmt.Printf("ELEMENT: %s\n", el.String())
						fmt.Printf("DIVIDER: %s\n", d.String())
						fmt.Printf("NEW ELEMENT: %s\n", v.String())
						fmt.Printf("MODIFIER: %s\n", m.String())
						fmt.Printf("MODIFYING ELEMENT: %s\n", mr[k].String())
						fmt.Printf("PRODUCTION: %s\n", mr[k].Multiply(m).String())
					}
					cr[k] = v
				}
				result.SetRow(i, cr)
			}
		}
	}
	for _, row := range *result {
		fmt.Println(row)
	}
	d := rational.ParseInt(1)
	for i, row := range *result {
		for j, element := range row {
			if i == j {
				d = d.Multiply(element)
			}
		}
	}
	return d, nil
}

func (m *Matrix) multyRow(i int, modifier *rational.Rational) []*rational.Rational {
	var result []*rational.Rational
	current := (*m)[i]
	for _, c := range current {
		result = append(result, c.Multiply(modifier))
	}
	return result
}
