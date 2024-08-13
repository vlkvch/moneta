package models

import (
	"fmt"
	"strconv"
)

type Currency struct {
	Code  string  `json:"Cur_Abbreviation"`
	Rate  float64 `json:"Cur_OfficialRate"`
	Scale int     `json:"Cur_Scale"`
}

func (c Currency) String() string {
	return fmt.Sprintf("%s BYN", strconv.FormatFloat(c.Rate, 'f', 4, 32))
}
