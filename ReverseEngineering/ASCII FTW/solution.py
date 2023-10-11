input = open("./input.txt", "r")
flag = ""
lines = input.readlines()
for line in lines:
    hexString = bytearray.fromhex(line.split(",")[1].strip()[2:]).decode()
    flag += hexString
print(flag)
    