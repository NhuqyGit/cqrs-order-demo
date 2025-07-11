# CQRS Order Demo with Golang, RabbitMQ, PostgreSQL, and MongoDB

<img width="3780" height="1890" alt="cqrs" src="https://github.com/user-attachments/assets/7991e4ee-5c6c-46ed-a109-ac52665d5770" />

A simple demonstration of the **CQRS (Command Query Responsibility Segregation)** architectural pattern using:

- **Go** for both `cmd-service` and `query-service`
- **PostgreSQL** as the write database
- **MongoDB** as the read database
- **RabbitMQ** as the event broker

---

## What is CQRS?

**CQRS (Command Query Responsibility Segregation)** is a design pattern that separates read and write operations into distinct models:

- **Commands**: Create, update, and delete operations (POST, PUT, DELETE)
- **Queries**: Read-only operations (GET)

In a CQRS-based system:
- The **write side** stores normalized, transactional data (typically using a relational database like PostgreSQL)
- The **read side** stores denormalized, query-optimized data (commonly with a NoSQL database like MongoDB)

---

## CQRS is useful when:

- Your system has **heavy read/write imbalance** (e.g., read-heavy systems)
- You need **performance-optimized queries** separate from the write model
- You want to **scale read and write operations independently**
- Especially useful for **read-heavy applications**, where the number of queries far exceeds the number of writes

---

### Services Overview

| Service         | Description                        | Port |
|----------------|------------------------------------|------|
| `cmd-service`   | Handles write operations (POST/PUT/DELETE) | 8080 |
| `query-service` | Handles read operations (GET)       | 8081 |
| PostgreSQL      | Stores write data                  | 5432 |
| MongoDB         | Stores read data                   | 27018 |
| RabbitMQ        | Message broker                     | 5671 (AMQP), 15671 (UI) |

---

### Running with Docker

```bash
docker-compose up --build
```

### Create a product
```bash
POST http://localhost:8080/api/products
{
  "name": "Product A",
  "description": "Product Test",
  "price": 120,
  "quantity": 3,
  "sku": "SKU-DEMO-0024"
}
```

### Get products
```bash
GET http://localhost:8081/api/products
```

### Access mongodb to test
```bash
mongosh mongodb://localhost:27018
```

### Project structure
```bash
├───cmd-service
│   ├───db
│   ├───event
│   │   └───publisher
│   ├───handler
│   ├───models
│   ├───repository
│   ├───routers
│   └───service
└───query-service
    ├───db
    ├───event
    │   └───consumer
    ├───handler
    ├───models
    ├───repository
    ├───routers
    └───service
```
