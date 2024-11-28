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
- **MySQL** installed (locally or using Docker): [Install MySQL](https://dev.mysql.com/doc/refman/8.0/en/installing.html)

### For MySQL with Docker (optional):
- Docker installed: [Install Docker](https://docs.docker.com/get-docker/)

## Setup and Installation

### 1. Clone the repository
```bash
git remote add origin https://github.com/Krishna868601/Week13.git
cd Week13
