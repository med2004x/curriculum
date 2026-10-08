#include <stdio.h>

struct Account { int balance; int open; };

int main(void) {
    struct Account a = {100, 1};
    if (a.open && a.balance >= 0) puts("valid");
    return 0;
}
