## Запуск с Docker Compose
1. Убедитесь, что Docker and Docker Compose установлены.
2. Откройте терминал и перейдите в директорию проекта.
3. Запустите Docker Compose:
   ```sh
   docker-compose up --build
   ```
4. Сервер запустится на http://localhost:8080, и БД PostgreSQL будет на localhost:5432.

## Функции
- При первом запуске приложения должны создаваться 10 кошельков с случайными
адресами и 100.0 у.е. на счету.
- Метож, который отправляет средства с одного из
кошельков на указанный кошелек.
- Метод возвращающий информацию о балансе кошелька в JSON-объекте.
- Метод возвращающий информацию о N последних по времени переводах средств.

## API Endpoints
### POST `/api/send`
Send - метод имеющий эндпоинт POST /api/send, который отправляет 
средства с одного из кошельков на указанный кошелек.
#### Request Body
```json
{
  "from": "<sender_wallet_address>",
  "to": "<receiver_wallet_address>",
  "amount": <amount_to_transfer>
}
```
#### Responses
- `200 OK` - Transaction completed successfully.
- `400 Bad Request` - Insufficient funds or invalid request.
- `404 Not Found` - Sender wallet not found.

### GET `/api/transactions?count=N`
Метод возвращающий информацию о N последних по времени переводах средств.
#### Query Parameters
- `count` - Количество переводов.
#### Response
```json
[
  {
    "id": 1,
    "from": "<sender_wallet_address>",
    "to": "<receiver_wallet_address>",
    "amount": <amount_transferred>,
    "time": "<transaction_timestamp>"
  }
]
```

### GET `/api/wallet/{address}/balance`
Метод возвращающий информацию о балансе кошелька в JSON-объекте.
#### Path Parameters
- `address` - Адрес кошелька.
#### Response
```json
{
  "address": "<wallet_address>",
  "balance": <wallet_balance>
}
```

## Локальный запуск
1. Убедитесь что установлен Gо (1.20 или последующие версии).
2. Клонируйте репозиторий и перейдите в директорию проекта.
3. Введите `go run main.go`.
4. Сервер запустится на `http://localhost:8080`.

## База данных
Приложение использует PostgreSQL для сохранения данных. Для хранения 
кошельков и транзакций создается файл с именем "transactions.db".

## Dependencies
- `github.com/lib/pq` - драйвер PostgreSQL.