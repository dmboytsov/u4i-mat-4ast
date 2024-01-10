Патеры/приемы при работе с параллельностью/асинхроностью, горутинами/мьютаксами/каналами/вейт группами
? - точно ли это патерн

# Канал Done
// TODO

# Mutex
// TODO пример

# Semaphore
- Примитив синхронизации. Ограничивает максимальное количество потоков/горутин, чтобы не получить большое количество горутин и не забить какой нибудь ресурс (например сеть, или перегрузить сторонний сервис)
- можно реализовать на буверезированном канале
```go
// метод что-то получает из какого то апи
func getDataFromAPI() {
    ...
}

// семафор размером 10. Маскимально 10 горутин одновременно
var sem = make(chan int, 10)
func main() {
   for _, employee := range employeeList {
       sem <- 1 // блочим местечко, или ждем пока освободится
       go func(){ 
           remoteDeleteEmployeeRPC(employee.ID)
           <-sem // освобождаем
       }()
   }
}
```
- https://pkg.go.dev/golang.org/x/sync/semaphore - взвешенный семафор, можно использовать какобычный семафор, но еще есть метод  **TryAcquire** который без блокировки проверяет можно ли запустить горутину


# Interlock
//? простые операции атомарны, аля пакет атомик в го

# Double checked locking (DCL)
// TODO ?

# Monitor
//? TODO не очень понял суть, погуглить, типа объект который следит за правильной работой остальных объектов

# Генераторы, 
//TODO
# Fan-In
//TODO
# Fan-Out
//TODO
# Управляющий канал и тд
//TODO

# Ссылки
- сборник патернов на английском - https://github.com/iamuditg/go-concurrency-patterns
- https://medium.com/@thejasbabu/concurrency-patterns-golang-5c5e1bcd0833
- семафор - https://medium.com/@deckarep/gos-extended-concurrency-semaphores-part-1-5eeabfa351ce
- взвешенный семафор - https://pkg.go.dev/golang.org/x/sync/semaphore
- http://www.golangpatterns.info/functional/closures
