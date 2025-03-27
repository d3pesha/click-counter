``` sh
git clone https://github.com/d3pesha/click-counter.git
cd click-counter
```

Заполнить .env или переименовать .env.example

```
POSTGRES_USER=your_user
POSTGRES_PASSWORD=your_password
POSTGRES_DB=your_db
POSTGRES_HOST=your_db_host
POSTGRES_PORT=5432
POSTGRES_SSLMODE=disable
```

Запустить сервис в docker compose 

```sh
docker-compose up --build
```

