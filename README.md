# Distributed Calculator
## Структура проекта

- **`cmd/`**
  - **`orchestrator/`**
    - **`main.go`**
  - **`agent/`**
    - **`main.go`**
- **`internal/`**
  - **`orchestrator/`**
    - **`orchestrator.go`**
  - **`agent/`**
    - **`agent.go`**
    - **`agent_test.go`**
- **`web/`**
  - **`index.html`**
- **`go.mod`**
- **`go.sum`**
- **`README.md`**
## Запуск проекта
1. **Установите Go версии 1.23.2**:
   Ссылка для этого: [Go Download](https://go.dev/doc/install)
3. **Клонируйте репозиторий**:
   ```
   git clone https://github.com/3SMA3/distributed-calculator.git
   ```
4. **Перейдите в папку с проектом**:
   ```
   cd distributed-calculator
   ```
5. **Установите зависимости**:
   ```
   go mod tidy
   ```
6. **Запустите оркестратор**:
   ```
   go run ./cmd/orchestrator/main.go
   ```
7. **Запустите агентов**:
   Откройте новый терминал и выполните:
   ```
   go run ./cmd/agent/main.go
   ```
8. **Откройте веб-интерфейс**:
   Перейдите по адресу `http://localhost:8080` в браузере.
## Примеры использования
### Добавление выражения
```
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```
### Получение списка выражений
```
curl --location 'localhost:8080/api/v1/expressions'
```
### Получение выражения по ID
```
curl --location 'localhost:8080/api/v1/expressions/:id'
```
## Тесты
В проекте реализованы unit-тесты для пакета `agent`, которые проверяют корректность вычислений и обработки задач. Тесты находятся в файле `internal/agent/agent_test.go`.
### Запуск тестов
**Выполните команду:**
```
go test ./internal/agent/...
```
### Что тестируется
Тесты покрывают следующие сценарии:
1. **Арифметические операции**:
   - Сложение (`+`).
   - Вычитание (`-`).
   - Умножение (`*`).
   - Деление (`/`).
   - Обработка деления на ноль.
   - Обработка неизвестной операции.
2. **Время выполнения операций**:
   - Проверка, что операции выполняются за указанное время.
3. **Корректность результатов**:
   - Проверка, что результаты вычислений соответствуют ожидаемым.
### Примеры тестов
- **Сложение**:
```
task := agent.Task{
    Arg1:          2.5,
    Arg2:          3.5,
    Operation:     "+",
    OperationTime: 100,
}
result, err := agent.Compute(task)
assert.NoError(t, err)
assert.Equal(t, 6.0, result)
```
- **Деление на ноль**:
```
task := agent.Task{
    Arg1:          10.0,
    Arg2:          0.0,
    Operation:     "/",
    OperationTime: 100,
}
result, err := agent.Compute(task)
assert.Error(t, err)
assert.Equal(t, "division by zero", err.Error())
assert.Equal(t, 0.0, result)
```
- **Неизвестная операция**:
```
task := agent.Task{
    Arg1:          10.0,
    Arg2:          5.0,
    Operation:     "%",
    OperationTime: 100,
}
result, err := agent.Compute(task)
assert.Error(t, err)
assert.Equal(t, "unknown operation", err.Error())
assert.Equal(t, 0.0, result)
```
