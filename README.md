# CLI-Go-Web-Backend

A RESTful backend API written in **Go** using the **Gin** framework - designed to power a CLI-style portfolio frontend in [Here](https://github.com/Niko-Cloud/Cli-Flutter-Web-Fe).  
It exposes endpoints to retrieve commands, profile info, skills, showcase projects, work experience, education, and contacts.

## Project Overview

This backend serves data for a CLI-style portfolio interface, exposing structured resources such as:

- `GET /api/commands` — list all CLI commands  
- `GET /api/commands/:name` — get help details for a specific command  
- `GET /api/profile` — developer profile information  
- `GET /api/skill` — technical skill list  
- `GET /api/showcase` — portfolio showcase list
- `GET /api/showcase/:id` — portfolio showcase list  
- `GET /api/education` — educational history  
- `GET /api/work` — work experience history
- `GET /api/work/:id` — work experience history
- `GET /api/contact` — contact links

It’s built with **Go**, organized following clean architecture principles, and ready for deployment.

---

## Features

✔ REST API with Gin  
✔ PostgreSQL data persistence  
✔ Migration support via `migrate`  
✔ Structured domain-repository-service layers  
✔ CORS support for frontend integration

---

## Tech Stack

| Component | Technology |
|-----------|------------|
| Language | Go (Golang) |
| Web Framework | Gin |
| Database | PostgreSQL |
| Config | Viper |
| Migrations | go-migrate |
| Hosting (suggested) | Railway / Cloud Run / Heroku |

---
## Getting Started

### Prerequisites

*   Go (version 1.24 or later)
*   PostgreSQL

### Installation & Setup

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/Niko-Cloud/CLI-Go-Web-Backend.git
    cd CLI-Go-Web-Backend
    ```

2.  **Install dependencies:**
    ```sh
    go mod tidy
    ```

3.  **Configure the application:**
    The application uses a `config/application-<env>.yaml` file for configuration. For local development, create `internal/config/application-dev.yaml`:

    ```yaml
    # internal/config/application-dev.yaml
    server:
      port: 8080

    database:
      uri: postgres://youruser:yourpassword@localhost:5432/yourdb?sslmode=disable
    ```
    Replace the `database.uri` with your PostgreSQL connection string.

4.  **Run the application:**
    ```sh
    go run cmd/api/main.go
    ```
    The server will start, automatically apply database migrations, and listen on the configured port (e.g., `http://localhost:8080`).

### Production Configuration

For production, set the `APP_ENV` environment variable to `prod`. The configuration will be loaded from `internal/config/application-prod.yaml`, which uses environment variables for sensitive data:

*   `APP_PORT`: The port for the server to run on.
*   `APP_DATABASE_URI`: The PostgreSQL database connection URL.

Example:
```sh
export APP_ENV=prod
export APP_PORT=8080
export APP_DATABASE_URI="postgres://user:pass@host:port/db?sslmode=require"
go run cmd/api/main.go
```
