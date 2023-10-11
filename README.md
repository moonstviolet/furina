# furina

某二字游戏角色圣遗物查询和评分计算Web程序, 使用Golang作为后端语言, 一键启动.

可以选择使用本地文件存储, 不依赖其他环境(比如Node.js和数据库), 方便部署.

芙门!

## 安装使用

### 配置项

配置文件为`config/config.yaml`

```yaml
Server:
  HttpPort: 80     # web端口, 缺省为80
Database:
  LocalStorageDir: # 本地存储目录, 缺省为项目根目录下local
  Username:        # 数据库用户名
  Password:        # 数据库密码
  Addr:            # mongo数据库地址, 形如"127.0.0.1:27017", 缺省使用本地文件存储
  DBName: furina   # 数据库名
```

### 一键启动

下载编译好的可执行文件, 如`windows-amd64`, 放在项目根目录下运行即可.

### 圣遗物属性权重

配置文件为`config/propertyWeight.json`, 可自行修改后重启程序应用

属性名称对应如下, 权重为0-100

```shell
"Hp":               生命
"Atk":              攻击
"Def":              防御
"Mastery":          精通
"CritRate":         暴击
"CritDmg":          暴伤
"Recharge":         充能
"PhysicalDmgBonus": 物伤加成
"PyroDmgBonus":     火伤加成
"ElectroDmgBonus":  雷伤加成
"HydroDmgBonus":    水伤加成
"DendroDmgBonus":   草伤加成
"AnemoDmgBonus":    风伤加成
"GeoDmgBonus":      岩伤加成
"CryoDmgBonus":     冰伤加成
"HealingBonus":     治疗加成
```

### Bugs(Features)

- 角色面板属性中, 没有将武器属性单独分列出来(被动计算比较麻烦)
- 伤害计算暂时空缺(武器特效部分比较麻烦)
- 部分角色面板有误(比如荧)

### Todo

- [ ] 手动输入圣遗物属性计算评分 

### 编译

安装配置好golang环境后, 按需填写配置文件, 使用`go build`编译即可.

## 免责声明

本项目的图片与其他素材均来自于网络，仅供交流学习使用，如有侵权请联系删除.

本项目初衷为本地使用的便捷工具, 如您将其部署于公网之上, 请确保您有足够的能力处理网络安全问题, 一切后果均应由您自行承担.

## 致谢

* [Enka.Network](https://enka.network/) 面板数据源
* [miao-plugin](https://github.com/yoimiya-kokomi/miao-plugin) UI和部分资源
