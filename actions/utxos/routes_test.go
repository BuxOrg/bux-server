package utxos

import (
	"net/http"
	"testing"

	"github.com/BuxOrg/bux-server/config"
	"github.com/stretchr/testify/assert"
)

// TestUtxoRegisterRoutes will test routes
func (ts *TestSuite) TestUtxoRegisterRoutes() {
	ts.T().Run("test routes", func(t *testing.T) {

		// get utxo
		handle, _, _ := ts.Router.HTTPRouter.Lookup(http.MethodGet, "/"+config.CurrentMajorVersion+"/utxo")
		assert.NotNil(t, handle)
	})
}
