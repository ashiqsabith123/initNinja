package ui

func Bold(textToBold string) string {

	bold := "\033[1m"

	reset := "\033[0m"

	return bold + textToBold + reset
}
