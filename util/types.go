package util

import "time"

type User struct {
	ID        string `json:"id,omitempty" bson:"id,omitempty"`
	Firstname string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

type DBDetails struct {
	DbName         string
	CollectionName string
}

type HttpOptions struct {
	timeout time.Duration
}
