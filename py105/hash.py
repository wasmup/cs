import hashlib
import os

# Define the hash function to use (SHA256 in this example)


def hash_file(filename):
    sha256 = hashlib.sha256()
    with open(filename, "rb") as f:
        while True:
            data = f.read(65536)
            if not data:
                break
            sha256.update(data)
    return sha256.hexdigest()


# Create a list of all files in the current directory and subdirectories
files = []
for root, dirs, filenames in os.walk("."):
    for filename in filenames:
        files.append(os.path.join(root, filename))

# Sort the list of files by size (largest first)
files.sort(key=lambda x: os.path.getsize(x), reverse=True)

# Group the files by their hash values
hash_groups = {}
for file in files:
    hash_value = hash_file(file)
    if hash_value not in hash_groups:
        hash_groups[hash_value] = []
    hash_groups[hash_value].append(file)

# Write the output to a file
with open("hash.txt", "w") as f:
    for hash_value, files in hash_groups.items():
        f.write(hash_value + "\n")
        for file in files:
            f.write('"' + file + '"\n')
        f.write("\n")
