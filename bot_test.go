package openwechat

import (
	"fmt"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
	bot := DefaultBot(Desktop)
	bot.LoginCallBack = func(body CheckLoginResponse) {
		t.Log("login")
	}
	if err := bot.Login(); err != nil {
		t.Error(err)
	}
}

func TestLogout(t *testing.T) {
	bot := DefaultBot(Desktop)
	bot.LoginCallBack = func(body CheckLoginResponse) {
		t.Log("login")
	}
	bot.LogoutCallBack = func(bot *Bot) {
		t.Log("logout")
	}
	bot.MessageHandler = func(msg interface{}) {
		if message, ok := msg.(*Message); ok { // 使用指针类型断言
			if message.IsText() && message.Content == "logout" {
				if err := bot.Logout(); err != nil {
					t.Error(err)
				}
			}
		}
	}
	if err := bot.Login(); err != nil {
		t.Error(err)
		return
	}
	_ = bot.Block()
}

func TestMessageHandle(t *testing.T) {
	bot := DefaultBot(Desktop)
	bot.MessageHandler = func(msg interface{}) {
		if message, ok := msg.(*Message); ok { // 使用指针类型断言
			if message.IsText() && message.Content == "ping" {
				if _, err := message.ReplyText("pong"); err != nil {
					t.Error(err)
				}
			}
		}
	}
	if err := bot.Login(); err != nil {
		t.Error(err)
		return
	}
	_ = bot.Block()
}

func TestFriends(t *testing.T) {
	bot := DefaultBot(Desktop)
	if err := bot.Login(); err != nil {
		t.Error(err)
		return
	}
	user, err := bot.GetCurrentUser()
	if err != nil {
		t.Error(err)
		return
	}
	friends, err := user.Friends()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(friends)
}

func TestGroups(t *testing.T) {
	bot := DefaultBot(Desktop)
	if err := bot.Login(); err != nil {
		t.Error(err)
		return
	}
	user, err := bot.GetCurrentUser()
	if err != nil {
		t.Error(err)
		return
	}
	groups, err := user.Groups()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(groups)
}

func TestPinUser(t *testing.T) {
	bot := DefaultBot(Desktop)
	if err := bot.Login(); err != nil {
		t.Error(err)
		return
	}
	user, err := bot.GetCurrentUser()
	if err != nil {
		t.Error(err)
		return
	}
	friends, err := user.Friends()
	if err != nil {
		t.Error(err)
		return
	}
	if friends.Count() > 0 {
		f := friends.First()
		if err = f.Pin(); err != nil {
			t.Error(err)
			return
		}
		time.Sleep(time.Second * 5)
		if err = f.UnPin(); err != nil {
			t.Error(err)
			return
		}
	}
}

func TestSender(t *testing.T) {
	bot := DefaultBot(Desktop)
	bot.MessageHandler = func(msg interface{}) {
		if message, ok := msg.(*Message); !ok {
			if message.IsSendByGroup() {
				fmt.Println(message.SenderInGroup())
			} else {
				fmt.Println(message.Sender())
			}
		}

	}
	if err := bot.Login(); err != nil {
		t.Error(err)
		return
	}
	_ = bot.Block()
}

// TestGetUUID
// @description: 获取登录二维码(UUID)
// @param t
func TestGetUUID(t *testing.T) {
	bot := DefaultBot(Desktop)

	uuid, err := bot.Caller.GetLoginUUID(bot.Context())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(uuid)
}

// TestLoginWithUUID
// @description: 使用UUID登录
// @param t
func TestLoginWithUUID(t *testing.T) {
	uuid := "oZZsO0Qv8Q=="
	bot := DefaultBot(Desktop, WithUUIDOption(uuid))
	err := bot.Login()
	if err != nil {
		t.Errorf("登录失败: %v", err.Error())
		return
	}
}
