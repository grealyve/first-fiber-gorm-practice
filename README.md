# Fiber Book API
This project consists of a simple book API written in Go. The API can interact with a Postgresql database and handle books using HTTP POST, GET, and DELETE methods. It runs on Docker with a connected Postgresql DB.

## Installation
```powershell
docker-compose up -d

go mod tidy
go run main.go

 ┌───────────────────────────────────────────────────┐
 │                   Fiber v2.48.0                   │
 │               http://127.0.0.1:3000               │
 │       (bound on host 0.0.0.0 and port 3000)       │
 │                                                   │
 │ Handlers ............. 6  Processes ........... 1 │
 │ Prefork ....... Disabled  PID ............. 00000 │
 └───────────────────────────────────────────────────┘
```

1. The API will be available at http://localhost:3000 by default.
2. The API can be accessed using the specified HTTP methods:
- To add a new book: POST http://localhost:3000/api/new_books
- To get all books: GET http://localhost:3000/api/books
- To get a specific book: GET http://localhost:3000/api/get_book/{book_id}
- To delete a book: DELETE http://localhost:3000/api/delete_book/{book_id}

## API Documentation
### Adding a New Book - POST /api/new_books
With this request, you can add a new book to the API. The request body should be in JSON format and include the following properties:
```json
{
  "author": "Author Name",
  "title": "Book Title",
  "publisher": "Publisher" 
}
```
![image](https://github.com/grealyve/first-fiber-gorm-practice/assets/41903311/d34360b4-3e6e-4faf-b573-809e444c8acc)



### Getting All Books - GET /api/books
With this request, you can retrieve all books from the API. The result will be in JSON format:
```json
[
  {
    "id": 1,
    "author": "Author 1",
    "title": "Book 1",
    "publisher": "Publisher 1" 
  },
  {
    "id": 2,
    "author": "Author 2",
    "title": "Book 2",
    "publisher": "Publisher 2" 
  }
]
```
![image](https://github.com/grealyve/first-fiber-gorm-practice/assets/41903311/d152107c-c038-495b-9961-b7172e11eb4c)

### Getting a Specific Book - GET /api/get_book/{book_id}
With this request, you can retrieve a specific book from the API. You need to include the {book_id} parameter in the request.
![image](https://github.com/grealyve/first-fiber-gorm-practice/assets/41903311/aa37afeb-f97d-4492-957d-4142ee04ec08)

### Deleting a Book - DELETE /api/delete_book/{book_id}
With this request, you can delete a specific book from the API. You need to include the {book_id} parameter in the request.
![image](https://github.com/grealyve/first-fiber-gorm-practice/assets/41903311/b055a35c-ee50-4a7d-a9dc-7ffc3f4af1ec)


### Referrence:
https://youtu.be/1XPktts9USg
