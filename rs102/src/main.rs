use rand::Rng;
use std::cmp::Ordering;
use std::io;

fn main() -> io::Result<()> {
    let mut a = 1;
    let mut b = 100;
    let v: Vec<u32> = std::env::args().skip(1).flat_map(|s| s.parse()).collect();
    if v.len() == 2 {
        a = v[0];
        b = v[1];
    }
    if v.len() == 1 {
        b = v[0];
    }

    let mut s = String::new();
    let t = io::stdin();

    let secret_number = rand::thread_rng().gen_range(a..=b);

    loop {
        println!("Please input your guess [{a}..{b}]:");
        s.clear();
        t.read_line(&mut s)?;

        let guess: u32 = match s.trim().parse() {
            Ok(num) => num,
            Err(_) => continue,
        };

        match guess.cmp(&secret_number) {
            Ordering::Less => println!("Too small!"),
            Ordering::Greater => println!("Too big!"),
            Ordering::Equal => {
                println!("You win!");
                break;
            }
        }
    }

    Ok(())
}
