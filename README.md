# Weather API in Go

This is a simple weather API implemented in Go. It fetches weather data from the Central Weather Bureau and serves it via a web API.

## Getting Started

### Prerequisites

- Go 1.20 or higher
- Gin framework

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/weather_go.git
    cd weather_go
    ```

2. Install the dependencies:

    ```bash
    go get -u github.com/gin-gonic/gin
    ```

### Running the Application

To run the application, use the following command:

    ```bash
    go run main.go
    ```

The server will start on `http://localhost:8080`.

### API Endpoints

- **Home**: `GET /`
  - Returns a welcome message.

- **Weather**: `GET /weather`
  - Fetches and returns weather data from the Central Weather Bureau.

### Project Structure

    ```plaintext
    weather_go/
    ├── go.mod
    ├── go.sum
    ├── main.go
    └── README.md
    ```

## License

This project is licensed under the MIT License.

### 1. Initialize Git repository and add files

In the project directory, initialize the Git repository, add files, and make the first commit:

    ```bash
    git init -b main
    git add .
    git commit -m "Initial commit"
    ```

### 2. Create GitHub repository

1. Open GitHub and log into your account.
2. Create a new repository named `weather_go`.
3. Do not check any initialization options (like README or .gitignore).

### 3. Connect to GitHub and push

In the project directory, link the local repository to the GitHub repository and push:

    ```bash
    git remote add origin https://github.com/yourusername/weather_go.git
    git push -u origin main
    ```

Please replace `yourusername` with your GitHub username.

### 4. Confirm successful push

Open the GitHub repository page, and you should see the project files and README.md file.

This README provides a complete guide to setting up and running the Go project, including the necessary commands within copyable code blocks. If you need further assistance, feel free to ask.
