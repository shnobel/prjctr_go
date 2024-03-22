package main

func main() {
	editor1 := NewEditor("text10.txt")
	editor2 := NewEditor("text100.txt")
	editor3 := NewEditor("text10k.txt")

	text := "duplicate"

	editor1.getRowsFast(text)
	editor2.getRowsFast(text)
	editor3.getRowsFast(text)

	editor1.getRowsSlow(text)
	editor2.getRowsSlow(text)
	editor3.getRowsSlow(text)
}
