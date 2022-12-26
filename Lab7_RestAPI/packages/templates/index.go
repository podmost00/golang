package templates

const (
	INDEX_PAGE = `<h3>Lab 7 RestAPI</h3><p><a href="/quadratic-equation/">Завдання 1: Вирішити квадратне рівняння</a></p> <p><a href="/slice/">Завдання 2: Робота зі зрізами</a></p>`

	SLICE_PAGE = `<h3>Робота зі зрізами</h3><div><p>За допомогою Get-запиту</p><form action="/slice/" method="GET"><p><label for="a">Введіть елементи зрізу, використовуючи роздільник пробіл: </label>
								<input placeholder="5.2 7.8" name="slice" id="slice"/></p><p><input type="submit" value="Розрахувати"></form></div>
								<div><p>За допомогою Post-запиту</p><form action="/slice/" method="POST"><p><label for="a">Введіть елементи зрізу, використовуючи роздільник пробіл: </label>
								<input placeholder="5.2 7.8" name="slice" id="slice"/></p><p><input type="submit" value="Розрахувати"></form></div>`

	QUADRATIC_EQUATION_PAGE = `<h3>Вирішення квадратного рівняння</h3><div><p>За допомогою Get-запиту</p><form action="/quadratic-equation/" method="GET"><p><label for="a">Введіть коефіцієнт а: </label>
								<input placeholder="7" name="a" id="a"/></p><p><label for="b">Введіть коефіцієнт b: </label><input placeholder="3" name="b" id="b"/></p>
								<p><label for="c">Введіть коефіцієнт c: </label><input placeholder="-9" name="c" id="c"/></p><input type="submit" value="Розрахувати"></form></div>
								<div><p>За допомогою Post-запиту</p><form action="/quadratic-equation/" method="POST"><p><label for="a">Введіть коефіцієнт а: </label><input placeholder="-4" name="a" id="a"/></p><p>
								<label for="b">Введіть коефіцієнт b: </label><input placeholder="2" name="b" id="b"/></p><p>
								<label for="c">Введіть коефіцієнт c: </label><input placeholder="3" name="c" id="c"/></p>
								<input type="submit" value="Розрахувати"></form><div>`

	PAGE_START = `<!DOCTYPE HTML><html><head><title>Lab7 RestAPI</title><style>.error{color:#FF0000;}</style></head>`
	ERROR_TAG  = `<p class="error">%s</p>`
	BODY_START = `<body>`
	BODY_END   = `</body>`
	PAGE_END   = `</html>`
)
