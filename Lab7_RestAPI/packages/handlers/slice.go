package handlers

import (
	"Lab7_RestAPI/packages/templates"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func Slice(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, templates.PAGE_START, templates.BODY_START, templates.SLICE_PAGE)

	var inputStr string
	var slice []float64
	var err error

	if r.Method == "GET" {
		inputStr = r.FormValue("slice")

		if inputStr == "" {
			return
		}

		fmt.Fprint(w, "<p>Обраховано за допомогою Get-запиту:</p>")
	}

	if r.Method == "POST" {
		err = r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, templates.ERROR_TAG, err.Error())
			return
		}

		post := r.PostForm
		inputStr = post.Get("slice")

		if inputStr == "" {
			return
		}

		fmt.Fprint(w, "<p>Обраховано за допомогою Post-запиту:</p>")
	}

	slice, err = readSlice(inputStr)

	if err != nil {
		fmt.Fprintf(w, templates.ERROR_TAG, err.Error())
		return
	}

	multiplication := calculateMultiplicationEvenItems(slice)
	sum := calculateSumBetweenFirstZeroItemAndLastZeroItem(slice)

	responseStr := "<p></p><p><big>Результати: </big></p>"
	responseStr += fmt.Sprintf("<p><i>Добуток елементів з парними номерами</i>: <b>%.3f</b></p>", multiplication)
	responseStr += fmt.Sprintf("<p><i>Сума елементів, розташованих між першим і останнім нульовими \nелементами</i>: <b>%.3f</b></p>", sum)
	fmt.Fprint(w, responseStr)

	fmt.Fprint(w, templates.BODY_END)
	fmt.Fprint(w, templates.PAGE_END)
}

func readSlice(sliceStr string) ([]float64, error) {
	if sliceStr == "" {
		return nil, errors.New("empty input string")
	}

	elementsStr := strings.Split(sliceStr, " ")
	var result []float64

	for _, element := range elementsStr {
		number, err := strconv.ParseFloat(element, 32)

		if err != nil {
			return nil, err
		}

		result = append(result, number)
	}

	return result, nil
}

// Обраховуємо парні елементи, рахуючи їх від 1, тобто елемент з 1 індексом - другий і парний і так далі
func calculateMultiplicationEvenItems(slice []float64) float64 {
	if len(slice) <= 1 {
		return 0
	}

	result := 1.0
	for i := 1; i < len(slice); i += 2 {
		result *= slice[i]
	}

	return result
}

func calculateSumBetweenFirstZeroItemAndLastZeroItem(slice []float64) float64 {
	countZeroItems, firstZeroItemId, lastZeroItemId := findIndexesZeroItems(slice)

	if countZeroItems <= 1 {
		return 0
	}

	result := 0.0
	for i := firstZeroItemId; i <= lastZeroItemId; i++ {
		result += slice[i]
	}

	return result
}

func findIndexesZeroItems(slice []float64) (int, int, int) {
	firstZeroItemId := -1
	for i := 0; i < len(slice); i++ {
		if slice[i] == 0 {
			firstZeroItemId = i
			break
		}
	}

	lastZeroItem := -1
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == 0 {
			lastZeroItem = i
			break
		}
	}

	if lastZeroItem == firstZeroItemId {
		if lastZeroItem == -1 {
			return 0, firstZeroItemId, lastZeroItem
		} else {
			return 1, firstZeroItemId, -1
		}
	}

	return 2, firstZeroItemId, lastZeroItem
}
