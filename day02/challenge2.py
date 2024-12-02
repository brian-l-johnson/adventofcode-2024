#!/usr/bin/env python3

import sys

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
    #read input from stdin line by line into a list of strings
    lines = [line.strip() for line in sys.stdin]
    #print(lines)
    safe = 0
    for line in lines:
        #split line into list of strings seperated by spaces and convert to integer
        parts = [int(part) for part in line.split(" ")]
        print(parts)
        if is_safe(parts):
            safe += 1
        else:
            for i in range(len(parts)):
                tp = parts.copy()
                print(f'removing item {i} from {parts}')
                tp.pop(i)
                if is_safe(tp):
                    safe += 1
                    break
    print(f'Safe: {safe}')
if __name__ == "__main__":
    main()
