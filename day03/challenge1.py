#!/usr/bin/env python3

import sys
import re

def main():
    #read input from stdin
    lines = [line.strip() for line in sys.stdin]

    #a regular expression that captures every instance of mul() with two 1-3 digit numbers as arguments
    valid_exp = r"(mul\(\d{1,3}\s*,\s*\d{1,3}\))"

    #a regular expression that extracts the two numeric arguments to mul()
    mul_exp = r"mul\(\s*(\d{1,3})\s*,\s*(\d{1,3})"

    sum = 0
    #find all instances of the above regex in the input
    for line in lines:
        matches = [match.groups() for match in re.finditer(valid_exp, line)]
        print(matches)
        for match in matches:
            #extract the arguments to mul from match
            mul_args = re.search(mul_exp, match[0]).groups()
        
            print(f'{mul_args}')
            sum += int(mul_args[0])*int(mul_args[1]) 


    print(sum)

if __name__ == "__main__":
    main()