package com.example.Application.Router;

import com.example.Application.Respond.Respond;
import com.example.Application.Respond.StudentListRespond;
import com.example.Application.Respond.StudentRespond;
import com.example.Application.dao.StudentDao;
import com.example.Application.domain.StudentDomain;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;


//@RestController 是一个组合注解 作用是将类中的方法返回的数据转换为JSON格式
// 并且自动添加到HTTP响应的body中
// @RequestMapping 是一个组合注解 作用是将一个或多个请求映射到一个控制器的处理方法上 这个是类级别的 可以给类中的方法添加前缀
// @GetMapping 是一个组合注解 作用是将一个GET请求映射到一个控制器的处理方法上
// @RequestAttribute 是一个注解 作用是从请求中获取一个属性 比如IP地址

//
@RestController
@RequestMapping("/test")
public class TestRouter {
  @GetMapping("/hello")
  public Respond hello() {
    return new Respond(200, "Hello World!");
  }
  
  // @PathVariable 是一个注解 作用是从请求路径中获取一个参数 /hello/{name}/xxx 其中 name就是我们的参数
  @GetMapping("/hello/{name}")
  public Respond hello(@PathVariable String name) {
    return new Respond(200, "Hello " + name);
  }
  
  // @RequestParam 是一个注解 作用是从请求参数中获取一个参数 ?name=xxx&age=xxx 这种
  @GetMapping("/query")
  public Respond query(@RequestParam String name) {
    return new Respond(200, "Hello " + name);
  }
  
  
  @Autowired
  private StudentDao studentDao;
  @GetMapping("/add")
  public Respond add(@RequestParam String name, @RequestParam int age) {
    // 这里就是我们的数据库操作了
    StudentDomain stu = new StudentDomain();
    stu.setName(name);
    stu.setAge(age);
    studentDao.addStudent(stu);
    return new Respond(200, "Hello " + name + " " + age +"\n your id is " + stu.getId());
  }
  
  
  // 返回JSON的信息
  @GetMapping("/get/{id}")
  public StudentDomain get(@PathVariable int id) {
    return studentDao.getStudent(id);
  }
  
  // 解析JSON的信息
  @PostMapping("/update")
  public StudentRespond update(@RequestBody StudentDomain stu) {
    studentDao.updateStudent(stu);
    StudentRespond respond = new StudentRespond(200, "update success", stu);
    return respond;
  }
  
  // 返回JSON列表
  @GetMapping("list")
  public StudentListRespond StudentList(){
    StudentListRespond res = new StudentListRespond();
    StudentDomain[]  stus= new StudentDomain[5];
    // 向其中添加
    for (int i = 0; i < 5; i++) {
      stus[i] = new StudentDomain("dzg",21);
      studentDao.addStudent(stus[i]);
    }
    res.students = stus;
    res.code=200;
    res.message = "测试返回信息列表";
    return res;
  }
}
