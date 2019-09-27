#include<iostream>

std::string s = "stressed";

int main(void){

  for(unsigned char i = 0, j = s.size() - 1; i < j; ++i, --j){
    char tmp = s[i];
    s[i] = s[j];
    s[j] = tmp;
  }
  std::cout << s << '\n';

  return 0;
}