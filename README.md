# Prime Number Generator

A simple Golang-based API service to generate prime numbers using different algorithms.

## Features

- Generate prime numbers in a given range
- Supports multiple algorithms:
  - Brute Force
  - Sieve of Eratosthenes
- REST API using Gin
- CLI support
- Execution time tracking
- Clean project structure using Strategy Pattern

---

## Run Project

Install dependencies:

```bash
go mod tidy
```

---

## Run Server

Start REST API server:

```bash
go run cmd/server/main.go
```

Server runs on:

```text
http://localhost:8080
```

Example API:

```http
GET /primes?start=1&end=100&algo=sieve
```

---

## Run CLI

Run prime generator from terminal:

```bash
go run cmd/cli/main.go --start=1 --end=100 --algo=brute
```

Example Output:

```text
Primes: [2 3 5 7 11 13 17 19]
```

---

## Sample API Response

```json
{
  "status": "success",
  "status_code": 200,
  "message": "Prime numbers fetched successfully",
  "data": {
    "algorithm": "sieve",
    "count": 25,
    "primes": [2,3,5,7,11],
    "time_us": 5
  }
}
```

---

## Algorithms

| Algorithm | Complexity |
|---|---|
| Brute Force | O(n√n) |
| Sieve | O(n log log n) |

---

## Tech Stack

- Golang
- Gin
- REST API
- Strategy Pattern

---

## Author

Jeel Jain