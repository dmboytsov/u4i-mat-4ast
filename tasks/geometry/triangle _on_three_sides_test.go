package geometry

import (
	"fmt"
	"testing"
)

// Условие.
// Даны три натуральных числа.
// Возможно ли построить треугольник с такими сторонами?
// Если это возможно, выведите строку YES, иначе выведите строку NO.
//
// Треугольник — это три точки, не лежащие на одной прямой.
//
// Решение.
// Сумма длин наименьших сторон должна быть больше длинны третьей стороны.
// Можно представить, что берем два наименьших отрезка, и чтобы подошел третий разварачиваем треугольник,
// увеличиваем угол, и вот если угол дошел до 180, значит отрезки встали в прямую и никакого треугольника не получиться.
//

func IsPossibleTriangle(a, b, c int) bool {
	// выбрать самую болььшую сторону, сумма двух оставшихся должна быть меньше самой большой

	// рандомно сетим
	maxSide := a
	oneSide := b
	twoSide := c

	if a > b {
		oneSide = b
		if a > c {
			twoSide = c
			maxSide = a
		} else {
			twoSide = a
			maxSide = c
		}
	} else {
		oneSide = a
		if b > c {
			twoSide = c
			maxSide = b
		} else {
			twoSide = b
			maxSide = c
		}
	}

	if (oneSide + twoSide) > maxSide {
		return true
	}

	return false
}

func TestIsPossibleTriangle(t *testing.T) {
	type tcase struct {
		a      int
		b      int
		c      int
		expect bool
	}

	cases := []tcase{
		{
			a:      3,
			b:      4,
			c:      5,
			expect: true,
		}, {
			a:      5,
			b:      4,
			c:      3,
			expect: true,
		}, {
			a:      5,
			b:      3,
			c:      4,
			expect: true,
		},
	}

	for _, c := range cases {
		got := IsPossibleTriangle(c.a, c.b, c.c)
		if got != c.expect {
			t.Errorf("case %v, got %v, want %v", c, got, c.expect)
		} else {
			fmt.Printf("success case %v", c)
		}
	}
}
