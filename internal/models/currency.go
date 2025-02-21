package models

import "fmt"

type Currency struct {
	Scale int     `json:"Cur_Scale"`
	Rate  float64 `json:"Cur_OfficialRate"`
	Code  string  `json:"Cur_Abbreviation"`
}

func (c Currency) String() string {
	return fmt.Sprintf("%.4f BYN", c.Rate)
}
