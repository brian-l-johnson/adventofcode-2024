#!/usr/bin/env python3
import sys
import re

def main():
    #read input from stdin
    lines = "".join([line.strip() for line in sys.stdin])

    #a regular expression that captures every instance of mul() with two 1-3 digit numbers as arguments
    valid_exp = r"(mul\(\d{1,3}\s*,\s*\d{1,3}\)|don\'t\(\)|do\(\))"

    #a regular expression that extracts the two numeric arguments to mul()
    mul_exp = r"mul\(\s*(\d{1,3})\s*,\s*(\d{1,3})"

    #a regular expression to remove any instructions between a don't() and a do()
    disable_exp = r"(don\'t\(.*\).*?do\(.*?\))"
    
    print(lines)

    #match valid expressions
    expressions = [match.groups() for match in re.finditer(valid_exp, lines)]
    mul = True
    sum = 0
    for exp in expressions:
        print(exp[0])
        if exp[0] == "don't()":
            mul = False
        elif exp[0] == "do()":
            mul = True
        else:
            if not mul:
                continue
            match = re.match(mul_exp, exp[0])
            sum+= int(match.group(1)) * int(match.group(2))


    print(expressions)
    print(sum)

if __name__ == "__main__":
    main()

