# share_website

标准的前后端分离项目，前端vue3脚手架搭建，后端使用go的gin搭建web服务器，存储是mysql，用到了腾讯云的oss
网站的功能：图片视频上传，单词薄，留言薄，登陆注册

### 前端部署方式
你需要安装node.js 和 npm
- 拉下代码之后在前端代码的根目录下执行  `npm install` 下载依赖
- 打包代码 `npm run build`生成dist文件夹
- 使用 静态资源服务器代理 dist文件夹，建议是使用nginx,类似这样
- 与后端通信是通过axios [代码](https://github.com/qq1141000259/share_website/blob/master/frontend/share-website/src/utils/request.js)
  默认后端地址是`http://localhost:8099`,你可以改成你的后端所在服务器再编译代码
```
  location / {
    alias /home/web/frontend/dist/;
  }
```
### 后端部署
需要把config目录下的config.yaml里的配置填上
之后直接 go build 编译就好了
