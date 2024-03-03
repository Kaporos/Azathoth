package components

func RenderPrompt(input string, typing bool) string {
	cursor := ""
	if typing {
		cursor += "█"
	}
	return "> " + input + cursor
}
