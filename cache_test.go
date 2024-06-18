package storagehelper_test

import (
	"testing"

	"github.com/chayim/storagehelper"
	"github.com/stretchr/testify/assert"
)

func TestSetGetDelete(t *testing.T) {
	c := storagehelper.NewCache()
	key := "iamakey"
	val := "iamavalue"
	err := c.Set(key, val)
	assert.Nil(t, err)

	found, err := c.Get(key)
	assert.Nil(t, err)
	assert.Equal(t, found, val)

	err = c.Delete(key)
	assert.Nil(t, err)

	keys, err := c.GetKeys()
	assert.Nil(t, err)
	assert.Len(t, keys, 0)

	c.Set(key, val)
	keys, _ = c.GetKeys()
	assert.Len(t, keys, 1)

	err = c.Flush()
	assert.Nil(t, err)
	keys, _ = c.GetKeys()
	assert.Len(t, keys, 0)
}

func TestExists(t *testing.T) {
	c := storagehelper.NewCache()
	key := "iamakey"
	val := "iamavalue"
	err := c.Set(key, val)
	assert.Nil(t, err)

	f := c.Exists(key)
	assert.True(t, f)

	f = c.Exists("notarealkey")
	assert.False(t, f)

}
