package utils

import (
	"bytes"
	"image/jpeg"
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
		maxSize  int64
	}

	testCases := []testCase{
		{"valid", "test-image.png", 100, true, 1024 * 1024},
		{"invalid_path", "test-image.png", 100, false, 1024 * 1024},
		{"invalid_pic", "test-image.png", 100, false, 1024 * 1024},
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

			compressedPic, err := CompressPictureUntilSize(picData, 1024*1024)
			if err != nil {
				if tc.expected {
					t.Fatalf("unexpected error: %v", err)
				}
				// error expected and received, test passed
				return
			}

			// ensure compressed picture is valid
			buffer := bytes.NewBuffer(compressedPic)
			_, err = jpeg.Decode(buffer)
			if err != nil {
				t.Errorf("Compressed picture is invalid: %v", err)
			}
			if len(compressedPic) > int(tc.maxSize) {
				t.Errorf("Compressed picture is too large: %v", len(compressedPic))
			}
		})
	}
}

// test for TransferPicDataToJpg
func TestTransferPicDataToJpg(t *testing.T) {
	// Setup test data
	imgData, err := os.ReadFile("test-image.png")
	if err != nil {
		t.Fatalf("failed to read test image file: %v", err)
	}
	jpgImgData, err := os.ReadFile("test-image-tran.jpg")
	if err != nil {
		t.Fatalf("failed to read test image file: %v", err)
	}
	testCases := []struct {
		name    string
		width   int64
		imgData []byte
	}{
		{"tran", 500, imgData},
		{"tran", 0, jpgImgData},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Invoke the function being tested.
			compressedData, err := TransferPicDataToJpg(imgData)

			// Check the result.
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(compressedData) == 0 {
				t.Fatalf("compressed data is empty")
			}
			create, err := os.Create("test-image-tran.jpg")
			if err != nil {
				return
			}
			defer create.Close()
			_, err = create.Write(compressedData)
		})

	}
}
