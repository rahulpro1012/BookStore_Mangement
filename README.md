# BookStore-Management Application

<h1>To Run the Application</h1>
<ul>
  <li>Install Docker Desktop</li>
  <li>Run docker-compose up</li>
</ul>

# Technologies Used -
•	Gorilla/mux<br>
•	CRUD operations<br>
•	GORM<br>
•	Middleware Session Management<br>
•	Docker <br>
•	Redis <br>

# Working –
•	The application uses session management as a middleware.<br>
•	Only the routes login and logout are middleware free and can be accessed by anyone.<br>
•	Other routes are checked for valid sessions which are only valid for certain time. <br> 
•	Gorilla mux is used to create routes for CRUD operations. <br>
•	GORM model is used to structure the data in the database perform CRUD operations in the database. <br>
•	Docker is used to containerize the application. <br>
•	Docker compose file is used to implement multiple containers like Redis and my application go-bookstore.<br>
•	Redis is used as container and stores the get all books data as cache memory to reduce DB call load <br>
