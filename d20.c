#include <stdlib.h>
#include <stdio.h>
int main() {
  int x=arc4random_uniform(20)+1;
  printf("%d\n",x);
  return x;
}
