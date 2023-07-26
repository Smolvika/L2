package main

type HMBuilder struct {
	Brand    string
	Color    string
	Material string
	Length   int
	Size     int
}

func (hm *HMBuilder) SetBrand() {
	hm.Brand = "H&M"
}

func (hm *HMBuilder) SetColor() {
	hm.Color = "blac"

}
func (hm *HMBuilder) SetMaterial() {
	hm.Material = "cotton"
}

func (hm *HMBuilder) SetLength() {
	hm.Length = 30
}

func (hm *HMBuilder) SetSize() {
	hm.Size = 44
}

func (hm *HMBuilder) GetShirt() Shirt {
	return Shirt{
		Brand:    hm.Brand,
		Color:    hm.Color,
		Material: hm.Material,
		Length:   hm.Length,
		Size:     hm.Size,
	}
}
