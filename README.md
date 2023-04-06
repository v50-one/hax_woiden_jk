# hax_woiden_jk
hax woiden 的库存监控(个人版)

<details>
  
<summary><code><strong>「 点击查看 邮件SMTP配置教程 」</strong></code></summary>
  
  
****
  [![ppoIN24.png](https://s1.ax1x.com/2023/04/06/ppoIN24.png)](https://imgse.com/i/ppoIN24)
  [![ppoIdM9.png](https://s1.ax1x.com/2023/04/06/ppoIdM9.png)](https://imgse.com/i/ppoIdM9)
  [![ppoIsIK.png](https://s1.ax1x.com/2023/04/06/ppoIsIK.png)](https://imgse.com/i/ppoIsIK)
  [![ppoI5Zt.png](https://s1.ax1x.com/2023/04/06/ppoI5Zt.png)](https://imgse.com/i/ppoI5Zt)
  [![ppoIXss.png](https://s1.ax1x.com/2023/04/06/ppoIXss.png)](https://imgse.com/i/ppoIXss)
  [![ppoIjLn.png](https://s1.ax1x.com/2023/04/06/ppoIjLn.png)](https://imgse.com/i/ppoIjLn)
  [![ppooBlQ.png](https://s1.ax1x.com/2023/04/06/ppooBlQ.png)](https://imgse.com/i/ppooBlQ)
****
  
  
</details>


<details>
  
<summary><code><strong>「 点击查看 配置文件填写教程 」</strong></code></summary>
  
  
****
  

## 配置文件填写教程
  
  
``` yaml
#smtp邮件发送
email:
#smtp服务器
  smtp: smtp.163.com #填写smtp服务器域名
#smtp服务器端口
  port: 25 #默认25 基本上不用更改
#发信邮箱
  from: 1@163.com
#发信邮箱密码 非邮箱密码而是创建的smtp的密码一般叫授权码
  key: 
#收信邮箱 可以多个 
  to:
   - 1@qq.com
   - 1@gmail.com
# pushplus推送 官网 https://www.pushplus.plus/
push_plus:
#push_plus 你的token
  token: 1
#推送通道 需要用哪个通道就把哪个通道的注释取消了
#通道名字不能更改
  channel:
#mail 邮箱
    - mail
#微信
    - wechat
#企业微信
    - webhook
#钉钉
    - cp
#短信
    - sms
  
```
  
  
****
  
  
</details>
