# go-unit-converter

Конвертер единиц измерения - это веб-приложение, позволяющее пользователям конвертировать значения между различными единицами измерения длины, веса и температуры.

## 1. Описание проекта

Этот проект представляет собой полноценное веб-приложение с бэкендом на Go и фронтендом на HTML/CSS/JavaScript. Приложение предоставляет удобный интерфейс для конвертации единиц измерения в трех категориях:

- **Длина**: миллиметры, сантиметры, метры, километры, дюймы, футы, ярды, мили
- **Вес**: граммы, килограммы, фунты, унции
- **Температура**: Цельсий, Фаренгейт, Кельвин

### Технологический стек

**Бэкенд:**
- Язык: Go (Golang)
- Фреймворк: Gin Gonic
- Логирование: Zap
- CORS: gin-contrib/cors

**Фронтенд:**
- HTML5
- CSS3
- JavaScript (Vanilla)

**Инфраструктура:**
- Docker и Docker Compose для контейнеризации
- Nginx в качестве веб-сервера и обратного прокси

### Архитектура

Проект следует чистой архитектуре с четким разделением на слои:

1. **Handler**: Обработка HTTP-запросов
2. **Service**: Бизнес-логика
3. **Util**: Утилиты для конвертации
4. **Model**: Модели данных
5. **Router**: Маршрутизация
6. **Config**: Конфигурация
7. **Middleware**: Промежуточное ПО

### Особенности реализации

- Поддержка двух окружений: разработка (dev) и продакшн (prod)
- Настраиваемый CORS
- Централизованное логирование
- Graceful shutdown сервера
- Контейнеризация с помощью Docker

## 2. Как установить и запустить проект

### Предварительные требования

- Docker и Docker Compose
- Go 1.25 (если запускаете без Docker)

### Установка и запуск

#### Вариант 1: С использованием Docker (рекомендуется)

```bash
# Клонировать репозиторий
git clone https://github.com/kavlan-dev/go-unit-converter.git
cd go-unit-converter

# Собрать и запустить контейнеры
docker-compose up -d --build
```

Приложение будет доступно по адресу:
- Фронтенд: http://localhost:8000
- Бэкенд API: http://localhost:8080

#### Вариант 2: Без Docker

```bash
# Клонировать репозиторий
git clone https://github.com/kavlan-dev/go-unit-converter.git
cd go-unit-converter

# Установить зависимости
go mod download

# Запустить бэкенд
go run cmd/app/main.go
```

### Переменные окружения

Вы можете настроить приложение с помощью следующих переменных окружения:

- `ENV`: Окружение (dev или prod), по умолчанию dev
- `CORS`: Разрешенные источники CORS, разделенные запятыми, по умолчанию *

Пример:
```bash
export ENV=dev
export CORS="http://localhost:8000,http://localhost:3000"
```

## 3. Как использовать проект

### Интерфейс пользователя

1. Откройте браузер и перейдите по адресу http://localhost:8000
2. Выберите вкладку с типом конвертации (Длина, Вес или Температура)
3. Введите значение для конвертации
4. Выберите исходную единицу измерения
5. Выберите целевую единицу измерения
6. Нажмите кнопку "Конвертировать"
7. Результат появится ниже формы

### API Endpoints

Бэкенд предоставляет следующие API-эндпоинты:

**POST /api/convert/length**
```json
{
  "from_value": 100,
  "from_unit": "meter",
  "to_unit": "kilometer"
}
```

**POST /api/convert/weight**
```json
{
  "from_value": 100,
  "from_unit": "kilogram",
  "to_unit": "pound"
}
```

**POST /api/convert/temperature**
```json
{
  "from_value": 100,
  "from_unit": "celsius",
  "to_unit": "fahrenheit"
}
```

Все эндпоинты возвращают JSON-ответ в формате:
```json
{
  "result": 0.1
}
```

### Примеры использования API

```bash
# Конвертация длины
curl -X POST http://localhost:8080/api/convert/length \
  -H "Content-Type: application/json" \
  -d '{"from_value": 100, "from_unit": "meter", "to_unit": "kilometer"}'

# Конвертация веса
curl -X POST http://localhost:8080/api/convert/weight \
  -H "Content-Type: application/json" \
  -d '{"from_value": 100, "from_unit": "kilogram", "to_unit": "pound"}'

# Конвертация температуры
curl -X POST http://localhost:8080/api/convert/temperature \
  -H "Content-Type: application/json" \
  -d '{"from_value": 100, "from_unit": "celsius", "to_unit": "fahrenheit"}'
```

## 4. Сведения о проекте

Проект создан с использованием следующих технологий и библиотек:
- [Gin Gonic](https://github.com/gin-gonic/gin) - Веб-фреймворк для Go
- [Zap](https://github.com/uber-go/zap) - Быстрое структурированное логирование
- [gin-contrib/cors](https://github.com/gin-contrib/cors) - Middleware для CORS

## 5. Лицензия

Этот проект лицензируется по лицензии MIT. Подробности смотрите в файле [LICENSE](LICENSE).
