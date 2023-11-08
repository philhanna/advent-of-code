#! /usr/bin/python
'''
Your device's communication system is correctly detecting packets, but
still isn't working. It looks like it also needs to look for messages.

A start-of-message marker is just like a start-of-packet marker, except
it consists of 14 distinct characters rather than 4.

Here are the first positions of start-of-message markers for all of the
above examples:

mjqjpqmgbljsphdztnvjfqwrcgsmlb: first marker after character 19
bvwbjplbgvbhsrlpgdmjqwftvncz: first marker after character 23
nppdvjthqldpwncqszvftbrmjlhg: first marker after character 23
nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg: first marker after character 29
zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw: first marker after character 26

How many characters need to be processed before the first
start-of-message marker is detected?
'''

LIMIT = 14

def main():
    ''' Mainline '''
    FILENAME = "input.txt"
    with open(FILENAME) as fp:
        data = fp.read()
    p, prev = find_marker(data)
    print(f"{p=}, {prev=}")

def find_marker(s: str) -> int:
    ''' Find the position (zero-based) at which the previous 14
    characters are all distinct '''
    p = 0
    for ch in s:
        if p >= LIMIT:
            prev = s[p-LIMIT:p]
            if not dups(prev):
                return p, prev
        p += 1
    return 0

def dups(s: str) -> bool:
    ''' Return True if there are duplicate characters in this string '''
    for i in range(len(s)):
        ci = s[i]
        for j in range(i+1, len(s)):
            cj = s[j]
            if ci == cj:
                return True
    return False

if __name__ == '__main__':
    main()
