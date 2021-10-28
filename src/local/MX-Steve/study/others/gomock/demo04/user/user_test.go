package user

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestReturn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockUserRepository(ctrl)
	repo.EXPECT().FindOne(1).Return(&User{Name: "lisi"}, nil)
	// repo.EXPECT().FindOne(2).Return(&User{Name: "wangwu"}, nil)
	// repo.EXPECT().FindOne(3).Return(nil, errors.New("user not found"))
	// log.Print(repo.FindOne(1))
	// log.Print(repo.FindOne(2))
	// log.Print(repo.FindOne(3))
	// log.Print(repo.FindOne(4))
}

// func TestOrder(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	repo := NewMockUserRepository(ctrl)
// 	o1 := repo.EXPECT().FindOne(1).Return(&User{Name: "张三"}, nil)
// 	o2 := repo.EXPECT().FindOne(2).Return(&User{Name: "李四"}, nil)
// 	o3 := repo.EXPECT().FindOne(3).Return(nil, errors.New("user not found"))
// 	gomock.InOrder(o1, o2, o3) //设置调用顺序
// 	// 按顺序调用，验证一下结果
// 	log.Println(repo.FindOne(1)) // 这是张三
// 	log.Println(repo.FindOne(2)) // 这是李四
// 	log.Println(repo.FindOne(3)) // user not found

// 	// 如果我们调整了调用顺序，将导致测试不通过：
// 	// log.Println(repo.FindOne(2)) // 这是李四
// 	// log.Println(repo.FindOne(1)) // 这是张三
// 	// log.Println(repo.FindOne(3)) // user not found
// }
