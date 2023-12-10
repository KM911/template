package util

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
)

// top     ╭ ─ ┬ ╮
// middle  ├ ─ ┼ ┤
// content │   │ │
// bottom  ╰ ─ ┴ ╯

const (
	MaxLength = 28
	MinLength = 6
)

var (
	BlockBuilder      = strings.Builder{}
	TopLineString     = [4]string{"╭", "─", "┬", "╮\n"}
	MiddleLineString  = [4]string{"├", "─", "┼", "┤\n"}
	ContentLineString = [4]string{"│", " ", "│", "│\n"}
	BottomLineString  = [4]string{"╰", "─", "┴", "╯\n"}
)

func DrawBlock(title string, content []string) {
	currentMax := 0
	for _, v := range content {
		if len(v) > currentMax && currentMax < MaxLength {
			currentMax = len(v)
		}
	}
	DrawTop(currentMax)
	DrawTitle(color.Note.Render(title), currentMax)
	DrawMiddle(currentMax)
	for _, v := range content {
		DrawContent(v, currentMax)
	}
	DrawBottom(currentMax)
	fmt.Println(BlockBuilder.String())
	BlockBuilder.Reset()
}

func DrawTop(length int) {
	BlockBuilder.WriteString(TopLineString[0])
	for i := 0; i < length; i++ {
		BlockBuilder.WriteString(TopLineString[1])
	}
	BlockBuilder.WriteString(TopLineString[3])
}

func DrawTitle(title string, length int) {
	BlockBuilder.WriteString(ContentLineString[0])
	BlockBuilder.WriteString(TitleMiddleAlign(title, length))
	BlockBuilder.WriteString(ContentLineString[3])
}

func DrawContent(content string, length int) {
	BlockBuilder.WriteString(ContentLineString[0])
	BlockBuilder.WriteString(ContentLeftAlign(content, length))
	BlockBuilder.WriteString(ContentLineString[3])
}

func DrawMiddle(length int) {
	BlockBuilder.WriteString(MiddleLineString[0])
	for i := 0; i < length; i++ {
		BlockBuilder.WriteString(MiddleLineString[1])
	}
	BlockBuilder.WriteString(MiddleLineString[3])
}

func DrawBottom(length int) {
	BlockBuilder.WriteString(BottomLineString[0])
	for i := 0; i < length; i++ {
		BlockBuilder.WriteString(BottomLineString[1])
	}
	BlockBuilder.WriteString(BottomLineString[3])
}
func TitleMiddleAlign(ctx string, length int) string {
	ctxLength := ColorTextLength(ctx)
	return fmt.Sprintf("%*s%s%*s", (length-ctxLength)/2, "", ctx, (length - (length-ctxLength)/2 - ctxLength), "")
}

func ContentLeftAlign(ctx string, length int) string {
	ctxLength := len(ctx)
	return fmt.Sprintf("%s%*s", ctx, length-ctxLength, "")
}

func ColorTextLength(ctx string) int {
	return len(ctx) - 11
}

func CheckDataVaild(title []string, content [][]string) bool {
	return true
}
