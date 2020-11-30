package ad_test

import (
	"github.com/labstack/echo/v4"
	adclient "github.com/personal-golang/ad-rest-api/pkg/clients/ad"
	"github.com/personal-golang/ad-rest-api/pkg/components/ad"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setup() echo.Context {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetPath("/group/:group_name")
	context.SetParamNames("group_name")
	context.SetParamValues("test-grouphandler")
	return context
}

func TestNewGroupHandler(t *testing.T) {
	adclient.InitConn()
	defer adclient.CloseConn()
	adclient.DeleteGroup("test-grouphandler")
	context := setup()
	err := ad.NewGroupHandler(context)
	assert.Equal(t, err, nil, "NewGroupHandler no error response should not throw an error")

	err = ad.NewGroupHandler(context)
	assert.Equal(t, err, nil, "NewGroupHandler error response should not throw an error")
}
