import pathlib
import functools

test_data = """
abc

a
b
c

ab
ac

a
a
a
a

b
"""


def solve(data):
    groups = data.split("\n\n")
    groups = [g for g in groups if g]
    groups = [[p for p in g.splitlines() if p] for g in groups]
    groups = [[{a for a in p} for p in g] for g in groups]
    groups = [functools.reduce(set.intersection, g) for g in groups]
    print(groups)
    lengths = [len(g) for g in groups]
    print(lengths)
    summation = sum(lengths)
    print(summation)


solve(test_data)
solve(pathlib.Path("input.txt").read_text())
