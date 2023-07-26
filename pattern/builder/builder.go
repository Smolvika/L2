package main

const (
	HMBuilderType   = "H&M"
	ZaraBuilderType = "Zara"
)

type Builder interface {
	SetBrand()
	SetColor()
	SetMaterial()
	SetLength()
	SetSize()
	GetShirt() Shirt
}

func GetBuilder(builderType string) Builder {
	switch builderType {
	case HMBuilderType:
		return &HMBuilder{}
	case ZaraBuilderType:
		return &ZaraBuilder{}
	default:
		return nil
	}
}
