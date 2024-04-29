class MyQueue:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.input_stack = []
        self.out_stack = []

    def push(self, x: int) -> None:
        """
        Push element x to the back of queue.
        """
        self.input_stack.append(x)



    def pop(self) -> int:
        """
        Removes the element from in front of queue and returns that element.
        """
        # 如果输出栈是空的时候就把所有的输出拿进来
        if not self.out_stack:
            for i in range(len(self.input_stack)):
                out = self.input_stack.pop()
                self.out_stack.append(out)
            return self.out_stack.pop()
        else:
            return self.out_stack.pop()

    def peek(self) -> int:
        """
        Get the front element.
        """
        if not self.out_stack:
            for i in range(len(self.input_stack)):
                out = self.input_stack.pop()
                self.out_stack.append(out)
            return self.out_stack[-1]
        else:
            return self.out_stack[-1]


    def empty(self) -> bool:
        """
        Returns whether the queue is empty.
        """
        # if self.input_stack ==[] and self.out_stack == []:
        #     return True
        # else:
        #     return False
        # 别人用一步替代我上面四步代码，用一个‘与’操作
        return (not self.inStack) and (not self.outStack)

    # Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()