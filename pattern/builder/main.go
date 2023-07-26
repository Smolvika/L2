package main

func main() {
	// Создание интерфейсов строителя
	hmBuilder := GetBuilder("H&M")
	zaraBuilder := GetBuilder("Zara")

	director := NewDirector(hmBuilder)

	// Создание Бренда
	zaraShir := director.CreateDirector()
	zaraShir.InfoShirt()

	// Измениение Бренда
	director.SetBuilder(zaraBuilder)
	hmShir := director.CreateDirector()
	hmShir.InfoShirt()
}
