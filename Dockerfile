# Используем официальное изображение Golang для сборки бинарного файла
FROM golang:1.22-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum и устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем все файлы исходного кода
COPY . .

# Собираем бинарный файл приложения
# Указываем путь к main.go в папке cmd
RUN go build -o /video-balancer ./cmd/main.go

# Используем минимальное изображение Alpine для запуска приложения
FROM alpine:latest

# Создаем директорию для приложения
WORKDIR /root/

# Копируем бинарный файл из стадии сборки
COPY --from=builder /video-balancer .

# Устанавливаем переменную окружения по умолчанию (можно изменить при запуске)
ENV CDN_HOST=storage.googleapis.com

# Открываем порт 8080
EXPOSE 8080

# Запускаем приложение
CMD ["./video-balancer"]
