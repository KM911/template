package com.example.Application.domain;


// 数据表的实体类 对应结构

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

//Repository 作用是将我们的实体类交给Spring管理 也就是说我们的实体类可以被Spring注入
@Data
@NoArgsConstructor
@AllArgsConstructor
public class StudentDomain {
  public int id;
  public String name;
  public int age;
  
  public StudentDomain(String name,int age){
    this.age = age;
    this.name = name;
  }
}
