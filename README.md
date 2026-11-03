*# GoRedis

A lightweight Redis-inspired in-memory key-value store built using Go. This project explores TCP communication, RESP protocol parsing, concurrent client handling, and basic Redis commands.

## Features

- TCP server-client communication
- RESP protocol parsing
- `SET` and `GET` commands
- Concurrent client handling using goroutines
- In-memory key-value storage
- Thread-safe data access

## Tech Stack

- **Language:** Go
- **Protocol:** RESP
- **Networking:** TCP
- **Library:** [tidwall/resp](https://github.com/tidwall/resp)

## Getting Started

Clone the repository:

```bash
git clone <your-repository-url>
cd goredisclone
```

Build the project:

```bash
make build
```

Run the server:

```bash
./bin/goredis --listenAddr :5001
```

## Testing

```bash
go test ./...
```

## Learning Outcomes

This project helped me understand Go fundamentals, TCP networking, RESP protocol, goroutines, channels, and concurrent programming.


I wrote up the full journey here: https://go-redis.surge.sh/
