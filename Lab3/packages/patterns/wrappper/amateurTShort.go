package wrappper

type AmateurTShort struct {
	Number  int
	Wrapper TShortWrapper
}

func (a AmateurTShort) GetPrice() float64 {
	var coefficient float64

	if a.Number <= 10 {
		coefficient = 1.7
	} else {
		coefficient = 1.25
	}

	return a.Wrapper.GetPrice() * coefficient
}
