import pathlib

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
    groups = [[p for p in g.splitlines() if p] for g in groups]
    groups = [''.join(g) for g in groups]
    groups = [{a for a in g} for g in groups]
    print(groups)
    lengths = [len(g) for g in groups]
    print(lengths)
    summation = sum(lengths)
    print(summation)


solve(test_data)
solve(pathlib.Path("input.txt").read_text())
