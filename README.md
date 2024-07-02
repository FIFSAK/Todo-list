# Proxy Server Documentation

## Routers

### Health Check
- **Endpoint:** `GET /health-check`
  - **Response:** `OK`

### Make Request
- **Endpoint:** `POST /`
  - **Body:**
    ```json
    {
      "method": "GET",
      "url": "http://google.com",
      "headers": {
        "Authentication": "Basic bG9naW46cGFzc3dvcmQ=",

      }
    }
    ```

### Get Response
- **Endpoint:** `GET /?id={response_id}`
  - **Response:**
    ```json
    {
      "id": "response_id",
      "status": "HTTP-статус ответа стороннего сервиса",
      "headers": { "массив заголовков из ответа стороннего сервиса" },
      "length": "длина содержимого ответа"
    }
    ```

## Models Structure

```sql
Request {
  method  string,
  url     string,
  headers json
}

Response {
  id      INTEGER,
  status  INTEGER,
  headers JSON,
  length  INTEGER
}
```

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/FIFSAK/Hl-task1
   cd Hl-task1
   ```
2. **Build the Docker images:**
   ```bash
   make build
   ```
3. **Start the Docker containers:**
   ```bash
   make up
   ```
4. **Check the health of the server:**
   Open your browser and go to http://localhost:8080/health-check to ensure the server is running properly.
   
6. **Stop the Docker containers:**
   ```bash
   make down
   ```
**LINK: https://hl-task1.onrender.com/**
