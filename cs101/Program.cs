string SpinWords(string sentence) => string.Join(" ", from s in sentence.Split(" ") select s.Length > 4 ? new string(s.Reverse().ToArray()) : s);

Console.WriteLine(SpinWords("This is rehtona test"));



/* 
Write a function that takes in a string of one or more words, and returns the same string, but with all five or more letter words reversed:
 */