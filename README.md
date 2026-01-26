# CryptoWallet Backend

Криптовалютный кошелек с функционалом управления аккаунтами, кошельками и транзакциями.

## Что делает программа

MVP Plan:

- Регистрация и аутентификация пользователей с JWT токенами
- Создание и управление криптовалютными кошельками
- Отправка и получение криптовалют
- История транзакций
- P2P обмен между пользователями
- Интеграция с блокчейнами (Bitcoin, Ethereum, Solana)
- Мониторинг и метрики в реальном времени

## Архитектура

Проект строится на микросервисной архитектуре с использованием Clean Architecture:

```
cmd/
├── api-gateway/          # HTTP REST API для фронтенда (Gin)
├── auth-service/         # Аутентификация и авторизация (gRPC)
├── wallet-service/       # Управление кошельками (gRPC)
├── transaction-service/  # Обработка транзакций (gRPC)
├── p2p-service/         # P2P обмен (gRPC)
└── exchange-service/    # Криптовалютный обмен (gRPC)

internal/
├── api-gateway/
│   ├── delivery/http/   # HTTP handlers (Gin) для фронтенда
│   ├── infrastructure/grpc/  # gRPC клиенты для вызова микросервисов
│   └── middleware/      # Логирование, метрики
├── auth/                 # Auth микросервис
│   ├── domain/          # Бизнес-логика (сущности, интерфейсы)
│   ├── infrastructure/  # Внешние системы (PostgreSQL, JWT, Password)
│   └── delivery/grpc/   # gRPC сервер
├── wallet/              # Wallet микросервис
└── ...

api/
└── proto/               # Proto файлы для gRPC
    └── auth.proto

pkg/                     # Общие компоненты
├── logger/              # Структурированное логирование
└── metrics/             # Метрики Prometheus
```

### Коммуникация между сервисами

- **Фронтенд ↔ API Gateway**: HTTP REST API (Gin)
- **API Gateway ↔ Микросервисы**: gRPC
- **Микросервисы ↔ Микросервисы**: gRPC

## Технологии

### Основной стек
- **Go 1.23** - основной язык программирования
- **Gin** - HTTP фреймворк для REST API (API Gateway)
- **gRPC** - межсервисная коммуникация между микросервисами
- **Protocol Buffers** - определение контрактов для gRPC
- **PostgreSQL** - основная база данных
- **Redis** - кэширование и сессии
- **Apache Kafka** - event-driven коммуникация между сервисами

### Мониторинг и наблюдаемость
- **Prometheus** - сбор метрик
- **Grafana** - визуализация метрик и дашборды
- **Zerolog** - структурированное логирование

### Инфраструктура
- **Docker Compose** - локальная разработка
- **golang-migrate** - миграции базы данных
- **JWT** - аутентификация и авторизация

## Запуск

1. Установите зависимости:
```bash
go mod download
```

2. Запустите инфраструктуру (PostgreSQL, Redis, Prometheus, Grafana):
```bash
make up
```

3. Примените миграции:

###TODO

4. Запустите Auth Service (gRPC сервер на порту 50051):

###TODO

5. Запустите API Gateway (HTTP REST API на порту 8080):
```bash
go run cmd/api-gateway/main.go
```

## API Endpoints

### TODO

## Разработка

### Структура проекта
Проект следует принципам Clean Architecture:
- **Domain слой** содержит бизнес-логику и интерфейсы (не зависит от внешних систем)
- **Infrastructure слой** реализует внешние зависимости (БД, gRPC клиенты/серверы)
- **Delivery слой** обрабатывает запросы (HTTP handlers для API Gateway, gRPC handlers для микросервисов)

### Микросервисы

Каждый микросервис:
- Имеет свою точку входа в `cmd/<service-name>/main.go`
- Содержит полную реализацию в `internal/<service-name>/`
- Предоставляет gRPC API для других сервисов
- Имеет свою базу данных (или схему в общей БД)

### Тестирование

### TODO

### Миграции
```bash
# Создание новой миграции
migrate create -ext sql -dir migrations -seq add_users_table

# Применение миграций
migrate -path migrations -database "postgres://cryptowallet:cryptowallet123@localhost:5432/cryptowallet?sslmode=disable" up
```

### Генерация gRPC кода
```bash
# После изменения proto файлов
protoc --go_out=. --go-grpc_out=. api/proto/*.proto
```

## Порты сервисов

- **API Gateway**: 8080 (HTTP REST API)
- **Auth Service**: 50051 (gRPC)
- **Wallet Service**: 50052 (gRPC)
- **Transaction Service**: 50053 (gRPC)
- **PostgreSQL**: 5432
- **Redis**: 6379
- **Prometheus**: 9090
- **Grafana**: 3000

## Цели проекта

- Изучение современных подходов к backend разработке
- Практика работы с микросервисной архитектурой
- Освоение gRPC для межсервисной коммуникации
- Освоение event-driven подходов
- Реализация мониторинга
- Применение BestPractice Go разработки
