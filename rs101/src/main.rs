fn main() {
    println!("{}", spin_words(",olleH !dlrow") == "Hello, world!");
}

fn spin_words(st: &str) -> String {
    let mut sb = String::new();
    let mut dlim = "";
    for s in st.split_whitespace() {
        sb += dlim;
        dlim = " ";
        if s.len() > 4 {
            let r: String = s.chars().rev().collect();
            sb += &r;
        } else {
            sb += s;
        }
    }
    sb
}

// Write a function that takes in a string of one or more words, and returns the same string, but with all five or more letter words reversed:
