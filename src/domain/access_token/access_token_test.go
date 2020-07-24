package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T){
	assert.EqualValues(t, 24,TokenExpireTime,"token expire time should be 24")
	
	//    if TokenExpireTime !=24{
	// 		t.Error("token expire time should be 24")
	// 	}

}
func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t,at.IsExpired(),"New token should not be expired")
	assert.EqualValues(t, "",at.AccessToken,"access token should not be empty")
	assert.True(t, at.UserID==0,"new  accesstoken should not have any user id")

	// if at.IsExpired(){
	// 	t.Error("New token should not be expired")
	// }
	// if at.AccessToken != "" {
	// 	t.Error(" access token should not be empty")
	// }
	// if at.UserID !=0 {
	// 	t.Error("new  accesstoken should not have any user id")
	// }

}

func TestAccessTokenIsExpired(t *testing.T){
	at := AccessToken {}
    assert.True(t, at.IsExpired(),"empty access token should be expiired by default" )
	at.Expires = time.Now().UTC().Add( 3* time.Hour).Unix()
	assert.False(t,at.IsExpired(),"access token expiring three hours from now should NOT be expired " )
	
	// if !at.IsExpired() {
	// 	t.Error("empty access token should be expiired by default")
	// }
	
	// if at.IsExpired() {
	// 	t.Error("access token expiring three hours from now should NOT be expired ")
	// }

}