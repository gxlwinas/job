package model

import (
	"github.com/jinzhu/gorm"
)

type jobtype uint8
type duration uint8

const (
	Short duration = iota
	Medium
	Long
)

const (
	Physical jobtype = iota
	Mental
)

type PartTimeJob struct {
	gorm.Model
	PublisherID uint
	Jobtype     jobtype
	Duration    duration
	Wages       string
}

type Job struct {
	gorm.Model
	PublisherID uint
	JobID       uint `json:"jobid"`
	Name        string
	Jobtype     jobtype
	Duration    duration
	Wages       string
	isEffective bool
}

type JobJson struct {
	gorm.Model
	PublisherID uint
	JobID       uint `json:"jobid"`
	Name        string
	Jobtype     jobtype
	Duration    duration
	Wages       string
	isEffective bool
}
