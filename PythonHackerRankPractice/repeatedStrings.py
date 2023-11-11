import math
import os
import random
import re
import sys

def repeatedString(s, n):
    resultCount = 0
    index = 0
    if len(s) == 1:
        if s[0] == 'a':
            return n
        else:
            return 0
    
    for index in range(0, len(s)):
        if s[index] == 'a':
            resultCount += 1
    iterations = math.floor(n/len(s))
    resultCount *= iterations
    remainder = n%len(s)
    if remainder != 0:
        index = 0
        for index in range(0, remainder):
            if s[index] == 'a':
                resultCount += 1
    return resultCount

if __name__ == '__main__':
    s = "aba"
    n = 10
    print(repeatedString(s, n))
