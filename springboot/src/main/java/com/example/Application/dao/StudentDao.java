package com.example.Application.dao;

import com.example.Application.domain.StudentDomain;
import org.apache.ibatis.annotations.*;

// @Mapper 是一个注解 作用是将接口标记为一个Mapper接口
// Mapper接口是一个特殊的接口 他的实现类是由Mybatis框架自动生成的
// 作用是将我们的接口中的方法映射到SQL语句上
// 也就是说我们的接口中的方法就是我们的数据库操作了


@Mapper
public interface StudentDao {
  
  // 这里就是我们的数据库操作了
  // 添加学生信息
  
  /**
   *  添加学习信息并且返回id
   * @param stu
   */
  @Options(useGeneratedKeys = true, keyProperty = "id")
  @Insert("insert into students(name, age) values(#{name}, #{age})")
  public void addStudent(StudentDomain stu);
  
  // 查看学生信息
  @Select("select * from students where id = #{id}")
  public StudentDomain getStudent(int id);
  
  // 更新学生信息
  @Update("update students set name = #{name}, age = #{age} where id = #{id}")
  public void updateStudent(StudentDomain stu);
}
