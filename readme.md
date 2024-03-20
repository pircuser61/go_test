### 1. Написание базового теста функции

\internal\someFunc\someFunc_test.go

- TestSum
- TestSumTable табличный тест
- TestSumParallelTable запуск табличных тестов параллельно

### 2. Тестирование ошибок

\internal\someFunc\someFunc_test.go

- TestError

### 3. Тестирование слайсов и мапов

\internal\someFunc\someFunc_test.go

- TestSlice

### 4. Использование подпакета testing/quick

\internal\someFunc\someFuncQuick_test.go

- TestSumQ

### 5. Тестирование HTTP-обработчик

\cmd\httpclient\main_test.go

### 6. Mocking

\internal\core\logic\logic_test.go

Генерация:
mockgen -destination=internal/core/mocks/mock_doer.go ^
-package=mocks ^
github.com/pircuser61/go_test/internal/core/doer ^
Doer

- TestLogicOk
- TestLogicErr

### 7. Тестирование времени

\internal\core\logic\logic_test.go

- TestLogicTime

### 8. Тестирование конкурентности

\internal\core\logic\logic_test.go

- TestLogicParallel
- TestLogicOrder

### 9. Тестирование конкретной ошибки

\internal\core\logic\logic_test.go

- TestLogicConcreteErr Проверка на равеснство переменной с ошибкой
- TestLogicErrType Проверка на тип ошибки

### 10. Тестирование HTTP-сервера

\cmd\httpserver\main_test.go
