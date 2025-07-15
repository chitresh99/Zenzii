# Zenzii

To run the files use the following commands

### Server

```bash
go run main.go server.go client.go -mode=server
```

### Client

```bash
go run main.go server.go client.go -mode=client
```

To build the binary and then run use

```bash
go build -o zenzii main.go server.go client.go


./zenzii -mode=server


./zenzii -mode=client
```