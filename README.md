## Project Structure

```plaintext
│   .env
│   docker-compose.yml
│   Dockerfile
│   go.mod
│   go.sum
│
├───cmd
│       main.go
│
├───config
│       config.go
│
├───internal
│   ├───api
│   │   └───handler
│   │           get_stats.go
│   │           increment_click.go
│   │           router.go
│   │
│   ├───database
│   │       postgres.go
│   │
│   ├───model
│   │   │   entities.go
│   │   │
│   │   └───api
│   │           model.go
│   │
│   ├───service
│   │       get_statistic.go
│   │       increment_click.go
│   │       service.go
│   │       worker.go
│   │
│   └───storage
│           banner.go
│           banner_click.go
│           memory.go
│           repository.go
│
├───migrations
│       000001_create_banner_clicks_table.down.sql
│       000001_create_banner_clicks_table.up.sql
│
└───seed
        seed.go
```

## Installation and Running

1. Clone the repository:
``` sh
git clone https://github.com/d3pesha/click-counter.git
cd click-counter
```

2. Install dependencies:
```sh
go mod download
```

3. Configure environment variables
```
POSTGRES_USER=your_user
POSTGRES_PASSWORD=your_password
POSTGRES_DB=your_db
POSTGRES_HOST=your_db_host
POSTGRES_PORT=5432
POSTGRES_SSLMODE=disable
```

4. Run the service with Docker Compose
```sh
docker-compose up --build
```

## API 

- Register click to banner
```http
GET /counter/<bannerID>
```

- Get banner's stat
```http
POST /stats/{bannerID}
Content-Type: application/json

{
    "from": "2025-03-27T10:00:00Z",
    "to": "2025-03-30T17:00:00Z"
}
```