# descrit-api

[![CircleCI](https://circleci.com/gh/nicewrk/design-brain-api.svg?style=svg)](https://circleci.com/gh/nicewrk/design-brain-api)

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
