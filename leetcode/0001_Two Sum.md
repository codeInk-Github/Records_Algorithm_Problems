## Problem Description 【问题描述】

Given an array of integers `nums` and an integer `target`, return the indices of the two numbers such that they add up
to `target`.

**Assumptions:**

- Each input has exactly one solution.
- You may not use the same element twice.
- The answer can be returned in any order.

---

### Examples

**Example 1:**

- **Input:** `nums = [2, 7, 11, 15]`, `target = 9`
- **Output:** `[0, 1]`
- **Explanation:** `nums[0] + nums[1] = 2 + 7 = 9`.

**Example 2:**

- **Input:** `nums = [3, 2, 4]`, `target = 6`
- **Output:** `[1, 2]`
- **Explanation:** `nums[1] + nums[2] = 2 + 4 = 6`.

**Example 3:**

- **Input:** `nums = [3, 3]`, `target = 6`
- **Output:** `[0, 1]`

---

### Constraints

- `2 <= nums.length <= 10⁴`
- `-10⁹ <= nums[i] <= 10⁹`
- `-10⁹ <= target <= 10⁹`
- Only one valid answer exists.

---

### Follow-up

Can you come up with an algorithm with time complexity less than \(O(n^2)\)?

## Solution 【解题思路】

## Code

> Support Language: Java,Python,GOLang

### JAVA

### Python

### GOLang

```go
package main

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

```