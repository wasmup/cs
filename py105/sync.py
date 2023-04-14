import os
import shutil

# Define the function to read the hash file and create a dictionary of files


def read_hash_file(hash_file):
    file_dict = {}
    with open(hash_file, "r") as f:
        lines = f.readlines()
        i = 0
        while i < len(lines):
            hash_value = lines[i].strip()
            i += 1
            file_list = []
            while i < len(lines) and lines[i].strip() != "":
                file_list.append(lines[i].strip('" \n'))
                i += 1
            file_dict[hash_value] = file_list
            i += 1
    return file_dict

# Define the function to sync files from the source folder to the destination folder


def sync_files(source_folder, destination_folder, hash_file):
    # Read the hash file and create a dictionary of files
    file_dict = read_hash_file(hash_file)

    # Copy new files from the source folder to the destination folder
    for root, dirs, filenames in os.walk(source_folder):
        for filename in filenames:
            file_path = os.path.join(root, filename)
            file_hash = hash_file(file_path)
            if file_hash not in file_dict:
                destination_path = os.path.join(destination_folder, os.path.relpath(file_path, source_folder))
                os.makedirs(os.path.dirname(destination_path), exist_ok=True)
                shutil.copy2(file_path, destination_path)


# Example usage:
hash_file = "hash.txt"
source_folder = "folder1"
destination_folder = "folder2"
sync_files(source_folder, destination_folder, hash_file)
