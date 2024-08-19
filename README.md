<div align="center">
  <img src="https://raw.githubusercontent.com/mrinjamul/gitignore-service/c41cdcb940554cd8f4908cc0b0cdc40aa75054c7/static/favicon.png" width="100px" alt="logo" />
  <h1><code>gitignore-service</code></h1>
  <p>
    <strong>A web service to help you generate [.gitignore] files for Git repositories.</strong>
  </p>
</div>

Just simply download gitignore file by,

```shell
curl -s "http://localhost:8080/api/gi/get?for=go" -o .gitignore
```

## Features

- **Generate `.gitignore` Files**: Easily generate `.gitignore` files for multiple languages and frameworks.
- **API Endpoints**: Access various endpoints to fetch `.gitignore` templates and check service health.
- **Health Check**: Monitor the service health with detailed health check endpoints.
- **Static File Serving**: Serve static files like images and documents from the `/static` endpoint.

## Installation

### Prerequisites

- Go 1.18 or later
- Git

### Clone the Repository

```shell
git clone https://github.com/mrinjamul/gitignore-service.git
cd gitignore-service
```

### Load Environment Variables

For Linux,

```sh
export $(cat .env | xargs)
```

For Windows,

```powershell
Get-Content .env | ForEach-Object { if ($_ -and $_ -notmatch '^\s*#') { $parts = $_ -split '=', 2; [System.Environment]::SetEnvironmentVariable($parts[0].Trim(), $parts[1].Trim()) } }
```

### Install Dependencies

```shell
go mod download
```

### Running the Application

```shell
go run .
```

## Usage

**Base URL**: `http://localhost:8080`

**Endpoints**:

- `GET /static/*filepath` --> Serves static files.
- `GET /` --> Main page.
- `GET /api/gi/` --> Fetch list of `.gitignore` files.
- `GET /api/gi?for=go` --> Fetch a specific `.gitignore` files.
- `GET /api/gi/get?for=go` --> Download a specific `.gitignore` file.
- `GET /api/health` --> Health check endpoint.
- `GET /api/health/full` --> Detailed health check.

## Contributing

Contributions are welcome! To contribute:

1. **Fork the repository**:

   Click the "Fork" button at the top right of the repository page on GitHub.

2. **Create a new branch**:

   ```shell
    git checkout -b feature-branch
   ```

3. **Commit your changes**:
   ```shell
   git commit -am 'Add new feature'
   ```
4. **Push to the branch**:
   ```shell
   git push origin feature-branch
   ```
5. **Create a new Pull Request**:

   Go to the repository on GitHub, switch to your branch, and click "New Pull Request".

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For any questions or issues, please open an issue on GitHub or contact the project maintainer.
