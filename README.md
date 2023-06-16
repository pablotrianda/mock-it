[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub release](https://img.shields.io/github/release/pablotrianda/mock-it.svg)](https://GitHub.com/pablotrianda/mock-it/releases/)
![status](https://github.com/pablotrianda/mock-it/actions/workflows/go.yml/badge.svg)


# MOCK-IT ⚡ Create a mock server on the fly

```sh
mock-it -v <VERB> -e <ENDPOINT> -d <RESPONSE_DATA> -s <STATUS_CODE>
```
![mockit](https://i.imgur.com/kZ6fvmY.gif)

## Download
* Download the latest version from [releases section](https://github.com/pablotrianda/mock-it/releases). 

## Flags
   * `-v` Http verb -> By default: `GET`
   * `-e` Endpoint  -> By default: `/`
   * `-d` Data to respond -> By default: `{"msg":"Hello from MOCKIT 🧉"}`
   * `-s` Response status  -> By default: `200`
   * `-p` Port -> By default: `3000`

## Suported Http Verbs 
   * GET
   * POST
   * DELETE
   * PUT
   * PATCH


## Examples:
   * Create a simple server, run: 
      ```
      $ mock-it
      ```
      This will create a new server with the enpoint:<br>
         `GET - http://localhost:3000/` and response with `{"msg":"Hello from MOCKIT 🧉"}` and 
         status `200`
         
         
   * Create a server and responde a data from file:
      ```
      $ mock-it -v post -e user -d data.json -s 202
      ```
      This will create a new server with the enpoint: <br>
      `POST - http://localhost:3000/user` and response with the data.json data and status `202`:<br>
      data.json
      ```JSON
      [
          {
              "user":123,
              "name":"foo"
          }
      ]

      ```

## Run on development mode
   1. Clone this repo
   2. Run `go run src/cmd/main.go`
   3. To create a new build with [dagger](https://dagger.io): `go run build`
