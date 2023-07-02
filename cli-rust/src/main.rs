use std::thread::sleep;

mod cli;


// 获取时间
struct TimeClock {
    start: std::time::Instant,

}

impl TimeClock {
    pub fn new() -> TimeClock {
        TimeClock {
            start: std::time::Instant::now(),
        }
    }
}
// 需要单独实现Drop trait
// trait 需要单独实现
impl Drop for TimeClock {
    fn drop(&mut self) {
        let end = std::time::Instant::now();
        let time = end.duration_since(self.start).as_millis();
        println!("cost time: {}ms", time);
    }
}


fn main() {
    // cli::app::Hello();

    let time = TimeClock::new();
    // 获取命令行参数
    let args: Vec<String> = std::env::args().collect();

    sleep(std::time::Duration::from_secs(1));


}