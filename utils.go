package ifonlyutils

import "fmt"

// GroupBySuit 通过花色进行分组
func GroupBySuit(cards []byte) [][]byte {
	result := make([][]byte, 5)
	// 初始化分组结果
	// 0: 方片, 1: 梅花, 2: 红桃, 3: 黑桃, 4: 王
	for i := range result {
		result[i] = []byte{}
	}

	// 遍历输入的牌面数据并分组
	for _, card := range cards {
		var suit int
		switch {
		case card >= 0x01 && card <= 0x0d: // 方片
			suit = 0
		case card >= 0x11 && card <= 0x1d: // 梅花
			suit = 1
		case card >= 0x21 && card <= 0x2d: // 红桃
			suit = 2
		case card >= 0x31 && card <= 0x3d: // 黑桃
			suit = 3
		case card == 0x4e || card == 0x4f: // 王
			suit = 4
		default:
			fmt.Println("...分组时产生未知值", card)
			continue // 如果有未知的值，跳过
		}
		// 将牌面值存储到对应的分组中
		result[suit] = append(result[suit], card)
	}
	return result
}

// GetStraightOnyByOne 一个接一个的获取
func GetStraightOnyByOne(cards []byte) (straight []byte, overCards []byte) {
	tempBytes := []byte{}
	for {
		if len(cards) <= 0 {
			break
		}

		if IsContinuous(cards) {
			// 是顺子
			return cards, tempBytes
		} else {
			tempBytes = append(tempBytes, cards[len(cards)-1])
			cards = cards[:len(cards)-1]
		}
	}
	return []byte{}, tempBytes
}

// IsContinuous 判断数据是否连续
func IsContinuous(nums []byte) bool {
	if len(nums) < 2 {
		return false
	}
	// TIPS: 这里无法判断Joker、是否连续
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			return false
		}
	}
	return true
}

// UniqueAndDuplicates 对byte切片进行去重，返回去重后的切片和重复的元素
func UniqueAndDuplicates(input []byte) (unique, duplicates []byte) {
	seen := make(map[byte]bool)
	addedDup := make(map[byte]bool)
	for _, b := range input {
		if !seen[b] {
			seen[b] = true
			unique = append(unique, b)
		} else if !addedDup[b] {
			duplicates = append(duplicates, b)
			addedDup[b] = true
		}
	}
	return
}
