package model

type WordBook struct {
	IntBase
	Word       string `json:"word" gorm:"column:c_word;size:32"`
	Annotation string `json:"annotation" gorm:"column:c_annotation;size:256"`
	Remark 	   string `json:"remark" gorm:"column:c_remark;size:256"`
}
