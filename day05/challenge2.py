#!/usr/bin/env python3 

import re
import sys

def is_valid(rules, pages):
    for pi in range(len(pages)):
        if pages[pi] in rules:
            for r in rules[pages[pi]]:
                if r in pages:
                    ri = pages.index(r)
                    if ri < pi:
                        return False
    return True

def get_rule_break(rules, pages):
    for pi in range(len(pages)):
        if pages[pi] in rules:
            for r in rules[pages[pi]]:
                if r in pages:
                    ri = pages.index(r)
                    if ri < pi:
                        return [pi, ri]

def main():
    print("Day 4 Challenge 1")

    #read input from stdin
    lines = sys.stdin.readlines()

    rules = True

    #regular epxression to match rules, rules are numbers seperated by pipes
    rule_re = re.compile(r'^(\d+)\|(\d+)$')
    #regular expression to match page orderings, page orderings are comma separated numbers
    ordering_re = re.compile(r'^((\d+),)+(\d+)$')

    rules = {}
    sum = 0
    for line in lines:
        line = line.strip()
        if rule_re.match(line):
            r = line.split('|')
            if r[0] not in rules:
                rules[r[0]] = []
            rules[r[0]].append(r[1].strip())
        elif ordering_re.match(line):
            pages = line.split(',')
            if(is_valid(rules, pages)):
                pass
                #print(f"page ordering {pages} is valid")
                #get the middle value from the list of pages and add it to the sum of the valid page orderings
                #sum += int(pages[len(pages)//2])
            else:
                #print(f"page ordering {pages} is invalid")
                fixed_pages = pages.copy()
                while(not is_valid(rules, fixed_pages)):
                    a, b = get_rule_break(rules, fixed_pages)
                    tmp = fixed_pages[a]
                    fixed_pages[a] = fixed_pages[b]
                    fixed_pages[b] = tmp
                #print(f"fixed line {pages} to {fixed_pages}")
                sum += int(fixed_pages[len(fixed_pages)//2])
        elif line == '':
            pass
        else:
            print("invalid input")
            exit(-1)
    
    print(f"sum of valid page orderings is {sum}")
 

if __name__ == "__main__":
    main()