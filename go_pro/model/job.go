package model

import (
	"github.com/jinzhu/gorm"
)

type jobtype uint8
type duration uint8

const (
	Leaflet     jobtype = iota //派发传单
	Examine                    //审核
	Full                       //充场
	Customer                   //客服
	Assistant                  //会场协助
	Education                  //家教
	Photography                //摄影
	Clip                       //剪辑
	Physics                    //物理
)

//type PartTimeJob struct {
//	gorm.Model
//	PublisherID uint
//	Jobtype     jobtype
//	Duration    duration
//	Wages       string
//}

type Job struct {
	gorm.Model
	Useremail   string  `json:"useremail"`
	Name        string  `json:"name"`
	Jobtype     jobtype `json:"jobtype"`
	Time        string  `json:"time"`
	Wages       string  `json:"wages"`
	Address     string  `json:"address"`
	FullAddress string  `json:"fullAddress"`
	isEffective bool    `json:"isEffective"`
}

type Joblink struct {
	ID         uint
	Jobid      int
	Applyemail string
}
