# Go Application

This is a Go application designed to interact with the Gemini API. The application is configured to run from `./cmd/main.go` and requires an `.env` file for proper functionality.

--- 

## Prerequisites

- Go installed (version 1.23 or higher recommended).
- A valid Gemini API key.

---

## Getting Started

### 1. Clone the Repository

```bash
git clone <repository-url>
cd <repository-directory>
```

### 2. Set Up the `.env` File

Create a `.env` file in the root directory and add the following environment variables:

```env
SERVER_HOST=0.0.0.0
SERVER_REST_PORT=8080
GEMINI_API_KEY=<your-gemini-api-key>
```

Replace `<your-gemini-api-key>` with your actual Gemini API key.

#### How to Obtain a Gemini API Key

1. Visit the [Gemini API Documentation](https://ai.google.dev/gemini-api/docs/api-key).
2. Log in to your Google account.
3. Create a new API key with the required permissions.
4. Copy the generated API key and paste it into the `GEMINI_API_KEY` field in your `.env` file.

### 3. Install Dependencies

Ensure all necessary Go modules are installed:

```bash
go mod tidy
```

### 4. Run the Application

Navigate to the `cmd` directory and run the application using:

```bash
go run ./cmd/main.go
```

---

## Configuration

### Environment Variables

| Variable          | Description                           | Default Value |
|-------------------|---------------------------------------|---------------|
| `SERVER_HOST`     | The host where the server runs        | `0.0.0.0`     |
| `SERVER_REST_PORT`| The port where the REST API listens   | `8080`        |
| `GEMINI_API_KEY`  | Your Gemini API key for authentication| None          |

---

## Troubleshooting

- Ensure the `.env` file exists and is properly configured.
- Check that your Gemini API key has the required permissions.
- Verify that the port `8080` is not already in use.

---

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

