package concurrent_test

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetSum(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMath := NewMockMath(ctrl)
	mockMath.EXPECT().Sum(3, 8).Return(11)
	// i := concurrent.GetSum(mockMath)
	// if i != 11 {
	// 	t.Errorf("Expected 11, got %d", i)
	// }

}
