import time, os
TEST_COUNT = 4 # set this to a lower number for faster results
LEN = 8
CHOICES = "0123456789"

def print_check():
    TEST_COUNT = 4 # set this to a lower number for faster results
    LEN = 8
    CHOICES = "0123456789"
    prefix = ""
    for c in CHOICES:
        cur = prefix+c+'0'*(LEN-len(prefix)-1)
        foo = 0
        for _ in range(TEST_COUNT): # repeat TEST_COUNT times and get the average to improve accuracy
            start = time.time()
            os.system(f"echo '{cur}' | ./pin_checker > /dev/null")
            foo += (time.time()-start)
        avgtime = foo/TEST_COUNT
        print(f"{cur}: {avgtime}")

def check(prefix):
  mxtime = -1
  res = ''
  for c in CHOICES:
    cur = prefix+c+'0'*(LEN-len(prefix)-1)
    foo = 0
    for _ in range(TEST_COUNT): # repeat TEST_COUNT times and get the average to improve accuracy
      start = time.time()
      os.system(f"echo '{cur}' | ./pin_checker > /dev/null")
      foo += (time.time()-start)
    avgtime = foo/TEST_COUNT
    if avgtime>mxtime:
      mxtime = avgtime
      res = c
  return res


def main():
  prefix = ''
  for i in range(7):
    c = check(prefix)
    prefix+=c
    print(c)
  # Check the final number(do not hide the results)
  for c in CHOICES:
    cur = prefix+c
    print(f"------ trying {cur} --------")
    os.system(f"echo '{cur}' | ./pin_checker")

print_check()
main()