// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewNode struct {
	Type       string                 `json:"type"`
	Attributes map[string]interface{} `json:"attributes"`
}

type Node struct {
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	Attributes map[string]interface{} `json:"attributes"`
}

type QueryNodes struct {
	Type string `json:"type"`
}