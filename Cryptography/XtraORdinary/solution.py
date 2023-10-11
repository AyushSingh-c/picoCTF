import itertools

with open("output.txt") as f:
    flag = bytes.fromhex(f.read())

def encrypt(ptxt, key):
    ctxt = b''
    for i in range(len(ptxt)):
        a = ptxt[i]
        b = key[i % len(key)]
        ctxt += bytes([a ^ b])
    return ctxt


random_strs = [
    b'my encryption method',
    b'is absolutely impenetrable',
    b'and you will never',
    b'ever',
    b'break it'
]

prefix = b"picoCTF{"

def apply_encryption(ctxt, truth_table):
    for i, to_xor in enumerate(truth_table):
        if to_xor:
            ctxt = encrypt(ctxt, random_strs[i])
    return ctxt

for p in itertools.product([True, False], repeat=len(random_strs)):
    ctxt = apply_encryption(flag, p)
    print(f"{p}: \t{encrypt(ctxt[:len(prefix)], prefix)}")
    
ctxt = apply_encryption(flag, (False, False, True, True, True))
print(encrypt(ctxt, b"Africa!"))