package mathematics

import (
	"reflect"
	"testing"
)

// https://leetcode.com/problems/4sum/
// - Дан массив целых чисел и таргет, нужно вернуть все сочитания по 4 элемента, которые в сумме дают таргет

// Вариант идти навстречу с лева и справа, а внутри (вложенным циклом) идти тоже навстречу, не работает,
// так как если двигать и правый и левый, то мы пропускаем сочетания, если двигать только правый, то пропускаем сочетания слева.
// Движение на встречу примениму только для 2sum

// Можно решить за квадрат.
// a+b+c+d = target
// a+b = target-(c+d)
// Собираем мапы сум, все сочетания двух чисел (тут будет квадрат).
// Вторым циклом идем по мапе, авчитаем из таргета, смотрим есть ли результат в мапе.
// TODO решить основная проблема с дубликатами, одни и теже элементы массива могут получать разные суммы.
// Попробовать сделать мапу где ключем будем массив (а так можно?)

func fourSum(list []int, target int) [][]int {
	result := make(map[[4]int]struct{}, len(list))
	if len(list) < 4 {
		return [][]int{}
	}

	// мапа ключ - сумма двух чисел, значение - мапа, где ключ минимальный элемент, значение максимальный элемент
	// вложенная мапа, чтобы сократить дубли
	twoSum := make(map[int]map[int]int, len(list)*len(list))
	for i := 0; i < len(list)-2; i++ {
		for j := i + 1; j < len(list)-1; j++ {
			s := list[i] + list[j]
			minKey := list[j]
			maxVal := list[i]
			if list[i] < list[j] {
				minKey = list[i]
				maxVal = list[j]
			}
			if _, ok := twoSum[s]; !ok {
				twoSum[s] = map[int]int{}
			}
			twoSum[s][minKey] = maxVal
		}
	}

	for s, m := range twoSum {
		s2 := target - s
		if m2, ok := twoSum[s2]; ok {
			// сочетания m и m2. одна и таже сумма может получаться из разных слагаемых, поэтому за суммой скарывается список пар слагаемых
			// сочетания этих пар будет нашим ответом
			for k, v := range m {
				for k2, v2 := range m2 {
					if k == k2 || v == v2 {
						continue
					}
					// составляем ключ для мапы, боремся с дублями
					mk := [4]int{k, k2, v, v2}
					if k > k2 {
						mk[0] = k2
						mk[1] = k
					}
					if v > v2 {
						mk[2] = v2
						mk[3] = v
					}
					result[mk] = struct{}{}
				}
			}
		}
		// удаляем s так как, когда первый цикл дойдет до суммы s2, то мы снова возьмем эту пару.
		// первый раз s и s2 а второй раз s2 и s.
		//delete(twoSum, s)
	}

	// мапу в слайс
	rl := make([][]int, 0, len(result))
	for mk := range result {
		rl = append(rl, mk[:])
	}

	return rl
}

// //
func TestFourSum(t *testing.T) {
	type tcase struct {
		list   []int
		target int
		result [][]int
	}

	cases := []tcase{
		{
			list:   []int{-1, -2, 0, 2, 1, 3, 4, 5, 6, 7, 8},
			target: 8,
			result: [][]int{
				[]int{-2, -1, 3, 8},
				[]int{-2, 0, 2, 8},
				[]int{-2, -1, 4, 7},
				[]int{-2, 0, 3, 7},
				[]int{-2, 1, 2, 7},
				[]int{-2, -1, 5, 6},
				[]int{-2, 0, 4, 6},
				[]int{-2, 1, 3, 6},
				[]int{-2, 1, 4, 5},
				[]int{-2, 2, 3, 5},
				[]int{0, 1, 3, 4},
				[]int{0, 1, 2, 5},
			},
		},
	}

	for _, tc := range cases {
		got := fourSum(tc.list, tc.target)
		if !reflect.DeepEqual(got, tc.result) {
			t.Errorf("got %v do not equal", got)
			t.Errorf("want %v", tc.result)
		}
	}
}
