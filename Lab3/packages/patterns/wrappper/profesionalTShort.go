package wrappper

type ProfessionalTShort struct {
	Number  int
	Name    string
	Wrapper TShortWrapper
}

func (p *ProfessionalTShort) GetPrice() float64 {
	var coefficient float64

	if p.Number <= 12 {
		coefficient = 1.9
	} else {
		coefficient = 1.55
	}

	return p.Wrapper.GetPrice() * coefficient * 1.3
}
