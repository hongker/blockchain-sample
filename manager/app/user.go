package app

import (
	"fmt"
	"github.com/manager/vendor"
)

// RegisterUser 用户注册
func RegisterUser() {
	setup := vendor.BaseSetupImpl{
		ConfigFile: "../" + vendor.ConfigFile,
	}

	setup, err := setup.InitConfig()
	if err != nil {
		fmt.Printf("Failed InitConfig [%s]\n", err)
		os.Exit(1)
	}
}

// EnrollUser 用户登录
func EnrollUser() {
	fmt.Println("Success")
}
