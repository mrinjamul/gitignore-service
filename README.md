<div align="center">
  <img src="static/favicon.png" width="100px" alt="logo" />
  <h1><code>gitignore-service</code></h1>
  <p>
    <strong>A web service to help you with [dot] gitgnore files for git repositories.</strong>
  </p>
</div>

## Development

First load environment variables into the shell.

For Linux,

```sh
export $(cat .env | xargs)
```

For Windows,

```powershell
Get-Content .env | ForEach-Object { if ($_ -and $_ -notmatch '^\s*#') { $parts = $_ -split '=', 2; [System.Environment]::SetEnvironmentVariable($parts[0].Trim(), $parts[1].Trim()) } }
```

To start the applications,

```shell
go run .
```
