#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Define a data structure for an Elf.
struct Elf {
    int number;
    int cals;
};

// Make an array to keep track of elf numbers and their total calories.
struct Elf elves[1000];

// Read the input data into an array, with a blank line separating the entries
// between the calorie records for each elf.
int main() {
    FILE *fp;
    char *line = NULL;
    size_t len = 0;
    ssize_t read;
    int i, j, k, cals, total;
    fp = fopen("input.dat", "r");
    if (fp == NULL) {
        printf("Cannot open file\n");
        exit(EXIT_FAILURE);
    }
    i = 0;
    while ((read = getline(&line, &len, fp)) != -1) {
        // New elf.
        struct Elf elf = {i + 1, 0};
        // Calculate all this elf's calories.
        char *token = strtok(line, "\n");
        while (token != NULL) {
            cals = atoi(token);
            elf.cals += cals;
            token = strtok(NULL, "\n");
        }
        // Add this elf to the array.
        elves[i] = elf;
        i++;
    }
    fclose(fp);
    if (line) {
        free(line);
    }
    int n = i;
    // Sort the elf array descending by total calories.
    for (i = 0; i < n; i++) {
        for (j = i + 1; j < n; j++) {
            if (elves[j].cals > elves[i].cals) {
                struct Elf temp = elves[j];
                elves[j] = elves[i];
                elves[i] = temp;
            }
        }
    }
    // Part 1.
    struct Elf topElf = elves[0];
    printf("Part1: Elf %d has %d calories\n", topElf.number, topElf.cals);
    // Part 2.
    total = 0;
    for (k = 0; k < 3; k++) {
        total += elves[k].cals;
    }
    printf("Part2: Top three elves have %d calories\n", total);
    return 0;
}
