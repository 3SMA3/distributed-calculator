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
    - **`parser/`**
      - **`parser.go`**
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
2. **Клонируйте репозиторий**:
   ```
   git clone https://github.com/3SMA3/distributed-calculator.git
   ```
3. **Перейдите в папку с проектом**:
   ```
   cd distributed-calculator
   ```
4. **Установите зависимости**:
   ```
   go mod tidy
   ```
5. **Запустите оркестратор**:
   ```
   go run ./cmd/orchestrator/main.go
   ```
   Оркестратор запустится на порту 8080
6. **Запустите агентов**:

   **Откройте новый терминал и выполните:**
   ```
   go run ./cmd/agent/main.go
   ```
8. **Откройте веб-интерфейс**:

Перейдите по адресу `http://localhost:8080` в браузере.
```
http://localhost:8080
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
   - Вычитание (`-`)(его нет в тестах, но оно работает).
   - Умножение (`*`).
   - Деление (`/`).
2. **Приоритет операций**:
   - Умножение выполняется перед сложением.
3. **Выражения со скобками**:
   - Корректное вычисление выражений с учётом скобок.
4. **Обработка ошибок**:
   - Деление на ноль.
   - Некорректные выражения.
### Примеры тестов
- **Простое сложение**:
```
t.Run("Simple addition", func(t *testing.T) {
    result, err := agent.ComputeExpression("2+2")
    assert.NoError(t, err)
    assert.Equal(t, 4.0, result)
})
```
- **Умножение перед сложением**:
```
t.Run("Multiplication before addition", func(t *testing.T) {
    result, err := agent.ComputeExpression("2+2*2")
    assert.NoError(t, err)
    assert.Equal(t, 6.0, result)
})
```
- **Выражение со скобками**:
```
t.Run("With parentheses", func(t *testing.T) {
    result, err := agent.ComputeExpression("(2+3)*4")
    assert.NoError(t, err)
    assert.Equal(t, 20.0, result)
})
```
- **Деление**:
```
t.Run("Division", func(t *testing.T) {
    result, err := agent.ComputeExpression("10/(2+3)")
    assert.NoError(t, err)
    assert.Equal(t, 2.0, result)
})
```
- **Деление на ноль**:
```
t.Run("Division by zero", func(t *testing.T) {
    _, err := agent.ComputeExpression("10/0")
    assert.Error(t, err)
    assert.Equal(t, "division by zero", err.Error())
})
```
- **Некорректное выражение**:
```
t.Run("Invalid expression", func(t *testing.T) {
    _, err := agent.ComputeExpression("2+*3")
    assert.Error(t, err)
})
```
