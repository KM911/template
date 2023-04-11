package com.example.Application.Respond;

import com.example.Application.domain.StudentDomain;
import lombok.Data;


@Data
public class StudentRespond extends Respond {
  public StudentDomain student;
  
  
  public StudentRespond(int code, String message, StudentDomain student) {
    super(code, message);
    this.student = student;
  }
}
