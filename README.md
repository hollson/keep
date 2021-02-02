# KEEP
- 一个精简版的TCP聊天服务

## 测试
- 测试keepalive：
```shell
# 查看网卡
ifconfig

# 卸载网卡，用于模拟网络中断
ifconfig lo0 down

# 启动网卡
ifconfig lo0 up
```