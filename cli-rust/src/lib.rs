

struct App {
    Name: String,
    Version: String,
    Author: String,
    Description: String,
    // UpdateTime: String,
    Commands: Vec<Command>,
    // Flags: Vec<Flag>,
}

// App的方法
impl App {
    // new

}




struct  Command{
    Name: String,
    Alias: String,
    Action: fn(),
    Description: String,
    // Flags: Vec<Flag>,
    // UpdateTime: String,
    // Commands: Vec<Command>,
}