#include <stdlib.h>
#include <stdio.h>
int main() {
  int x=arc4random_uniform(12)+1;
  printf("%d\n",x);
  return x;
}