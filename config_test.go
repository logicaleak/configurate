package configurate

import (
	"testing"
	//"flag"

	"github.com/stretchr/testify/assert"
)

func TestMustLoadShouldSucceedWithDifferentProfiles(t *testing.T) {
	MustLoad()

	value := GetString("key8")
	assert.Equal(t, "8", value)

	*profile = "local"
	MustLoad()

	value = GetString("key8")
	assert.Equal(t, "9", value)

	//Change it back to normal profile
	*profile = ""
}

func TestGetFloat(t *testing.T) {
	MustLoad()
	floatValue := GetFloat("key1.key4")
	assert.Equal(t, 3.5, floatValue)
}

func TestGetInt(t *testing.T) {
	MustLoad()
	intValue := GetInt("key1.key3")
	assert.Equal(t, 1, intValue)
}


func TestNestedGetString(t *testing.T) {
	MustLoad()
	stringValue := GetString("key5.key6.key7")
	assert.Equal(t, "7", stringValue)
}

