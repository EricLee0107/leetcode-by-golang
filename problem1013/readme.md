# 1010. 总持续时间可被 60 整除的歌曲

给你一个整数数组 A，只有可以将其划分为三个和相等的非空部分时才返回 true，否则返回 false。

形式上，如果可以找出索引 `i+1 < j` 且满足 `(A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1])` 就可以将数组三等分。

 

**示例 1：**
```
输入：[0,2,1,-6,6,-7,9,1,2,0,1]
输出：true
解释：0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1
```

**示例 2：**
```
输入：[0,2,1,-6,6,7,9,-1,2,0,1]
输出：false
```
**示例 3：**
```
输入：[3,3,6,5,-2,2,5,1,-9,4]
输出：true
解释：3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4
```

**提示：**

3 <= A.length <= 50000
-10^4 <= A[i] <= 10^4


# 题解
**细节是魔鬼：**

题目中说需要分为三段完全相等的,才会返回true，如果大于或小于三段都返回false。

但是有一个特殊情况就是每段总和为0，因为如果每段组合为0，此时段数大于3段也可以将多段0组合成一个段，因为`0+0=0`  （ps:好讽刺的加法计算） 所以当组合值为0时段数可以大于3段，其他情况段数必须严格为3才会返回true


