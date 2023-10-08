Welcome to the Ultraviolet Coding Test Project - completed by Don Jackson

Instructions:
1) Create a folder and unzip the file contents into that folder

2) Open a terminal and navigate to the project directory.

3) Make sure Docker is running on your machine

4) Start the services by running: 
docker-compose up

*** Optional ***
If there are problems, here are some hints:
- In another terminal, you may make sure both services (ultraviolettest-app AND mysql:latest) are running: 
docker ps
- Make sure Docker is running. You may need to stop and restart it
- Sometimes you may want to delete all containers in Docker and start from scratch. If any containiners won't delete, run these first:
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
You can build the application separatlely by running:
docker build -t myapp .
- You may see the program image by running
docker images
*****************

5) Test the application using Postman or curl. The application is reachable on localhost:8080. Here are the available endpoints:
a) Retrieve a list of books:
   - GET
   - URL: http://localhost:8080/books
   curl --location 'localhost:8080/books'

b) Retrieve a specific book by ID (replace {id} with an actual bookID):
   - GET
   - URL: http://localhost:8080/books/{id}
   curl --location 'localhost:8080/books/1'

c) Add a book:
   - POST
   - URL: http://localhost:8080/books
   - BODY
      {
         "title": "book name here",
         "author": "author name here",
         "PublicationDate": "YYYY-MM-DD",
         "ISBN": "text here",
         "PageCount": 123
      }
   curl --location 'localhost:8080/books' \
   --header 'Content-Type: application/json' \
   --data '{
      "title": "The Cat in the Hat",
      "author": "Dr. Suese",
      "PublicationDate": "2023-03-04",
      "ISBN": "12344",
      "PageCount": 12
   }
   '

d) Update a book: (replace {id} with an actual bookID):
   - PUT
   - URL: http://localhost:8080/books/{id}
   - BODY
      {
         "title": "book name here",
         "author": "author name here",
         "PublicationDate": "YYYY-MM-DD",
         "ISBN": "text here",
         "PageCount": 123
      }
      curl --location --request PUT 'localhost:8080/books/1' \
      --header 'Content-Type: application/json' \
      --data '{
         "title": "book name - updated",
         "author": "Mark Twain",
         "PublicationDate": "2023-03-04",
         "ISBN": "12344",
         "PageCount": 12
      }
      '

e) Delete a book: (replace {id} with an actual bookID):
   - DELETE
   - URL: http://localhost:8080/books/{id}
   curl --location --request DELETE 'localhost:8080/books/1' \
      --data ''


Please let me know if you have any questions or encounter any issues during setup or testing.
Don Jackson
(408) 858-6490
donj91711@yahoo.com
