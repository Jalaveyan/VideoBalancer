# VideoBalancer

VideoBalancer - это gRPC-сервис для балансировки запросов видео между оригинальным сервером и CDN.

## Предварительные требования

Для запуска проекта вам нужно:

- Docker
- Docker Compose
- Git (для клонирования репозитория)

## Как начать

### Клонирование репозитория

Для начала склонируйте репозиторий на ваш пк:

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

## Для того бы протестировать без cdn, то достаточно убрать содержимое данной строки cdnHost = "storage.googleapis.com"

## По нагрузочным тестам около 60к вызовов
