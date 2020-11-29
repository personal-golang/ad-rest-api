package ad

import "github.com/personal-golang/ad-rest-api/pkg/clients/ad"

func CreateNewGroup(groupName string) error {
	return adclient.CreateGroup(groupName)
}
