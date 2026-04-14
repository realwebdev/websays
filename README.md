# ⏱️ Time Tracking RESTful JSON API

A backend service built in **Go** that allows multiple users to create projects, track time spent on them, and retrieve summarized reports.

This project demonstrates clean API design, layered architecture, and multiple persistence strategies under time constraints.

---

## 🚀 Overview

This is a **RESTful JSON API** for a time tracking application with authentication and project based time logging.

### Core Features

* User authentication using JWT
* Create and manage projects
* Track time spent on projects
* Sign in / sign out workflow
* Summary report of total time per project
* CRUD operations for additional entities with different persistence mechanisms

---

## 🧱 Tech Stack

* **Language:** Go
* **Framework:** Gin
* **Database:** PostgreSQL
* **ORM:** GORM
* **Authentication:** JWT
* **Configuration Management:** `.env` files
* **CI/CD:** GitHub Actions
* **Build Automation:** Makefile

---

## 📦 Architecture & Design

The application is structured with separation of concerns in mind:

* Clear handler / service / repository boundaries
* Minimal interlayer coupling
* Dependency isolation where possible
* Different persistence strategies per entity to demonstrate flexibility

The goal was to show modular design even under time constraints.

---

## 🔐 Authentication

* JWT-based authentication
* Protected routes require valid Bearer token
* Token issued at login

---

## 🗂 Entities & Persistence Strategies

### 1️⃣ Project & Time Tracking (PostgreSQL)

* Managed via GORM
* Supports:

  * Create project
  * Start tracking time
  * Stop tracking time
  * Generate summary report per project

### 2️⃣ Category (In-Memory Persistence)

* Full CRUD
* Data stored in application memory
* Demonstrates abstraction over storage layer

### 3️⃣ Article (Text File Persistence)

* Full CRUD
* Data stored in text files
* Demonstrates file based persistence strategy

---

## 🧪 Testing

* Includes example Go unit test
* Demonstrates testing structure
* Full test coverage not implemented due to time constraints

---

## ⚙️ Setup & Installation

### 1. Clone Repository

```bash
git clone https://github.com/your-username/your-repo.git
cd your-repo
```

### 2. Configure Environment

Create a `.env` file:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=timetracker
JWT_SECRET=your_secret_key
PORT=8080
```

### 3. Run Database

Ensure PostgreSQL is running locally or via Docker.

### 4. Run Application

```bash
make run
```

Or:

```bash
go run cmd/main.go
```

---

## 🛠 Makefile Commands

```bash
make run        # Run the application
make test       # Run tests
make build      # Build binary
```

---

## 🔄 CI/CD

GitHub Workflows configured for:

* Build validation
* Test execution

---

## 📊 Example API Endpoints

### Authentication

* `POST /register`
* `POST /login`

### Projects

* `POST /projects`
* `GET /projects`
* `POST /projects/:id/start`
* `POST /projects/:id/stop`
* `GET /projects/summary`

### Category (Memory)

* `POST /categories`
* `GET /categories`
* `PUT /categories/:id`
* `DELETE /categories/:id`

### Article (File-Based)

* `POST /articles`
* `GET /articles`
* `PUT /articles/:id`
* `DELETE /articles/:id`

---

## 📌 Design Considerations

* Demonstrates multiple persistence strategies in one service.
* Shows understanding of:

  * REST principles
  * JWT authentication
  * ORM integration
  * Dependency separation
  * Basic CI/CD
* Designed under time constraints while prioritizing clarity and isolation.

---

## ⚠️ Limitations

* Only partial unit test coverage
* No full integration tests
* No production-grade logging or observability
* Simplified error handling in some areas

This project focuses primarily on demonstrating backend fundamentals and architectural thinking.

---

## 🧠 Future Improvements

* Add full unit and integration test coverage
* Introduce structured logging
* Add Swagger/OpenAPI documentation
* Dockerize the entire setup
* Improve time tracking edge case handling
* Add rate limiting and request validation middleware

---

## 👨‍💻 Author

To demonstrate backend engineering fundamentals in Go.

