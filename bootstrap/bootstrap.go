package bootstrap

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/869413421/wechatbot/handlers"
	"github.com/eatmoreapple/openwechat"
)

func Run() {
	for {
		runOnce()
		time.Sleep(1 * time.Second)
		// 删除 storage.json 文件
		err := os.Remove("storage.json")
		if err != nil {
			log.Printf("failed to delete storage.json: %v", err)
		}
	}
}

func runOnce() {
	//bot := openwechat.DefaultBot()
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式，上面登录不上的可以尝试切换这种模式

	// 发送url到钉钉机器人
	// sendURLToDingTalkBot("https://www.example.com", "钉钉机器人的webhook地址")

	// 注册消息处理函数
	bot.MessageHandler = handlers.Handler
	// 注册登陆二维码回调
	// bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	bot.UUIDCallback = SendQrcodeUrlToDingding

	// 创建热存储容器对象
	reloadStorage := openwechat.NewJsonFileHotReloadStorage("storage.json")
	// 执行热登录
	err := bot.HotLogin(reloadStorage)
	if err != nil {
		if err = bot.Login(); err != nil {
			log.Printf("login error: %v \n", err)
			return
		}
	}
	// 使用通道和协程来捕获退出事件
	exit := make(chan error)
	go func() {
		exit <- bot.Block()
	}()

	select {
	case err := <-exit:
		log.Printf("bot exited: %v", err)
	}

}

// 发送url到钉钉机器人
func SendQrcodeUrlToDingding(uuid string) {
	// println("访问下面网址扫描二维码登录")
	qrcodeUrl := "https://login.weixin.qq.com/qrcode/" + uuid
	// println(qrcodeUrl)

	dingTalkRobotUrl := "https://oapi.dingtalk.com/robot/send?access_token="
	message := fmt.Sprintf("GPT通知\n请扫描以下二维码登录微信：\n%s", qrcodeUrl)
	payload := fmt.Sprintf(`{"msgtype": "text", "text": {"content": "%s"}}`, message)
	resp, err := http.Post(dingTalkRobotUrl, "application/json", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		log.Printf("failed to send message to DingTalk robot: %v", err)
		return
	}
	defer resp.Body.Close()

}
