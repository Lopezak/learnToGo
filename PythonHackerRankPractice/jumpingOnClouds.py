#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'jumpingOnClouds' function below.
#
# The function is expected to return an INTEGER.
# The function accepts INTEGER_ARRAY c as parameter.
#

def jumpingOnClouds(c):
    jumps = 0
    index = 0

    while index < len(c) - 1: # 0 < 7
        if index + 2 < len(c) and c[index + 2] == 0: 
            index += 2
        else:
            index += 1
        
        jumps += 1
    
    return jumps

def chatGPTJumping(c):
    jumps = 0
    index = 0

    while index < len(c) - 1:
        # Check if it's safe to jump 2 steps
        if index + 2 < len(c) and c[index + 2] == 0:
            index += 2
        else:
            index += 1

        jumps += 1

    return jumps


if __name__ == '__main__':
    c = [0, 0, 1, 0, 0, 1, 0] # len = 7
    #    0  1  2  3  4  5  6
    d = [0, 0, 0, 1, 0, 0] # len = 6
    #    0  1  2  3  4  5
    print("Test 0 : ", jumpingOnClouds(c))
    print("Test 1 : ", jumpingOnClouds(d))

#0+1 <= 7 and c[1] == 0 || jumps = 1|| index = 1 
#1+2 <= 7 and c[3] == 0 || jumps = 2|| index = 3
#3+1 <= 7 and c[4] == 0 || jumps = 3|| index = 4
#4+2 <= 7 and c[6] == 0 || jumps = 4|| index = 