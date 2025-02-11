# Glofox Studio Management

## Overview
This project implements an API for managing studio class creation and bookings. The API allows Studio owners to create classes and Studio members to book classes,


## Project Setup

### Prerequisites

Before setting up the project, ensure you have the following installed:

- Go (version 1.18 or later)

### Getting Started

1. **Clone the repository:**
    ```
    git clone https://github.com/saikumar-neelam/glofox_studio.git

    cd glofox_studio
    ```

2. **Initialize Go Modules:**
    ```
    If there are no go.mod and go.sum files, then initialize as per below statment

    go mod init github.com/saikumar-neelam/glofox_studio
    ```

3. **Install Dependencies:**
    ```
    Run below command to download and install all go dependencies

    go mod tidy
    ```

4. **Run the Application:**
    ```
    To run the application, use the following command:

    go run cmd/glofox/main.go
    ```

5. **Running Unit Tests::**
    ```
    To run unit tests, use the following command:

    go test ./...
    ```

## Folder Structure
- `cmd/glofox/`: Module entry point which has main
- `api/handlers`: HTTP handlers for Classes and Bookings
- `api/routers`: Routes
- `internal/structs/`: Structs representing entities (e.g., Class, Booking)
- `internal/processors/`: Business logic for managing classes and bookings

## Endpoints
### POST `/classes`
Book a class by providing member details and the class date.

Request body:
```json
{
    "class_name": "yoga",
    "start_date": "2025-02-13",
    "end_date": "2025-02-28",
    "capacity": 100
}
```

### POST `/bookings`
Book a class by providing class details, member details and the class date.

Request body:
```json
{
  "class_name": "yoga",
  "member_name": "Sai Kumar",
  "class_date": "2025-02-15"
}
```

### GET `/bookings/{bookingDate(YYYY-MM-DD)}`
Retrieve the number of bookings done on particular date for different classes.**(Optional Developed for testing)**

Request body:
None
