import math
import os
import random
import re
import sys

#
# Complete the 'minimumBribes' function below.
#
# The function accepts INTEGER_ARRAY q as parameter.
#

'''
Conceptually:
Step 1, check if the place in line matches our index + 1 (cause 0 based)
Step 2, if does not match, how far from proper position?
    each step away from proper position would count as a bribe
logic check:
    [2,1,5,3,4]
    i = 0, position = 1:
        does not match, 1 away from position POSITIVELY (+1 bribe)
    i = 1, position = 2:
        does not match, -1 away from position, bribed (+0 bribe)
    i = 2, position = 3:
        does not match, 2 away from position, bribing (+2 bribes)
    i = 3, position = 4: 
        does not match, -1 away from position, bribed (+0 bribe)
    i = 4, position = 5:
        does not match, -1 away from position, bribed (+0 bribe)

    1 2 3 4 5 6 7 8
    1 2 5 3 4 6 7 8 +2
    1 2 5 3 7 4 6 8 +2
    1 2 5 3 7 8 4 6 +2
    1 2 5 3 7 8 6 4 +1    = 7 bribes
'''
def minimumBribes(q):
    sizeOfLine = len(q)
    bribeCount = 0
    for i in range(0, sizeOfLine):
        if q[i] - i+1 > 0:
            if q[i] - i+1 > 2:
                print("Too Chaotic")
                return
            bribeCount += 1
    
    print(bribeCount)

def minimumBribes2(q):
    n = len(q)
    numSwaps = 0

    for i in range(n):
        if q[i] - i > 3:
            # q[i] - i == 3 means i moved forward twice
            print("Too chaotic")
            return
        
        # to count bribes we count the number of elements to the
        # left of i that is greater than q[i]
        # but get timeout if we iterate all left sub-array (0:i-1) 
        
        # let k = index of left-most number > q[i] where k < i
        # k is index to the left of q[i] on original sorted array
        # Ex: [1, 2, 3, 4, 5] -> left-most possible number > 4 is
        # 5 at index 2, so we only iterate starting at index 2
        
        k = q[i] - 2
        
        if q[i] == 1:
            # we can't have k = -1 since its an index
            k = 0
        for j in range(k, i):
            # we add a bribe for every number > q[i]
            if q[j] > q[i]:
                numSwaps += 1

    print(numSwaps)
    return


if __name__ == '__main__':
    #q1 = [1,2,3,5,4,6,7,8] #1 bribe
    #q2 = [4,1,2,3] #Too Chaotic (at most 2 bribes are allowed per person)
    q1 = [2,1,5,3,4]
    q2 = [2,5,1,3,4]
    minimumBribes(q1)
    minimumBribes(q2)
    minimumBribes2(q1)
    minimumBribes2(q2)
    