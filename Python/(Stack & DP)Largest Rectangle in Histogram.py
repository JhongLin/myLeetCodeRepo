# https://leetcode.com/problems/largest-rectangle-in-histogram/

import numpy as np

class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        #print(len(heights))
        res = 0
        heights.append(0)
        stack = [-1]
        for i in range(len(heights)):
            while heights[i]<heights[stack[-1]]:
                yoko = heights[stack.pop()]
                tate = i - stack[-1] -1
                res = max(res, yoko*tate)
            stack.append(i)
        return res