// nolint
package testhelpers

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/ory/kratos-client-go"

	"github.com/ory/x/ioutilx"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/kratos/driver"
	"github.com/ory/kratos/driver/config"
	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/selfservice/flow/recovery"
	"github.com/ory/kratos/x"
)

func NewRecoveryUIFlowEchoServer(t *testing.T, reg driver.Registry) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		e, err := reg.RecoveryFlowPersister().GetRecoveryFlow(r.Context(), x.ParseUUID(r.URL.Query().Get("flow")))
		require.NoError(t, err)
		reg.Writer().Write(w, r, e)
	}))
	reg.Config(context.Background()).MustSet(config.ViperKeySelfServiceRecoveryUI, ts.URL+"/recovery-ts")
	t.Cleanup(ts.Close)
	return ts
}

func GetRecoveryFlow(t *testing.T, client *http.Client, ts *httptest.Server) *kratos.RecoveryFlow {
	publicClient := NewSDKCustomClient(ts, client)

	res, err := client.Get(ts.URL + recovery.RouteInitBrowserFlow)
	require.NoError(t, err)
	require.NoError(t, res.Body.Close())

	rs, _, err := publicClient.PublicApi.GetSelfServiceRecoveryFlow(context.Background()).Id(res.Request.URL.Query().Get("flow")).Execute()
	require.NoError(t, err, "%s", res.Request.URL.String())
	assert.Empty(t, rs.Active)

	return rs
}

func InitializeRecoveryFlowViaBrowser(t *testing.T, client *http.Client, ts *httptest.Server) *kratos.RecoveryFlow {
	publicClient := NewSDKCustomClient(ts, client)

	res, err := client.Get(ts.URL + recovery.RouteInitBrowserFlow)
	require.NoError(t, err)
	require.NoError(t, res.Body.Close())

	rs, _, err := publicClient.PublicApi.GetSelfServiceRecoveryFlow(context.Background()).Id(res.Request.URL.Query().Get("flow")).Execute()
	require.NoError(t, err)
	assert.Empty(t, rs.Active)

	return rs
}

func InitializeRecoveryFlowViaAPI(t *testing.T, client *http.Client, ts *httptest.Server) *kratos.RecoveryFlow {
	publicClient := NewSDKCustomClient(ts, client)

	rs, _, err := publicClient.PublicApi.InitializeSelfServiceRecoveryViaAPIFlow(context.Background()).Execute()
	require.NoError(t, err)
	assert.Empty(t, rs.Active)

	return rs
}

func GetRecoveryFlowMethodConfig(t *testing.T, rs *kratos.RecoveryFlow, id string) *kratos.RecoveryFlowMethodConfig {
	require.NotEmpty(t, rs.Methods[id])
	require.NotEmpty(t, rs.Methods[id].Config)
	require.NotEmpty(t, rs.Methods[id].Config.Action)
	c := rs.Methods[id].Config
	return &c
}

func RecoveryMakeRequest(
	t *testing.T,
	isAPI bool,
	f *kratos.RecoveryFlowMethodConfig,
	hc *http.Client,
	values string,
) (string, *http.Response) {
	require.NotEmpty(t, f.Action)

	res, err := hc.Do(NewRequest(t, isAPI, "POST", f.Action, bytes.NewBufferString(values)))
	require.NoError(t, err)
	defer res.Body.Close()

	return string(ioutilx.MustReadAll(res.Body)), res
}

// SubmitRecoveryForm initiates a registration flow (for Browser and API!), fills out the form and modifies
// the form values with `withValues`, and submits the form. If completed, it will return the flow as JSON.
func SubmitRecoveryForm(
	t *testing.T,
	isAPI bool,
	hc *http.Client,
	publicTS *httptest.Server,
	withValues func(v url.Values),
	method identity.CredentialsType,
	expectedStatusCode int,
	expectedURL string,
) string {
	hc.Transport = NewTransportWithLogger(hc.Transport, t)
	var f *kratos.RecoveryFlow
	if isAPI {
		f = InitializeRecoveryFlowViaAPI(t, hc, publicTS)
	} else {
		f = InitializeRecoveryFlowViaBrowser(t, hc, publicTS)
	}

	time.Sleep(time.Millisecond) // add a bit of delay to allow `1ns` to time out.

	config := GetRecoveryFlowMethodConfig(t, f, method.String())
	formPayload := SDKFormFieldsToURLValues(config.Nodes)
	withValues(formPayload)

	b, res := RecoveryMakeRequest(t, isAPI, config, hc, EncodeFormAsJSON(t, isAPI, formPayload))
	assert.EqualValues(t, expectedStatusCode, res.StatusCode, "%s", b)
	assert.Contains(t, res.Request.URL.String(), expectedURL, "%+v\n\t%s", res.Request, b)

	return b
}
