package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var stickersAmount int
	var initialSticker string
	
	fmt.Fscanf(in, "%s\n", &initialSticker)

	fmt.Fscan(in, &stickersAmount)

	for stickerIndex := 0; stickerIndex < stickersAmount; stickerIndex++ {
		var stickerStart, stickerEnd int
		var sticker string

		fmt.Fscan(in, &stickerStart, &stickerEnd)

		fmt.Fscanf(in, "%s\n", &sticker)

		initialSticker = initialSticker[:stickerStart-1] + sticker + initialSticker[stickerEnd:]
	}

	fmt.Fprintln(out, initialSticker)

	defer out.Flush()
}
