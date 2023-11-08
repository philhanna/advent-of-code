#! /usr/bin/python
'''
The device will send your subroutine a datastream buffer (your puzzle
input); your subroutine needs to identify the first position where the
four most recently received characters were all different. Specifically,
it needs to report the number of characters from the beginning of the
buffer to the end of the first such four-character marker.

For example, suppose you receive the following datastream buffer:

mjqjpqmgbljsphdztnvjfqwrcgsmlb

After the first three characters (mjq) have been received, there have
not been enough characters received yet to find the marker. The first
time a marker could occur is after the fourth character is received,
making the most recent four characters mjqj. Because j is repeated, this
is not a marker.

The first time a marker appears is after the seventh character arrives.
Once it does, the last four characters received are jpqm, which are all
different. In this case, your subroutine should report the value 7,
because the first start-of-packet marker is complete after 7 characters
have been processed.
'''

def main():
    ''' Mainline '''
    FILENAME = "input.txt"
    with open(FILENAME) as fp:
        data = fp.read()
    p, prev = find_marker(data)
    print(f"{p=}, {prev=}")

def find_marker(s: str) -> int:
    ''' Find the position (zero-based) at which the previous four
    characters are all distinct '''
    p = 0
    for ch in s:
        if p >= 4:
            prev = s[p-4:p]
            if not dups(prev):
                return p, prev
        p += 1
    return 0

def dups(s: str) -> bool:
    ''' Return True if there are duplicate characters in this string '''
    for i in range(4):
        ci = s[i]
        for j in range(i+1, 4):
            cj = s[j]
            if ci == cj:
                return True
    return False

if __name__ == '__main__':
    main()
