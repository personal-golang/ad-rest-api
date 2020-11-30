package ad_test

import (
	adclient "github.com/personal-golang/ad-rest-api/pkg/clients/ad"
	"github.com/personal-golang/ad-rest-api/pkg/components/ad"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewGroup(t *testing.T) {
	adclient.InitConn()
	defer adclient.CloseConn()
	adclient.DeleteGroup("test-group001")
	err := ad.CreateNewGroup("test-group001")
	assert.Equal(t, err, nil, "Creating a new ad group should not throw an error")
}