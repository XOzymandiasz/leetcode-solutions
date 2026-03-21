#include "uthash.h"

typedef struct  {
    int key;
    int value;
    UT_hash_handle hh;
} HashItem;

int* twoSum(int* nums, int numsSize, int target, int* returnSize) {
    HashItem* map = NULL;

    for (int i = 0; i < numsSize; ++i) {
        int complement = target - nums[i];
        
        HashItem* item;
        HASH_FIND_INT(map, &complement, item);
        
        if (item != NULL) {
            int* result = malloc(2 * sizeof(int));
            result[0] = item->value;
            result[1] = i;
            *returnSize = 2;
            return result;
        }
        
        item = malloc(sizeof(HashItem));
        item->key = nums[i];
        item->value = i;
        HASH_ADD_INT(map, key, item);
    }

    
    *returnSize = 0;
    return NULL;
}