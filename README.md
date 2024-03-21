# BookStore-Management Application

# Technologies Used -
•	Gorilla/mux
•	CRUD operations
•	GORM
•	Middleware Session Management
•	Docker 
•	Redis 

# Working –
•	The application uses session management as a middleware.<br>
•	Only the routes login and logout are middleware free and can be accessed by anyone.
•	Other routes are checked for valid sessions which are only valid for certain time.  
•	Gorilla mux is used to create routes for CRUD operations. 
•	GORM model is used to structure the data in the database perform CRUD operations in the database. 
•	Docker is used to containerize the application. 
•	Docker compose file is used to implement multiple containers like Redis and my application go-bookstore.
•	Redis is used as container. 
Redis stores the get all books data as cache memory to reduce DB call load 
