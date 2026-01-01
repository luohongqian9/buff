package flags

import (
	"errors"
	"fmt"
	"os"
	"server/global"
	"server/model/appTypes"
	"server/model/database"
	"server/utils"
	"syscall"

	"github.com/google/uuid"
	"golang.org/x/crypto/ssh/terminal"
)

func Admin() error {
	var user database.User

	fmt.Println("请输入邮箱:")
	var email string
	_, err := fmt.Scanln(&email)
	if err != nil {
		return fmt.Errorf("输入错误: %v", err)
	}
	user.Email = email

	fd := int(syscall.Stdin)
	oldState, err := terminal.GetState(fd)
	if err != nil {
		return fmt.Errorf("获取终端状态失败: %v", err)
	}
	defer terminal.Restore(fd, oldState)

	fmt.Println("请输入密码:")
	password, err := readPassword()
	fmt.Println()
	if err != nil {
		return fmt.Errorf("读取密码失败: %v", err)
	}

	fmt.Println("请再次输入密码:")
	rePassword, err := readPassword()
	fmt.Println()
	if err != nil {
		return fmt.Errorf("读取密码失败: %v", err)
	}
	if password != rePassword {
		return errors.New("两次输入的密码不一致")
	}
	if len(password) < 8 || len(password) > 26 {
		return errors.New("密码长度必须在8-26个字符之间")
	}

	user.UUID = uuid.Must(uuid.NewV6())
	user.Username = global.Config.Website.Name
	user.Password = utils.BcryptHash(password)
	user.RoleID = appTypes.Admin
	user.Avatar = "/images/avatar.png"
	user.Address = global.Config.Website.Address

	if err := global.DB.Create(&user).Error; err != nil {
		return fmt.Errorf("创建管理用户失败: %v", err)
	}
	fmt.Println("创建管理用户成功")
	return nil
}

func readPassword() (string, error) {
	var password string
	var buf [1]byte
	for {
		_, err := os.Stdin.Read(buf[:])
		if err != nil {
			return "", err
		}
		char := buf[0]
		if char == '\n' || char == '\r' {
			break
		}
		password += string(char)
	}
	return password, nil
}
