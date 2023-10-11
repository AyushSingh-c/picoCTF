from pwn import *

with open("encoded.bmp", "rb") as b:
    b.seek(2000)
    bin_str = ""
    for j in range(50 * 8):
        bin_str += str(ord(b.read(1)) & 1)

char_str = unbits(bin_str, endian = 'little')
print("".join(map(lambda c: chr(c + 5), char_str)))