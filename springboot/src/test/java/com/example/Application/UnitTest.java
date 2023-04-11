package com.example.Application;

import com.example.Application.dao.StudentDao;
import com.example.Application.domain.StudentDomain;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

@SpringBootTest
class UnitTest {
// @Resource 的作用就是可以自动注入 BookDao
  
  // 测试访问我们的/hello
  @Test
  public void testhello() {
    // 访问 localhost:8080/hello
    // 会返回 hello world
    
    // 这里其实可以作为就是数据库的调用测试地方就是说 哈哈哈
    
  }
  
  
//  @Autowired 还有一个作用就是可以自动注入我们的实体类
  // 包括我们的数据库操作类
  
  
  @Autowired
  private StudentDao studentDao;
  
  @Test
  public void TestDatabase() {
    StudentDomain studentDomain = new StudentDomain();
    studentDomain.setName("张三");
    studentDomain.setAge(18);
    studentDao.addStudent(studentDomain);
    
//    System.out.println(studentDao.getStudent(1));
    
  
  }
}
