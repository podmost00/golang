package handlers

import (
	"Lab7_RestAPI/packages/templates"
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func QuadraticEquation(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, templates.PAGE_START, templates.BODY_START, templates.QUADRATIC_EQUATION_PAGE)

	var a float64
	var b float64
	var c float64
	var err error

	if r.Method == "POST" {
		err = r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, templates.ERROR_TAG, err.Error())
			return
		}

		post := r.PostForm

		aStr := post.Get("a")
		bStr := post.Get("b")
		cStr := post.Get("c")

		if aStr == "" && bStr == "" && cStr == "" {
			return
		}

		a, err = strconv.ParseFloat(aStr, 32)
		b, err = strconv.ParseFloat(bStr, 32)
		c, err = strconv.ParseFloat(cStr, 32)
		if err != nil {
			fmt.Fprintf(w, templates.ERROR_TAG, err.Error())
			return
		}

		fmt.Fprint(w, "<p>Обраховано за допомогою Post-запиту:</p>")
	}

	if r.Method == "GET" {
		aString := r.FormValue("a")
		bString := r.FormValue("b")
		cString := r.FormValue("c")

		if aString == "" && bString == "" && cString == "" {
			return
		}

		a, err = strconv.ParseFloat(aString, 32)
		b, err = strconv.ParseFloat(bString, 32)
		c, err = strconv.ParseFloat(cString, 32)

		if err != nil {
			fmt.Fprintf(w, templates.ERROR_TAG, err.Error())
			return
		}

		fmt.Fprint(w, "<p>Обраховано за допомогою Get-запиту:</p>")
	}

	resultStr := makeResultString(a, b, c)
	fmt.Fprintf(w, resultStr)

	fmt.Fprint(w, templates.BODY_END)
	fmt.Fprint(w, templates.PAGE_END)
}

func makeResultString(a float64, b float64, c float64) string {
	var result string

	roots, x1, x2 := calculate(a, b, c)
	if roots == 2 {
		result = fmt.Sprintf("<p><b>Результат:</b> <i>корінь1</i> = <b>%.3f</b> <i>корінь2</i> = <b>%.3f</b></p>", x1, x2)
	} else if roots == 1 {
		result = fmt.Sprintf("<p><b>Результат:</b> <i>корінь1</i> = <b>%.3f</b>", x1)
	} else {
		result = "<p><b>Результат:</b> розв'язків немає (D < 0)"
	}

	return result
}

func calculate(a float64, b float64, c float64) (int, float64, float64) {
	discriminator := b*b - 4*a*c

	var roots int

	var firstRoot float64
	var secondRoot float64
	if discriminator > 0 {
		roots = 2

		firstRoot = (-b + math.Sqrt(discriminator)) / (2 * a)
		secondRoot = (-b - math.Sqrt(discriminator)) / (2 * a)
	} else if discriminator == 0 {
		roots = 1

		firstRoot = (-b + math.Sqrt(discriminator)) / (2 * a)
	}

	return roots, firstRoot, secondRoot
}
