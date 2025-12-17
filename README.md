# HTTP to TFTP Bridge

<div align="center">

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![License](https://img.shields.io/badge/license-MIT-green?style=for-the-badge)

**A high-performance bridge that enables TFTP clients to fetch files from HTTP URLs.**

[Report Bug](https://github.com/yourusername/http-to-tftp/issues) • [Request Feature](https://github.com/yourusername/http-to-tftp/issues)

</div>

---

## 📑 Table of Contents

- [About The Project](#about-the-project)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

---

## 📖 About The Project

The **HTTP to TFTP Bridge** is a specialized proxy server designed to modernize legacy workflows. It allows standard TFTP clients—often found in networking hardware, embedded systems, and PXE boot environments—to retrieve files transparently from modern HTTP servers.

Instead of maintaining a separate TFTP server with local files, this bridge dynamically fetches resources from the web as they are requested.

## 🚀 Features

*   **⚡ Protocol Translation**: Seamlessly proxies TFTP read requests (`RRQ`) to HTTP `GET` requests.
*   **🚀 High Performance**: Uses a 64KB block size (`blksize`) to maximize throughput and minimize protocol overhead.
*   **🛠️ Standard Compliance**: Fully supports the `tsize` (transfer size) option, ensuring compatibility with strict TFTP clients.
*   **🐳 Docker Ready**: Comes with a multi-stage Dockerfile for instant, lightweight deployment.

---

## 🏁 Getting Started

### Prerequisites

*   **Go**: Version 1.21 or higher (for local build)
*   **Docker**: Optional, for containerized execution

### Installation

#### Option 1: Local Build

1.  **Clone the repository**
    ```bash
    git clone https://github.com/yourusername/http-to-tftp.git
    cd http-to-tftp
    ```

2.  **Run the server**
    ```bash
    go run main.go
    ```
    *Note: Binding to port 69 typically requires administrator/root privileges.*

#### Option 2: Docker Deployment

1.  **Build the image**
    ```bash
    docker build -t tftp-bridge .
    ```

2.  **Run the container**
    ```bash
    docker run -d -p 69:69/udp --name tftp-server tftp-bridge
    ```

---

## ⚡ Usage

Once running, the server listens on UDP port `69`. The "filename" you request via TFTP is interpreted as the target HTTP URL.

### Example

Using a standard command-line TFTP client:

```bash
# Connect to the local bridge
tftp localhost

# Request a file from the web
tftp> get http://example.com/file.txt
```

### How It Works

1.  **TFTP Request**: Client sends a read request for `http://example.com/file.txt`.
2.  **HTTP Fetch**: The bridge parses the filename and triggers an HTTP GET to `http://example.com/file.txt`.
3.  **Streaming**: The HTTP response body is streamed back to the TFTP client in standard TFTP data packets.

---

## ⚙️ Configuration

| Parameter | Value | Description |
| :--- | :--- | :--- |
| **Port** | `69` | Standard TFTP UDP port (hardcoded). |
| **Block Size** | `65464` | Negotiated block size for optimal speed. |
| **Transfer Size** | `Enabled` | Sends file size (`tsize`) to client before transfer. |

---

## 🤝 Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1.  Fork the Project
2.  Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3.  Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4.  Push to the Branch (`git push origin feature/AmazingFeature`)
5.  Open a Pull Request

---

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.
