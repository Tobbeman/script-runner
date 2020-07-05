package runner

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testDir = "../../tests/scripts/accessible"

func TestRunner(t *testing.T) {
	r := New(testDir)
	res, err := r.Run("echo_lol.sh", []string{})
	if assert.Nil(t, err) {
		assert.Equal(t, "lol\n", res)
	}

	rCmd, err := r.RunAsync("sleep_echo_lol.sh", []string{})
	if assert.Nil(t, err) {
		assert.False(t, rCmd.CheckDone())
	}

	rCmd2, err := r.RunAsync("sleep_echo_lol.sh", []string{})
	assert.Nil(t, err)

	res, err = r.Run("sleep_echo_lol.sh", []string{})
	if assert.Nil(t, err) {
		assert.Equal(t, "lol\n", res)
	}

	res, err = rCmd2.Wait()
	if assert.Nil(t, err) {
		assert.Equal(t, "lol\n", res)
	}

	//Should always be true, since we wait on the same command above
	if assert.True(t, rCmd.CheckDone()) {
		res, err = rCmd.Collect()
		if assert.Nil(t, err) {
			assert.Equal(t, "lol\n", res)
		}
	}

	res, err = r.Run("does_not_exist.sh", []string{})
	assert.NotNil(t, err)

	res, err = r.Run("../not_accessible.sh", []string{})
	assert.NotNil(t, err)
}
