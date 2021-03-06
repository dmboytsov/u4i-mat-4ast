
# Алгоритмы

## QuickSelect. Поиск k-ого элемента в не отсортированном массиве
- k-ый имеется ввиду что если бы массив был отсортирован то k это индекс элемента в отсортированном массиве
- в среднем алгоритм отрабатывает за O(n) но в худшем случаем за O(n^2)
- Идея в том что массив разбивается на две части, где в левой части элементы которые меньше *pivot (опорный элемент)* в правой все больше, 
разбиение алгоритма на две части назвывается *partitioning* результат partitioning - индекс элемента левее которого все меньше, а правее все больше. 
После partitioning у нас есть индекс если это не k-ый мы берем левую или правую часть (в зависимости k больше или меньше) и тд пока не найдем k-ый
- понятное объяснение partitioning - https://www.youtube.com/watch?v=MZaf_9IZCrc
  - выбираем опорный элемент, последний или случайный
  - идем с начала два индекса i-ый и j-ый. i-ый указывает на конец левой последовательности (конец меньших элементов), j-ый идет по массиву
  - если элемент больше опорного ничего не делаем, если элемент меньше опорного инкрементим i и меняем i-ый с j-ым - swap. (получается что в начале массива копятся элементы меньше j)
  - когда дошли до конца i-ый указывает на конец последовательности, i+1 меняем с опорным. Получили индекс элемента где слева все меньше опрного а справа все больше
  
## Очередь с приоритетом
- Такая очередь поддерживает (как минимум) две операции вставка и извлечение максимального (приоритетного), подходит для задач где нужно не по порядку брать, а в зависиости от чего-то, например распределение задач по процессорам, брать сначала долгие задачи.
- Такую очередь удобно строить на [бинарнной куче](https://github.com/dmboytsov/u4i-mat-4ast/blob/main/algorithms/data-structures.md#%D0%B1%D0%B8%D0%BD%D0%B0%D1%80%D0%BD%D0%B0%D1%8F-%D0%BA%D1%83%D1%87%D0%B0-binary-heap) при добавление в кучу элемент встает на свое место в дереве, вставка за логарифмическое время, извлечение максимального - это взять корень, и перестроить дерево (выбрать кто из потомков теперь корень)
