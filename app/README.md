<h3 align="center">
  Backend application
</h3>

## About
This app have a CLI interface and REST endpoints.

## Architecture
 * This is an attempt to implement a clean architecture, that can be founded here https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

## Setup Environment
 * Create a database with a name: rentcar
 * Set environment variables: 
 ```
  export DATABASE_HOST=localhost
  export DATABASE_PORT=5432
  export DATABASE_NAME=rentcar
  export DATABASE_USER=postgres
  export DATABASE_PASSWORD=postgres
  export WEB_PORT=3000
 ```

## Running
 ```
  cd infrastructure/cmd/server/
  go run main.go
 ```