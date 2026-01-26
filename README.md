# Система управления аптекой

## Демонстрация

Демонстрация недоступна.

## Описание

Комплексный REST API для управления платформой онлайн-аптеки, разработанный на Go и PostgreSQL. Система обрабатывает каталоги продуктов, учетные записи пользователей, корзины покупок, заказы, платежи и отзывы клиентов с масштабируемой, чистой архитектурой.

## Технологический стек

**Язык программирования**
- Go 1.25.0

**Фреймворк бэкенда**
- Gin (веб-фреймворк HTTP для создания REST API)

**База данных**
- PostgreSQL
- GORM (ORM для операций с БД)
- pgx (драйвер PostgreSQL)

**Сборка и разработка**
- Go Modules (управление зависимостями)

## Скриншоты

> TODO: Добавьте скриншоты документации API или примеры ответов API. Рассмотрите добавление:
> - Документация Swagger/OpenAPI
> - Примеры структур JSON запросов и ответов
> - Диаграмма схемы базы данных
> - Диаграмма обзора архитектуры

## Project Structure

```
.
├── cmd/
│   └── pharmacy/
│       └── main.go                 # Точка входа приложения
├── internal/
│   ├── config/
│   │   └── database.go             # Подключение и конфигурация БД
│   ├── models/                     # Доменные сущности
│   │   ├── category.go
│   │   ├── subcategory.go
│   │   ├── user.go
│   │   ├── pharmacy.go             # Товары
│   │   ├── cart.go
│   │   ├── orders.go
│   │   ├── payments.go
│   │   ├── rewievs.go              # Отзывы
│   │   └── promocode.go
│   ├── repository/                 # Слой доступа к данным
│   │   ├── user_repository.go
│   │   ├── pharmacy_repository.go
│   │   ├── category_repository.go
│   │   ├── subcategory_repository.go
│   │   ├── cart-repository.go
│   │   ├── order_repository.go
│   │   ├── payment_repository.go
│   │   ├── rewievs-repository.go
│   │   └── promocode-repository.go
│   ├── service/                    # Слой бизнес-логики
│   │   ├── user-service.go
│   │   ├── pharmacy_service.go
│   │   ├── catergory_service.go
│   │   ├── subcategory_service.go
│   │   ├── cart-service.go
│   │   ├── order_service.go
│   │   ├── payment_service.go
│   │   ├── rewievs-service.go
│   │   └── promocode-service.go
│   └── transport/                  # HTTP обработчики и маршруты
│       ├── user-handler.go
│       ├── pharmacy_handler.go
│       ├── category_handler.go
│       ├── subcategory_handler.go
│       ├── cart-handler.go
│       ├── order_hendler.go
│       ├── payment_hendler.go
│       ├── reviews_handler.go
│       ├── promocode-handler.go
│       └── routes.go               # Регистрация маршрутов
├── go.mod                          # Определение Go модуля
├── go.sum                          # Контрольные суммы зависимостей
└── .env.example                    # Шаблон переменных окружения

```

## Архитектура

Проект реализует паттерн **чистой архитектуры** с трехслойным дизайном:

1. **Слой представления** (`internal/transport/`) - HTTP обработчики и управление запросами/ответами
2. **Слой услуг** (`internal/service/`) - бизнес-логика, валидация и оркестрация
3. **Слой репозитория** (`internal/repository/`) - доступ к БД и сохранение данных

Такое разделение обеспечивает слабую связанность, легкое тестирование и поддерживаемость.

## Основные возможности

- **Управление пользователями** - создание, чтение, обновление и удаление профилей пользователей
- **Каталог товаров** - управление товарами аптеки с иерархической категоризацией
- **Управление категориями** - организация товаров по категориям и подкатегориям
- **Корзина покупок** - добавление, обновление и управление товарами в корзине
- **Обработка заказов** - создание заказов из товаров корзины и отслеживание статуса
- **Обработка платежей** - обработка платежей с поддержкой нескольких методов
- **Отзывы клиентов** - пользователи могут оставлять рейтинги и отзывы о товарах
- **Промокоды** - применение кодов скидок к заказам (фиксированные или процентные)

## API Endpoints

### Пользователи
- `POST /users` - Регистрация нового пользователя
- `GET /users` - Получить всех пользователей
- `GET /users/:id` - Получить информацию о пользователе
- `PATCH /users/:id` - Обновить информацию пользователя
- `DELETE /users/:id` - Удалить учетную запись пользователя
- `GET /users/:id/cart` - Получить корзину пользователя
- `GET /users/:id/orders` - Получить историю заказов пользователя

