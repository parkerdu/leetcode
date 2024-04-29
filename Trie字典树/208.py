# coding:utf-8
class Trie:
    """实现一个字典树"""

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.root = {}
        self.end_of_word = '#'

    def insert(self, word: str) -> None:
        """
        Inserts a word into the trie.
        """
        node = self.root
        for w in word:
            # 我唯一不知道的就是这个地方，怎么插入的
            # 法一 node.setdefault(w,{})的意思是，如果node字典中有key w就把w对应得value返回出来，如果没有这个key就建立w:{}的键值对，并且返回当前的value
            # 如果原来node = {}
            # 使用node.setdefault之后，node={w:{}}
            # node = node.setdefault(w,{})

            # 法二 node.get方法的意思是，如果node字典中有key w就把w对应得value返回出来，如果没有这个key也不会建立w:{}的键值对，只会返回默认的value {}
            # 如果没有那就返回没有，但是他不会给你新建一个，所以如果原来node={}
            # 使用get之后，node={}
            node[w] = node.get(w,{})
            node = node[w]

        node[self.end_of_word] = self.end_of_word

    def search(self, word: str) -> bool:
        """
        Returns if the word is in the trie.
        """
        if not self.root:
            return False
        node = self.root
        for w in word:
            if not node.get(w,0):
                return False
            node = node[w]
        # 下面这4行代码换成一行就可以了
        # if node.get(self.end_of_word,0) == self.end_of_word:
        #     return True
        # else:
        #     return False
        return self.end_of_word in node

    def startsWith(self, prefix: str) -> bool:
        """
        Returns if there is any word in the trie that starts with the given prefix.
        """
        if not self.root:
            return False
        node = self.root
        for i in prefix:
            if not node.get(i,0):
                return False
            node = node[i]
        # 出来之后后面要么是结尾符，要么是其他字母，这两种情况就是true,所以只要出来就可以返回true
        return True


# Your Trie object will be instantiated and called as such:
# obj = Trie()
# obj.insert(word)
# param_2 = obj.search(word)
# param_3 = obj.startsWith(prefix)
if __name__ == "__main__":
    trie = Trie()
    trie.insert("hell")
    trie.insert("hello")
    print(trie.search("hell"))
    print(trie.search("helloa"))
    print(trie.search("hello"))

    print(trie.startsWith("hell"))
    print(trie.startsWith("helloa"))
    print(trie.startsWith("hello"))

    # trie.insert("app")
    # print(trie.search("app"))
    # print(trie.startsWith("app"))


