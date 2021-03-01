package analysis

// Analysis holds methods for generating API responses
type Analysis struct{}

// CreateService creates an instance of the analysis service object
func CreateService() *Analysis {
	service := &Analysis{}

	return service
}
