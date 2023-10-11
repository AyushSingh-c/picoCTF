from sage.all import *
import math
import hashlib
import sys

ITERS = int(2e7)
P = (10**10000)
VERIF_KEY = "96cc5f3b460732b442814fd33cf8537c"
ENCRYPTED_FLAG = bytes.fromhex("42cbbce1487b443de1acf4834baed794f4bbd0dfe7d7086e788af7922b")

def m_func(i):
    A = matrix([[0, 1, 0, 0], [0, 0, 1, 0], [0, 0, 0, 1], [55692, -9549, 301, 21]]) 
    B = pow(A, i, P)
    value = B * matrix([[1], [2], [3], [4]])
    return value[0]

# Decrypt the flag
def decrypt_flag(sol):
    sol = sol % P
    sol = str(sol)
    sol_md5 = hashlib.md5(sol.encode()).hexdigest()

    if sol_md5 != VERIF_KEY:
        print("Incorrect solution")
        sys.exit(1)

    key = hashlib.sha256(sol.encode()).digest()
    flag = bytearray([char ^ key[i] for i, char in enumerate(ENCRYPTED_FLAG)]).decode()

    print(flag)

if __name__ == "__main__":
    sol = m_func(ITERS)
    decrypt_flag(sol)