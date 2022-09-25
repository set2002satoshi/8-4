package models

import (
	"time"
	"errors"
)

type Options struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}


func NewOptions(
	createdAt time.Time,
	updatedAt time.Time,
)(*Options, error) {
	o := &Options{}
	if err := o.setCreatedAt(createdAt); err != nil {
		return nil, errors.New("setCreatedAtがエラー")
	}

	if err := o.setUpdatedAt(updatedAt); err != nil {
		return nil, errors.New("setCreatedAtがエラー")
	}
	return o, nil

}



func (o *Options) GetCreatedAt() time.Time {
	return o.CreatedAt 
} 

func (o *Options) GetUpdatedAt() time.Time {
	return o.UpdatedAt 
} 


func (o *Options) setCreatedAt(createdAt time.Time) error {
	o.CreatedAt = createdAt
	return nil
}

func (o *Options) setUpdatedAt(UpdatedAt time.Time) error {
	o.UpdatedAt = UpdatedAt
	return nil
}

