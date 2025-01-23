# ADG Multi-Watchlist API

This project provides a robust API to manage multiple watchlists, allowing users to add, delete, and retrieve watchlists under a single API endpoint. Built with Go, Gin, and GORM, it integrates seamlessly with PostgreSQL to store and manage watchlist data efficiently.

---

## Features

- **Add** scripts to multiple watchlists in one request.
- **Delete** scripts from multiple watchlists.
- **Get** scripts from specific watchlists or watchlists associated with a script.
- Single endpoint to handle all actions: `ADD`, `DELETE`, and `GET`.

---

## Project Structure

### Key Directories and Files

- **main.go**: Initializes the database and starts the server.
- **routes/AdgWatchListRouter.go**: Sets up routes and binds handlers to endpoints.
- **handler/Handler.go**: Processes API requests and sends responses.
- **business/AdgWatchListService.go**: Contains business logic for managing watchlists.
- **repositories/WatchlistRepository.go**: Handles data operations with the database.
- **database/WatchlistDatabase.go**: Initializes the PostgreSQL database and defines migrations.
- **models/**: Defines the database schema and request/response models.

---

## Prerequisites

1. **Go** (>= 1.18)
2. **PostgreSQL**
3. **Postman** for testing the API

---

## Setup and Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/SahilWadhwani/ADG-To-MultiWatchList.git
   cd ADG-To-MultiWatchList

2. Install dependencies:

    ```bash 
    go mod tidy

3. Configure the database:
   - Update the connection string in src/database/WatchlistDatabase.go:
 
   

    ```go
   dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"


  Replace with your PostgreSQL credentials.

4. Run the application:
    ```bash
    go run main.go
   ```


The server will start at http://localhost:8080.

---

## API Endpoints

### Base URL
    http://localhost:8080

### Endpoint
#### POST ```/watchlist/multi-adg```

#### Request Body
```json
{
  "action": "ADD | DELETE | GET",
  "scriptId": 1,
  "watchListId": [1, 2],
  "userId": 123
}
```                          

#### Responses
- 200 OK: Operation successful.
- 400 Bad Request: Invalid request.
- 404 Not Found: Script or watchlist not found.
- 500 Internal Server Error: Operation failed.

---

### Testing with Postman 

1. Open Postman and create a new request.

2. Set the request method to POST.

3. Use the URL:

```bash 
http://localhost:8080/watchlist/multi-adg
```

4. Go to the Body tab and select raw. Set the format to JSON.

5. Enter a request body based on the action you want to test, e.g.,:

**Add Script**
```josn
{
  "action": "ADD",
  "scriptId": 1,
  "watchListId": [1, 2],
  "userId": 123
}
```

**Delete Script**

```json
{
  "action": "DELETE",
  "scriptId": 1,
  "watchListId": [1, 2],
  "userId": 123
}
```

**Get Scripts***
```json
{
  "action": "GET",
  "scriptId": 1,
  "watchListId": [1, 2],
  "userId": 123
}
```

6. Send the request and verify the response.

---

### Future Enhancements 

- Implement authentication and authorization.
- Add comprehensive unit tests.
- Enhance error handling.
- Support bulk operations for better performance.

---
