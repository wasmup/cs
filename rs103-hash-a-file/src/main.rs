use ring::digest::{Context, SHA256};
use std::fs::File;
use std::io::{BufReader, Read};

fn main() {
    let file_path = "src/main.rs";
    let file = match File::open(&file_path) {
        Ok(file) => file,
        Err(e) => {
            eprintln!("Error opening file: {}", e);
            return;
        }
    };
    let mut reader = BufReader::new(file);
    let mut context = Context::new(&SHA256);
    let mut buffer = [0; 4096];
    loop {
        let count = match reader.read(&mut buffer) {
            Ok(0) => break,
            Ok(n) => n,
            Err(e) => {
                eprintln!("Error reading file: {}", e);
                return;
            }
        };
        context.update(&buffer[..count]);
    }
    let digest = context.finish();
    let digest_string = digest
        .as_ref()
        .iter()
        .map(|b| format!("{:02x}", b))
        .collect::<String>();
    println!("{}", digest_string);
}
