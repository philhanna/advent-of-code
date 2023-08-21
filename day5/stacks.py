from .stack import Stack


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
