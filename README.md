# Golang struct generator for PostgreSQL

CLI tool for generation Golang structures by table definition from PostgreSQL.

## Build
Install Go vendor tool [govendor](https://github.com/kardianos/govendor) and run
```
make build
```


## Dependencies
### For generated files
`import "github.com/guregu/null/v5"`
Provides support for `null` values.

`import "github.com/satori/go.uuid"`
Provides support for `uuid.UUID` values.

## Help
```
Usage:
  go-pg-generator [OPTIONS]

Application Options:
  -s, --server=         Server name or IP address (default: 127.0.0.1)
  -p, --port=           Port (default: 5432)
  -u, --user=           Database user.
  -w, --password=       Database password.
  -d, --database=       Database name.
  -t, --tables=         Tables to export.
      --ssl=            SSL mode (require|verify-full|verify-ca|disable) (default: disable)
  -f, --file-per-table  Save each structure to its own .go file.
      --package=        Package name for generated files.

Help Options:
  -h, --help            Show this help message
```
