

# 功能强大的 校园 失物招领和寻物启事 后端项目

### 项目目的介绍

为了解决校园物品遗失以及失物找回困难的问题，开发的一个后端项目，通过发帖子展示捡到的东西和发帖表示自己丢失了啥，方便学生之间互相联系而设计。同时独创了推送系统，在失物招领时有相关人员的实名信息，将直接进行推送

### 技术说明

##### 选择

关于后端的编程语言选择了性能更好的，开发方便的go语言，其天生支持并发的特性，使得无论是高并发访问还是持续量巨大都可以轻松应对。同时采用了gin框架和gorm框架，gin是一个简单的http框架，相较于beego或者java的大部分框架都能做到更好的性能，且它更接近底层，寻找问题和解决问题也更方便。gorm是一种orm即对象关系映射，来应对持久化需求，做到将数据使用者和数据实体相关联。gorm具有高度拓展性和极佳的性能，可以说是go语言orm的不二之选，可以做到完全的底层自定义。关于持久化对象的选择，我使用了mysql，可以说mysql就是这个最佳答案。

##### 项目架构

│  go.mod
│  go.sum
│  main
│  README.md
│
├─api
│      post.go
│      user.go
│
├─assets
│  ├─lost_and_found
│  ├─notice_for_lost
│  └─user
│          default.png
│
├─cmd
│      main.go
│
├─internal
│      dealfuncation.go
│      middleware.go
│      post.go
│      user.go
│      u_test.go
│
├─model
│      add_test.go
│      comment.go
│      file.go
│      init.go
│      lost_and_found.go
│      notice_for_lost.go
│      user.go
│
├─service
│      file.go
│      post_test.go
│      post_thing.go
│      user.go
│
└─util
        jwt.go
        ptr.go

mvc架构的一种变种，util相当于工具包，是自己写的一些小玩意，model属于持久化，service是对model的封装，使得架构更加清晰和降低耦合度，internal是我的处理逻辑，api就是我所开放的接口，assets是图像区域，也就是我自己做的图床，用的加载静态资源文件做的服务器文件。cmd是程序入口，main是编译出来的linux二进制可执行文件

##### 技术细节

对于同样都是帖子的失物招领和寻物启事来说，我直接用了post接口进行抽象，使得拓展性和耦合度大幅降低。由于go是强类型编程语言，对于使用麻烦的原生各种转换例如：int.int64,uint和string之间的相互转换我写了一个自己的工具包简化操作。我选择了jwt作为token验证，工具包中的jwt.go文件也有我的加密算法，防止用户伪造数据，同时也能过滤掉大部分的网络攻击。我将assets单独作为静态资源服务器，实现自己的图床。我也实现了弱网环境下的图片上传和下载，例如断点续传。同时为了做到实时的推送，我做了一个监听器，监听所有的发帖并做到主动推送给用户是否有你的失物被找到。

### 接口文档

[postman链接](https://documenter.getpostman.com/view/12213031/TVmJheKX#eeb01fb2-a1c6-4ca7-bb27-afb2ee7b2cb8) ，我使用的postman测试，点开就可以查看接口数据