### Товары (Аптека)
- `POST /pharmacy` - Добавить новый товар
- `GET /pharmacy` - Получить все товары
- `GET /pharmacy/:id` - Получить информацию о товаре
- `PATCH /pharmacy/:id` - Обновить информацию о товаре
- `DELETE /pharmacy/:id` - Удалить товар

### Категории и подкатегории
- `POST /category` - Создать новую категорию
- `GET /category` - Получить все категории
- `GET /category/:id/subcategory` - Получить подкатегории по категории
- `POST /category/:id/subcategory` - Добавить подкатегорию в категорию

### Корзина покупок
- `POST /users/:id/cart/items` - Добавить товар в корзину
- `PATCH /users/:id/cart/items/:item_id` - Обновить количество товара
- `DELETE /users/:id/cart/items/:item_id` - Удалить товар из корзины
- `DELETE /users/:id/cart` - Очистить корзину

### Заказы
- `POST /users/:id/orders` - Создать новый заказ из корзины
- `GET /orders/:id` - Получить информацию о заказе
- `GET /users/:id/orders` - Получить заказы пользователя
- `PATCH /orders/:id/status` - Обновить статус заказа

### Платежи
- `POST /orders/:id/payments` - Создать платеж для заказа
- `GET /orders/:id/payments` - Получить платежи для заказа
- `GET /payments/:id` - Получить информацию о платеже

### Отзывы
- `POST /pharmacy/:id/reviews` - Создать отзыв о товаре
- `GET /pharmacy/:id/reviews` - Получить отзывы о товаре
- `PATCH /reviews/:id` - Обновить отзыв
- `DELETE /reviews/:id` - Удалить отзыв

### Промокоды
- `POST /promocodes` - Создать новый промокод
- `GET /promocodes` - Получить все промокоды
- `GET /promocodes/:id` - Получить информацию о промокоде
- `PUT /promocodes/:id` - Обновить промокод
- `DELETE /promocodes/:id` - Удалить промокод

## Локальная установка и запуск

### Требования

- Go 1.25.0 или выше
- PostgreSQL 12 или выше
- Git

### Шаги установки

1. **Клонируйте репозиторий**
   ```bash
   git clone https://github.com/kuduzow/team-3-pharmacy.git
   cd team-3-pharmacy-1
   ```

2. **Установите зависимости Go**
   ```bash
   go mod download
   ```

3. **Настройте переменные окружения**
   ```bash
   cp .env.example .env
   ```

   Отредактируйте `.env` с вашими данными PostgreSQL:
   ```bash
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=ваш_пароль
   DB_NAME=pharmacy
   DB_SSLMODE=disable
   PORT=8888
   ```

4. **Создайте базу данных PostgreSQL**
   ```bash
   createdb pharmacy
   ```

5. **Запустите приложение**
   ```bash
   go run cmd/pharmacy/main.go
   ```

   API сервер запустится на `http://localhost:8888`

### Рабочий процесс разработки

```bash
# Запуск с автоматической перезагрузкой (требует инструмента 'air')
air

# Запуск тестов
go test ./...

# Сборка исполняемого файла
go build -o pharmacy cmd/pharmacy/main.go
```

## Схема базы данных

Приложение использует PostgreSQL с автоматическим созданием схемы через GORM миграции. Основные сущности:

- **users** - Учетные записи и профили пользователей
- **categories** - Категории товаров
- **sub_categories** - Вложенные подкатегории товаров
- **pharmacies** - Инвентарь товаров
- **carts** - Данные корзины покупок
- **cart_items** - Товары в корзине
- **orders** - Заказы клиентов
- **order_items** - Товары в заказах
- **payments** - Транзакции платежей
- **reviews** - Отзывы и рейтинги клиентов
- **promocodes** - Коды скидок и промокодов

## Contributors

- [Alsiev Yusup](https://github.com/isaalsiev)
- [Strannik-chr](https://github.com/Strannik-chr)
- [kuruevimran17](https://github.com/kuruevimran17)
- [zaynalabuev](https://github.com/zaynalabuev)
- [tsuruevimran17](https://github.com/tsuruevimran17)

## Лицензия

Этот проект разработан как часть командного портфолио. Пожалуйста, см. файл LICENSE или свяжитесь с владельцем репозитория для получения информации о лицензии.

---

**Статус:** ✅ Готово к производству | **Последнее обновление:** январь 2026
