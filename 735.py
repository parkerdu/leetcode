# coding:utf-8
class Solution:
    """行星碰撞"""
    def asteroidCollision(self, asteroids):
        stack = []
        for i in range(len(asteroids)):
            if not stack:
                stack.append(asteroids[i])
            else:
                # 同号或者左负右正，没有碰撞
                if stack[-1] * asteroids[i] > 0 or (stack[-1] <0 and asteroids[i] > 0):
                    stack.append(asteroids[i])
                # 剩下就是左正右负，只有这时候会发生碰撞
                else:
                    if abs(stack[-1]) < abs(asteroids[i]):
                        # 循环结束之后，栈里面所有可能爆炸的炸了，除了asteroids[i]的相反数
                        while stack:
                            if stack[-1] < 0:
                                stack.append(asteroids[i])
                                break

                            elif stack[-1] + asteroids[i] == 0:
                                stack.pop()
                                break
                            elif stack[-1] > abs(asteroids[i]):
                                break
                            stack.pop()
                        if not stack:
                            stack.append(asteroids[i])
                    elif abs(stack[-1]) == abs(asteroids[i]):
                        stack.pop()

        return stack

if __name__ == "__main__":
    asteroids = [1,-2,-2,-2]
    su = Solution()
    print(su.asteroidCollision(asteroids))
