#  Context
// TODO

# Приведение типов
// TODO

# Примитивы синхронизации
## RWMutex - можно давать читать, но не давать писать, 
```
var m sync.RWMutex
m.RLock() // блокировка на запись, можно читать в других горутнинах
m.RUnlock() // снимаеи блокировку
```
## Каналы
## Горутины
## Паттерны конкурентного программирования
- Генераторы, Fan-In, Fan-Out, Управляющий канал и тд
