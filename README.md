# wsl-Ip

wsl2每次启动都会分配一个不同的ip地址，这给我们从windows主机访问wsl2主机造成了一定的不便，使用本工具，可以通过为Windows设置hosts的方式来方便的使用wsl域名访问wsl2主机。

原理：
1. 使用ifconfig 命令获取wsl本机ip地址
2. 将获得的ip地址写入Windows主机的hosts中，分配域名wsl

使用方式：
```bash
curl -o wsl https://github.com/yangchnet/wsl-Ip/releases/download/v0.0.1/wsl
chmod +x wsl
sudo mv wsl /usr/local/bin/wsl
```
此时可以使用`wsl`命令来为Windows设置hosts.

当然也可以将其配置为system服务并设置开机启动：
```bash
sudo vim /etc/systemd/system/wsl.service
```
写入以下内容
```
[Unit]
Description= wsl-Ip
After=network.target

[Service]
Type=simple
ExecStart=/usr/local/bin/wsl

[Install]
WantedBy=multi-user.target
```

```bash
systemctl daemon-reload  # 重载systectm
systemctl enable wsl # 设置开机启动
```

有任何问题，欢迎反馈，有用请给个star.