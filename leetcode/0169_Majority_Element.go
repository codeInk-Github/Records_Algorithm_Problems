package main

import "sort"

/**
<p>Given an array <code>nums</code> of size <code>n</code>, return <em>the majority element</em>.</p>

<p>The majority element is the element that appears more than <code>⌊n / 2⌋</code> times. You may assume that the majority element always exists in the array.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>
<pre><strong>Input:</strong> nums = [3,2,3]
<strong>Output:</strong> 3
</pre>
<p><strong class="example">Example 2:</strong></p>
<pre><strong>Input:</strong> nums = [2,2,1,1,1,2,2]
<strong>Output:</strong> 2
</pre>
<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
 <li><code>n == nums.length</code></li>
 <li><code>1 &lt;= n &lt;= 5 * 10<sup>4</sup></code></li>
 <li><code>-10<sup>9</sup> &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
</ul>

<p>&nbsp;</p>
<strong>Follow-up:</strong> Could you solve the problem in linear time and in
<code>O(1)</code> space?

<div><div>Related Topics</div><div><li>Array</li><li>Hash Table</li><li>Divide and Conquer</li><li>Sorting</li><li>Counting</li></div></div><br><div><li>👍 21249</li><li>👎 743</li></div>
*/

// leetcode submit region begin(Prohibit modification and deletion)
func majorityElementSort(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func majorityElementHashMap(nums []int) int {
	sMap := make(map[int]int)
	for _, num := range nums {
		if idx, found := sMap[num]; found {
			sMap[num] = idx + 1
		} else {
			sMap[num] = 1
		}
	}
	var flag, max = 0, 0
	for k, v := range sMap {
		if flag < v {
			flag = v
			max = k
		}
	}
	return max
}

func majorityElementMooreVoting(nums []int) int {
	var max, vote = 0, 0
	for _, num := range nums {
		if vote == 0 {
			max = num
		}
		if max == num {
			vote++
		}
		if max != num {
			vote--
		}
	}

	return max
}

//leetcode submit region end(Prohibit modification and deletion)
