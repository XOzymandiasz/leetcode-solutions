#include <stdio.h>
#include <stdlib.h>
#include "solution.h"

void test_two_sum(int* nums, int numsSize, int target) {
    int returnSize = 0;
    int* result = twoSum(nums, numsSize, target, &returnSize); 

    if (result != NULL && returnSize == 2 && nums[result[0]] + nums[result[1]] == target){
        printf("PASS\n");
    } else {
        printf("FAIL\n");
        if (result != NULL) {
            printf("I've got: [%d, %d]\n", result[0], result[1]);
        }
    }

    free(result);
}

void test_no_solution(int* nums, int numsSize, int target) {
    int returnSize = 0;
    int* result = twoSum(nums, numsSize, target, &returnSize); 

    if (result == NULL && returnSize == 0){
        printf("PASS\n");
    } else {
        printf("FAIL\n");
        if (returnSize == 2) {
            printf("I've got: [%d, %d]\n", result[0], result[1]);
        }
    }
    free(result);
}

int main() {
    int nums1[] = {2, 7, 11, 15};
    test_two_sum(nums1, 4, 9);

    int nums2[] = {3, 2, 4};
    test_two_sum(nums2, 3, 6);

    int nums3[] = {3, 3};
    test_two_sum(nums3, 2, 6);
    
    int nums4[] = {7, 3, 4, -1};
    test_two_sum(nums4, 4, 2);
    
    int nums5[] = {7};
    test_no_solution(nums5, 1, 2);

    return 0;
}