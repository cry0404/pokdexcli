package main

import "testing"

func TestCleanInput(t *testing.T){
	cases := []struct {
		input string
		expected []string
	}{
		{ " hello world ", []string{"hello", "world"},},
		{"  hello world  ", []string{"hello", "world"}},
		{"Go is awesome", []string{"go", "is", "awesome"}},
		{"  leading and trailing spaces ", []string{"leading", "and", "trailing", "spaces"}},
		{"multiple   spaces", []string{"multiple", "spaces"}},
		{"  ", []string{}}, // 空白字符串
		{"", []string{}},   // 空字符串
		{"  word  ", []string{"word"}},
		{"Word", []string{"word"}}, // 测试大小写
		{"Word", []string{"word"}}, // 测试标点符号
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// 首先检查切片的长度是否一致
		if len(actual) != len(c.expected) {
			t.Errorf("For input '%s', expected %d words but got %d. Expected: %v, Got: %v",
				c.input, len(c.expected), len(actual), c.expected, actual)
			// 如果长度不一致，后面的逐个比较就没有意义了，可以直接跳过当前测试用例的剩余部分
			continue
		}
		// 逐个比较切片中的元素
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				// 使用 t.Errorf 报告错误，并提供详细信息
				t.Errorf("For input '%s', word at index %d is '%s', but expected '%s'. Full expected: %v, Full actual: %v",
					c.input, i, word, expectedWord, c.expected, actual)
				// 注意：t.Errorf 会标记测试失败，但会继续执行当前测试用例的剩余部分。
				// 如果你希望在发现第一个不匹配时就停止当前测试用例，可以使用 t.Fatalf。
				// t.Fatalf("For input '%s', word at index %d is '%s', but expected '%s'. Full expected: %v, Full actual: %v",
				// 	c.input, i, word, expectedWord, c.expected, actual)
			}
		}
	}
}