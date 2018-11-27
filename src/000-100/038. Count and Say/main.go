package main

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	lastStr := "1"
	for i := 2; i <= n; i++ {
		var res []byte
		lastByte := lastStr[0]
		num := 1
		lenght := len(lastStr)
		for j := 1; j < lenght; j++ {
			if lastStr[j] == lastByte {
				num++
			} else {
				res = append(res, byte(num+48))
				res = append(res, lastByte)
				lastByte = lastStr[j]
				num = 1
			}
		}
		res = append(res, byte(num+48))
		res = append(res, lastByte)
		lastStr = string(res)
	}
	return lastStr
}

func main() {

}
