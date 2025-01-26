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
- `address` - The address of the wallet.
#### Response
```json
{
  "address": "<wallet_address>",
  "balance": <wallet_balance>
}
```

## Running Locally
1. Install Go (1.20 or later).
2. Clone the repository and navigate to the project directory.
3. Run `go run main.go`.
4. The server will start on `http://localhost:8080`.

## Database
The application uses PostgreSQL to persist data. A file named `transactions.db` is created to store wallets and transactions.

## Dependencies
- `github.com/lib/pq` - PostgreSQL driver.