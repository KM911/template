package com.example.Application.Controller;

import com.dtflys.forest.annotation.Post;
import org.springframework.boot.configurationprocessor.json.JSONArray;
import org.springframework.boot.configurationprocessor.json.JSONException;
import org.springframework.boot.configurationprocessor.json.JSONObject;
import org.springframework.web.bind.annotation.*;

@RestController
public class Hello {
  @GetMapping("/hello")
  public String hello() {
    return "Hello World!";
  }

  @PostMapping("/list")
  public String list(@RequestBody String body) throws JSONException {
    JSONObject json = new JSONObject(body);
    JSONArray array = json.getJSONArray("list");
    for (int i = 0; i < array.length(); i++) {
      System.out.println(array.get(i));
    }

    return "Hello World!";
  }
}
