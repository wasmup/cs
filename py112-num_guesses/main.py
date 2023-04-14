import random

# Generate a random number between 1 and 100
number = random.randint(1, 100)

# Ask the user to guess the number
guess = int(input("Guess the number (between 1 and 100): "))

# Keep track of the number of guesses
num_guesses = 1

# Keep guessing until the user gets it right
while guess != number:
    # If the guess is too low, tell the user to guess higher
    if guess < number:
        print("Too low! Guess higher.")
    # If the guess is too high, tell the user to guess lower
    else:
        print("Too high! Guess lower.")
        
    # Ask the user to guess again
    guess = int(input("Guess again: "))
    
    # Increment the number of guesses
    num_guesses += 1
    
# If the user guessed the number, tell them how many tries it took
print("Congratulations! You guessed the number in", num_guesses, "tries.")
