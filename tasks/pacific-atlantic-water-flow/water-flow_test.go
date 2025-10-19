package main

import (
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
	// начинаем обход с середины и по спирали расскручиваем
	points[len(heights)-1][0].p = 1
	points[len(heights)-1][0].a = 1
	dfs(int(len(points)/2), int(len(points[0])/2), points)
	return getWaterFlow(points)
}

// с нижне левого угла нужно идти по часовой стрелке, видимо напрлавлене нужно передавать, если уперлись в точку которая уже пройдена,
// то меняем направление
func dfs(r, c int, points [][]Point) {
	curP := &points[r][c]
	var rightP, leftP, topP, bottomP *Point
	if curP.p != 0 && curP.a != 0 {
		// эту точку уже обошли
		return
	}
	// смотрим точки в округе, если высота меньше всех точек в округе то нельзя ничего из этой точки достичь
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
	// left
	if c == 0 { // это крайняя левая точка и с нее достигаем ытлантический океан
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

	// если ниже всех, то не достежим
	if rightP != nil && leftP != nil && topP != nil && bottomP != nil {
		if curP.h < rightP.h && curP.h < leftP.h && curP.h < topP.h && curP.h < bottomP.h {
			curP.a = -1
			curP.p = -1
		}
	}

	if rightP != nil {
		dfs(rightP.r, rightP.c, points)
	}
	if topP != nil {
		dfs(topP.r, topP.c, points)
	}
	if leftP != nil {
		dfs(leftP.r, leftP.c, points)
	}
	if bottomP != nil {
		dfs(bottomP.r, bottomP.c, points)
	}

	// обошли соседей, и из никого не получилось достичь океанов ставим -1
	if curP.p == 0 {
		curP.p = -1
	}
	if curP.a == 0 {
		curP.a = -1
	}
}

// копируем доступность океанов с болле низкоц
func swapOceans(aPoint, bPoint *Point) {
	if bPoint.h >= aPoint.h { // b выше a значит все что доступно a то и доступно b
		if aPoint.p == 1 {
			bPoint.p = 1
		}
		if aPoint.a == 1 {
			bPoint.a = 1
		}
	} else { // cur выше. Значит a выше копируем доступность с b
		if bPoint.p == 1 {
			aPoint.p = 1
		}
		if bPoint.a == 1 {
			aPoint.a = 1
		}
	}
}

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

func getWaterFlow(points [][]Point) [][]int {
	res := make([][]int, len(points))
	for i, row := range points {
		for j, point := range row {
			if point.a == 1 && point.p == 1 {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func TestWaterFlow(t *testing.T) {
	type tcase struct {
		heights [][]int
		expect  [][]int
	}
	cases := []tcase{
		{
			heights: [][]int{
				{1, 2, 3, 2, 1},
				{1, 2, 3, 2, 1},
				{1, 2, 3, 2, 1},
			},
			expect: [][]int{
				{0, 2}, {1, 2}, {2, 2},
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
