package techplay_test

import (
	"testing"

	"github.com/hiko1129/go-techplay/techplay"
	"github.com/stretchr/testify/assert"
)

// 	change it to a valid token to run test
const token = ""

func TestTechplay(t *testing.T) {
	c, err := techplay.New(token)
	assert.Nil(t, err)

	_, err = c.GetEventRanking()
	assert.Nil(t, err)
}
