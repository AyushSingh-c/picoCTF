from math import log, ceil
rotation_red = "39722847074734820757600524178581224432297292490103995912415595360101562905"
rotations_blue = int(rotation_red)*5
n_bytes_needed = ceil(log(rotations_blue, 2**8))
print(rotations_blue.to_bytes(n_bytes_needed, byteorder="big"))