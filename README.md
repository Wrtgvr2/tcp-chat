# CLI TCP Chat
*(Made and published without major reason)*
## Description
Small and simple CLI chat built with Go.
## Features
- Server and client
- Username choosing
- Easy-to-add chat commands
## Get started
1. Clone repo:
```bash
git clone https://github.com/Wrtgvr2/tcp-chat.git
cd tcp-chat
```
2.1 Run server without Docker:
- Run server `.go` file:
```bash
go run ./server/
```
or
- Build server `.exe` file:
```bash
go build ./server/
```
2.2 Run server with Docker:
```bash
docker-compose up
```
3. Run client:
- Run client `.go` file:
```bash
go run ./client/
```
or
- Build client `.exe` file:
```
go build ./client/
```
