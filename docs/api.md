# Design Brain API

## `/healthcheck`

```shell
curl http://dev.design-brain.com:8080/healthcheck

{
	"status": 200
}
```

## `/users`

### GET `/users/:username`

```shell
curl http://dev.design-brain.com:8080/users/aja

{
	"data": {
		"uid": "5d8a20b4-6678-4b34-a9a4-e866ad3dce2d",
		"email": "a@nicewrk.com",
		"username": "aja",
		"is_verified": false,
		"created_at": "2018-01-14T17:29:44.157654Z",
		"updated_at": "2018-01-14T21:43:36.755943Z"
	},
	"status": 200
}
```

### POST `/users`

```shell
curl http://dev.design-brain.com:8080/users \
-d '{"email": "a@nicewrk.com", "username": "aja"}'

{
	"error": "unable to create user: email already exists",
	"status": 422
}
```

```shell
curl http://dev.design-brain.com:8080/users \
-d '{"email": "s@nicewrk.com", "username": "sara"}'

{
	"data": {
		"uid": "812bdf91-8201-4daa-8ddc-7a0ede842dc4",
		"email": "s@nicewrk.com",
		"username": "sara",
		"is_verified": false,
		"created_at": "2018-01-14T22:27:21.300514Z",
		"updated_at": "2018-01-14T22:27:21.300514Z"
	},
	"status": 201
}
```
