package pkg

import (
	"fmt"
)

// Scrape data
func Scrape(facultyCodes []string, startYear int, endYear int) {
	resultCh := make(chan *user, 50)
	fetchCount := 0
	for year := startYear; year <= endYear; year++ {
		for _, code := range facultyCodes {
			failedCount := 0
			for i := 1; i <= 999; i++ {
				nim := fmt.Sprintf("%s%02d%03d", code, year, i)
				go fetch(nim, resultCh)
				fetchCount++
				// waiting result, only when not reach max fetch
				deadlock := true
				for deadlock {
					select {
					case result := <-resultCh:
						if result == nil {
							failedCount++
							fmt.Println(failedCount)
						} else {
							// toleransi banyak failed berturut2
							if failedCount > 0 {
								failedCount--
							}
							fmt.Println(result.NimTPB)
							save(result)
						}
						fetchCount--
					default:
						if fetchCount <= 100-failedCount {
							deadlock = false
						}
						if failedCount > 50 {
							deadlock = false
						}
					}
				}
				if failedCount > 50 {
					break
				}
			}
		}
	}
	// waiting for all result done
	for fetchCount > 0 {
		println(fetchCount)
		result := <-resultCh
		if result != nil {
			fmt.Println(result.NimTPB)
			save(result)
		}
		fetchCount--
	}
}
