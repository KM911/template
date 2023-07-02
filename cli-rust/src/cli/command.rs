pub struct  Command{
    Name: String,
    Alias: String,
    Action: fn(),
    Description: String,
    // Flags: Vec<Flag>,
    // UpdateTime: String,
    // Commands: Vec<Command>,
}

// impl Command{
//     pub Execute(&self){
//         (self.Action)();
//     }
// }