import socket
import struct
import os

server_address = ('0.0.0.0',69)
server_socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
server_socket.bind(server_address)
print("TFTP server is listening on port 69")
data, client_address = server_socket.recvfrom(1024)
print(f"Client is connected. Client address is {client_address}")
file_size =os.path.getsize("netboot.xyz-arm64.efi")
server_socket.sendto(b'\x00\x06tsize\x00'+str(file_size).encode()+b'\x00',client_address)
data, client_address = server_socket.recvfrom(1024)
n=1
print('''File name is netboot.xyz-arm64.efi
File sending is started''')
with open("netboot.xyz-arm64.efi", 'rb') as file:
	while True:
		data, client_address = server_socket.recvfrom(1024)
		data_chunk = file.read(512)
		if not data_chunk:
			break
		header = struct.pack('!HH', 3, n)
		packet = header +data_chunk
		server_socket.sendto(packet,client_address)
		n=n+1
server_socket.close()
print(f'''File is sent succesfully
File Size is {file_size}''')