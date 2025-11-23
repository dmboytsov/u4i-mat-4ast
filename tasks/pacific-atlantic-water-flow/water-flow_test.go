package main

import (
	"reflect"
	"testing"
)

// https://leetcode.com/problems/pacific-atlantic-water-flow/description/
//
// Данна матрица, где значения это высоты.
// Слева и сверху аталнтический океан, снизу и справа тихий, нужно вычислить водораздел.
// Водораздел это элементы матрицы, с которых вода может стечь и в атлантический и в тихий океан.
//
//  1 1 2
//  1 2 1
//  1 2 1  => [ [0,2], [1,1] [2,1], [2,0] ]
//
// Решение.
// Делаем две матрицы одну для тихого океана, вторую для атлантического
// Матрицы булевые, TRUE - значит что с этой точки достежим океан.
// Матрицы проинициализированны фалсами
//
// Обходим матрицы для кажого океана.
// Для обода тихого океана двигаемся по первому столбцу, и по первой строке,
// в качестве prevHeight - передаем высоту точку, в итоге в dfs получается,
// что сравниваем высоту саму с собой и проставляем TRUE.
// (первый столбей и первая строка достигают тихий океан так как с ним граничат)
// Аналагично для атлантического, но для старта dfs последняя строка и последний столбец
//
// Когда собрали две матрицы выбираем точки коткорые TRUE b и для тихого и атлантического океана.

func waterFlow(heights [][]int) [][]int {
	if len(heights) == 0 {
		return nil
	}

	// init  oceans
	rows, cols := len(heights), len(heights[0])
	pacific := make([][]bool, rows)
	atlantic := make([][]bool, rows)

	for i := range pacific {
		pacific[i] = make([]bool, cols)
		atlantic[i] = make([]bool, cols)
	}

	var dfs func(int, int, [][]bool, int)
	//  заполняет visited - true значит что океан доступен
	dfs = func(row, col int, visited [][]bool, prevHeight int) {
		if row < 0 || col < 0 || row >= rows || col >= cols {
			// проверка выход ща границы матрицы
			return
		}
		if visited[row][col] {
			// уже вычеслили что точка есть в водоразделе
			return
		}
		if heights[row][col] < prevHeight {
			// высота ниже, предыдущего.
			return
		}
		visited[row][col] = true
		// Высота больше соседа, но значит ли это что океан достежим?
		// да, так как проваливаясь по рекурсии в глубь (dfs)
		// мы шли, по большим высотам. То есть, высота больше пред шага,
		//а у пред шага было выше пред пред, и в итоге получается что путь есть.

		// пошли по соседям
		dfs(row+1, col, visited, heights[row][col])
		dfs(row-1, col, visited, heights[row][col])
		dfs(row, col+1, visited, heights[row][col])
		dfs(row, col-1, visited, heights[row][col])
	}

	for i := 0; i < rows; i++ {
		dfs(i, 0, pacific, heights[i][0])            // первый столбец
		dfs(i, cols-1, atlantic, heights[i][cols-1]) // последний столбец
	}

	for j := 0; j < cols; j++ {
		dfs(0, j, pacific, heights[0][j])            // первая строка
		dfs(rows-1, j, atlantic, heights[rows-1][j]) // последняя строка
	}

	result := make([][]int, 0, len(atlantic))
	for i, row := range atlantic {
		for j, fl := range row {
			if pacific[i][j] && fl {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

// Запуск и проверка кейсов
func TestWaterFlow(t *testing.T) {
	type tcase struct {
		heights [][]int
		expect  [][]int
	}
	cases := []tcase{
		{
			heights: [][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			},
			expect: [][]int{
				{0, 2}, {1, 1}, {2, 0},
			},
		},
		{
			heights: [][]int{
				{1, 3, 1},
				{2, 3, 2},
				{1, 3, 1},
			},
			expect: [][]int{
				{0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1},
			},
		},
		{ // со всех точек достежимы и тихий и атлантический океан, так как значения в столбике одинаковые вода может литься как вверх так и вниз
			heights: [][]int{
				{1, 2, 3, 2, 1},
				{1, 2, 3, 2, 1},
				{1, 2, 3, 2, 1},
			},
			expect: [][]int{
				{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4},
				{1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4},
				{2, 0}, {2, 1}, {2, 2}, {2, 3}, {2, 4},
			},
		},
		// из примера - https://algo.monster/liteproblems/pacific-atlantic-water-flow
		{
			heights: [][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			expect: [][]int{
				{0, 4},
				{1, 3}, {1, 4},
				{2, 2},
				{3, 0}, {3, 1},
				{4, 0},
			},
		},
		{
			heights: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
			expect: [][]int{
				{0, 2},
				{1, 0}, {1, 1}, {1, 2},
				{2, 0}, {2, 1}, {2, 2}, // 2,1 потому что с 6 на 5,  а с 5 на 4, с 4 на 3.
			},
		},
		{
			heights: [][]int{
				{5, 5, 5, 5},
				{4, 4, 4, 4},
				{5, 5, 5, 5},
			},
			expect: [][]int{
				{0, 0}, {0, 1}, {0, 2}, {0, 3},
				{1, 0}, {1, 1}, {1, 2}, {1, 3},
				{2, 0}, {2, 1}, {2, 2}, {2, 3},
			},
		},
	}

	for _, c := range cases {
		got := waterFlow(c.heights)
		if !reflect.DeepEqual(got, c.expect) {
			t.Errorf("case %v, got %d, want %d", c, got, c.expect)
		}
	}
}
