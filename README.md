
# Wallet

Это приложение является REST API для финансового учреждения, где он предоставляет своим партнёрам услуги электронного кошелька. У него есть два типа учетных записей электронного кошелька: идентифицированные и неидентифицированные.

API может поддерживать несколько клиентов, и следует использовать только методы http, post с json в качестве формата данных. Клиенты должны быть аутентифицированы через http параметр заголовок X-UserId и X-Digest.
## Установка

Для начала работы с проектом убедитесь, что у вас установлен Go и PostgreSQL.

## Настройка

Создайте файл `.env` в корневой директории проекта и добавьте следующие переменные окружения:

```env
CONFIG_PATH=config/config.yml
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=password
POSTGRES_DATABASE=wallet
SECRET_KEY=secret
```
## Запуск

Для запуска приложения используйте следующую команду:

```sh
make run
```

Эта команда выполнит `go run cmd/wallet/main.go`, запустив сервер на порту, указанном в config файле (по умолчанию 54321).

## Миграции

Для применения миграций базы данных используйте следующие команды:

### Применить миграции:

```sh
make migrate_up
```
`//NOTE: Замените конфигурационные данные БД в Makefile на свои`
### Откатить миграции:

```sh
make migrate_down
```

## API Роуты

Клиенты аутентифицированы через http параметр заголовок X-UserId и X-Digest.

### Exists

Проверка существует ли кошельёк:

```
POST http://localhost:54321/api/v1/wallets/exists
```

Успешний ответ:
```json
{
  "message": "wallet exists"
}
```
Неуспешний ответ: 
```json
{
  "error": "wallet not found"
}
```

### Balance

Проверка баланса кошелька:

```
POST http://localhost:54321/api/v1/wallets/balance
```

Успешний ответ:
```json
{
  "balance": 5003.03
}
```
Неуспешний ответ:
```json
{
  "error": "wallet not found"
}
```

### Deposit

Пополнение кошелька:

```
POST http://localhost:54321/api/v1/wallets/deposit
```

Успешний ответ:
```json
{
  "message": "balance added"
}
```
Неуспешний ответ:
```json
{
  "error": "wallet not found"
}
```
Превышение лимита:
```json
{
  "error": "max balance exceeded"
}
```

### total-deposits

Получение общего количества и суммы операций пополнения за текущий месяц:

```
POST http://localhost:54321/api/v1/wallets/total-deposits
```

Успешний ответ:
```json
{
  "total_count": 5,
  "total_sum": 15002.02
}
```
Неуспешний ответ:
```json
{
  "error": "wallet not found"
}
```

## Описание Таблиц

### Таблица `wallets`

```sql
CREATE TABLE wallets (
    id BIGSERIAL PRIMARY KEY,
    balance BIGINT DEFAULT 0,
    user_id       VARCHAR(255),
    is_identified BOOLEAN DEFAULT FALSE
);
```

### Таблица `transactions`

```sql
create table transactions(
    id BIGSERIAL PRIMARY KEY,
    wallet_id BIGINT REFERENCES wallets(id),
    amount    BIGINT NOT NULL,
    time      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```
