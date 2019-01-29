package main

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	total := (C-A)*(D-B) + (G-E)*(H-F)

	if min(C, G) > max(A, E) && min(D, H) > max(B, F) {
		return total - (min(C, G)-max(A, E))*(min(D, H)-max(B, F))
	} else {
		return total
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {

}
