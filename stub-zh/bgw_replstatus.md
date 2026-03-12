


## 用法

> [bgw_replstatus: 用于负载均衡器的轻量级复制状态检查器](https://github.com/mhagander/bgw_replstatus)

bgw_replstatus 是一个后台工作进程，它监听 TCP 端口并立即报告节点是 `MASTER` 还是 `STANDBY`。专为负载均衡器健康检查设计，无需建立完整的 PostgreSQL 连接。

### 工作原理

任何到配置端口的 TCP 连接都会立即收到文本响应（`MASTER` 或 `STANDBY`），然后连接关闭。无需发送请求。

### 快速测试

```bash
nc localhost 5400
# 输出：MASTER（或 STANDBY）
```

### HAProxy 配置示例

```
frontend test
    bind 127.0.0.1:5999
    default_backend pgcluster

backend pgcluster
    mode tcp
    option tcp-check
    tcp-check expect string MASTER
    server s1 127.0.0.1:5500 check port 5400
    server s2 127.0.0.1:5501 check port 5401 backup
    server s3 127.0.0.1:5502 check port 5402 backup
```

### 配置

| 参数 | 默认值 | 描述 |
|-----------|---------|-------------|
| `bgw_replstatus.port` | 5400 | 监听的 TCP 端口 |
| `bgw_replstatus.bind` | (通配符) | 绑定的 IP 地址 |

**安全提示**：没有源地址验证。请使用主机防火墙保护该端口。
