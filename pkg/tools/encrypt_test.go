package tools

import "testing"

func TestPasswordEncryptAndCompare(t *testing.T) {
	// 测试用例1：正常密码加密和验证
	password1 := "password123"
	encrypted1 := PasswordEncrypt(password1)
	if !PasswordCompare(password1, encrypted1) {
		t.Errorf("Test case 1 failed: Password encryption and comparison failed")
	}

	// 测试用例2：空密码加密和验证
	password2 := ""
	encrypted2 := PasswordEncrypt(password2)
	if !PasswordCompare(password2, encrypted2) {
		t.Errorf("Test case 2 failed: Password encryption and comparison failed")
	}

	// 测试用例3：特殊字符密码加密和验证
	password3 := "!@#$%^&*()"
	encrypted3 := PasswordEncrypt(password3)
	if !PasswordCompare(password3, encrypted3) {
		t.Errorf("Test case 3 failed: Password encryption and comparison failed")
	}

	// 测试用例4：密码长度小于8位加密和验证
	password4 := "123"
	encrypted4 := PasswordEncrypt(password4)
	if !PasswordCompare(password4, encrypted4) {
		t.Errorf("Test case 4 failed: Password encryption and comparison failed")
	}

	// 测试用例5：密码长度大于20位加密和验证
	password5 := "12345678901234567890"
	encrypted5 := PasswordEncrypt(password5)
	if !PasswordCompare(password5, encrypted5) {
		t.Errorf("Test case 5 failed: Password encryption and comparison failed")
	}
}
