#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the minimumSwaps function below.
'''
Function intention is to take in an array of ints with no duplicates (no need to check)
How many swaps does it take to get this array sorted in asc order?
Looking for the minimum number of swaps:
    First thought is looking for a directed graph
    How do I determine the shortest way without having to check every single solution
    (bubble sort?)
    ex: step through
    [7,2,3,1,4,5,6] init
    [6,2,3,1,4,5,7] swap(0,6) 1
    [5,2,3,1,4,6,7] swap(0,5) 2
    [4,2,3,1,5,6,7] swap(0,4) 3
    [1,2,3,4,5,6,7] swap(0,3) 4

    ex2: step through
    [7,2,1,3,6,5,4] init
    [4,2,1,3,6,5,7] swap(0,6) 1
    [3,2,1,4,6,5,7] swap(0,3) 2
    [1,2,3,4,6,5,7] swap(0,2) 3
    
    for i in range(len(arr)-1):
        while arr[i] < arr[i+1] and arr[i] != len(arr)-1:

'''
def minimumSwaps(arr):
    swapCount = 0
    for i in range(len(arr)-1):
        while arr[i] != i + 1:
            temp = arr[i]
            arr[i] = arr[arr[i]-1]
            arr[temp-1] = temp
            swapCount += 1
    return swapCount


if __name__ == '__main__':
    arr = [4, 3, 1, 2]
    arr2 = [7,2,1,3,6,5,4]
    arr3 = [2,3,4,1,5]
    print("[4, 3, 1, 2] - ", minimumSwaps(arr)) 
    print("[7,2,1,3,6,5,4] - ", minimumSwaps(arr2)) 
    print("[2,3,4,1,5] - ", minimumSwaps(arr3)) 