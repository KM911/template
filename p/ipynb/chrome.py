# 这个的意义就是作为我们的selenium的模板 最好可以写一些说明文字不是吗? 这样就可以非常好地帮助你了


from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.chrome.options import Options
import time

# 手动设置chrome 路径

# 可以设置这里的config 吗? 

# def 

def CreateChromeSlient()-> webdriver.Chrome:
  # Chrome_Path = "D:\SOFT\Chrome\chrome.exe"
# 设置无头浏览器
  chrome_options = Options()
  # chrome_options.binary_location = Chrome_Path
  chrome_options.add_argument('--headless')
  # chrome_options.add_argument('--log-level=3')
  chrome_options.add_experimental_option('excludeSwitches', ['enable-logging'])



  # url = "https://www.kuwo.cn/play_detail/228908"
  driver = webdriver.Chrome(options=chrome_options)
  return driver
