Here’s a well-documented `README.md` file for your backend setup instructions! 🎉

```markdown
# Go Backend Setup Guide 🚀

This guide will walk you through setting up and running the Go backend project on **Windows**, **macOS**, and **Linux**.

---

## 📂 Project Structure

- **controllers/**: Contains all the API logic.
- **middleware/**: Middleware for handling requests and responses.
- **models/**: Data models and schemas.
- **storage/**: Persistent storage utilities.
- **utils/**: Helper functions.
- **main.go**: The entry point of the application.

---

## 🛠 Prerequisites

Before running this project, ensure you have the following installed:

1. **Go** (version 1.19+): [Download Go](https://golang.org/dl/)
2. **Redis** (for caching or database needs): Installation instructions below.
3. **Git** (to clone the repository): [Download Git](https://git-scm.com/downloads)

---

## 🚀 Setup Instructions

### 1️⃣ Clone the Repository

```bash
git clone https://github.com/your-username/your-repository.git
cd your-repository
```

---

### 2️⃣ Install Dependencies

Run the following command to install all required Go modules:

```bash
go mod tidy
```

---

### 3️⃣ Install and Run Redis

Redis is required for this backend to work. Follow the instructions for your operating system:

#### **Linux** (Ubuntu/Debian)
```bash
sudo apt update
sudo apt install redis -y
sudo systemctl start redis
sudo systemctl enable redis
```

Verify Redis is running:
```bash
redis-cli ping
# Output: PONG
```

#### **macOS**
Install Redis using Homebrew:
```bash
brew install redis
brew services start redis
```

Verify Redis is running:
```bash
redis-cli ping
# Output: PONG
```

#### **Windows**
Download and install Redis from the official repository: [Redis for Windows](https://github.com/microsoftarchive/redis/releases)

Start the Redis server by navigating to the installation folder and running:
```bash
redis-server
```

Verify Redis is running:
```bash
redis-cli ping
# Output: PONG
```

---

### 4️⃣ Run the Backend

Start the application by running:

```bash
go run main.go
```

If everything is set up correctly, you should see logs indicating the server is running (e.g., `Server running on localhost:8080`).

---

## 🛠 Troubleshooting

### Error: `dial tcp [::1]:6379: connect: connection refused`
- This error occurs if Redis is not running or cannot be reached. 
- Ensure Redis is installed and running on `localhost:6379`. Follow the steps in the **Install and Run Redis** section above.

---

## 🐋 Running Redis with Docker (Optional)

If you prefer using Docker, you can run Redis with the following command:

```bash
docker run --name redis -p 6379:6379 -d redis
```

This starts a Redis server accessible at `localhost:6379`.

---

## 📊 Testing the API

Use tools like **Postman** or `curl` to test the API.

Example `curl` command:
```bash
curl -X GET http://localhost:8080/api/example
```

---

## 📦 Building for Production

To build a standalone binary for production:

```bash
go build -o app
./app
```

---

## 📝 Notes

- Update the `Redis` connection configuration in the code if you're using a custom host, port, or password. Look for something like:
  ```go
  rdb := redis.NewClient(&redis.Options{
      Addr:     "localhost:6379", // Update if needed
      Password: "",               // Add password if Redis is secured
      DB:       0,                // Use the appropriate DB index
  })
  ```
- Always restart the Redis server if changes are made to its configuration.

---

## 💡 Need Help?

If you encounter any issues or have questions, feel free to open an issue in this repository or reach out to the maintainer.

---

**Happy Coding!** 💻✨
```

This `README.md` includes detailed setup instructions for all operating systems, troubleshooting steps, and even optional Docker usage. Let me know if you want to customize it further! 🚀
