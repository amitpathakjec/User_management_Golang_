# User Management System

This is a simple User Management System built with Go (Golang). It supports basic CRUD operations for managing users, with a PostgreSQL database for persistence. The application is containerized with Docker for easy deployment.

## Features

- Create, Read, Update, and Delete (CRUD) user records.
- RESTful API endpoints.
- PostgreSQL integration.
- Docker support for seamless setup and deployment.

## Prerequisites

- Go (version 1.18 or higher)
- Docker
- PostgreSQL database

## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/<your-username>/<your-repo>.git
cd <your-repo>
```

### 2. Configure Environment Variables

Create a `.env` file by copying the `.env.example`:

```bash
cp .env.example .env
```

Update the `.env` file with your database credentials.

### 3. Run Locally Without Docker

1. Install dependencies:
   ```bash
   go mod tidy
   ```
2. Start the application:
   ```bash
   go run main.go
   ```
3. Access the server at `http://localhost:8080`.

### 4. Run with Docker

1. Build the Docker image:
   ```bash
   docker build -t user-management-app .
   ```
2. Run the container:
   ```bash
   docker run -p 8080:8080 --env-file .env user-management-app
   ```
3. Or use Docker Compose:
   ```bash
   docker-compose up --build
   ```

## API Endpoints

### Base URL

`http://localhost:8080`

### Create User

**POST** `/users/create`

```json
{
  "first_name": "John",
  "last_name": "Doe",
  "email": "john.doe@example.com",
  "phone_number": "1234567890",
  "account_type": "savings",
  "initial_balance": 1000
}
```

### Get User

**GET** `/users/get?id={id}`

### Update User

**PUT** `/users/update`

```json
{
  "id": 1,
  "first_name": "Rachel",
  "last_name": "Stewards",
  "email": "r.stewards@example.com",
  "phone_number": "123456789",
  "account_type": "current"
}
```

### Delete User

**DELETE** `/users/delete?id={id}`

### List Users

**GET** `/users/list`

## Project Structure

```
├── controllers/       # Handlers for API endpoints
├── db/                # Database connection logic
├── models/            # Data structures and request/response models
├── main.go            # Application entry point
├── Dockerfile         # Docker build file
├── docker-compose.yml # Docker Compose configuration
├── .env.example       # Sample environment variables
└── README.md          # Project documentation
```

## License

--
