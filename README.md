# quickdbClient

quickdbClient is a powerful command-line interface (CLI) tool written in Go for interacting with QuickDB. It enables users to perform efficient database operations directly from the terminal.

## Features
- **Lightweight and fast** – optimized for quick execution.
- **Simple command-line interface** – easy to use.
- **Supports PING,GET,SET,HGET,HSET,HGETALL** – It support all the operation that are mentioned and can be improved accordingly 

## Installation

### Prerequisites
- Go 1.20 or later installed.
- Git installed on your system.

### Build from Source

To install quickdbClient, follow these steps:

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yagyagoel1/quickdbClient.git
   cd quickdbClient
   ```

2. **Build the binary:**
   ```sh
   go build -o quickdbClient
   ```

3. **Move the binary to a directory in your `$PATH`** (optional but recommended):
   ```sh
   mv quickdbClient /usr/local/bin/
   ```

4. **Verify the installation:**
   ```sh
   quickdbClient --help
   ```

5. **example usage**
   ```sh
   quickdbClient SET key value
   quickdbClient GET key 
   quickdbClient HSET map key value
   quickdbClient HGET map key
   quickdbClient HGETALL map
   quickdbClient PING
   ```

