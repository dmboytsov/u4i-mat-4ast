package main

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
// С крайних точек вода может стекать в океан. То есть как минимум для каждой крайней точки достежим хотябы один океан.
// Первой строчке достежим атлантический океан
// Последней строчке достежим тихий
// Первый элементам в строке достежим атлантический
// Поселдним элементам достежим тихий
// Нижний левый элемент, с него доступен  и тихий и атлантический, поэтому с него начинам обход (можно и в верхнего правого)
// Идем налево и в верх, (не ходим по кругу, чтобы не зациклиться)
// На кажом шаге смотрим на соседей сравниваем ввысоты и копируем доступность океанов если элемент выше (по высоте),
// значит ему доступно все то что для нижнего.

import (
	"fmt"
	"reflect"
	"testing"
)

type Point struct {
	r, c int // indexes
	p, a int // 0 - dо not visited, 1 - achieve ocean, -1 - do not achieve, p - pacific, a - atlantic
	h    int //heights
}

func waterFlow(heights [][]int) [][]int {
	if len(heights) == 0 {
		return nil
	}

	points := heightsToPoints(heights)
	//printPoints(points) //debug
	// обход матрицы, идем снизу слева
	dfs(len(points)-1, 0, points)
	//printPoints(points) //debug
	return getWaterFlow(points)
}

// Метод для отдалки
func printPoints(points [][]Point) {
	for _, row := range points {
		for _, point := range row {
			fmt.Printf("[a:%d, p:%d, h:%d] ", point.a, point.p, point.h)
		}
		fmt.Println()
	}
}

// Метод обхода матрицы
// с нижне левого угла нужно идти, так как с него всегда жоступны два океана
// используем направеление в право и вверх (влево и вниз не ходим чтобы не зациклиться)
// ?по часовой стрелке, видимо напрлавлене нужно передавать, если уперлись в точку которая уже пройдена,
// то меняем направление
func dfs(r, c int, points [][]Point) {
	curP := &points[r][c]
	var rightP, leftP, topP, bottomP *Point

	// смотрим точки в округе, если высота меньше всех точек в округе, то нельзя ничего из этой точки достичь
	// left // сначала смотрим на левую точку, так как мы идем слева на право и левая точка уже вычеслина на пред шаге
	if c == 0 { // это крайняя левая точка и с нее достигаем атлантический океан
		curP.a = 1
	} else {
		leftP = &points[r][c-1]
		swapOceans(curP, leftP)
	}

	// bottom
	if r == len(points)-1 {
		curP.p = 1 // нижняя точка достежим тихий океан
	} else {
		bottomP = &points[r+1][c]
		swapOceans(curP, bottomP)
	}

	// right
	if c == len(points[r])-1 { // это крайняя правая точка и с нее достигаем тихий океан
		curP.p = 1
	} else {
		rightP = &points[r][c+1]
		swapOceans(curP, rightP)
	}
	// top
	if r == 0 {
		curP.a = 1 // верхняя точка достежим атлантисеский океан
	} else {
		topP = &points[r-1][c]
		swapOceans(curP, topP)
	}

	// Если ниже всех, то не достежим
	if rightP != nil && leftP != nil && topP != nil && bottomP != nil {
		if curP.h < rightP.h && curP.h < leftP.h && curP.h < topP.h && curP.h < bottomP.h {
			curP.a = -1
			curP.p = -1
		}
	}

	// идем с лева на право
	if rightP != nil {
		dfs(rightP.r, rightP.c, points)
	}
	// снизу вверх
	if topP != nil {
		dfs(topP.r, topP.c, points)
	}

	// обошли соседей, и из никого не получилось достичь океанов ставим -1
	if curP.p == 0 {
		curP.p = -1
	}
	if curP.a == 0 {
		curP.a = -1
	}
}

// копируем доступность океанов с болле низкой
func swapOceans(aPoint, bPoint *Point) {
	if bPoint.h >= aPoint.h { // b выше  a, значит все что доступно a. то и доступно b
		if aPoint.p == 1 {
			bPoint.p = 1
		}
		if aPoint.a == 1 {
			bPoint.a = 1
		}
	}
	if aPoint.h >= bPoint.h { // a выше  b, значит все что доступно b, то и доступно a
		if bPoint.p == 1 {
			aPoint.p = 1
		}
		if bPoint.a == 1 {
			aPoint.a = 1
		}
	}
}

// Матрицу высот переобразуем в матрицу структур Point. Крайним вершинам проставляем доступность океана
func heightsToPoints(heights [][]int) [][]Point {
	res := make([][]Point, len(heights))
	for i, row := range heights {
		points := make([]Point, len(row))
		for j, height := range row {
			p := Point{
				h: height,
				r: i,
				c: j,
				p: 0, // do not visit
				a: 0, // do not visit
			}
			if (i == len(heights)-1) || (j == len(row)-1) {
				p.p = 1 // низ и правая сторона могут стекать в тихий океан
			}
			if (i == 0) || (j == 0) {
				p.a = 1 // верх и левая сторона могут стекать в атлантический океан
			}
			points[j] = p
		}
		res[i] = points
	}
	return res
}

// выбирает точки с коткорых доступны тихий и атлантический океан
func getWaterFlow(points [][]Point) [][]int {
	res := make([][]int, 0, len(points))
	for i, row := range points {
		for j, point := range row {
			if point.a == 1 && point.p == 1 {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
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
		//{
		//	heights: [][]int{
		//		{1, 3, 1},
		//		{2, 3, 2},
		//		{1, 3, 1},
		//	},
		//	expect: [][]int{
		//		{0, 1}, {0, 2}, {1, 1}, {2, 0}, {2, 1},
		//	},
		//},
		//{
		//	heights: [][]int{
		//		{1, 2, 3, 2, 1},
		//		{1, 2, 3, 2, 1},
		//		{1, 2, 3, 2, 1},
		//	},
		//	expect: [][]int{
		//		{0, 2}, {1, 2}, {2, 0}, {2, 1}, {2, 2},
		//	},
		//},
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
	}

	for _, c := range cases {
		got := waterFlow(c.heights)
		if !reflect.DeepEqual(got, c.expect) {
			t.Errorf("case %v, got %d, want %d", c, got, c.expect)
		}
	}
}
