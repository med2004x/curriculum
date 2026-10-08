#include <stdio.h>

typedef struct {
    double rows;
    double cost;
} Plan;

static Plan estimate(double rows, double selectivity, double page_cost) {
    /* TODO: encode the cost assumptions from the lecture, then compare two plans. */
    (void)rows;
    (void)selectivity;
    (void)page_cost;
    return (Plan){0, 0};
}

int main(void) {
    Plan p = estimate(100000, 0.01, 1.0);
    printf("rows=%.0f cost=%.2f\n", p.rows, p.cost);
    return 0;
}
