
class Solution():
    def hammingWeight(self, n):
        """
        暴力解决法
        :type n: int
        :rtype: int
        """
        cnt = 0
        while(n > 0):
            if n&0x01 >0:
                cnt+=1
            n >>=1
        return cnt

    def hammingWeight1(self,n):
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
            cnt+=1
            n=n&(n-1)
        return cnt

    def hammingWeight2(self,n):
        """
        空间换时间了,列出8位所有的位表,应该是最快的
        :type n: int
        :rtype: int
        """
        pass




    

