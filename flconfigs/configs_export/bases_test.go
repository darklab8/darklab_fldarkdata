package configs_export

import (
	"testing"

	"github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped"
	"github.com/stretchr/testify/assert"
)

func TestExportBases(t *testing.T) {
	configs := configs_mapped.TestFixtureConfigs()
	exporter := NewExporter(configs)

	bases := exporter.Bases()
	assert.Greater(t, len(bases), 0)
	assert.NotEqual(t, bases[0].Nickname, bases[1].Nickname)
}
