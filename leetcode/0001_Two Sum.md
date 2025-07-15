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
	// 声明一个哈希表，用来存储
	s_map := make(map[int]int, len(nums))
	for i, num := range nums {
		// 临时变量 2: 计算差值
		complement := target - num

		// 检查差值是否在哈希表中
		if idx, found := s_map[complement]; found {
			return []int{idx, i} // 找到解
		}

		// 将当前数字存入哈希表
		s_map[num] = i
	}
	return []int{0, 0}
}

```