# wepay——微信支付gosdk
根据微信支付官方文档开发的go程序，提供支持**普通商户**微信支付的功能

1. 使用go语言开发微信支付相关接口  [微信支付文档](https://pay.weixin.qq.com/wiki/doc/api/index.html)
2. 使用echo框架开发restful API  [echo简介](http://go-echo.org/)
3. 使用mongodb存储数据  [mgo驱动官网](http://labix.org/mgo)
4. 使用dep进行包管理 [dep包管理工具](https://github.com/golang/dep)
5. 项目结构

   ```
    wepay/  
    |---main.go    main函数 
    |---config.go  配置函数 初始化配置文件
    |---router.go  路由函数 提供api
    |---db.go      数据函数 提供数据库连接
    |---.gitignore git忽略文件 
    |---.env       env配置文件 提供config.go读取的配置文件
    |---Dockerfile dockerfile 构建docker镜像
    ```
 6. 使用到的第三方库

    echo : 一款go写的高效的web框架  https://github.com/labstack/echo

    godotenv : 一个go写的读取.env文件的库 https://github.com/joho/godotenv

    logrus : 一个go写的高效的日志库 https://github.com/sirupsen/logrus

    multiconfig : 一个go写的多重载入配置文件的库 https://github.com/koding/multiconfig
    
    mgo : 一个mongodb的go语言的驱动 https://labix.org/mgo
6. 由于使用的包有被GFW墙掉的，所以我没有ignore掉vendor文件夹
# Tips
打包成可执行程序并压缩     
1.使用go build -ldflags '-w -s'进行代码编译，得到.exe文件    
2.使用upx小工具进行压缩，使得.exe文件大幅度缩小 [upx官网](https://upx.github.io/ "点击upx下载")