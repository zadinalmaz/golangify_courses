package main

import (
	"github.com/mowshon/moviego"
)

func main() {
	first, _ := moviego.Load("forest.mp4")

	// Создаем обычный скриншот.
	first.Screenshot(2, "simple-screen.png")

	// Применяем эффект FadeIn и FadeOut и скриншотим момент применения фильтра.
	first.FadeIn(0, 6).FadeOut(5).Screenshot(2, "screen.png")
}
