# design-brain-api

[![CircleCI](https://circleci.com/gh/nicewrk/design-brain-api.svg?style=shield)](https://circleci.com/gh/nicewrk/design-brain-api) [![Coverage Status](https://coveralls.io/repos/github/nicewrk/design-brain-api/badge.svg?branch=master)](https://coveralls.io/github/nicewrk/design-brain-api?branch=master)

## Development

### Docker
`.env`:
```shell
DB_HOST=designbrainapi_db_1
DB_NAME=postgres
DB_PASS=
DB_PORT=5432
DB_USER=postgres
```

Run:
```shell
make up
```

Remove containers/images:
```shell
docker rm $(docker ps -a -q)
docker rmi $(docker images -q)
```

### Database

#### Create new migration
Example:
```shell
make name=users-table migration
```

### Testing

```shell
make test
```

```shell
go test -v ./store
```

```shell
go test -v ./store -run TestConnURI
```

#### Coverage
Check for race conditions:
`go test -race -covermode=atomic -coverprofile=$TMP_OUTFILE $package`

#### Idiomatic Go and Development Tools
Many editors now offer support to make development in Golang easier (e.g., [using `goimports`](https://godoc.org/golang.org/x/tools/cmd/goimports)).

In lieu of having such a local configuration, please feel free to use `make check` and `make fmt`.

#### Circle CI
When making modifications to Circle CI configuration, be sure to [validate your YAML](https://codebeautify.org/yaml-validator) and otherwise [validate your configuration](https://circleci.com/docs/2.0/local-jobs/).
