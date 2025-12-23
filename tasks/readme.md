# Каталог с задачами

- Top interview questions
https://leetcode.com/explore/interview/card/top-interview-questions-medium/

# Яндекс
- https://coderun.yandex.ru/
- про собесы - https://yandex.ru/jobs/pages/dev_interview
- список задач- https://www.techinterviewhandbook.org/grind75/?difficulty=Easy&difficulty=Medium&weeks=8
- Задача про подмассив. [sub_array_test.go](arrays/sub_array_test.go)
Дан массив целых чисел и некоторое целое неотрицательное число K. Необходимо определить длину максимального подмассива из единиц,
  который можно получить заменой не более K любых элементов массива. 

# Списки
https://leetcode.com/problems/merge-k-sorted-lists/
https://leetcode.com/problems/linked-list-cycle/
https://leetcode.com/problems/add-two-numbers/
https://leetcode.com/problems/reverse-linked-list/

# binary search
https://leetcode.com/problems/binary-search/
https://leetcode.com/problems/guess-number-higher-or-lower/
https://leetcode.com/problems/search-a-2d-matrix/
https://leetcode.com/problems/search-in-rotated-sorted-array/
https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
https://leetcode.com/problems/search-in-rotated-sorted-array-ii/

# hash table
https://leetcode.com/problems/single-number/ - решить за O(1) по памяти
https://leetcode.com/problems/two-sum/
- есть решение через мапу
- есть решение отсортировать и двигаться с двух краев

https://leetcode.com/problems/4sum/
- Дан массив целых чисел и таргет, нужно вернуть все сочитания по 4 элемента, коткорые в сумме дают таргет
- Можно решить за квадрат.
- Делаем два цикла вложенный, находим все суммы все сочетания i j.
- Дальше также два вложенных цикла проверяем, что target - numc[l] - nums[k] = s ищем s в мапе сумм.
- 
- Надо попробовать такой вариант 
- делаем мапу s3 ключ  target - el, значение el
- делаем s2 ключ s3 - el значение el
- s2 - el = el2 если el2 есть, то по мапам находим пред варианты
- нужно чтобы это был не один и тот же el
- кейсы с дублями

https://leetcode.com/problems/group-anagrams/
https://leetcode.com/problems/valid-anagram/
https://leetcode.com/problems/find-all-anagrams-in-a-string/

# queue/stack
https://leetcode.com/problems/valid-parentheses/

# sort
https://leetcode.com/problems/merge-intervals/


----
# heap/hash
https://leetcode.com/problems/top-k-frequent-words/
https://leetcode.com/problems/top-k-frequent-elements/

# two pointers
https://leetcode.com/problems/container-with-most-water/
https://leetcode.com/problems/partition-labels/
# sliding window:
https://leetcode.com/problems/sliding-window-median/
https://leetcode.com/problems/sliding-window-maximum/
https://leetcode.com/problems/longest-repeating-character-replacement/

# tree
https://leetcode.com/problems/same-tree/
https://leetcode.com/problems/symmetric-tree/
https://leetcode.com/problems/balanced-binary-tree/
https://leetcode.com/problems/path-sum-ii/

# greedy problems
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

# Графы

dfs/bfs:
https://leetcode.com/problems/number-of-islands/
https://leetcode.com/problems/remove-invalid-parentheses/

# Реализации рейт лимитов