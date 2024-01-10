use std::io::Read;

fn main() {
    // 获取命令行参数
    // exe path length
    let args: Vec<String> = std::env::args().collect();
    let filepath = &args[1];
    // 转成数字
    // 这样的错误处理不够好
    // 错误处理
    match args[2].parse::<usize>() {
        Ok(length) => {
               println!("length: {}", length);
            println!("content: {}", ReadFileContent(filepath, &length));
        },
        Err(e) => {
            print!("文件路径错误或者文件不存在");
            return;
        }
    };
    println!("filepath: {}", filepath);


    // println!("{:?}", args);
}
fn ReadFileContent(filepath: &str , length: &usize) -> String {
    let mut file = std::fs::File::open(filepath).unwrap();
    let mut content = String::new();
    file.read_to_string(&mut content).unwrap();
    content
}