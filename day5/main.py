import json
import os
from pathlib import Path

class Stack:
    """ Represents one stack of crates """
    def __init__(self, number: int, contents: list[str]):
        self.number = number
        self.contents: list[str] = list()
        for stack_id in contents:
            self.contents.append(stack_id)

    def __str__(self) -> str :
        sb = f"StackNumber {self.number}: {','.join(self.contents)}"            
        return sb
        


# ============================================================
# Mainline
# ============================================================
if __name__ == '__main__':

    thisFile = Path(__file__).absolute()
    day5 = Path(thisFile).parent
    testdataFile = Path(day5).joinpath("testdata.json")
    with open(testdataFile) as fp:
        data = json.load(fp)

    # Load the stacks
    stacks = {}
    for stack_number, contents in data['stacks'].items():
        stacks[stack_number] = Stack(stack_number, contents)

