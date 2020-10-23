#!/usr/bin/env python3

import math

def process(mass):
    # to find the fuel required for a module, take its mass, divide by three, round down, and subtract 2.
    return math.floor(int(mass) / 3) - 2

with open('input.txt', 'rt') as f:
    fuel = 0
    for l in f.readlines():
        fuel += process(l)
print(fuel)
