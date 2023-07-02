package com.example.Application.Http;


import com.dtflys.forest.annotation.Get;

public interface MyClient {
  @Get("http://localhost:8080/hello")
  String helloForest();
}

