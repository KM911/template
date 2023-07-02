use crate::cli::command;
// mod crate::cli::command;



pub fn Hello(){
    println!("hello world")
}



struct App {
    Name: String,
    Version: String,
    Author: String,
    Description: String,
    // UpdateTime: String,
    Commands: Vec<command::Command>,
    // Flags: Vec<Flag>,
}

// App的方法
impl App {
    // new

}




