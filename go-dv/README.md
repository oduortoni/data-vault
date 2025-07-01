# Go Data Vault (go-dv)

This is the Go implementation of the Data Vault application. It is a web server that provides APIs for user registration, authentication, and data management.

## Setup and Running

1.  **Prerequisites**:
    *   Go 1.18+

2.  **Navigate to the project directory**:
    ```bash
    cd go-dv
    ```

3.  **Configuration**:
    The application is configured via environment variables. You can create a `.env` file in this directory and use a tool like `godotenv` or export them directly.

    ```plaintext
    # Server configuration
    PORT=9000
    HOST="0.0.0.0"

    # Database configuration
    DATABASE_DSN="database.sqlite"

    # Auth configuration (use a strong, randomly generated secret)
    JWT_SECRET="your-super-strong-and-long-jwt-secret"
    ```

4.  **Install Dependencies**:
    ```bash
    go mod tidy
    ```

5.  **Run the application**:

You can go back to the root directory and run:

    ```bash
    cd ..
    make run-go
    ```

The server will start, and you can access it at `http://localhost:9000` (or the port you configured).

## Architecture

The application follows a layered architecture, separating concerns into `cmd`, `internal`, `mvc`, and `pkg` directories for clarity and maintainability.