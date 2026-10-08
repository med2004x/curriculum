#include <stdio.h>

struct User {
    int id;
    int active;
};

int main(void) {
    struct User u = {42, 1};
    printf("%d %d\n", u.id, u.active);
    return 0;
}
