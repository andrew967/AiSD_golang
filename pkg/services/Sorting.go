package services

type Sorting interface {
	Sort(arr []float32) ([]float32, error)
}
