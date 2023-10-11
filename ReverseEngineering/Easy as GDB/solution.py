import gdb
import string
from queue import Queue, Empty

MAX_FLAG_LEN = 0x200

class Checkpoint(gdb.Breakpoint):
    def __init__(self, queue, target_hitcount, *args):
        super().__init__(*args)
        self.silent = True
        self.queue = queue
        self.target_hitcount = target_hitcount
        self.hit = 0

    def stop(self):
        res = []
        self.hit += 1
        #print(f"\nhit {self.hit}/{self.target_hitcount}")
        if self.hit == self.target_hitcount:
            al = gdb.parse_and_eval("$al")
            dl = gdb.parse_and_eval("$dl")
            self.queue.put(al == dl)
        return False

class Solvepoint(gdb.Breakpoint):
    def __init__(self, *args):
        super().__init__(*args)
        self.silent = True
        self.hit = 0

    def stop(self):
        #gdb.execute("q")
        self.hit += 1
        return False


gdb.execute("set disable-randomization on")
gdb.execute("delete")
sp = Solvepoint("*0x56555a71")
queue = Queue()


flag = ""
ALPHABET = string.ascii_letters + string.digits + "{}_"

for i in range(len(flag), MAX_FLAG_LEN):
    for c in ALPHABET:
        bp = Checkpoint(queue, len(flag) + 1, '*0x5655598e')
        gdb.execute("run <<< {}{}".format(flag, c))
        try:
            result = queue.get(timeout = 1)
            bp.delete()
            if result:
                flag += c
                print("\n\n{}\n\n".format(flag))
                break
        except Empty:
            print("Error: Empty queue!")
            gdb.execute("q")

    if sp.hit > 0:
        print("Found flag: {}".format(flag))
        gdb.execute("q")

gdb.execute("q")