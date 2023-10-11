filename = 'output.bmp'
with open(filename, 'rb') as f:
    content = f.read()
noise = content.hex()
data = ""
out = open("extracted.zip", 'wb')
for i in range(280, len(noise), 8):
    data += noise[i:i+4]
out.write(bytes.fromhex(data))