package routing

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/vpenando/piggy/piggy"
	"gorm.io/gorm"
)

var (
	//database            *gorm.DB
	operationController *piggy.OperationController
	categoryController  *piggy.CategoryController
)

func InitControllers(db *gorm.DB) {
	operationController, _ = piggy.NewOperationController(db)
	categoryController, _ = piggy.NewCategoryController(db)
}

func HandleError(w http.ResponseWriter, err error, status int) {
	if err == nil {
		return
	}
	log.Println("Error:", err)
	http.Error(w, err.Error(), status)
}

func ParseVarYear(vars map[string]string) (int, error) {
	return parseVarKey("year", vars)
}

func ParseVarMonth(vars map[string]string) (time.Month, error) {
	var month time.Month
	m, err := parseVarKey("month", vars)
	if err != nil {
		return month, err
	}
	if m < 1 || m > 12 {
		err = fmt.Errorf("invalid month %d", m)
	}
	return time.Month(m), err
}

func ServeImage(w http.ResponseWriter, r *http.Request, image string) {
	r.Header.Set("Content-Type", "image/png")
	http.ServeFile(w, r, image)
}

func parseVarKey(key string, vars map[string]string) (int, error) {
	if v, ok := vars[key]; ok {
		return strconv.Atoi(v)
	}
	return 0, errors.New("unexisting key")
}

func isPNG(rawFile []byte) bool {
	magic := []byte{0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a}
	if len(rawFile) < len(magic) {
		return false
	}
	for i := 0; i < len(magic); i++ {
		if rawFile[i] != magic[i] {
			return false
		}
	}
	return true
}

// RawCategory contains a category name and
type RawCategory struct {
	Name string
	Icon []byte
}

func saveCategoryIcon(rawCategoryIcon []byte, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	buffer := new(bytes.Buffer)
	err = binary.Write(buffer, binary.BigEndian, rawCategoryIcon)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(file, buffer)
	if err != nil {
		return err
	}
	return nil
}