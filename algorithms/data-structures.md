# Куча (heap) // дерево
Куча — это специализированная структура данных типа дерево, которая удовлетворяет свойству кучи: если B является узлом-потомком узла A, то ключ(A) ≥ ключ(B). 
Из этого следует, что элемент с наибольшим ключом всегда является корневым узлом кучи, поэтому иногда такие кучи называют max-кучами.

## Бинарная куча (binary heap)
https://habr.com/ru/post/112222/
- Бинарная куча - бинарное дерево, у узла два потомка, в корне элемент с наивысшим приоритетом.
- Двоичную кучу удобно хранить в виде одномерного массива, причем левый потомок вершины с индексом i имеет индекс 2*i+1, а правый 2*i+2. 
Корень дерева – элемент с индексом 0. Высота двоичной кучи равна высоте дерева, то есть log2 N, где N
[Рисунок как куча храниться в массиве](https://ru.wikipedia.org/wiki/%D0%9A%D1%83%D1%87%D0%B0_(%D1%81%D1%82%D1%80%D1%83%D0%BA%D1%82%D1%83%D1%80%D0%B0_%D0%B4%D0%B0%D0%BD%D0%BD%D1%8B%D1%85)#/media/%D0%A4%D0%B0%D0%B9%D0%BB:Max-Heap.svg)
- Вставка происходит за логарефмическое время. Вставляем элемент в конец, смотрим на родителя, 
если приоритет элемента больше родителя меняем местами, и тд пока не найдем место.