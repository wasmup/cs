#include <algorithm>
#include <iostream>
#include <sstream>
#include <string>

std::string spinWords(const std::string &s) {
  std::stringstream result("");
  std::stringstream ss(s);
  std::string word;
  std::string d = "";
  while (ss >> word) {
    if (word.length() > 4)
      std::reverse(word.begin(), word.end());
    result << d << word;
    d = " ";
  }
  return result.str();
}

int main() {
  std::cout << (spinWords("Welcome") == "emocleW") << std::endl;       // 1
  std::cout << spinWords("to").compare("to") << std::endl;             // 0
  std::cout << spinWords("CodeWars").compare("sraWedoC") << std::endl; // 0
}