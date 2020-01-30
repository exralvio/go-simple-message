# go-simple-message
## Build with
- [Echo Framework](https://echo.labstack.com/)
- [Gorilla WebSocket](https://github.com/gorilla/websocket)

## Installation
- Make sure `Go` already installed in your machine
- Clone the project **outside** $GOPATH directory
- Go to project root directory 
- Simply call `go run cmd/http/main.go`

## Feature
- WebSocket / Show real time messages `[GET] http://127.0.0.1:3000/`
- Create new message `[POST] http://127.0.0.1:3000/api/message`
- Show inMemory messages `[GET] http://127.0.0.1:3000/api/message`
- Unittest
