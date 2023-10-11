import gdb
from queue import Queue

ALPHABET = '1234567890abcdef'
EXPECTED = 'occdpnkibjefihcgjanhofnhkdfnabmofnopaghhgnjhbkalgpnpdjonblalfciifiimkaoenpealibelmkdpbdlcldicplephbo'

class Checkpoint(gdb.Breakpoint):
    def __init__(self, queue, *args):
        super().__init__(*args)
        self.silent = True
        self.queue = queue

    def stop(self):
        index = self.queue.get(timeout = 10)
        a = gdb.parse_and_eval("((char*)($rbp - 0x70))[{}]".format(index))
        self.queue.put(chr(a))
        print(f"\t{a}")
        
        return False

gdb.execute("set disable-randomization on")
gdb.execute("delete")
queue = Queue()
bp = Checkpoint(queue, '*0x55555540099d')

with open("flag.txt") as f:
    flag = f.read()

key = ""

for i in range(len(key), len(EXPECTED)):
    for c in ALPHABET:
        queue.put(i)
        gdb.execute("run {}".format(key + c))
        result = queue.get(timeout = 10)
        if result == EXPECTED[i]:
            key += c
            print(key)
            break

bytes_key = bytes.fromhex(key)
bytes_flag = bytes.fromhex(flag)

xor = bytes(a ^ b for a, b in zip(bytes_flag, bytes_key)).decode("ascii")

print("\n")

print(f"flag: {flag}")
print(f"key : {key}")
print(f"xor : {xor}")