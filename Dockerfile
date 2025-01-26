# Используем официальный образ Golang для сборки
FROM golang:1.20 AS builder

WORKDIR /app

# Копируем файлы зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение для Linux
RUN go build -o main .

# Используем минимальный образ Alpine для запуска приложения
FROM alpine:3.16

WORKDIR /root/

# Устанавливаем необходимые библиотеки
RUN apk add --no-cache libc6-compat

# Копируем собранное приложение
COPY --from=builder /app/main .

# Устанавливаем права на выполнение
RUN chmod +x /root/main

# Открываем порт для приложения
EXPOSE 8080

# Команда для запуска приложения
CMD ["./main"]