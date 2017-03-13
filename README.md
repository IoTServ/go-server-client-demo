#chuantou 内网穿透
###内网穿透，可以让全世界访问家用电脑里的网站。0.1测试项目。为下一个多线程版本做准备
##软件作用：让你的家用电脑中建立的网站全世界都能访问
##原理（client为家用电脑，内装有自己的网站）
![输入图片说明](http://git.oschina.net/uploads/images/2017/0313/235933_fd3a3ee6_891703.png "在这里输入图片标题")
##使用方法：
###1、配置好go语言环境，
###2、把server.go上传到公网服务器上。运行go run server.go
###3、把client.go放在家用电脑上（无公网ip，家用电脑80端口可以访问到本地的网站）。运行go run client.go -host 第二步中的服务器ip(比如服务器ip为1.1.1.1 则命令为go run client.go -host 1.1.1.1 )
###4、访问公网ip:3000即可访问到家用电脑中的网站。（如服务器ip为1.1.1.1则访问1.1.1.1:3000）
##测试截图
###我是在本地测试，所以服务器ip也为127.0.0.1
![输入图片说明](http://git.oschina.net/uploads/images/2017/0314/000649_d5039279_891703.png "在这里输入图片标题")
###我的家用电脑网站地址127.0.0.1:80。网站中有phpinfo.php这个文件
![输入图片说明](http://git.oschina.net/uploads/images/2017/0314/000659_42fc2a04_891703.png "在这里输入图片标题")
#也就是说全世界任何地方访问 服务器ip:3000 就会访问到家用电脑的127.0.0.1:80
##注意事项：
1.这只是实验，为下一个版本做测试。可能有一些bug
2.此版本只支持单线程。不支持并发。
3.做微信开发时可以使用，不用开发一会传到服务器测试
