#!/usr/bin/env python3

import sys

def main():
    input = sys.stdin.readlines()
    if len(input) > 1:
        print("Too many lines")
        return

    #print(f'{input=}')
    sum = 0
    files = []
    gaps = []
    infiles = True
    for c in input[0]:
        i = int(c)
        sum += i
        if infiles:
            files.append(i)
            infiles = False
        else:
            gaps.append(i)
            infiles = True

    gi = 0
    #for i in range(len(files)):
    #    for j in range(files[i]):
    #        print(f"{i} ", end="")
    #    if gi < len(gaps):
    #        for j in range(gaps[gi]):
    #            print(". ", end="")
    #        gi += 1
    #print("")

    #print(f'{sum=}')
    #print(f"{files=}")
    #print(f"{gaps=}")

    gi = 0
    ef = len(files) - 1
    checksum = 0
    bc = 0
    #partial = ""
    for i in range(len(files)):
        if files[i] > 0 and i < ef:
            #checksum += bc * i * files[i]
            pc = 0
            for k in range(files[i]):
                pc += i*(bc+k)
            checksum+= pc
            bc += files[i]
            #print((f"{i}")*files[i], end="")
            #partial = partial + (f"{i}")*files[i]
            #print(partial)
            files[i] = 0
        if gi < len(gaps):
            for j in range(gaps[gi]):
                if files[ef] > 0:
                    #print(f"{ef}", end="")
                    checksum += bc * ef
                    bc+=1
                    #partial = partial + str(ef)
                    files[ef] -= 1
                    if files[ef] == 0:
                        ef -= 1
            gi +=1
        #print(partial)
        #print(files)
    print(f"\n{checksum}")
                


if __name__ == "__main__":
    main()