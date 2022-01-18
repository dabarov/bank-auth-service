# Сервис авторизации 

## Спринты

### Спринт 1

- [x] Sign up - Registration - могут все Создание пользователя
- [x] Sign in - Login Получение токена (авторизация)
- [x] Валидация токена (через secret)
- [x] Подготовка инфраструктуры и технологии для реализации

### Спринт 2

- [x] Добавить в сервис авторизации endpoint для получения информации о клиенте. Пока что это только ИИН, login, дата регистрации.

### Спринт 3

- [x] ~~Изменить endpoint в сервисе авторизации для получения информации о юзере. Добавить информацию о кошельках со второго сервиса.~~

### Спринт 4

- [x] Обернуть приложение в Docker
- [x] Запускать оба сервиса через docker compose
- [x] Написать unit тесты для сервиса авторизации

## Использованные технологии

 - [PostgreSQL](https://www.postgresql.org/)
 - [Redis](https://redis.io/)
 - [Docker](https://www.docker.com/)
 - [GORM](https://gorm.io/)
 - [go-clean-arch](https://github.com/bxcodec/go-clean-arch)
 - [fasthttp](https://github.com/valyala/fasthttp)
 - [fasthttprouter](https://github.com/buaazp/fasthttprouter)
 - [sqlmock](https://github.com/DATA-DOG/go-sqlmock)
 - [redismock](https://github.com/go-redis/redismock)
 - [testify](https://github.com/stretchr/testify)