package main

import (
	"strconv"
)

func manager(successChan chan bool, nimChan chan string) {
	for year := 16; year <= 20; year++ {
		for _, faculty := range code {
			for ratusan := 0; ratusan <= 9; ratusan++ {
				var fail = false
				for puluhan := 0; puluhan <= 9; puluhan++ {
					for satuan := 0; satuan <= 9; satuan++ {
						nimChan <- faculty + strconv.Itoa(year) + strconv.Itoa(ratusan) + strconv.Itoa(puluhan) + strconv.Itoa(satuan)
					}
					//evaluation
					var failCount = 0
					for i := 0; i <= 9; i++ {
						if !<-successChan {
							failCount++
						}
					}
					if failCount > 5 {
						fail = true
						break
					}
				}
				if fail {
					break
				}
			}
		}
	}
	close(nimChan)
}

var code = []string{
	"160", // FMIPA
	"161", // SITH
	"162", // SF
	"163", // FITB
	"164", // FTTM
	"165", // STEI
	"166", // FTSL
	"167", // FTI
	"168", // FSRD
	"169", // FTMD
	"190", // SBM
	"199", // SAPPK
}
