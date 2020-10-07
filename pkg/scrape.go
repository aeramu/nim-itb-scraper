package pkg

import (
	"fmt"
)

// Scrape data
func Scrape(facultyCodes []string, startYear int, endYear int) {
	resultCh := make(chan *user, 50)
	for _, code := range facultyCodes {
		for year := startYear; year <= endYear; year++ {
			fetchCount := 0
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
							go save(result)
						}
						fetchCount--
					default:
						if fetchCount <= 100-failedCount {
							deadlock = false
						}
						if failedCount == 100 {
							deadlock = false
						}
					}
				}
				if failedCount == 100 {
					break
				}
			}
		}
	}
}
