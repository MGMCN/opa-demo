package middlewares

import (
	"bytes"
	"context"
	"fmt"
	"github.com/open-policy-agent/opa/sdk"
	sdktest "github.com/open-policy-agent/opa/sdk/test"
)

type OpaServer struct {
	ctx    context.Context
	server *sdktest.Server
	config []byte
}

func NewOpaServer(ctx context.Context) *OpaServer {
	return &OpaServer{ctx: ctx}
}

func (s *OpaServer) InitServer() error {
	var err error
	// create a mock HTTP bundle server
	s.server, err = sdktest.NewServer(sdktest.MockBundle("/bundles/bundle.tar.gz", map[string]string{
		"example.rego": `
				package authz

				import future.keywords.if
				
				default allow := false
				
				allow if input.body.type == "admin"
			`,
	}))
	if err != nil {
		return err
	}
	// provide the OPA configuration which specifies
	// fetching policy bundles from the mock server
	// and logging decisions locally to the console
	s.config = []byte(fmt.Sprintf(`{
		"services": {
			"test": {
				"url": %q
			}
		},
		"bundles": {
			"test": {
				"resource": "/bundles/bundle.tar.gz"
			}
		},
		"decision_logs": {
			"console": true
		}
	}`, s.server.URL()))
	return nil
}

func (s *OpaServer) Check(inputData map[string]interface{}) bool {
	// create an instance of the OPA object
	opa, err := sdk.New(s.ctx, sdk.Options{
		ID:     "opa-test-admin-1",
		Config: bytes.NewReader(s.config),
	})
	if err != nil {
		// handle error.
	}

	defer opa.Stop(s.ctx)

	// get the named policy decision for the specified input
	if result, err := opa.Decision(s.ctx, sdk.DecisionOptions{Path: "/authz/allow", Input: inputData}); err != nil {
		// handle error.
		return false
	} else if decision, ok := result.Result.(bool); !ok || !decision {
		// handle error.
		return false
	} else {
		return result.Result.(bool)
	}

}
