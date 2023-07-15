# Delos Task Assignment

## Getting Started

Before running the service, make sure you have the following prerequisites installed:

- Docker (https://docs.docker.com/get-docker/)

Or

- Go (https://go.dev/learn/)
- MySQL (https://www.mysql.com/)

## Installation (Required If Not Using Docker)

1. Install the required dependencies
```
go mod download
```
2. Update the `config.json` file with your HTTP port and database connection details
3. Run the `db.sql` in `database` directory to your MySQL database to set up the initial database tables

## Usage

If using Docker, run the following command
```
docker-compose up
```
Otherwise, run the following command
```
go run main.go
```

## API Addresses
You can access it via `http://localhost:{http_port}`. Replace the `http_port` with your HTTP port in `config.json`

## API Endpoints

-  Farm
    - `/api/farm` (POST): Create Farm

        payload: {
            "name": string
        }

    - `/api/farm` (GET): Get Farm

    - `/api/farm/:id` (GET): Get Farm By Id

    - `/api/farm/:id` (PUT): Update Farm

        payload: {
            "name": string
        }

    - `/api/farm/:id` (DELETE): Delete Farm

-  Pond
    - `/api/pond` (POST): Create Pond

        payload: {
            "name": string,
            "farm_id": int
        }

    - `/api/pond` (GET): Get Pond

    - `/api/pond/:id` (GET): Get Pond By Id

    - `/api/pond/:id` (PUT): Update Pond

        payload: {
            "name": string,
            "farm_id": int
        }

    - `/api/pond/:id` (DELETE): Delete Pond

-  Statistics
    - `/api/statistics` (GET): Get Statistics

## API Documentation

https://api.postman.com/collections/21473149-1af2d273-e316-45a5-866a-9c00fc60a45f?access_key=PMAT-01H5DMHWCEC3FRB3EGMD85A0NX