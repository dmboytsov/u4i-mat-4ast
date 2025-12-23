package arrays

import (
	"fmt"
	"testing"
)

// Дан массив целых чисел и некоторое целое неотрицательное число K.
// Необходимо определить длину максимального подмассива из единиц,
// который можно получить заменой не более K любых элементов массива.
//
// Алгоритм.
// Два индекса на начало подмассива и на конец. l-left r-right
// Сначала двигаем правый указатель до тех пор пока не закончатся замены (декрементим k пока он не уйдет в отрицательный)
// Когда закончились замены, начинаем двигать левый указатель, пока не увеличим K1, нужно пройти хотябы один элемент, который потратил замену
//
// Сложност.
// Получается O(n)
//
// Примеры:
// FindMaxLengthOfOnes([]int{1,-2,3,-4,1}, 2) == 3
// FindMaxLengthOfOnes([]int{1,1,1,1}, 1) == 4
// FindMaxLengthOfOnes([]int{1,3,1,3,1,7}, 2) == 5
// FindMaxLengthOfOnes([]int{1,3,4,1,3,3,1,7,1,1}, 2) == 4

func FindMaxLengthOfOnes(sl []int, k int) int {
	maxLen := 0
	l, r := 0, 0
	k1 := k
	for r < len(sl) {
		if sl[r] != 1 {
			k1 = k1 - 1 //уменьшаем K1 (копия K), уменьшаем запас замен на 1
		}

		for k1 < 0 { // запас замен закончился,
			// начинаем двигать левую границу, пока замен не станет больше 0, если 1 то замен не добавляется
			if sl[l] != 1 {
				k1++
			}
			l++
		}

		r++              // правую границу двигаем всегда, на каждой итерации
		cl := r - l      // длина подстроки
		if cl > maxLen { // запоминаем максимальную длинну
			maxLen = cl
		}
	}

	return maxLen
}

func TestFindMaxLengthOfOnes(t *testing.T) {
	type tcase struct {
		sl     []int
		k      int
		expect int
	}

	cases := []tcase{
		{
			sl:     []int{1, -2, 3, -4, 1},
			k:      2,
			expect: 3,
		}, {
			sl:     []int{1, 1, 1, 1},
			k:      1,
			expect: 4,
		}, {
			sl:     []int{0, 1, 1, 1},
			k:      1,
			expect: 4,
		}, {
			sl:     []int{1, 1, 1, 0},
			k:      1,
			expect: 4,
		}, {
			sl:     []int{1, 1, 1, 0},
			k:      1,
			expect: 4,
		}, {
			sl:     []int{1, 3, 1, 3, 1, 7},
			k:      2,
			expect: 5,
		}, {
			sl:     []int{1, 3, 4, 1, 3, 3, 1, 7, 1, 1},
			k:      2,
			expect: 5, // 3, 1, 7, 1, 1
		},
	}

	for _, c := range cases {
		got := FindMaxLengthOfOnes(c.sl, c.k)
		if got != c.expect {
			t.Errorf("case %v, got %v, want %v", c, got, c.expect)
		} else {
			fmt.Printf("success case %v", c)
		}
	}
}
