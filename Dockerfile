# === ЭТАП 1: Сборка бинарника ===
FROM golang:1.22-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы зависимостей
COPY go.mod go.sum ./

# Скачиваем зависимости (кешируется Docker-ом)
RUN go mod download

# Копируем весь остальной исходный код проекта
COPY . .

# Компилируем Go-приложение в один бинарный файл static-bot
# CGO_ENABLED=0 отключает зависимости от C-библиотек (делает бинарник полностью автономным)
# GOOS=linux гарантирует, что бинарник скомпилирован под Linux, даже если собираешь на Windows
RUN CGO_ENABLED=0 GOOS=linux go build -o static-bot .

# === ЭТАП 2: Минимальный финальный образ ===
# Используем alpine вместо scratch, так как нам нужны SSL-сертификаты для запросов к Telegram API
FROM alpine:3.19

# Устанавливаем ca-certificates (критично для работы с HTTPS/Telegram API)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем скомпилированный бинарник из первого этапа (builder)
COPY --from=builder /app/static-bot .

# Команда для запуска бота при старте контейнера
CMD ["./static-bot"]