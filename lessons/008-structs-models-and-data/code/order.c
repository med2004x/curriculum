#include <stdio.h>

enum Status { PENDING, PAID, CANCELLED };
struct Order { int id; enum Status status; };

int main(void) {
    struct Order order = {7, PAID};
    if (order.status == PAID) puts("ready");
    return 0;
}
