# Find The Students Average (With A Database)

A simple program that creates a PostgreSQL database and a class room table that stores the students name, id and average grade, as well as having CRUD operations.

# What I learned

- SQL queries and fundamental understanding of PostgreSQL and in general SQL based concepts.
- Using the database/SQL library from the go standard library.
- Creating a database with a table using Go.
- Creating functions that interact with the database, and implementing backed logic with it.
- Creating tests that require user inputs and being able to mock them.
- Reinforced usage of pointers when creating a student data structure.
- Merging a existing project to this project.
- Creating a terminal interface that interacts with an database, while allowing CRUD based operations.

# What I struggled with

- Creating code that does not repeat for when opening the database.
- Usage of the proper functions from the database/SQL library (Like when to open and close a database, as well proper parameters for queries).
- Creating test for the SQL based functions.

# Credit

- Gave me inspirations for when creating test for user input based functions https://stackoverflow.com/questions/17456901/how-to-write-tests-against-user-input-in-go

# TODO

- [] Using Docker for this project, so anyone can download it without needing to download and configure PostgreSQL.
