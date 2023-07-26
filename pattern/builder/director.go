package main

type Director struct {
	Builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{Builder: builder}
}

func (d *Director) SetBuilder(builder Builder) {
	d.Builder = builder
}

func (d *Director) CreateDirector() Shirt {
	d.Builder.SetBrand()
	d.Builder.SetColor()
	d.Builder.SetMaterial()
	d.Builder.SetLength()
	d.Builder.SetSize()
	return d.Builder.GetShirt()
}
