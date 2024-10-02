

# URL Shortening Service API

This project is a simple RESTful API for a URL shortening service. It allows users to shorten long URLs, retrieve original URLs from short codes, update or delete existing short URLs, and get statistics on the number of times a short URL has been accessed.

## Features

- Create a new short URL from a long URL.
- Retrieve the original URL from a short URL.
- Update an existing short URL.
- Delete an existing short URL.
- Get statistics for a short URL (e.g., access count).

## Tech Stack

- **Backend**: Go (Golang)
- **Database**: MongoDB (NoSQL)
- **Hashing**: MD5 for URL shortening
- **Framework**: Standard library with HTTP package
## API Endpoints

### 1. Create Short URL
**Method**: `POST /shorten`

**Request Body**:
```json
{
  "url": "https://www.example.com/some/long/url"
}
```

**Response**:
- **201 Created**:
```json
{
  "id": "1",
  "url": "https://www.example.com/some/long/url",
  "shortCode": "abc123",
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:00:00Z"
}
```

### 2. Retrieve Original URL
**Method**: `GET /shorten/{shortCode}`

**Response**:
- **200 OK**:
```json
{
  "id": "1",
  "url": "https://www.example.com/some/long/url",
  "shortCode": "abc123",
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:00:00Z"
}
```

### 3. Update Short URL
**Method**: `PUT /shorten/{shortCode}`

**Request Body**:
```json
{
  "url": "https://www.example.com/some/updated/url"
}
```

**Response**:
- **200 OK**:
```json
{
  "id": "1",
  "url": "https://www.example.com/some/updated/url",
  "shortCode": "abc123",
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:30:00Z"
}
```

### 4. Delete Short URL
**Method**: `DELETE /shorten/{shortCode}`

**Response**:
- **204 No Content**: If the short URL was successfully deleted.

### 5. Get URL Statistics
**Method**: `GET /shorten/{shortCode}/stats`

**Response**:
- **200 OK**:
```json
{
  "id": "1",
  "url": "https://www.example.com/some/long/url",
  "shortCode": "abc123",
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:00:00Z",
  "accessCount": 10
}
```

## Setup & Installation

### Prerequisites
- Go (tested on 1.23.1)
- MongoDB server.
- Git.

### Steps to run locally:

1. **Clone the repository**:
    ```bash
    git clone https://github.com/yourusername/url-shortening-service.git
    cd url-shortening-service
    ```

2. **Install dependencies** (Go modules should be automatically resolved):
    ```bash
    go mod tidy
    ```

3. **Configure MongoDB**:
    Ensure that MongoDB is running locally or remotely, and update the connection string in your `database` package.

4. **Run the application**:
    ```bash
    go run main.go
    ```

5. **Test the API**:
    You can test the endpoints using Postman, Curl, or any API client.
