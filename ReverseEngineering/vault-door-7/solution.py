from pwn import *
from math import ceil, log2
a = [1096770097,1952395366,1600270708,1601398833,1716808014,1734304867,942695730,942748212]
# print("".join(p32(x, endian='big') for x in a))
print("".join(x.to_bytes((x.bit_length() + 7) // 8, 'big').decode('utf-8') for x in a))