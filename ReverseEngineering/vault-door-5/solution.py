import base64
import urllib.parse
flag_base_encoded = "JTYzJTMwJTZlJTc2JTMzJTcyJTc0JTMxJTZlJTY3JTVm" + "JTY2JTcyJTMwJTZkJTVmJTYyJTYxJTM1JTY1JTVmJTM2" + "JTM0JTVmJTY1JTMzJTMxJTM1JTMyJTYyJTY2JTM0"
flag_url_encoded = base64.b64decode(flag_base_encoded)
print(flag_url_encoded)
flag = "picoCTF{" + urllib.parse.unquote(flag_url_encoded) + "}"
print(flag)