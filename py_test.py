'''
Descripttion: 
version: 
Author: xiaoshuyui
email: guchengxi1994@qq.com
Date: 2021-11-12 18:57:31
LastEditors: xiaoshuyui
LastEditTime: 2021-11-13 09:03:52
'''
import os
from pathlib import Path

s = "c:\\docs\\a.txt"
s2 = "c:\\docs\\a.tx\\t"
s3 = "a.txt"
s4 = "a..txt"

print(os.path.splitext(s))

print(os.path.splitext(s2))

print(os.path.splitext(s3))

print(os.path.splitext(s4))

print("====================================")

print(os.path.split(s))

print(os.path.split(s2))

print(os.path.split(s3))

print(os.path.split(s4))

print("====================================")

print(os.path.splitdrive(s)) 

print("==========================================")

print(os.getcwd())

print(Path.cwd())

print("==========================================")

print(os.path.islink("D:\github_repo\go_example\go.mod - 快捷方式.lnk"))
print(os.utime("D:\github_repo\go_example\go.mod"))