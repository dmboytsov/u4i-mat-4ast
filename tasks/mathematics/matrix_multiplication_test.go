package mathematics

import (
	"fmt"
	"reflect"
	"testing"
)

/**
 * Функция для умножения двух матриц.
 *
 * Строка первой матицы умножается на столбей второй матрицы.
 * В результирующей матрице будет строк как в первой а столбцов как во второй
 * Матрицы можно перемножать, если кол-во столбцов первой матрицы равно кол-ву строк во второй
 */
func MatrixMultiplication(a [][]int, b [][]int) ([][]int, error) {
	if len(a) == 0 || len(b) == 0 {
		return nil, fmt.Errorf("some matrix is empty")
	}

	// Матрицы можно перемножать, если кол-во столбцов первой матрицы равно кол-ву строк во второй
	// m - будем проверять, что в каждой строке одинаковое кол-во элементов, и кол-во строк в b
	m := len(a[0])
	if len(b) != m {
		return nil, fmt.Errorf("sizes do not match")
	}

	res := make([][]int, 0, m)
	for i := range a { // идем по строкам a
		res = append(res, make([]int, len(b[0])))

		for j := 0; j < len(b[0]); j++ { //по столбцам b
			c := 0
			for ai := range a[i] {
				c += a[i][ai] * b[ai][j]
			}
			res[i][j] = c
		}
	}

	return res, nil
}

func TestMatrixMultiplication(t *testing.T) {
	type tcase struct {
		a    [][]int
		b    [][]int
		want [][]int
	}

	cases := []tcase{
		{
			a: [][]int{
				{1, 2},
				{3, 4}},
			b: [][]int{
				{5, 6},
				{7, 8}},
			want: [][]int{
				{19, 22},
				{43, 50}},
		}, {
			a: [][]int{
				{1, 2},
				{2, 1}},
			b: [][]int{
				{3, 3, 3},
				{4, 4, 4}},
			want: [][]int{
				{11, 11, 11},
				{10, 10, 10}},
		}, {
			a: [][]int{
				{1, 2},
				{2, 1}},
			b: [][]int{
				{3, 3, 3},
				{4, 4, 4}},
			want: [][]int{
				{11, 11, 11},
				{10, 10, 10}},
		}, {
			a: [][]int{{3, 4}},
			b: [][]int{
				{9},
				{8}},
			want: [][]int{
				{59}},
		}, {
			a: [][]int{
				{0},
				{2}},
			b: [][]int{
				{1, 2, 8}},
			want: [][]int{
				{0, 0, 0},
				{2, 4, 16}},
		}, {
			a: [][]int{
				{6},
				{1},
				{3},
				{3},
				{1},
				{9}},
			b: [][]int{
				{10, 2, 0, 3}},
			want: [][]int{
				{60, 12, 0, 18},
				{10, 2, 0, 3},
				{30, 6, 0, 9},
				{30, 6, 0, 9},
				{10, 2, 0, 3},
				{90, 18, 0, 27}},
		},
	}

	for _, tc := range cases {
		got, _ := MatrixMultiplication(tc.a, tc.b)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got %v; want %v", got, tc.want)
			return
		}
	}
}
