package com.example.Application.Router;


import com.example.Application.Respond.Respond;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.context.request.RequestContextHolder;
import org.springframework.web.context.request.ServletRequestAttributes;

import javax.servlet.http.HttpServletRequest;

@RestController
@RequestMapping("/proxy")
public class Proxy {
  @GetMapping("/iptest")
  public Respond ipTest() {
    //
    HttpServletRequest request = ((ServletRequestAttributes) RequestContextHolder.getRequestAttributes()).getRequest();
    String ip = request.getRemoteAddr();
    return new Respond(200, "your ip is " + ip);
  }
}
