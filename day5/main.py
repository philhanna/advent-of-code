import json
import os
from pathlib import Path


class Stack:
    """ Represents one stack of crates """

    def __init__(self, contents: list[str]):
        self.contents: list[str] = list()
        for stack_id in contents:
            self.contents.append(stack_id)

    def pop(self) -> str:
        return self.contents.pop()
    
    def push(self, item: str):
        self.contents.append(item)


# ============================================================
# Mainline
# ============================================================
if __name__ == '__main__':
    this_file = Path(__file__).absolute()
    day5 = Path(this_file).parent
    testdata = Path(day5).joinpath("testdata.json")
    with open(testdata) as fp:
        data = json.load(fp)

    # Load the stacks
    stacks = {}
    for stack_number, contents in data['stacks'].items():
        stacks[stack_number] = Stack(contents)

    # Apply the moves
    for count, from_stack_number, to_stack_number in data['moves']:
        from_stack = stacks[str(from_stack_number)]
        to_stack = stacks[str(to_stack_number)]
        for i in range(count):
            item = from_stack.pop()
            to_stack.push(item)

