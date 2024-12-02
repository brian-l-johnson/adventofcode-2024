#!/usr/bin/env python3
import sys

safe = 3

def is_safe(readings):
    dir = 0
    safe = 3
    for i in range(1,len(readings)):
        v = readings[i] - readings[i-1]
        if i == 1:
            if v > 0:
                dir = 1
            else:
                dir = -1
        if v*dir > safe or v*dir <= 0:
            return False
    return True

def main():
    #read input from stdin line by line and split into a list of strings
    lines = [line.strip() for line in sys.stdin]
    #process data line by line, splitting line into a list of integers split by spaces
    safe_lines = 0
    for line in lines:
        #split the line into a list of strings and convert to integers
        parts = [int(x) for x in line.split(" ")]
        
        if is_safe(parts):
            safe_lines += 1

    print(safe_lines)

if  __name__ == "__main__":
    main()
