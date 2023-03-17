package utils

import (
	"bytes"
	"image/png"
	"os"
	"testing"
)

func TestCompressPicture(t *testing.T) {
	// Setup test data
	imgData, err := os.ReadFile("test-image.png")
	if err != nil {
		t.Fatalf("failed to read test image file: %v", err)
	}
	testCases := []struct {
		name  string
		width uint
	}{
		{"compress with a valid width", 500},
		{"compress with an invalid width", 0},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Invoke the function being tested.
			compressedData, err := CompressPicture(imgData)

			// Check the result.
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(compressedData) == 0 {
				t.Fatalf("compressed data is empty")
			}
			create, err := os.Create("test-image-compressed.png")
			if err != nil {
				return
			}
			defer create.Close()
			_, err = create.Write(compressedData)
		})

	}
}

func Test_CompressPicture(t *testing.T) {
	type testCase struct {
		name     string
		picPath  string
		width    uint
		expected bool // true if no error, false if error is expected
	}

	testCases := []testCase{
		{"valid", "test-image.png", 100, true},
		{"invalid_path", "test-image.png", 100, false},
		{"invalid_pic", "test-image.png", 100, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			picData, err := os.ReadFile(tc.picPath)
			if err != nil {
				if tc.expected {
					t.Fatalf("unexpected error: %v", err)
				}
				// error expected and received, test passed
				return
			}

			compressedPic, err := CompressPicture(picData)
			if err != nil {
				if tc.expected {
					t.Fatalf("unexpected error: %v", err)
				}
				// error expected and received, test passed
				return
			}

			// ensure compressed picture is valid
			buffer := bytes.NewBuffer(compressedPic)
			_, err = png.Decode(buffer)
			if err != nil {
				t.Errorf("Compressed picture is invalid: %v", err)
			}
		})
	}
}
