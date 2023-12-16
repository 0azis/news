package pkg

import "github.com/lib/pq"

// Convert pqArray to []int
func ConvertToArray(categories pq.Int64Array) []int {
	var result []int

	for n := range categories {
		result = append(result, int(categories[n]))
	}
	return result
}
