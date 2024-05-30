package main

import (
	"container/list"
	"fmt"
)

// 检查两个单词是否只有一个字符不同
func oneCharDiff(a, b string) bool {
	count := 0
	for i := range a {
		if a[i] != b[i] {
			count++
			if count > 1 {
				return false
			}
		}
	}
	return count == 1
}

// 寻找从开始单词到目标单词的最短路径长度
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	if !wordSet[endWord] {
		return 0
	}

	queue := list.New()
	queue.PushBack(beginWord)
	visited := make(map[string]bool)
	visited[beginWord] = true

	level := 1
	for queue.Len() > 0 {
		levelSize := queue.Len()
		for i := 0; i < levelSize; i++ {
			currentWord := queue.Remove(queue.Front()).(string)
			if currentWord == endWord {
				return level
			}

			for word := range wordSet {
				if oneCharDiff(currentWord, word) && !visited[word] {
					visited[word] = true
					queue.PushBack(word)
				}
			}
		}
		level++
	}

	return 0
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println("Minimum chain length:", ladderLength(beginWord, endWord, wordList))
}
