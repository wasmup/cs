#include <fmt/core.h>
#include <string>
#include <mutex>
#include <thread>
#include <chrono>
using namespace std::chrono_literals;

struct Shared {
  int value;
  std::mutex mux;
};
Shared shared{0,{}};


int main() {

  {
    auto t1 = std::jthread([&shared]{
      for (int i=0; i<10; i++){
        std::unique_lock lock(shared.mux);
        shared.value += 10;
      }
    });
    auto t2 = std::jthread([&shared]{
      for (int i=0; i<10; i++){
        std::unique_lock lock(shared.mux);
        shared.value += 1;
      }
    });
  }

  std::string a[] = {"0", "1", "12", "121", "2", "1000"};
  std::sort(a->begin(), a->end());

  for (auto v : a) {
    fmt::print("{} ", v);
  }
  fmt::print("\n");
  // 0 1 12 121 2 1000

    fmt::print("{} \n", shared.value);
}