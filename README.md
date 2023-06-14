[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub release](https://img.shields.io/github/release/pablotrianda/mock-it.svg)](https://GitHub.com/pablotrianda/mock-it/releases/)
![status](https://github.com/pablotrianda/mock-it/actions/workflows/go.yml/badge.svg)


# MOCK-IT

Create new mock server

```
mock-it -v <VERB> -e <ENDPOINT> -d <RESPONSE_DATA> -s <STATUS_CODE>
```

## Flags
    * `-v` Http verb -> By default: `GET`
    * `-e` Endpoint  -> By default: `/`
    * `-d` Data to response -> By default: `{"msg":"Hello from MOCKIT 🧉"}`
    * `-s` Response status  -> By default: `200`
    * `-p` Port -> By default: `3000`

## Suported Http Verbs 
    * GET
    * POST
    * DELETE
    * PUT
    * PATCH


## Examples:
```
$ mock-it
```
This will create a new server with the enpoint:
`GET - http://localhost:3000/` and response with `{"msg":"Hello from MOCKIT 🧉"}` and status `200`



```
$ mock-it -v post -e user -d data.json -s 202
```
This will create a new server with the enpoint:
`POST - http://localhost:3000/user` and response with the data.json data and status `202`:
data.json
```
[
    {
        "user":123,
        "name":"foo"
    }
]

```



```
$ mock-it -v get -e users -d data.json -s 200
```
This will create a new server with the enpoint:
`GET - http://localhost:3000/users` and response with the data.json data and status `200`:
data.json
```
[
    {
        "user":123,
        "name":"foo"
    },
    {
        "user":456,
        "name":"bar"
    }
]

```
