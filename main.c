#include <stdio.h>
#include "STRING.h"

int main() {
    char* str1 = "hello";
    char* str2 = "hello";
    
    if (Equals(str1, str2)) {
        printf("Exact match\n");
    } else {
        printf("No match\n");
    }



    char* text = "The quick brown fox jumps over the lazy dog";
    char* word = "fox";
    
    if (Contains(text, word)) {
        printf("Text contains 'fox'\n");
    }


    char* strings[] = {"Hello", " ", "world", "!", NULL};
    
    // Concatenate without separator
    char* result1 = ConcatAll(strings);
    printf("Concatenated: %s\n", result1);
    free(result1);
    
    
    return 0;
}


/*#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "STRING.h"

int main() {
    int count = 4;
    StringArray arr = NewStringArray(count);
    if (!arr) {
        fprintf(stderr, "Failed to allocate string array\n");
        return 1;
    }

    // Assign strings (must copy them)
    arr[0] = strdup("Manzana");
    arr[1] = strdup("Banana");
    arr[2] = strdup("Uvas");
    arr[3] = strdup("Naranjas");
    // arr[4] is already NULL from NewStringArray

    int size = GetStringArraySize(arr);
    printf("Array contains %d elements:\n", size);

    for (int i = 0; i < size; i++) {
        printf("- %s\n", arr[i]);
    }

    FreeStringArray(arr); // Free all memory
    return 0;
}*/


/*#include <stdio.h>
#include "STRING.h"

int main() {
    char* str = "Hello,world,from,C";
    char* sep = ",";
    
    StringArray arr = Split(str, sep);
    if (!arr) {
        printf("Split failed\n");
        return 1;
    }
    
    int size = GetStringArraySize(arr);
    printf("Split into %d parts:\n", size);
    
    for (int i = 0; arr[i] != NULL; i++) {
        printf("%d: %s\n", i, arr[i]);
    }
    
    FreeStringArray(arr);
    return 0;
}*/