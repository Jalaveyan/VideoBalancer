# VideoBalancer

VideoBalancer - это gRPC-сервис для балансировки запросов видео между оригинальным сервером и CDN.

## Предварительные требования

Для запуска проекта вам понадобятся:

- Docker
- Docker Compose
- Git (для клонирования репозитория)

## Как начать

### Клонирование репозитория

Для начала склонируйте репозиторий на вашу локальную машину:

```bash
git clone https://github.com/Jalaveyan/VideoBalancer.git
cd VideoBalancer

Сборка и запуск с помощью Docker Compose
docker-compose up --build

Для локальной разработки и тестирования вы можете запустить сервис вне Docker, если у вас установлен Go:
cd cmd
go run main.go

Убедитесь, что все зависимости установлены:
go mod tidy
