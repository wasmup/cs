package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"
)

func main() {
	t0 := time.Now()
	output := "hash.txt"
	dirs := os.Args[1:]
	if len(dirs) > 1 && dirs[0] == "-o" {
		output = dirs[1]
		dirs = dirs[2:]
	}
	if len(dirs) == 0 {
		dirs = append(dirs, ".")
	}

	err := run(output, dirs...)
	if err != nil {
		fmt.Println(err)
	}
	log.Println(time.Since(t0))
}

func run(output string, dirs ...string) error {
	scanner := &FileScanner{}

	r, err := scanner.Scan(dirs...)
	if err != nil {
		return fmt.Errorf("Scan failed:%w", err)
	}

	fmt.Println("\nWriting output to file:", output)
	f, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("Create file failed:%w", err)
	}

	for _, p := range r {
		fmt.Fprintln(f, p.Hash)
		for _, path := range p.Files {
			fmt.Fprintf(f, "%q\n", path)
		}
		fmt.Fprintln(f) // Empty line between hash groups
	}
	return nil
}

type FileInfo struct {
	Size  int64
	Hash  string
	Files []string
}

type FileScanner struct {
	m map[string]*FileInfo
}

func (f *FileScanner) visit(path string, info fs.DirEntry, err error) error {
	if err != nil || info.IsDir() {
		return err
	}

	fmt.Printf("scanning file %q\n", path)

	t0 := time.Now()
	hash, err := ComputeHash(path)
	log.Println(time.Since(t0))
	if err != nil {
		return err
	}

	p, exist := f.m[hash]
	if !exist {
		fi, err := info.Info()
		if err != nil {
			return err
		}

		p = &FileInfo{Size: fi.Size(), Hash: hash}
		f.m[hash] = p
	}

	p.Files = append(p.Files, path)
	return nil
}

func (f *FileScanner) Scan(directories ...string) ([]*FileInfo, error) {
	f.m = map[string]*FileInfo{}

	for _, directory := range directories {
		err := filepath.WalkDir(directory, f.visit)
		if err != nil {
			return nil, err
		}
	}

	a := make([]*FileInfo, 0, len(f.m))
	for _, p := range f.m {
		a = append(a, p)
	}

	f.m = nil

	if len(a) == 0 {
		return nil, nil
	}

	sort.Slice(a, func(i, j int) bool { return a[i].Size > a[j].Size })

	return a, nil
}

func ComputeHash(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
