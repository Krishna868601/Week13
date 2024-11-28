# Go API for Current Toronto Time with MySQL Logging

This project implements a simple API in Go that returns the current time in Toronto in JSON format. Each request made to the API logs the current time into a MySQL database. The application uses the `time` package for time zone conversion and MySQL for data persistence.

## Features

- **API Endpoint**: `/current-time` that returns the current time in Toronto.
- **Database Logging**: Each API request logs the current time into a MySQL database.
- **Time Zone Handling**: Automatically adjusts to Toronto's local time (`America/Toronto`).
- **MySQL Integration**: Logs time to a MySQL database running locally or in a Docker container.

## Prerequisites

To run this project, you will need:

- **Go** (Golang) installed: [Install Go](https://golang.org/doc/install)


## Setup and Installation

### 1. Clone the repository
```bash
git remote add origin https://github.com/Krishna868601/Week13.git
cd Week13
```
### 2. Set up MySQL Database

#### a. Using MySQL locally:

1. **Install MySQL**: 
   - Download and install MySQL from [here](https://dev.mysql.com/doc/refman/8.0/en/installing.html).
   - Start the MySQL service on your machine.

2. **Connect to MySQL**:
   - Open **MySQL Workbench** or open a command line terminal (e.g., Command Prompt or PowerShell).
   - Use the following command to connect to MySQL as the root user:

     ```bash
     mysql -u root -p
     ```

   - Enter your MySQL root password when prompted.

3. **Create a new database**:
   - After successfully logging in, create a new database named `toronto_time`:

     ```sql
     CREATE DATABASE toronto_time;
     USE toronto_time;
     ```

4. **Create the `time_log` table**:
   - Create a table named `time_log` that will store the timestamp of each API request:

     ```sql
     CREATE TABLE IF NOT EXISTS time_log (
         id INT AUTO_INCREMENT PRIMARY KEY,
         timestamp DATETIME NOT NULL
     );
     ```

   - This will create a table with two columns:
     - `id`: An auto-incrementing integer to serve as the primary key.
     - `timestamp`: A `DATETIME` field that stores the timestamp of each logged request.

### 5. Run the Application

To start the Go application, run the following command:

```bash
go run main.go
```
The API will start and listen on port 8090

### 6. Test the API
You can now test the API by making a GET request to:

```bash
http://localhost:8090/current-time
```
This will return the current time in Toronto in JSON format,
Each request will also log the current time into the time_log table in MySQL.

![image](https://github.com/user-attachments/assets/382ee830-6237-4bed-8599-0642f8bb2c3a)

### Database Query Output:
![image](https://github.com/user-attachments/assets/626b360c-589a-4564-a49c-e1fd65c34cbb)

