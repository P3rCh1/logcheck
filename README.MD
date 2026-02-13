# logcheck — Линтер для проверки лог-сообщений

### Репозиторий представляет тестовое задание для отбора на стажировку `Selectel`
Поддерживает логгеры **`log/slog`** и **`go.uber.org/zap`**.

---

## Проверяемые правила
- `capitalize`:  лог-сообщения должны начинаться со строчной буквы 
- `english`:  лог-сообщения должны быть только на английском языке
- `symbols`:  лог-сообщения не должны содержать спецсимволы или эмодзи
- `sensitive`:  логи не должны содержать потенциально чувствительные данные

## Особенности реализации

### Обработка сообщений (`capitalize`, `english`, `symbols`)
- Анализируется **только первый аргумент** логирующей функции (является сообщением, дальше - пары ключ значение или аргументы форматирования)
- Поддержка конкатенации строк: `"user " + username + " logged in"`

### Поиск чувствительных данных (`sensitive`)
- Анализ **имён переменных, полей структур, функций, слайсов и мап**: `user.Password`, `getApiKey()` ...
- Игнорируем строковые литералы: `"secret was changed"`
- Обработка всех аргументов: 
```
zap.Debug("auth", "password", pwd)
slog.Info("auth", slog.String("password", passwords[username]))
"api_key" + key
```

### Базовые правила
- **`symbols`**: разрешены только буквы, цифры, пробелы и символы из `allowed-symbols`
- **`sensitive`**: дефолтный набор `["token", "key", "pwd", "bearer", "password", "secret"]`

## Бонусные задания
### 1) Добавлена возможность конфигурации
*Пример конфигурации golangci.yaml:*
``` 
version: "2"
linters:
  default: none
  enable:
    - logcheck
  settings:
    custom:
      logcheck:
        type: "module"
        settings:
          enabled-rules: ["capitalize", "symbols", "sensitive"]
          sensitive-words: ["token", "bearer"]
          allowed-symbols: "%._"
```

### 2) Сделаны автоисправления для `capitalize`

### 3) Сделана возможность задавать кастомный список слов для `capitalize` *(см пример выше)*

### 4) Добавлен Github Workflow для автотестов и линтера

## Запуск
- Для сборки необходим установленный golengci-lint
### Для удобства был создан Mafefile:
- `make` или `make build` - сборка линтера `bin/logcheck`
- `make clean` - очищает `bin`
- `make example` - демонстрирует работу для тестового пакета `internal/analyzer/testdata`, в пакете есть свой конфиг (`internal/analyzer/testdata/.golangci.yaml`)

## Тестирование
- Реализованы unit тесты для всех правил (`internal/rules/rules_test`)
- Реализованы интеграционные тесты для набора testdata (`internal/analyzer/testdata`) с помощью `analysistest`
Для корректной работы интеграционных CI-тестов добавлен `internal/analyzer/testdata/vendor` (`go mod vendor`)


