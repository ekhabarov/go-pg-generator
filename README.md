## Go struct generator for PostgreSQL

CLI tool creates Golang structs by table description from PostgreSQL database. 

### Dependencies
#### For build
```
  go get github.com/jessevdk/go-flags
  go get github.com/lib/pq
  go get github.com/jinzhu/inflection
```
#### For generated files
`import "gopkg.in/guregu/null.v3"`
Provides support for `null` values.

`import "github.com/satori/go.uuid"`
Provides support for `uuid.UUID` values.

### Arguments 

```
Usage:
  go-pg-generator [OPTIONS]

  Application Options:
    -s, --server=   Server name or IP address (default: 127.0.0.1)
    -p, --port=     Port (default: 5432)
    -u, --user=     Username
    -d, --database= Database name
    -t, --tables=   Tables to export
        --ssl=      SSL mode (require|verify-full|verify-ca|disable) (default: disable)
    -o, --output=   Output filename

  Help Options:
    -h, --help      Show this help message
```
