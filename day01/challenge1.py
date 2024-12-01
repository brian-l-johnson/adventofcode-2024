#!/usr/bin/env python3
import sys
import re

def main():
    #read input from stdin
    input = sys.stdin.readlines()
    #remove whitespace characters like `\n` at the end of each line
    input = [x.strip() for x in input]
    #use a regular expression to change multiple spaces into one space
    input = [re.sub(r' +', ' ', x) for x in input]
    #iterate over input, converting to an int and storing into two lists, left and right
    left = []
    right = []
    for line in input:
        l,r = line.split(' ')
        left.append(int(l))
        right.append(int(r))
    #sort left and right lists
    left.sort()
    right.sort()
    #iterate over lists, comparing values at the same index
    diff = []
    sum = 0
    for i in range(len(left)):
        sum = sum + abs(left[i] - right[i])
    print(sum)

    

if __name__ == "__main__":
    main()
