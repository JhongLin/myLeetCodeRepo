# https://leetcode.com/problems/rotate-array/

class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        n = len(nums)
        if n==1:
            return
        k %= n
        if k==0:
            return
        begins = math.gcd(n,k)
        for i in range(begins) :
            val = nums[i]
            cur = (i+k)%n
            while cur!=i:
                nums[cur], val = val, nums[cur]
                cur = (cur+k)%n
            nums[i] = val
                
        