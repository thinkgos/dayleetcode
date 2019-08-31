import unittest
import time


class Solution():
    bits = [0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4,  # 0
            1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,  # 1
            1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,  # 1
            2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,  # 2
            1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,  # 1
            2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,  # 2
            2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,  # 2
            3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,  # 3
            1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,  # 1
            2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,  # 2
            2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,  # 2
            3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,  # 3
            2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,  # 2
            3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,  # 3
            3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,  # 3
            4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8]  # 4

    def wasteTime(func):
        def wrapper(self, *args):
            mark = time.time()
            result = func(self, *args)
            now = time.time()
            print("waste time: {:.4} s".format((now - mark)))
            return result

        return wrapper

    @wasteTime
    def hammingWeight(self, n: int) -> int:
        """
        暴力解决法
        :type n: int
        :rtype: int
        """
        cnt = 0
        while (n > 0):
            if n & 0x01 > 0:
                cnt += 1
            n >>= 1
        return cnt

    @wasteTime
    def hammingWeight1(self, n: int) -> int:
        """
        借位,主要利用2进制,向高位借位,而把1计算出来
        n = 10001000
        n-1 = 10000111
        n&(n-1) = 10000000
        :type n: int
        :rtype: int
        """
        cnt = 0
        while n != 0:
            cnt += 1
            n = n & (n - 1)
        return cnt

    @wasteTime
    def hammingWeight2(self, n: int) -> int:
        """
        空间换时间了,列出8位所有的位表,应该是最快的
        :type n: int
        :rtype: int
        """
        return self.bits[n & 0xff] + self.bits[(n >> 8) & 0xff] + self.bits[(n >> 16) & 0xff] + self.bits[
            (n >> 24) & 0xff]

    @wasteTime
    def hammingWeight3(self, n: int) -> int:
        """
        leetcode上看的,比上面还快,全内置函数！！！！
        :type n: int
        :rtype: int
        """
        s = bin(n)[2:]
        return s.count("1")


class TestSolution(unittest.TestCase):
    def setUp(self) -> None:
        self.obj = Solution()

    def test_hammingWeight(self):
        self.assertEqual(self.obj.hammingWeight(0x55), 4)
        self.assertEqual(self.obj.hammingWeight(0x8f), 5)
        self.assertEqual(self.obj.hammingWeight1(0x55), 4)
        self.assertEqual(self.obj.hammingWeight1(0x8f), 5)
        self.assertEqual(self.obj.hammingWeight2(0x55), 4)
        self.assertEqual(self.obj.hammingWeight2(0x8f), 5)
        self.assertEqual(self.obj.hammingWeight3(0x55), 4)
        self.assertEqual(self.obj.hammingWeight3(0x8f), 5)


if __name__ == '__main__':
    unittest.main()
