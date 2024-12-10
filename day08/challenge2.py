#!/usr/bin/env python3

import sys

def get_distance(p1, p2):
    return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1])
def get_slope(p1, p2):
    return (p1[0] - p2[0]), (p1[1] - p2[1])

def get_antipodes(p1, p2, max_row, max_col):
    ret = []
    s1 = get_slope(p1, p2)
    p = (p1[0]+s1[0], p1[1]+s1[1])
    ret.append(p1)
    ret.append(p2)
    while(point_is_in_board(p, max_row, max_col)):
        ret.append(p)
        p = (p[0]+s1[0], p[1] + s1[1])
    #ret.append((p1[0] + s1[0], p1[1] + s1[1]))
    s2 = get_slope(p2, p1)
    p = (p2[0]+s2[0], p2[1]+s2[1])

    while(point_is_in_board(p, max_row, max_col)):
        ret.append(p)
        p = (p[0]+s2[0], p[1] + s2[1])
    #ret.append((p2[0] + s2[0], p2[1] + s2[1]))
    return ret

def point_is_in_board(point, max_row, max_col):
    if(point[0] >= 0  and point[0] <= max_row and point[1]  >= 0 and point[1] <= max_col):
        return True
    else:
        return False


def main():
    antennas = {}
    input = sys.stdin.readlines()
    input = [x.strip() for x in input]
    r = 0
    for row in input:
        c = 0
        for s in row:
            if s != ".":
                if s not in antennas.keys():
                    antennas[s] = []
                antennas[s].append((r,c))
            c += 1
        r += 1

    max_row =r-1
    max_col =c-1
    print(antennas)
    print(f'max row: {max_row}, max col: {max_col}')

    antipode_set = {}
    for antena in antennas.keys():
        antenna_set = antennas[antena]
        print(f'antena: {antena}, count: {len(antenna_set)}')
        for i in range(0, len(antenna_set)):
            for j in range(i+1, len(antenna_set)):
                p1 = antenna_set[i]
                p2 = antenna_set[j]
                print(f'p1: {p1}, p2: {p2}')
                #print(f'distance: {get_distance(p1, p2)}')
                print(f'slope: {get_slope(p1, p2)}')
                antipodes = get_antipodes(p1, p2, max_row, max_col)
                for antipode in antipodes:
                    print(f'antipode: {antipode}, in board? {point_is_in_board(antipode, max_row, max_col)}')
                    if point_is_in_board(antipode, max_row, max_col):
                        print(f'antipode:  {antipode}')
                        antipode_set[antipode] = True

    print(len(antipode_set))



if __name__ == "__main__":
    main()

