package adclient

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"log"
)

var ldapConn *ldap.Conn
var err error

func InitConn() {
	ldapUrl := "ldap://localhost:10389"
	ldapConn, err = ldap.DialURL(ldapUrl)
	if err != nil {
		log.Fatal(err)
	}
	err = ldapConn.Bind("uid=admin,ou=system", "secret")
	if err != nil {
		log.Fatal(err)
	}
}

func CloseConn() {
	ldapConn.Close()
}

func CreateGroup(groupName string) error {
	addReq := ldap.NewAddRequest(fmt.Sprintf("cn=%s,ou=roles,dc=wimpi,dc=net", groupName), []ldap.Control{})

	addReq.Attribute("objectClass", []string{"top", "groupOfNames"})
	addReq.Attribute("cn", []string{groupName})
	addReq.Attribute("member", []string{"uid=test,ou=users,dc=wimpi,dc=net"})

	err := ldapConn.Add(addReq)

	return err
}
