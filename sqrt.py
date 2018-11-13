import sys
import math

def sqrt(a):
    return math.sqrt(a)

def output(result):
    sys.stdout.write("%f\n" % result)
    sys.stdout.flush()

while True:
    line = sys.stdin.readline()
    result = sqrt(float(line))
    output(result)
