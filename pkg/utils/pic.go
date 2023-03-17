package utils

import (
	"bytes"
	"context"
	"github.com/nfnt/resize"
	"github.com/sirupsen/logrus"
	"image"
	"image/png"
)

// picCompress Modify size to compress pictures.
//func picCompress(picData []byte) (newPicData []byte, err error) {
//	oldBuf := bytes.NewBuffer(picData)
//	pic, _, err := image.Decode(oldBuf)
//	if err != nil {
//		return
//	}
//	newPic := resize.Resize(uint(pic.Bounds().Dx()/2), 0, pic, resize.Lanczos3)
//	var newBuf bytes.Buffer
//	err = png.Encode(&newBuf, newPic)
//	if err != nil {
//		return
//	}
//	newPicData, err = ioutil.ReadAll(&newBuf)
//	if err != nil {
//		return
//	}
//	return
//}

// CompressPicture the picture by resizing it to the specified width.
func CompressPicture(picData []byte) ([]byte, error) {
	// Decode the input image.
	oldBuffer := bytes.NewBuffer(picData)
	inputImg, _, err := image.Decode(oldBuffer)
	if err != nil {
		return nil, err
	}
	// Resize the image to the specified width.
	outputImg := resize.Resize(uint(inputImg.Bounds().Dx()/2), 0, inputImg, resize.Lanczos3)
	// Encode the output image as PNG and write it to a buffer.
	newBuffer := new(bytes.Buffer)
	err = png.Encode(newBuffer, outputImg)
	if err != nil {
		return nil, err
	}
	// Return the compressed image as a byte slice.
	return newBuffer.Bytes(), nil
}

// CompressPictureUntilSize 2*1024*1024
func CompressPictureUntilSize(ctx context.Context, picData []byte, maxSize int) (res []byte, err error) {
	if len(picData) <= maxSize {
		return picData, nil
	}
	picDataSize := len(picData)
	for round := 0; round < 5; round++ {
		if picDataSize > maxSize {
			res, err = CompressPicture(picData)
			if err != nil {
				logrus.WithContext(ctx).Errorln(err)
				break
			}
			picDataSize = len(res)
		}
	}
	return
}
