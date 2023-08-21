import json
from pathlib import Path


class Stack:
    """ Represents one stack of crates """

    def __init__(self, contents: list[str]):
        self.contents: list[str] = []
        for stack_id in contents:
            self.contents.append(stack_id)

    def empty(self) -> bool:
        return len(self.contents) <= 0

    def pop(self) -> str:
        if self.empty():
            raise ValueError("Stack is empty")
        return self.contents.pop()

    def push(self, item: str):
        self.contents.append(item)

    def __str__(self):
        return " ".join(self.contents)


class Stacks:
    """ Represents all nine stacks """

    def __init__(self, data: dict[str, list[str]]):
        stacks: dict[int, Stack] = {}
        for k, contents in data.items():
            s: int = int(k)
            stacks[s] = Stack(contents)
        self.stacks = stacks

    def apply_move(self, move_number: int, move: list[int]):
        count: int
        index_f: int
        index_t: int
        count, index_f, index_t = move
        stack_f: Stack = self.stacks[index_f]
        stack_t:   Stack = self.stacks[index_t]

        # Verify that the "from" stack has enough items
        if len(stack_f.contents) < count:
            print(
                f"Error on move {move_number}: Stack {index_f} has too few items:")
            self.show_state()
            exit(-1)

        # Do the move
        for _ in range(count):
            item: str = stack_f.pop()
            stack_t.push(item)

    def show_state(self, label: str):
        print(label)
        for s in sorted(self.stacks):
            print(f"{s}: {self.stacks[s]}")
        print()


# ============================================================
# Mainline
# ============================================================
if __name__ == '__main__':

    # Get this directory
    this_file = Path(__file__).absolute()
    day5 = Path(this_file).parent

    # Get the input file path
    input_file = Path(day5).joinpath("input.txt")

    # Get the test data file file path
    testdata = Path(day5).joinpath("testdata.json")

    # Convert input to testdata
    with open(input_file) as fp:
        
        with open(testdata, "w") as out:
            pass


    # Load the stacks and moves from the test data file
    with open(testdata) as fp:
        data = json.load(fp)

    # Load the stacks
    stacks = Stacks(data["stacks"])
    stacks.show_state("Initial state:")

    # Apply the moves
    for i, move in enumerate(data["moves"]):
        stacks.apply_move(i, move)
        stacks.show_state(
            f"After {i} moves: move {move[0]} from stack {move[1]} to stack {move[2]}:")

    stacks.show_state("Final state:")

    # Display the elf label
    elf_label = ""
    for i in range(1, 10):
        stack = stacks.stacks[i]
        top_item = stack.pop()
        elf_label += top_item

    print(elf_label)
