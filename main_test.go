package main

import (
	"testing"
	// "time"
	// "github.com/davecgh/go-spew/spew"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	res := Call("Hello")
	assert.Equal(t, "Hello there.", res, "they should be equal")
	// assert.Equal(t, "", "", "they should be equal")
}