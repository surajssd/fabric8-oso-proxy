package middlewares

import (
	"fmt"
	"testing"
)

func TestAuthShouldX(t *testing.T) {
	witURL := "https://api.prod-preview.openshift.io/api"
	authURL := "https://auth.prod-preview.openshift.io/api"
	srvAccID := "sa1"
	srvAccSecret := "secret"

	a := NewOSIOAuth(witURL, authURL, srvAccID, srvAccSecret)
	fmt.Println(a)
}
