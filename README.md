# CLI-Go-Web-Backend

A RESTful backend API written in **Go** using the **Gin** framework — designed to power a CLI-style portfolio frontend and other clients.  
It exposes endpoints to retrieve commands, profile info, skills, showcase projects, work experience, education, and contacts.

## 🧠 Project Overview

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

## 🚀 Features

✔ REST API with Gin  
✔ PostgreSQL data persistence  
✔ Migration support via `migrate`  
✔ Structured domain-repository-service layers  
✔ CORS support for frontend integration

---

## 🛠️ Tech Stack

| Component | Technology |
|-----------|------------|
| Language | Go (Golang) |
| Web Framework | Gin |
| Database | PostgreSQL |
| Config | Viper |
| Migrations | go-migrate |
| Hosting (suggested) | Railway / Cloud Run / Heroku |

---
