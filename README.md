# loglint

[![Go Version](https://img.shields.io/badge/Go-1.25%2B-blue)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

loglint — это статический анализатор (линтер) для Go, который проверяет лог-записи на соответствие лучшим практикам.

## Возможности

Линтер проверяет лог-сообщения в вызовах следующих логгеров:
- Стандартный `log`
- `log/slog`
- `go.uber.org/zap`

### Правила проверки

1. **Регистр букв** — сообщения должны начинаться со строчной буквы
   ```go
   log.Info("starting server")     // правильно
   log.Info("Starting server")      // ошибка
   ```
2. **Английский язык** — только ASCII символы
```go
log.Info("server started")       // правильно
log.Info("сервер запущен")       // ошибка
```
3. **Спецсимволы** — запрещены !@#$%^&*() и другие
```go
log.Info("request completed")    // правильно
log.Info("request completed!")   // ошибка
```
4. **Чувствительные данные** — проверка по ключевым словам
```go
log.Info("user authenticated")    // правильно
log.Info("user password: 1234")   // ошибка
```
## Установка
### Вариант 1: Через go install
```bash
go install github.com/ваш-username/loglint/cmd/loginit@latest
```

После этого команда loginit будет доступна в терминале.

### Вариант 2: Сборка из исходников
```bash
git clone https://github.com/weizis/loglint.git
cd loglint
make build
./loglint ./...
```

## Использование
### Запуск из командной строки
```bash
# Проверить текущую папку
go run ./cmd/loginit ./...
# Проверить пример
go run ./cmd/loginit ./example
```

### Запуск через Makefile
``` bash
make run         # запустить на всем проекте
make test        # запустить тесты
make build       # собрать бинарник
make plugin      # собрать плагин для Linux/macOS
make install     # установить в систему
```

## Интеграция с golangci-lint
loglint может работать как плагин для golangci-lint.

### Для Linux/macOS (плагин)
Собери плагин:
```bash
make plugin
```
Создай файл .golangci.yml в корне проекта:
```yaml
linters-settings:
  custom:
    loglint:
      path: ./loglint.so
      description: Log message linter
      original-url: github.com/weizis/loglint

linters:
  enable:
    - loglint
```
Запусти golangci-lint:
```bash
golangci-lint run ./...
```
### Для Windows (отдельный бинарник)
На Windows плагины не поддерживаются, используй отдельный бинарник:

Собери бинарник:
```bash
go build -o loglint.exe ./cmd/loginit
```
Создай файл .golangci.yml:
```yaml
linters-settings:
  custom:
    loglint:
      type: "exec"
      path: ./loglint.exe
      description: Log message linter
      original-url: github.com/weizis/loglint

linters:
  enable:
    - loglint
```
Запусти golangci-lint:
```bash
golangci-lint run ./...
```
**Примечание:** Убедись, что версия golangci-lint собрана с Go не ниже версии твоего проекта (1.25+).

## Конфигурация 
Для изменения списка проверяемых слов создайте файл .loglint.yaml в корне проекта:
```yaml
sensitive_words:
  - password
  - token
  - api_key
  - secret
  - key
  - credential
  - private_key
  - access_token
  - auth
```
Если файл отсутствует — используются значения по умолчанию.

### Авто-исправление
Для ошибок с большой буквы в начале сообщения линтер может автоматически исправить код. В VS Code наведите на ошибку и нажмите "Quick Fix" (лампочка).

## Пример работы
```bash
$ go run ./cmd/loginit ./example
D:\loginit\example\example.go:9:2: log message should use only English (ASCII) characters
D:\loginit\example\example.go:11:2: log message may contain sensitive data (e.g., 'password')
D:\loginit\example\example.go:13:2: log message should not contain special characters or punctuation
exit status 3
```

## Тестирование
```bash
go test ./...
```

## Структура проекта
```
.
├── cmd/
│   ├── loginit/           # точка входа для go run
│   │   └── main.go
│   └── loglint/           # точка входа для отдельного бинарника
│       └── main.go
├── internal/
│   └── analyzer/          # логика линтера
│       ├── analyzer.go
│       ├── config.go
│       ├── extract.go
│       ├── rules.go
│       ├── analyzer_test.go
│       └── testdata/
│           └── src/
│               └── testpkg/
│                   └── logs.go
├── plugin/
│   └── golangci-lint/     # плагин для golangci-lint
│       └── main.go
├── example/                # примеры
│   └── example.go
├── .loglint.yaml           # настройки
├── .golangci.yml           # конфиг для golangci-lint
├── .gitignore
├── Makefile
├── go.mod
├── go.sum
└── README.md
```
## Требования
- Go 1.25 или выше
- golangci-lint (опционально, для работы плагина)

## Бонусные возможности
- Конфигурация через YAML — настройка списка чувствительных слов
- Авто-исправление — Quick Fix для ошибок регистра
- Кастомные паттерны — расширяемый список чувствительных данных
- Интеграция с golangci-lint — работает как плагин или отдельный бинарник
  
## Лицензия
MIT

## Автор
Анастасия Михеева
