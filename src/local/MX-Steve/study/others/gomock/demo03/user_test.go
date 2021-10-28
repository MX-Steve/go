package d1

import (
	"testing"

	user_mock "local/MX-Steve/study/others/gomock/demo03/mock"

	"github.com/golang/mock/gomock"
)

func TestUser1(t *testing.T) {
	mockUser := user_mock.NewMockUser(gomock.NewController(t))
	mockUser.EXPECT().V(gomock.Any(), "abc").Return("aaaa", nil).AnyTimes()
	var u User = mockUser
	a, err := u.V(1, "abc")
	t.Log(a, err)
	mockUser.EXPECT().Name().Return("lisi").AnyTimes()
	name := u.Name()
	t.Log(name)
}
