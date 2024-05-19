package service

import "github.com/jinzhu/copier"

type CopierInterface interface {
	Copy(toValue interface{}, fromValue interface{}) (err error)
}

type Copier struct{}

func (c Copier) Copy(toValue interface{}, fromValue interface{}) (err error) {
	return copier.Copy(toValue, fromValue)
}
