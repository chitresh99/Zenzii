# Zenzii

Zenzii is a file transfer tool written in python and go

This project demonstrates how to transfer files over a TCP network socket, with two separate implementations

- **Python**: A simple and minimal version
- **Go**: A command-line based, modular version

## Features

- Send files from a client to a server over TCP
- Local file transfer with minimal dependencies

## Execution

- **Python**
```
python3 tcp-server.py
```

```
python3 tcp-client.py
```

- **Go**

```
cd zenzii-go
```

```
go run . -mode=server
```

```
go run . -mode=client
```
