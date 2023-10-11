# coding=utf8

def bitstring_to_bytes(s):
    return int(s, 2).to_bytes((len(s) + 7) // 8, byteorder='big')

text = '                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                '

firstType = ' '
secondType =  ' '
binaryString = ''

for char in text: #Foreach char
	if char == firstType: #Check if it is the first type
		binaryString += '0' #Mark it as 0
	else:
		binaryString += '1' #Mark it as 1

# print(binaryString) #Print result
print(bitstring_to_bytes(binaryString))