#include <fmt/core.h>
#include <string>

int main() {
  std::string a[] = {"0", "1", "12", "121", "2", "1000"};
  std::sort(a->begin(), a->end());

  for (auto v : a) {
    fmt::print("{} ", v);
  }
  fmt::print("\n");
  // 0 1 12 121 2 1000
}