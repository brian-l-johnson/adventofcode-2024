#!/usr/bin/env python3
import sys, re

def main():
    #read input from stdin
    input = sys.stdin.readlines()
    #remove whitespace characters like `\n` at the end of each line
    input = [x.strip() for x in input]
    #use a regular expression to change multiple spaces into one space
    input = [re.sub(r' +', ' ', x) for x in input]
    #iterate over input, converting to an int and storing left in a list and right in a map with the count of instances
    left = []
    right = {}
    for line in input:
        l,r = line.split(' ')
        l = int(l)
        r = int(r)
        left.append(l)
        if r in right:
            right[r] = right[r]+1
        else:
            right[r] = 1
    print(right)
    #calculate similarity score
    sum = 0
    for v in left:
        if v in right:
            sum = v*right[v] + sum
    print(sum)

if __name__ == '__main__':
    main()