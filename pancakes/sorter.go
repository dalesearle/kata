package pancakes

func SortStack(stack string) int {
	flips := 0
	runes := []rune(stack)
	curr := runes[0]
	for _, tmp := range runes {
		if curr != tmp {
			flips++
			curr = tmp
		}
	}
	if curr == '-' {
		flips++
	}
	return flips
}
