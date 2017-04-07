package accord

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHistoryStack(t *testing.T) {
	os.RemoveAll("history.stack")
	defer os.RemoveAll("history.stack")
	stack, err := OpenHistoryStack("history.stack")
	assert.Nil(t, err)

	assert.Zero(t, stack.Size())
	err = stack.Push(&Message{Payload: []byte{1}})
	assert.Nil(t, err)
	err = stack.Push(&Message{Payload: []byte{2}})
	assert.Nil(t, err)
	err = stack.Push(&Message{Payload: []byte{3}})
	assert.Nil(t, err)

	msg, err := stack.Peek()
	assert.Nil(t, err)

	assert.Equal(t, []byte{3}, msg.Payload)
	assert.Equal(t, uint64(3), stack.Size())

	msg, err = stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, []byte{3}, msg.Payload)
	assert.Equal(t, uint64(2), stack.Size())

	msg, err = stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, []byte{2}, msg.Payload)

	assert.Equal(t, uint64(1), stack.Size())

	msg, err = stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, []byte{1}, msg.Payload)

	assert.Equal(t, uint64(0), stack.Size())

	msg, err = stack.Pop()
	assert.Nil(t, err)
	assert.Nil(t, msg)

	msg, err = stack.Peek()
	assert.Nil(t, err)
	assert.Nil(t, msg)

	assert.Equal(t, uint64(0), stack.Size())

}

func TestHistoryStackClear(t *testing.T) {
	os.RemoveAll("history.stack")
	defer os.RemoveAll("history.stack")
	stack, err := OpenHistoryStack("history.stack")
	assert.Nil(t, err)

	err = stack.Push(&Message{Payload: []byte{1}})
	assert.Nil(t, err)

	assert.Equal(t, uint64(1), stack.Size())

	err = stack.Clear()
	assert.Nil(t, err)

	assert.Equal(t, uint64(0), stack.Size())
	err = stack.Push(&Message{Payload: []byte{1}})
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), stack.Size())

}
