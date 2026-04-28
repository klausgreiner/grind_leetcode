package exercises

import (
	"fmt"
	"reflect"
)

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	originalColor := image[sr][sc]
	if originalColor == color {
		return image
	}
	return fill(image, sr, sc, originalColor, color)
}
func fill(image [][]int, r int, c int, originalColor int, newColor int) [][]int {
	if r < 0 || r >= len(image) || c < 0 || c >= len(image[0]) {
		return image
	}

	if image[r][c] != originalColor {
		return image
	}

	image[r][c] = newColor

	fill(image, r+1, c, originalColor, newColor) // Down
	fill(image, r-1, c, originalColor, newColor) // Up
	fill(image, r, c+1, originalColor, newColor) // Right
	fill(image, r, c-1, originalColor, newColor) // Left
	return image
}

func clone2DIntSlice(src [][]int) [][]int {
	if src == nil {
		return nil
	}
	dst := make([][]int, len(src))
	for i := range src {
		dst[i] = append([]int(nil), src[i]...)
	}
	return dst
}

// Test cases for Flood Fill (733)
var floodFillTestCases = []struct {
	image    [][]int
	sr       int
	sc       int
	color    int
	expected [][]int
}{
	{
		image:    [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
		sr:       1,
		sc:       1,
		color:    2,
		expected: [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
	},
	{
		image:    [][]int{{0, 0, 0}, {0, 0, 0}},
		sr:       0,
		sc:       0,
		color:    0,
		expected: [][]int{{0, 0, 0}, {0, 0, 0}},
	},
	{
		image:    [][]int{{0, 0, 0}, {0, 1, 1}},
		sr:       1,
		sc:       1,
		color:    1,
		expected: [][]int{{0, 0, 0}, {0, 1, 1}},
	},
	{
		image:    [][]int{{1}},
		sr:       0,
		sc:       0,
		color:    2,
		expected: [][]int{{2}},
	},
	{
		image:    [][]int{{1, 1, 1}, {1, 1, 1}},
		sr:       0,
		sc:       0,
		color:    9,
		expected: [][]int{{9, 9, 9}, {9, 9, 9}},
	},
}

// RunFloodFill executes the test cases for the floodFill function.
func RunFloodFill() {
	fmt.Println("--- Running Flood Fill (733) Tests ---")
	for _, tc := range floodFillTestCases {
		inputImage := clone2DIntSlice(tc.image)
		result := floodFill(inputImage, tc.sr, tc.sc, tc.color)

		fmt.Printf(
			"Input image: %v\nsr=%d sc=%d color=%d\nExpected:    %v\nGot:         %v\n",
			tc.image, tc.sr, tc.sc, tc.color, tc.expected, result,
		)

		if !reflect.DeepEqual(result, tc.expected) {
			fmt.Println("   ^^^ TEST FAILED ^^^ ")
		}
		fmt.Println("-----------------------")
	}
	fmt.Println("--- Finished Flood Fill (733) Tests ---")
}
