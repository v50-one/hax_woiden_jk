# hax_woiden_jk
hax woiden 的库存监控(个人版)

<details>
  
<summary><code><strong>「 点击查看 邮件SMTP配置教程 」</strong></code></summary>
  
  
****
  [![ppoIN24.png](https://s1.ax1x.com/2023/04/06/ppoIN24.png)](https://imgse.com/i/ppoIN24)
  [![ppoIdM9.png](https://s1.ax1x.com/2023/04/06/ppoIdM9.png)](https://imgse.com/i/ppoIdM9)
  [![ppoIsIK.png](https://s1.ax1x.com/2023/04/06/ppoIsIK.png)](https://imgse.com/i/ppoIsIK)
  [![ppo7qzD.png](https://s1.ax1x.com/2023/04/06/ppo7qzD.png)](https://imgse.com/i/ppo7qzD)
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
#yaml配置文件格式提醒
#1.冒号后面有空格
#2.缩进不能用tab，只能用空格
#3.-后面有空格
#4.字符串不用加引号
# yaml格式校验 https://www.bejson.com/validators/yaml_editor/
#smtp邮件发送 email: 后面不填任何东西
email:
#smtp服务器
  smtp: smtp.163.com
#smtp服务器端口
  port: 25
#发信邮箱
  from: 1@163.com
#发信邮箱密码 非邮箱密码而是创建的smtp的密码
  key: 1
#收信邮箱 可以多个
  to:
   - 1@qq.com
   - 1@gmail.com
# pushplus推送 官网 https://www.pushplus.plus/
push_plus:
#push_plus 你的token
  token: eed2d4c5b3f240b38a2d4e1f2c2fee7e
#推送通道
#不需要的通道可以注释掉 注释使用 # 开头
# 全部通道被注释之后关闭pushplus推送
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
