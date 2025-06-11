import socket

if __name__ == "__main__":
    host = "127.0.0.1"
    port = 8080
    totalClients = int(input("Enter number of clients: "))

    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    sock.bind((host, port))
    sock.listen(totalClients)

    connections = []
    print("Initiating clients...")

    for i in range(totalClients):
        conn = sock.accept()
        connections.append(conn)
        print("Connected with client", i + 1)


    for idx, conn in enumerate(connections): # idx -> index
        data = conn[0].recv(1024).decode()
        if not data:
            continue

        filename = f"output{idx}.txt"
        with open(filename, 'w') as fo:
            while data:
                fo.write(data)
                data = conn[0].recv(1024).decode()

        print(f"\nReceiving file from client {idx + 1}")
        print(f"Received successfully! New filename is: {filename}\n")

    # Closing connections
    for conn in connections:
        conn[0].close()
