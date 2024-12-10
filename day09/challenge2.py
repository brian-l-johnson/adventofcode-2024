#!/usr/bin/env python3

import sys

class BlockGroup:
    def __init__(self, isFile: bool, size: int, value: int):
        self.isFile = isFile
        self.size = size
        self.value = value
    def __str__(self):
        if self.isFile:
            return f'{self.value}'*self.size
        else:
            return f'{"."}'*self.size

def PrintBlockGroup(bg):
    for b in bg:
        print(b, end="")
    print()
def scoreBlockGroup(bg):
    score = 0
    bc = 0
    for b in bg:
        if b.isFile:
            for i in range(b.size):
                score += (b.value * bc)
                bc += 1
                #print(f"{bc} * {b.value}")
        else:
            bc += b.size
    return score

def main():
    input = sys.stdin.readlines()
    if len(input) > 1:
        print("Too many lines")
        return

    infiles = True
    fc = 0
    blockRep = []
    for c in input[0]:
        i = int(c)
        if infiles:
            blockRep.append(BlockGroup(infiles, i, fc))
            fc+=1
            infiles = False
        else:
            blockRep.append(BlockGroup(infiles, i, -1))
            infiles = True
    
    i = len(blockRep)-1
    #PrintBlockGroup(blockRep)
    while i > 0:
        if blockRep[i].isFile:
            reloc = False
            for j in range(i):
                if not blockRep[j].isFile:
                    if blockRep[i].size <= blockRep[j].size:
                        blockRep[j].size -= blockRep[i].size
                        b = blockRep.pop(i)
                        blockRep.insert(j,b)
                        blockRep.insert(i, BlockGroup(False, b.size, -1))
                        #print(f"relocated block {i} to {j}")
                        reloc = True
                        break
            #PrintBlockGroup(blockRep)
        
        #if not reloc:
        #    print(f"unabled to relocate block {i}")
        i -= 1
    print("defragged")

    print(scoreBlockGroup(blockRep))


if __name__ == "__main__":
    main()