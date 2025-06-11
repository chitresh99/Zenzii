import socket
import os

if __name__ == "__main__":
    host="127.0.0.1"
    port= 8080
    sock = socket.socket(socket.AF_INET,socket.SOCK_STREAM)
    sock.connect((host,port))
    

    filename = input("Enter the name of the file : ")
    
    if not os.path.isfile(filename):
        print("File doesn't exist")
        sock.close()
        exit()

    try:
        with open(filename,'r') as fi:
            while True:
                data = fi.read(1024) #reading in chunks
                if not data:
                    break
                sock.send(data.encode('utf-8'))
        print("File sent successfully")

    except IOError:
        print("Error opening or reading the file.")
    finally:
        sock.close()