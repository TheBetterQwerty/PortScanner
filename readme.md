# Go Port Scanner

This is a simple Go-based port scanner that checks the open ports of a given target. It works by attempting to make TCP connections to ports 1 through 65535 and reports which ports are open.

## Features

- Scans TCP ports from 1 to 65535.
- Uses goroutines for concurrent port checking, making the scan faster.
- Displays the progress of the scan.
- Lists open ports with simple output.

## Prerequisites

- Go 1.16 or higher

## Installation

1. Clone the repository or download the `main.go` file.

   ```bash
   git clone https://github.com/yourusername/port-scanner.git
   ```

2. Navigate to the project folder:

   ```bash
   cd port-scanner
   ```

3. Install the necessary Go packages (if any):

   ```bash
   go mod tidy
   ```

4. Build the Go project:

   ```bash
   go build -o port-scanner main.go
   ```

## Usage

To use the port scanner, run the following command with the `-t` flag followed by the target domain or IP address:

```bash
./port-scanner -t <target>
```

Example:

```bash
./port-scanner -t example.com
```

This will scan the target domain (`example.com`) for open ports from 1 to 65535 and display the results.

### Flags

- `-t` : The target to scan (e.g., `example.com`, `192.168.1.1`).

## Example Output

```bash
Progress: 100.00%
[+] Port 22 is Opened!
[+] Port 80 is Opened!
[+] Port 443 is Opened!
```

If no open ports are found:

```bash
Progress: 100.00%
[!] No Port Was Opened!
```

## Contributing

Feel free to fork this project and submit pull requests for improvements or fixes. If you encounter any issues, please open an issue on GitHub.

## License

This project is open-source and available under the [MIT License](LICENSE).

---

Let me know if you'd like to add or adjust anything!
