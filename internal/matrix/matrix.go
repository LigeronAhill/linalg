package matrix

import (
	"errors"
	"math"
)

type Number interface {
	int | float64
}

type Matrix[T Number] [][]T

func (m *Matrix[T]) Get(row, col int) (T, error) {
	if m.Rows() < row || m.Cols() < col {
		return 0, errors.New("wrong arguments")
	} else {
		return (*m)[row][col], nil
	}
}

func (m *Matrix[T]) Set(row, col int, newValue T) error {
	if m.Rows() < row || m.Cols() < col {
		return errors.New("wrong arguments")
	} else {
		(*m)[row][col] = newValue
		return nil
	}
}

func New[T Number](rows, cols int) *Matrix[T] {
	var m Matrix[T] = make([][]T, rows)
	for i := range m {
		m[i] = make([]T, cols)
	}
	return &m
}

func (m *Matrix[T]) Rows() int {
	return len(*m)
}

func (m *Matrix[T]) Cols() int {
	return len((*m)[0])
}

func (m *Matrix[T]) SetRow(index int, row []T) *Matrix[T] {
	copy((*m)[index], row)
	return m
}

func (a *Matrix[T]) Product(b *Matrix[T]) (*Matrix[T], error) {
	if a.Rows() != b.Cols() {
		return nil, errors.New("wrong matrices sizes")
	}
	n := b.Rows()
	c := New[T](a.Rows(), b.Cols())
	for i := 0; i < a.Rows(); i++ {
		for j := 0; j < b.Cols(); j++ {
			for l := 0; l < n; l++ {
				aVal, err := a.Get(i, l)
				if err != nil {
					return nil, err
				}
				bVal, err := b.Get(l, j)
				if err != nil {
					return nil, err
				}
				cVal, err := c.Get(i, j)
				if err != nil {
					return nil, err
				}
				v := cVal + aVal*bVal
				err = c.Set(i, j, v)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return c, nil
}

func (a *Matrix[T]) Add(b *Matrix[T]) (*Matrix[T], error) {
	if a.Rows() != b.Rows() || a.Cols() != b.Cols() {
		return nil, errors.New("wrong matrices sizes")
	}
	c := New[T](a.Rows(), a.Cols())
	for i, row := range *a {
		for j, element := range row {
			bVal, _ := b.Get(i, j)
			err := c.Set(i, j, element+bVal)
			if err != nil {
				return nil, err
			}
		}
	}
	return c, nil
}

func (a *Matrix[T]) Scalar(p T) *Matrix[T] {
	result := New[T](a.Rows(), a.Cols())
	for i, row := range *a {
		for j, element := range row {
			result.Set(i, j, element*p)
		}
	}
	return result
}

func (m *Matrix[T]) DeterminantClassic() (T, error) {
	if m.Rows() != m.Cols() {
		return -1, errors.New("matrix must be n x n")
	}
	if m.Rows() == 2 {
		a, _ := m.Get(0, 0)
		b, _ := m.Get(1, 1)
		c, _ := m.Get(0, 1)
		d, _ := m.Get(1, 0)
		return a*b - c*d, nil
	} else {
		var d T = 0
		for col, element := range (*m)[0] {
			minorMatrix := m.minor(col)
			sign := T(math.Pow(-1.0, float64(col)))
			minorDet, err := minorMatrix.DeterminantClassic()
			if err != nil {
				return -1, err
			}
			d += sign * element * minorDet
		}
		return d, nil
	}
}

func (m *Matrix[T]) minor(col int) *Matrix[T] {
	noFirstRowMatrix := New[T](m.Rows()-1, m.Cols())

	for i, row := range *m {
		if i != 0 {
			noFirstRowMatrix.SetRow(i-1, row)
		}
	}

	result := New[T](noFirstRowMatrix.Rows(), m.Cols()-1)

	for i, row := range *noFirstRowMatrix {
		var newRow []T
		for j, element := range row {
			if j != col {
				newRow = append(newRow, element)
			}
		}
		result.SetRow(i, newRow)
	}
	return result
}

func (m *Matrix[T]) Transpose() *Matrix[T] {
	result := New[T](m.Cols(), m.Rows())
	for j, row := range *m {
		for i, element := range row {
			result.Set(i, j, element)
		}
	}
	return result
}

func (m *Matrix[T]) Determinant() (T, error) {
	if m.Rows() != m.Cols() {
		return -1, errors.New("matrix must be n x n")
	}
	result := New[float64](m.Rows(), m.Cols())
	for i, row := range *m {
		for j, element := range row {
			result.Set(i, j, float64(element))
		}
	}
	for j := 0; j < result.Cols(); j++ {
		for i := 1; i < result.Rows(); i++ {
			if j < i {
				if element, _ := result.Get(i, j); element == 0.0 {
					continue
				}
				top, _ := result.Get(i, j)
				bottom, _ := result.Get(j, j)
				m := top / bottom
				cur := result.multyRow(i, 1.0)
				mod := result.multyRow(j, m)
				for i, el := range mod {
					mod[i] = cur[i] - el
				}
				result.SetRow(i, mod)
			}
		}
	}
	d := 1.0
	for i, row := range *result {
		for j, element := range row {
			if i == j {
				d *= element
			}
		}
	}
	e, _ := m.Get(0, 0)
	switch any(e).(type) {
	case int:
		return T(math.Round(d)), nil
	default:
		return T(d), nil
	}
}

func (m *Matrix[T]) multyRow(i int, modifier float64) []float64 {
	var result []float64
	current := (*m)[i]
	for _, c := range current {
		result = append(result, float64(c)*modifier)
	}
	return result
}
