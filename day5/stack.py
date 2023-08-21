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

