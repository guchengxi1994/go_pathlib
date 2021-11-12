import os

s = "c:\\docs\\a.txt"
s2 = "c:\\docs\\a.tx\\t"
s3 = "a.txt"
s4 = "a..txt"

print(os.path.splitext(s))

print(os.path.splitext(s2))

print(os.path.splitext(s3))

print(os.path.splitext(s4))