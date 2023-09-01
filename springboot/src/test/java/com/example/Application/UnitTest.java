package com.example.Application;


import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;
import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;

import javax.annotation.Resource;

// @SpringBootTest 的作用是
// 会自动加载SpringBoot的配置文件，然后启动SpringBoot的应用
// 也就是说，这个测试类是在SpringBoot的应用环境下运行的
// 我们可以直接去使用SpringBoot的一些特性
// 比如：自动注入
//@SpringBootTest
class UnitTest {

  @Test
  public void MainTest() throws JSONException {
    JSONArray jsonArray = new JSONArray();
    jsonArray.put(1);
    jsonArray.put("22");
    System.out.println(jsonArray);
  }
}
