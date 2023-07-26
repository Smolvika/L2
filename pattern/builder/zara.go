package main

type ZaraBuilder struct {
	Brand    string
	Color    string
	Material string
	Length   int
	Size     int
}

func (z *ZaraBuilder) SetBrand() {
	z.Brand = "Zara"
}

func (z *ZaraBuilder) SetColor() {
	z.Color = "white"

}
func (z *ZaraBuilder) SetMaterial() {
	z.Material = "cotton"

}

func (z *ZaraBuilder) SetLength() {
	z.Length = 50
}

func (z *ZaraBuilder) SetSize() {
	z.Size = 42
}

func (z *ZaraBuilder) GetShirt() Shirt {
	return Shirt{
		Brand:    z.Brand,
		Color:    z.Color,
		Material: z.Material,
		Length:   z.Length,
		Size:     z.Size,
	}
}
