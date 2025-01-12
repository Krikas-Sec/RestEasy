# RestEasy - A Simple API Testing CLI

RestEasy is a lightweight CLI tool built with Go for testing APIs via HTTP/HTTPS requests. The tool supports methods like GET, POST, PUT, DELETE, and allows users to include custom headers and body content. Additionally, configurations can be saved and loaded for later reuse.

## Features

- **HTTP Methods**: Supports GET, POST, PUT, DELETE.
- **Custom Headers**: Add headers as key-value pairs.
- **Request Body**: Send JSON or other text payloads.
- **Save and Load Configurations**: Save request settings to a JSON file and reload them when needed.
- **Readable Response**: Displays status, headers, and body of the API response.

## Prerequisites

- Go 1.16 or later

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/RestEasy.git
   cd RestEasy
   ```
2. Build the CLI tool:
   ```bash
   go build -o resteasy main.go
   ```
3. Run the tool:
   ```bash
   ./resteasy
   ```

## Usage

### Basic Usage

#### GET Request
```bash
./resteasy -method GET -url "https://jsonplaceholder.typicode.com/posts/1"
```

#### POST Request with Headers and Body
```bash
./resteasy -method POST -url "https://jsonplaceholder.typicode.com/posts" \
-headers "Content-Type:application/json" \
-body '{"title":"foo","body":"bar","userId":1}'
```

#### Save Request Configuration
```bash
./resteasy -method POST -url "https://jsonplaceholder.typicode.com/posts" \
-headers "Content-Type:application/json" \
-body '{"title":"foo","body":"bar","userId":1}' \
-save request.json
```

#### Load Request Configuration
```bash
./resteasy -load request.json
```

## Command-Line Options

| Option       | Description                              | Default   |
|--------------|------------------------------------------|-----------|
| `-method`    | HTTP method to use (GET, POST, etc.)     | `GET`     |
| `-url`       | API endpoint to call                    | `""`      |
| `-headers`   | HTTP headers in key:value,key:value format | `""`      |
| `-body`      | Request body content                    | `""`      |
| `-save`      | Save request configuration to a file    | `""`      |
| `-load`      | Load request configuration from a file  | `""`      |

## Example APIs for Testing

1. **JSONPlaceholder (Fake Online REST API)**
   - URL: `https://jsonplaceholder.typicode.com/posts`
   - Supports GET, POST, PUT, DELETE.

2. **Public APIs Directory**
   - Browse various free APIs at [public-apis.io](https://public-apis.io).

## Contributing

Contributions are welcome! If you'd like to contribute to this project:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Submit a pull request with a detailed explanation of your changes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Author

Created by [Krikas-Sec](https://Krikas-Sec.github.io). Feel free to reach out for questions or suggestions.

## Support

If you found this project helpful, consider supporting its development by [buying me a coffee](https://buymeacoffee.com/Tempcoder).
