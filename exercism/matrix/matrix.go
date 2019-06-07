package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func (m *Matrix) Rows() [][]int {
	n := make([][]int, len(*m))
	for idx1, row := range *m {
		n[idx1] = make([]int, len(row))
		for idx2, val := range row {
			n[idx1][idx2] = val
		}
	}
	return n
}

func (m *Matrix) Cols() [][]int {
	var colNum int
	if len(*m) > 0 {
		colNum = len((*m)[0])
	}
	n := make([][]int, colNum)
	for _, row := range *m {
		for idx, col := range row {
			n[idx] = append(n[idx], col)
		}
	}
	return n
}

func (m *Matrix) Set(r, c, val int) bool {
	if r < 0 || c < 0 || r >= len(*m) || (len(*m) > 0 && c >= len((*m)[0])) {
		return false
	}
	(*m)[r][c] = val
	return true // I.e. ok
}

func New(s string) (*Matrix, error) {
	var m Matrix
	for _, line := range strings.Split(s, "\n") {
		fields := strings.Fields(line)
		if len(m) > 0 && len(fields) != len(m[len(m)-1]) {
			return nil, errors.New("rows of unequal length")
		}
		row := make([]int, len(fields))
		for idx, field := range fields {
			val, err := strconv.Atoi(field)
			if err != nil {
				return nil, errors.New("invalid chunk component")
			}
			row[idx] = val
		}
		m = append(m, row)
	}
	return &m, nil
}
