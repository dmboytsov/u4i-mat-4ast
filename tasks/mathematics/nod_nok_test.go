package mathematics

import (
	"testing"
)

// https://coderun.yandex.ru/selections/quickstart/problems/gcd-and-lcm
// Вам даны 2 натуральных числа a и b
// Необходимо посчитать НОД(a, b) и НОК(a, b).
// НОД (a, b) − наибольшее натуральное число, на которое числа
// a, b делятся без остатка.
//
// НОК(a,b)− наименьшее натуральное число, которое делится на числа
// a, b без остатка.

// GCD вычисляет наибольший общий делитель (алгоритм Евклида)
// https://foxford.ru/wiki/matematika/algoritm-evklida?srsltid=AfmBOorZyfMSjrgc1QKdxq9IMi_cXWx4Kzz67gAf-VNN9pnp_yw2Gp7-&utm_referrer=https%3A%2F%2Fwww.google.com%2F
func GCD(a, b int) int {
	for b != 0 { // значит что a/b получилось без остатка.
		a, b = b, a%b //
		// a % b возвращает остаток от деления a на b.
		// например если a = 20 b = 8 (20/8 = 2.5) 2 целых. 20 - 8*2 = 4 //  остаток от деления
	}
	return a
}

// LCM вычисляет наименьшее общее кратное через формулу LCM(a, b) = |a*b| / GCD(a, b)
func LCM(a, b int) int {
	// Оптимизация для предотвращения переполнения: (a / GCD) * b
	gcd := GCD(a, b)
	return (a / gcd) * b
}

func TestGCDLCM(t *testing.T) {
	type tcase struct {
		a   int
		b   int
		nod int
		nok int
	}
	cases := []tcase{
		{
			a:   20,
			b:   8,
			nod: 4,
			nok: 40,
		}, {
			a:   2,
			b:   3,
			nod: 1,
			nok: 6,
		}, {
			a:   5,
			b:   15,
			nod: 5,
			nok: 15,
		},
	}

	for _, tc := range cases {
		nod := GCD(tc.a, tc.b)
		nok := LCM(tc.a, tc.b)
		if nod != tc.nod {
			t.Errorf("got %v; want %v", nod, tc.nod)
		}
		if nok != tc.nok {
			t.Errorf("got %v; want %v", nok, tc.nok)
		}
	}
}
