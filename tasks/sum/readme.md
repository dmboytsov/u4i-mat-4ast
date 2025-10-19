/*
Дан массив целых чисел, нужно найти непустой подотрезок (непрерывную подпоследовательность)
с заданной суммой target либо сказать, что это невозможно.
find_target([9, -6, 5, 1, 4, -2], 10) -> (2, 4)
*/

func findTarget(arr []int, target int) (int, int) {
sumMap := make(map[int]int, len(arr))

    s := 0
    lastIndex:=0
    for i, el  := range arr {
        k := s+el
        if k => target{
            lastIndex = i
        }
       
        sumMap[i] = k
        if k == target {
            return 0, i
        }

        delta := target - k
        if firtstIndex, ok := sumMap[delta] && ok {
            return firtstIndex, i
        }
    }

    return 0,0
}