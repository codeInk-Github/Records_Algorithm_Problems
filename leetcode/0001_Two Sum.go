package main

//<p>Given an array of integers <code>nums</code>&nbsp;and an integer <code>target</code>, return <em>indices of the two numbers such that they add up to <code>target</code></em>.</p>
//
//<p>You may assume that each input would have <strong><em>exactly</em> one solution</strong>, and you may not use the <em>same</em> element twice.</p>
//
//<p>You can return the answer in any order.</p>
//
//<p>&nbsp;</p>
//<p><strong class="example">Example 1:</strong></p>
//
//<pre>
//<strong>Input:</strong> nums = [2,7,11,15], target = 9
//<strong>Output:</strong> [0,1]
//<strong>Explanation:</strong> Because nums[0] + nums[1] == 9, we return [0, 1].
//</pre>
//
//<p><strong class="example">Example 2:</strong></p>
//
//<pre>
//<strong>Input:</strong> nums = [3,2,4], target = 6
//<strong>Output:</strong> [1,2]
//</pre>
//
//<p><strong class="example">Example 3:</strong></p>
//
//<pre>
//<strong>Input:</strong> nums = [3,3], target = 6
//<strong>Output:</strong> [0,1]
//</pre>
//
//<p>&nbsp;</p>
//<p><strong>Constraints:</strong></p>
//
//<ul>
// <li><code>2 &lt;= nums.length &lt;= 10<sup>4</sup></code></li>
// <li><code>-10<sup>9</sup> &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
// <li><code>-10<sup>9</sup> &lt;= target &lt;= 10<sup>9</sup></code></li>
// <li><strong>Only one valid answer exists.</strong></li>
//</ul>
//
//<p>&nbsp;</p>
//<strong>Follow-up:&nbsp;</strong>Can you come up with an algorithm that is less than
//<code>O(n<sup>2</sup>)</code>
//<font face="monospace">&nbsp;</font>time complexity?
//
//<div><div>Related Topics</div><div><li>Array</li><li>Hash Table</li></div></div><br><div><li>👍 57208</li><li>👎 1997</li></div>

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	// nums []int 声明了一个类型为[]int 的切片（slice）变量。切片是Go语言中动态数组的抽象
	// 声明一个哈希表，用来存储
	sMap := make(map[int]int, len(nums))

	//range nums 会遍历切片 nums 的每个元素。
	// 每次迭代返回两个值：
	// i：当前元素的下标（从0开始）
	// num：当前元素的值拷贝（修改num不会影响原始切片）
	//如果切片为 nil，循环会安全执行 0 次（不会报错）。
	for i, num := range nums {

		// 计算差值
		complement := target - num

		// 检查差值是否在哈希表中
		// Go语言中 判断map中键是否存在的特殊写法，格式如下:
		// value, ok := map[key]
		// if 格式
		// if 初始化语句; 布尔表达式 {}   初始化语句未必就是定义变量， 如 println("init") 也是可以的。

		if idx, found := sMap[complement]; found {
			return []int{idx, i} // 找到解
		}

		// 将当前数字存入哈希表
		// 当前下标，存为value；当前对应的值，存为key
		sMap[num] = i
	}
	return []int{0, 0}
}

//leetcode submit region end(Prohibit modification and deletion)
