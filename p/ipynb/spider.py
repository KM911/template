import json
import openpyxl
import requests
from bs4 import BeautifulSoup
import os
import gevent
from gevent import queue


def Copy(str):
    os.system(f'echo {str} | clip')

# 解析json header cookie url 三兄弟
def PrintJson(json)-> str:
    # string builder
    string_builder = '{'
    # 我需要知道是否是最后一个 这样不应该输出最后的,
    pc = 0
    for key, value in json.items():
        if (pc < len(json) - 1):
            string_builder += f'"{key}": "{value}",\n'
        else:
            string_builder += f'"{key}": "{value}"\n'
        pc += 1
    string_builder += '}'
    return string_builder

# 输入一个url
# 将url 和 query 返回
def ParseUrl(url):
    url, params = url.split('?', 1)
    params = params.split('&')
    params = [param.split('=') for param in params]
    params = {key: value for key, value in params}
    # print默认输出' 将其改为"即可
    print(f'url: "{url}"')
    print('params:')
    PrintJson(params)


def ParseHeader(header:str)-> str:
    # 移除第一行
    lines = header.replace("\"", "'").split('\n')
    Json = {}
    for line in lines:
        if line:
            key, value = line.split(':', 1)
            Json[key] = value.strip()
    return PrintJson(Json)


def ParseCookie(str):
    lines = str.replace("\"", "'").split(';')
    Json = {}
    for line in lines:
        if line:
            key, value = line.split('=', 1)
            Json[key.strip()] = value.strip()
    PrintJson(Json)


# excel 保存
# 只有单表
def SaveAsExcel(filename="test.xlsx", data=[[]], sheetname="sheet",auto_fit=True,auto_center=True):
    # 如果文件存在将其删除
    if os.path.exists(filename):
        os.remove(filename)
    if not filename.endswith(".xlsx"):
        filename = filename + ".xlsx"
    wb = openpyxl.Workbook()
    ws = wb.active
    ws.title = sheetname
    for row in data:
        ws.append(row)
        
    # 让文字居中
    if auto_center:
        for row in ws.rows:
            for cell in row:
                cell.alignment = openpyxl.styles.Alignment(horizontal='center', vertical='center')
    # 自动调整列宽
    if auto_fit:
            #  设置宽度
        max_length = [0 for i in range(len(data[0]))]
        for row in data:
            for i in range(len(row)):
                length = StringLength(row[i])
                if length > max_length[i]:
                    max_length[i] = length
        for i in range(len(max_length)):
            # column_dimensions最大的宽度是255
            # 还是自动调整 这里不想设置了 太麻烦了
            if max_length[i] > 255:
                max_length[i] = 255
            ws.column_dimensions[chr(ord('A') + i)].width = max_length[i] 
    # 保存
    wb.save(filename)
# 一个中文字符宽度为2 一个英文字符宽度为1
# 这样就可以计算出每一列的宽度了
def StringLength(str):
    length = 2
    for c in str:
        if ord(c) <= 256:
            length += 1
        else:
            length += 2
    return length 



# 常见的文件操作
def SaveFile(filename , data,append=False):
  if append:
    file = open(filename,"a",encoding="utf-8")
  else:
     file = open(filename,"w",encoding="utf-8")
  file.write(data)
  file.close()
# 将本地的html文件为bs4对象
def LoadHtml(filename):
    if not filename.endswith(".html"):
      print("文件格式不正确")
    file = open(filename,"r",encoding="utf-8")
    html = file.read()
    file.close()
    return BeautifulSoup(html,"html.parser")



# def GetResponse(url,headers=headers,params=params):
#   res = requests.get(url,headers=headers,params=params)
#   if res.status_code == 200:
#     # 其实也未必可以增
#     if res.text.startswith('{'):
#       return json.loads(res.text)
#     else:
#       return BeautifulSoup(res.text,'html.parser')
#   else: 
#     print("error, status_code : ",res.status_code)