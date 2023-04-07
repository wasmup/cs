use std::fs::File;
use std::io::{BufReader, Read, Result};
use std::path::{Path, PathBuf};

struct FileReader<'a> {
    file_path: &'a Path,
}

impl<'a> FileReader<'a> {
    fn new(file_path: &'a Path) -> Self {
        FileReader { file_path }
    }

    fn read_file(&self) -> Result<Vec<u8>> {
        let file = File::open(self.file_path)?;
        let mut buf_reader = BufReader::new(file);
        let mut file_contents = Vec::new();
        buf_reader.read_to_end(&mut file_contents)?;
        Ok(file_contents)
    }
}

trait Hasher {
    fn hash_file(&self, file_path: &Path) -> Result<Vec<u8>>;
}

struct Sha256Hasher;

impl Hasher for Sha256Hasher {
    fn hash_file(&self, file_path: &Path) -> Result<Vec<u8>> {
        let file_contents = FileReader::new(file_path).read_file()?;
        Ok(ring::digest::digest(&ring::digest::SHA256, &file_contents)
            .as_ref()
            .to_vec())
    }
}

struct FileSystemWalker<'a> {
    hasher: &'a dyn Hasher,
}

impl<'a> FileSystemWalker<'a> {
    fn new(hasher: &'a dyn Hasher) -> Self {
        FileSystemWalker { hasher }
    }

    fn walk_directory(&self, path: &Path) -> Result<Vec<(Vec<u8>, PathBuf)>> {
        let mut file_list: Vec<(Vec<u8>, PathBuf)> = Vec::new();
        for entry in path.read_dir()? {
            let entry = entry?;
            let file_path = entry.path();
            if file_path.is_dir() {
                let subdirectory_files = self.walk_directory(&file_path)?;
                file_list.extend(subdirectory_files);
            } else {
                let hash = self.hasher.hash_file(&file_path)?;
                file_list.push((hash, file_path));
            }
        }
        Ok(file_list)
    }
}

fn main() -> Result<()> {
    let hasher = Sha256Hasher {};
    let walker = FileSystemWalker::new(&hasher);
    let file_list = walker.walk_directory(Path::new("src"))?;
    let mut file_list = file_list
        .iter()
        .map(|(hash, path)| (path, hash))
        .collect::<Vec<_>>();
    file_list.sort_by(|a, b| {
        b.0.metadata()
            .unwrap()
            .len()
            .cmp(&a.0.metadata().unwrap().len())
    });

    let mut current_hash = Vec::new();
    for (path, hash) in file_list {
        if hash.ne(&current_hash) {
            println!();
            current_hash = hash.clone();
            println!("{}", hex::encode(&hash));
        }
        println!("{}", path.display());
    }
    Ok(())
}
