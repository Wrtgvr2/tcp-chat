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
2. Run server:
  - Without Docker:
    - Build server `.exe` file and run it:
    ```bash
    go build ./server/
    ```
  - With Docker:
    ```bash
    docker-compose up
    ```
3. Run client:
  - Build client `.exe` file and run it:
  ```
  go build ./client/
  ```
