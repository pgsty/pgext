## 用法

来源：

- [上游 README](https://github.com/AbdulYadi/pgsocket/blob/c52b088ae1bf9cb356e55b5c63715147c4f2e18b/README.md)
- [扩展 SQL](https://github.com/AbdulYadi/pgsocket/blob/c52b088ae1bf9cb356e55b5c63715147c4f2e18b/pgsocket--1.0.sql)
- [扩展 control 文件](https://github.com/AbdulYadi/pgsocket/blob/c52b088ae1bf9cb356e55b5c63715147c4f2e18b/pgsocket.control)

`pgsocket` 在 PostgreSQL 后端中提供 TCP 客户端。`pgsocketsend` 用于发送字节，`pgsocketsendrcvstxetx` 则发送字节并返回以 STX 和 ETX 分帧的响应：

```sql
CREATE EXTENSION pgsocket;

SELECT pgsocketsend(
  '127.0.0.1', 9090, 30,
  decode('48656c6c6f', 'hex')
);

SELECT pgsocketsendrcvstxetx(
  '127.0.0.1', 9090, 30, 40,
  decode('48656c6c6f', 'hex')
);
```

### 兼容性与安全

版本 `1.0` 还定义了专用请求/响应函数 `pgsocketgetimage`。README 表示项目在 Linux 上针对 PostgreSQL 10 编译，且没有提供更新版本的兼容矩阵，因此需要进行移植测试。这些函数允许数据库会话发起出站网络连接；应限制 `EXECUTE` 权限、审查网络出口策略，并避免使用不可信的目标或载荷输入。
