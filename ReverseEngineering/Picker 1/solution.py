from pwn import *

#####################################################change directory#################################3
ip = "saturn.picoctf.net:54880"
#######################################start process#############################################
r = remote(ip.split(':')[0], ip.split(':')[1])
r.sendlineafter("==>", "win")
input = str(r.recvline()).split(" ")[1:-2]
flag = ""
for i in input:
    flag += bytearray.fromhex(i.strip()[2:]).decode()
print(flag)
r.close()