def spin_words(sentence):
    return " ".join([s if len(s) < 5 else s[::-1] for s in sentence.split(' ')])


print(spin_words(",olleH !dlrow") == "Hello, world!")

# Write a function that takes in a string of one or more words, and returns the same string, but with all five or more letter words reversed:
