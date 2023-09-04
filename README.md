# furina

Golang版轻量级游戏角色圣遗物查询计算评分Web程序, 可使用本地文件存储, 不依赖其他软件.

## 安装使用

### 配置项

配置文件为`config/config.yaml`

```yaml
Server:
  HttpPort: # web端口, 缺省为80

Databse:
  Username: # 数据库用户名
  Password: # 数据库密码
  Host: # 数据库主机, 缺省使用本地文件存储.
  DBName: # 表名
```

### 一键启动

下载编译好的可执行文件, 如`windows-amd64`, 放在项目根目录下运行即可.

### 编译

安装配置好golang环境后, 按需填写配置文件, 使用`go build`编译即可.

## 免责声明

本项目的图片与其他素材均来自于网络，仅供交流学习使用，如有侵权请联系删除.

本项目初衷为本地使用的便捷工具, 如您将其部署于公网之上, 请确保您有足够的能力处理网络安全问题, 一切后果均应由您自行承担.

## 致谢

* [Enka.Network](https://enka.network/) 面板数据源
* [miao-plugin](https://github.com/yoimiya-kokomi/miao-plugin) UI