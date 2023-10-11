
flag = ""

with open("rev_this", "rb") as data:
    for i in range(8):
        flag += chr(data.read(1)[0])
    for i in range(8, 23):
        if (i & 1) == 0:
            flag += chr(data.read(1)[0] - 5)
        else:
            flag += chr(data.read(1)[0] + 2)
    flag+= chr(data.read(1)[0])

print("Flag: {}".format(flag))