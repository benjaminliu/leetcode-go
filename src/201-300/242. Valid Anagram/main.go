package main

func isAnagram(s string, t string) bool {
	lenS, lenT := len(s), len(t)
	if lenS != lenT {
		return false
	}
	maps := make(map[byte]int)
	for i := 0; i < lenS; i++ {
		b := s[i]
		if value, ok := maps[b]; ok {
			maps[b] = value + 1
		} else {
			maps[b] = 1
		}
	}

	for i := 0; i < lenS; i++ {
		b := t[i]
		if value,ok := maps[b];ok{
			value --
			if value == 0{
				delete(maps,b)
			}else {
				maps[b] = value
			}
		}else {
			return false
		}
	}
	return true
}

func main() {

}
