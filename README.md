# Find The Students Average (With A Database)

A simple program that creates a PostgreSQL database and a class room table that stores the students name, id and average grade, as well as having CRUD operations.

# Purpose

The main purpose of this project was to really learn more about how to use a database with Go in a simple environment and building on top of my previous project. As well to tinker with
the database/SQL library and not using a ORM. Overall its a super simple project that just creates a database and a table, which then allows the user to operate crud operations with the data. Nothing too fancy just wanted to experiment with simple back-end logic, mixed in with a database queries.  

# Install 
Just clone the repo and run go mod tidy, after that go into the starter.go file and change the default const to your credentials.
Will need a running PostgreSQL instance running, after all that you can compile and run it.

# What I learned

- SQL queries and fundamental understanding of PostgreSQL and in general SQL based concepts.
- Using the database/SQL library from the go standard library.
- Creating a database with a table using Go.
- Creating functions that interact with the database, and implementing backed logic with it.
- Creating tests that require user inputs and being able to mock them.
- Reinforced usage of pointers when creating a student data structure.
- Merging a existing project to this project.
- Creating a terminal interface that interacts with a database, while allowing CRUD based operations.

# What I struggled with

- Creating code that does not repeat for when opening the database.
- Usage of the proper functions from the database/SQL library (Like when to open and close a database, as well proper parameters for queries).
- Creating test for the SQL based functions (Shot my self in the foot in the way I designed the SQL functions. Next project will not have the same design).

# Credit

- Gave me inspirations for when creating test for user input based functions https://stackoverflow.com/questions/17456901/how-to-write-tests-against-user-input-in-go
