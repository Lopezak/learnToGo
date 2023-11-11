import math
import os
import random
import re
import sys

#
# Complete the 'countingValleys' function below.
#
# The function is expected to return an INTEGER.
# The function accepts following parameters:
#  1. INTEGER steps
#  2. STRING path
#

def countingValleys(steps, path):
    valleyCount = 0
    position = 0
    inValley = False
    step = 0
    for step in range(0, steps):
        if path[step] == 'U':
            position += 1
        elif path[step] == 'D':
            position += -1
        
        if position <= -1 and not inValley:
            valleyCount += 1
            inValley = True
        elif position >= 0 and inValley:
            inValley = False
    return valleyCount

if __name__ == '__main__':
    steps = 10
    #path = "UDDDUDUU"
    #path = "UUUDUUDUDDDD"
    path = "DUDDDUUDUU"
    print(countingValleys(steps, path))