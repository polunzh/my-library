package main

import (
	"errors"
	"fmt"
	"image"
	"io"
	"log"
	"mime/multipart"

	"io/ioutil"
	"net/http"

	"image/jpeg"
	"image/png"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"

	"github.com/polunzh/my-library/dal"
)

func parseISBN(img image.Image) (string, error) {
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return "", err
	}

	reader := oned.NewEAN13Reader()
	hints := map[gozxing.DecodeHintType]interface{}{
		gozxing.DecodeHintType_TRY_HARDER: true,
	}

	result, err := reader.Decode(bmp, hints)
	if err != nil {
		return "", err
	}

	return result.String(), err
}

func decodeImage(imageType string, src io.Reader) (image.Image, error) {
	var img image.Image
	var err error
	switch imageType {
	case "image/jpeg":
		img, err = jpeg.Decode(src)
		break
	case "image/png":
		img, err = png.Decode(src)
		break
	default:
		err = errors.New("invalid image type")
	}

	return img, err
}

func detectImageType(content multipart.File) (string, error) {
	isbnBytes, err := ioutil.ReadAll(content)
	if err != nil {
		return "", err
	}

	imageType := http.DetectContentType(isbnBytes)
	content.Seek(0, 0)

	return imageType, nil
}

func parseISBNHandler(c *gin.Context) {
	isbnFile, err := c.FormFile("isbn")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("read image err:%s", err.Error()))
		return
	}

	imageContent, err := isbnFile.Open()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("read image err:%s", err.Error()))
	}
	defer imageContent.Close()

	imageType, err := detectImageType(imageContent)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("detect image content error:%s", err.Error()))
	}

	img, err := decodeImage(imageType, imageContent)

	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("decode image error:%s", err.Error()))
	}

	isbn, err := parseISBN(img)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("read image err:%s", err.Error()))
	}

	c.String(http.StatusOK, isbn)
}

func getBookByISBNHandler(c *gin.Context) {
	isbn := c.Param("isbn")

	// isbnKey := os.Getenv("ISBN_KEY")
	// resp, err := http.Get(fmt.Sprintf("https://http://feedback.api.juhe.cn/ISBN&sub=9781985086593&key=%s", isbnKey))
	// if err != nil {
	// 	c.String(http.StatusInternalServerError, fmt.Sprintf("request isbn error: %s", err.Error()))
	// }

	// resp.Body.Close()

	book, err := dal.FindByISBN(isbn)
	if err != nil {
		c.String(http.StatusInternalServerError, "find book by isbn error:%s", err.Error())
		return
	}

	c.JSON(http.StatusOK, book)
}

func addBook(c *gin.Context) {
	title := c.PostForm("title")
	isbn := c.PostForm("isbn")
	remark := c.PostForm("remark")
	_, err := dal.Insert(&dal.Book{Title: title, Isbn: isbn, Remark: remark})
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("add book error: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, "ok")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env")
	}

	router := gin.Default()
	router.MaxMultipartMemory = 1 << 20

	router.POST("/parse-isbn", parseISBNHandler)

	router.GET("/books/:isbn", getBookByISBNHandler)

	router.POST("/books", addBook)

	router.Run(":8080")
}
