from scapy.all import *
from collections import namedtuple
import random

first_port = 0
seed = 1614044650
###############################################pickup payload packets##############################################

packets = rdpcap('./capture.pcapng')
Packet = namedtuple("Packet", "src_port payload")
payload_packets = []
for packet in packets:
    if not packet.haslayer(IP):
        continue
    if not packet.haslayer(UDP):
        continue
    if not packet[IP].src == "172.17.0.2":
        continue
    if not packet[IP].dst == "172.17.0.3":
        continue
    if not packet[UDP].dport == 56742:
        continue
    if not len(packet[UDP].payload) == 1:
        continue
    if first_port == 0:
        first_port = packet[UDP].sport
    payload_packets.append(Packet(packet[UDP].sport, bytes(packet[UDP].payload)))
    
#############################################check seed#########################################################

random.seed(seed)
random.shuffle([None]*len(payload_packets))
sport=random.randrange(65536)
assert(sport, first_port)

###############################################retrieving the flag########################################
random.seed(seed)
locations = list(range(len(payload_packets)))
random.shuffle(locations)
output = [None] * len(payload_packets)

for i, packet in enumerate(payload_packets):
    srcport = random.randrange(65536)

    if(srcport != packet.src_port):
        raise RuntimeError(f"Iteration #{i}: srcport ({srcport}) != port ({packet.src_port})")

    output[locations[i]] = ord(packet.payload) ^ random.randrange(256)
    
with open("output.bin", "wb") as o:
    for b in output:
        o.write(bytes([b]))