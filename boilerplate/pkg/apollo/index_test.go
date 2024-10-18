package apollo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrganizationIndex(t *testing.T) {
	orgIndex := NewOrganizationIndex()

	orgIndex.add("app1", "channel1", 1)
	orgIndex.add("app1", "channel1", 2)
	orgIndex.add("", "", 3)

	assert.Equal(t, 3, len(orgIndex.query("app1", "channel1")), "Expected 3")
	assert.Equal(t, 1, len(orgIndex.query("", "")), "Expected 1")

	stats := orgIndex.stat()
	assert.Equal(t, 1, stats[KConstGlobal], "Expected 1")
	assert.Equal(t, 2, stats["app1_channel1"], "Expected 2")
}

func TestGlobalOrganizationIndex(t *testing.T) {
	globalOrgIndex := NewGlobalOrganizationIndex()

	globalOrgIndex.add("scene1", 1)
	globalOrgIndex.add("scene1", 2)
	globalOrgIndex.add("", 3)

	assert.Equal(t, 3, len(globalOrgIndex.query("scene1")), "Expected 3")
	assert.Equal(t, 1, len(globalOrgIndex.query("")), "Expected 1")

	stats := globalOrgIndex.stat()
	assert.Equal(t, 1, stats[KConstGlobal], "Expected 1")
	assert.Equal(t, 2, stats["scene1"], "Expected 2")
}

func TestProductIndex(t *testing.T) {
	productIndex := NewProductIndex()

	productIndex.add(KConstGlobal, "scene1", "app1", "channel1", 1)
	productIndex.add(KConstGlobal, "scene1", "app1", "channel1", 2)
	productIndex.add("org1", "scene2", "app2", "channel2", 3)

	assert.Equal(t, 2, len(productIndex.query(KConstGlobal, "scene1", "app1", "channel1")), "Expected 2")
	assert.Equal(t, 1, len(productIndex.query("org1", "scene2", "app2", "channel2")), "Expected 1")

	stats := productIndex.stat()
	assert.Equal(t, 0, stats[KConstGlobal].(map[string]interface{})[KConstGlobal], "Expected 0")
	assert.Equal(t, 1, stats["org1"].(map[string]interface{})["app2_channel2"], "Expected 1")
}
