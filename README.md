# MOCKIT

Create new mock server

```
mock-it -v <VERB> -e <ENDPOINT> -d <RESPONSE_DATA> -s <STATUS_CODE>
```

## Flags
    * `-v` Http verb -> By default: `GET`
    * `-e` Endpoint  -> By default: `/`
    * `-d` Data to response -> By default: `{"msg":"ok"}`
    * `-s` Response status  -> By default: `200`
    * `-p` Port -> By default: `3000`

## Suported Http Verbs 
    * GET
    * POST
    * DELETE
    * PUT
    * PATCH


levanta un server con el endpoint:
GET -> `localhost:9090/users` -> return -> data in `data.json` -> status `200`

Ejemplos:

```
mock-it post user -d data.json
```
data.json
```
[
    {
        "user":123,
        "name":"foo"
    },
    {
        "user":123,
        "name":"bar"
    }
]
```

```
mock-it get users -s 402
```

