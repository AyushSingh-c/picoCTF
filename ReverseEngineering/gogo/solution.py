from pwn import *
#####################################################change directory#################################3
ip = "mercury.picoctf.net:47423"
value1 = 0x3836313833366631336533643632376466613337356264623833383932313465
value2 = 0x4a53475d414503545d025a0a5357450d05005d555410010e4155574b45504601
value3 = value1 ^ value2
#######################################start process#############################################
r = remote(ip.split(':')[0], ip.split(':')[1])
payload = bytearray.fromhex(hex(value3)[2:])
print(payload)
r.sendlineafter(":", payload)
r.sendlineafter("?", "goldfish")
r.interactive()