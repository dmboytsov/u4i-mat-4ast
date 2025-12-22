package matrix

import "testing"

// Дана матрица, элемент матрицы - вес, сколько стоит пройти через эту клетку
// Нужно найти вес/стоимость самого дешевого пути.
// Путь из левого верхнего угла до правого нижнего угла
// Двигаться можно направо и вниз.
//
// Решение.
// Делаем еще одну матрицу, в которой каждый элемент будет минимальая  `цена` (сумму весов)
// сколько стоит до него дойти.
// За проход заполним всю матрицу
// Для первой строчки `цена` будет левый сосед + вес текущей ячейки
// Для первой строки `цена` будет верхний сосед + вес текущей ячейки
// Для остальных нужно выбрать минимальный из верхнего и левого и добавить свой вес.

func CheapWay(matrix [][]int) int {
	// матрица в которой будем хранить самый дешевый путь до i-ой j-ой клетки
	wp := make([][]int, 0, len(matrix))
	for i, line := range matrix {
		wpline := make([]int, len(line))
		wp = append(wp, wpline)
		for j, weight := range line {
			if i == 0 && j == 0 { // первая клетка, нету соседей
				wp[i][j] = weight
				continue
			}

			if i == 0 { // первая строка, но уже не первый элемент
				wp[i][j] = weight + wp[i][j-1]
				continue
			}

			if j == 0 { // первый столбец, но уже не первая строка
				wp[i][j] = weight + wp[i-1][j]
				continue
			}

			top := wp[i-1][j]
			left := wp[i][j-1]
			if top < left {
				wp[i][j] = top + weight
			} else {
				wp[i][j] = left + weight
			}
		}
	}

	return wp[len(matrix)-1][len(matrix[len(matrix)-1])-1]
}

func TestCheapWay(t *testing.T) {
	type tcase struct {
		matrix [][]int
		want   int
	}

	tcases := []tcase{
		{
			matrix: [][]int{
				{1, 1, 1, 1, 1},
				{3, 100, 100, 100, 100},
				{1, 1, 1, 1, 1},
				{2, 2, 2, 2, 1},
				{1, 1, 1, 1, 1},
			},
			want: 11,
		},
	}

	for _, tc := range tcases {
		got := CheapWay(tc.matrix)
		if got != tc.want {
			t.Errorf("CheapWay(%v) = %d; want %d", tc.matrix, got, tc.want)
		}
	}
}
