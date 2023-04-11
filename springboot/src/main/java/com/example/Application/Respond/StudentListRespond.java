package com.example.Application.Respond;

import com.example.Application.domain.StudentDomain;
import lombok.NoArgsConstructor;


@NoArgsConstructor
public class StudentListRespond extends Respond{
  public StudentDomain[] students;
  
  public StudentListRespond(int code, String message, StudentDomain[] students) {
    super(code, message);
    this.students = students;
  }
}
