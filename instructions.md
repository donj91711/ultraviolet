# GoLang Tech Challenge: Build a RESTful API with GORM and Docker

## Objective:

Your task is to develop a simple RESTful API using Golang to manage a collection of books. This system should use the GORM ORM to interact with a database and you should containerize your solution using Docker and Docker Compose.

You may use any routing, web framework, or native library you are comfortable with, the only requirement is that you use GORM and Docker.

## Requirements:

1. **API Endpoints**:

   - `GET /books`: Retrieve a list of all books.
   - `GET /books/{id}`: Retrieve details of a specific book by its ID.
   - `POST /books`: Add a new book to the collection.
   - `PUT /books/{id}`: Update details of a specific book by its ID.
   - `DELETE /books/{id}`: Remove a book by its ID.

2. **Book Entity**:

   - ID (Unique identifier)
   - Title
   - Author
   - PublishedDate
   - ISBN
   - PageCount

3. **Database**:

   - Use any relational database of your choice (e.g., PostgreSQL, MySQL, SQLite).
   - Set up GORM to handle database operations.

4. **Docker**:

   - Write a `Dockerfile` for your application.
   - Use `docker-compose` to set up the application and the database as separate services (this will not apply if you choose to use SQLite).

5. **Bonus**:
   - Add a middleware for logging incoming requests.
   - Implement migrations using a tool such as Atlas to create the necessary tables.
   - Use the repository pattern to implement GORM-specific operations
   - Include unit tests for your endpoints.

## Deliverables:

1. Source code of your Go application.
2. The database schema and migrations as necessary.
3. Dockerfile and docker-compose.yml.
4. A README file with instructions on how to run and test your solution.

## Evaluation Criteria:

1. Functionality: The application should work as described.
2. Code Quality: Your code should be well-organized and readable.
3. Database Interaction: Proper usage of GORM.
4. Docker Setup: Ability to simply run `docker-compose up` to get the whole system running.
5. Bonus Points: Any additional features or improvements beyond the basic requirements.