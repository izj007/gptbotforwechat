
# 程序介绍
本程序是使用wechatbot二开，并对其进行了修改，以实现更好的功能。目前对原程序进行了如下修改：

修复了加好友异常程序退出的问题。
更换了当前的引擎到gpt-3.5-turbo，以提供更好的性能和体验。

### 目前实现了以下功能
 + 群聊@回复
 + 私聊回复
 + 自动通过回复

# 使用方法
````
# 获取项目
git clone https://github.com/869413421/wechatbot.git

# 进入项目目录
cd wechatbot

# 复制配置文件
copy config.dev.json config.json

# 启动项目
go run main.go

启动前需替换config中的api_key
```

技术细节
本程序使用了wechatbot二开作为基础框架，通过修改源码实现了加好友异常退出的修复。同时，为了提供更好的性能和体验，我们将当前的引擎更换为gpt-3.5-turbo，该引擎在自然语言处理领域有着很好的表现。

# 参考链接
https://github.com/djun/wechatbot