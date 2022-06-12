# 抖音项目文档
## 抖音项目接口文档
[抖音极简版API文档](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145)
## 技术栈
- web框架：Gin 
- jwt认证授权：gin-jwt
- 数据库：MySQL（ORM框架：Gorm)
- 云存储：OSS存储

## 成员分工
- 登录注册+用户信息、jwt鉴权、全局异常捕捉：茅津菁、谭嘉睿
- 视频流、视频投稿、发布列表：肖文卓、章梓豪
- 点赞、点赞列表、评论、评论列表：袁杰、汪列伟
- 关注、粉丝：李登龙、徐晓明 
  
## 项目结构
  本项目采取Controller + Service + DAO 三层架构模式：
  ![structure](https://tva1.sinaimg.cn/large/e6c9d24ely1h35qxwshlej207e0xjq44.jpg)
  
- config：项目配置（基于viper读取yaml文件）
- constants：项目中使用到的字符串常量
- controller：处理接口逻辑
- dal：数据库DAO层使用gorm进行增删改查操作
- errno：全局封装异常处理类，便于统一处理
- handlers：接口中传参DTO对象
- middlewares：用户认证中间件和全局异常处理中间件
- router：HTTP路由配置
- service：根据DAO层提供的操作封装成对应的服务接口
- utils：项目工具类，包括jwt工具类、云存储oss工具类、Bcrypt加密工具类

## 数据库表设计
表关系设计

![relation](https://tva1.sinaimg.cn/large/e6c9d24ely1h35qx3y5kbj20g50ctwes.jpg)

表字段设计

用户表：user

![user](https://tva1.sinaimg.cn/large/e6c9d24ely1h35qyiup5bj20g3075jst.jpg)

关注表：follow

![follow](https://tva1.sinaimg.cn/large/e6c9d24ely1h35r4o55gtj20fc05o756.jpg)

视频表：video

![video](https://tva1.sinaimg.cn/large/e6c9d24ely1h35qze5hboj20gb072gn1.jpg)

点赞表：favor

![favor](https://tva1.sinaimg.cn/large/e6c9d24ely1h35qyyxb7pj20ev05rjsc.jpg)

评论表：comment

![comment](https://tva1.sinaimg.cn/large/e6c9d24ely1h35r0zrkktj20f506hjsi.jpg)

## 功能实现设计思路
### jwt鉴权
- 说明：使用gin-jwt来配置midlleware插件，自定义其中的身份认证，认证失败相应、登录成功响应内容，根据接口文档在formdata中进行token认证要求去改写增强原有gin-jwt源码，支持formdata方式认证。
- 使用位置：登录接口使用该插件的LoginHandler；其他路由若是需要进行token认证，则可以再对应路由添加插件的MiddlewareFunc来达到认证的效果。

### 异常捕捉
- 说明：统一来接收service层抛出的异常来通过recover()函数来进行捕捉。
- 使用位置：router初始化后安装该异常插件即可。


## 基础接口
### 视频流：/douyin/feed GET
- 业务思路：①取出token中的userID。 ②获取lastTime的值。 ③查询视频流列表，返回按投稿时间倒序的视频列表。④遍历循环，为视频流列表的每条视频填充点赞等内容。 ⑤封装响应格式。

### 用户注册：/douyin/user/register POST
- 业务思路： ①取出请求中的username和password参数  ②检查用户名和密码是否过长或为空 ③检查用户名是否已经注册  ④若用户名未注册，创建新用户  ⑤用全局中间件gin-jwt获取token，封装响应格式。

### 用户登录：/douyin/user/login/ POST
- 实现方式：使用的是gin-jwt中提供的loginHandler，对应我只需要实现gin-jwt中的Authenticator身份认证以及PayloadFunc对应要携带的载荷函数以及成功响应、失败响应自定义函数即可。
- 业务思路：①检测用户是否存在。②密码是否正确。③以上都通过后来设置token中需要携带的载荷信息来使用jwt生成token返回。

### 获取用户信息：/douyin/user/ GET
- 业务思路：①根据指定用户id来获取目标用户信息。②取出token中的userID以及目标用户id去查询关注表是否关注。③将目标用户信息以及是否关注组合成对应的response对象进行返回。【若是中间出现异常直接抛出由异常捕捉插件统一进行异常响应】

### 投稿接口：/douyin/publish/action POST
- 业务思路：①使用阿里云OSS进行视频存储 编写相应工具类封装返回Bucket以及上传文件接口。②解析token获取userID 视频标题 以及视频文件。 ③调用工具类上传OSS，返回视频url以及通过视频截帧返回视频封面url，存入数据库。 ④封装响应格式。

### 发布列表： /douyin/publish/list/ GET
- 业务思路：①获取userID。②查询该作者的投稿视频。 ③遍历循环，为视频流列表的每条视频填充内容。 ④封装响应格式。

## 扩展接口-I
### 点赞接口：/douyin/favorite/action/ POST
- 业务思路：①根据token解析userID②校验传参videoID、actionType是否合法。③判断是点赞还是取消点赞操作④若是点赞操作，首先检查是否已点赞过，若已点赞过直接返回，否则进行点赞操作。若是取消点赞操作则将点赞记录设置为逻辑删除。

### 点赞列表：/douyin/favorite/list/ GET
- 业务思路：①校验传参userID是否合法②根据点赞记录查询对应视频，封装成数组对象③封装响应格式。

### 评论操作：/douyin/comment/action/ POST
- 业务思路：①根据token解析userID②校验传参videoID、actionType是否合法。③判断是评论还是删除评论操作④进行评论和删除评论操作。

### 评论列表：/douyin/comment/list/ GET
- 业务思路：①校验传参videoID是否合法②根据点赞记录查询对应视频，封装成数组对象③封装响应格式。


## 扩展接口-II
### 关注操作：/douyin/relation/action/  POST
- 业务思路：通过在中间表中增加记录实现，每一条记录对应两个用户ID（分别是被关注者ID和粉丝ID）
### 关注列表：/douyin/relation/follow/list/ GET
- 业务思路：通过用户ID寻找对应粉丝ID的所有关注者ID，再通过关注者ID循环查询用户ID信息
- 优化逻辑：每个ID并发查询用户信息
### 粉丝列表：/douyin/relation/follower/list/ GET
- 业务思路：通过用户ID寻找对应关注者ID下的所有粉丝ID，再通过粉丝ID循环查询用户ID信息
