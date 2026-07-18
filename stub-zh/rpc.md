## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/simon-engledew/postgresql-rpc/blob/9a21b7825e7c0879b1f6b20a92ec6e4a722af350/README.md)
- [1.0 版本 SQL 声明](https://github.com/simon-engledew/postgresql-rpc/blob/9a21b7825e7c0879b1f6b20a92ec6e4a722af350/extension/rpc--1.0.sql)
- [Unix socket C 客户端](https://github.com/simon-engledew/postgresql-rpc/blob/9a21b7825e7c0879b1f6b20a92ec6e4a722af350/extension/rpc.c)
- [扩展 control 文件](https://github.com/simon-engledew/postgresql-rpc/blob/9a21b7825e7c0879b1f6b20a92ec6e4a722af350/extension/rpc.control)

`rpc` 向 Unix 域 socket 发送 JSON 文档，并返回对端回复的 JSON 文档。每个帧都以大端 32 位长度作为前缀。

```sql
CREATE EXTENSION rpc;
SET rpc.path = '/var/run/postgresql/rpc.sock';
SET rpc.timeout = 5;
SELECT rpc('{"operation":"validate","value":42}'::json);
```

默认 socket 为 `/var/run/postgresql/rpc.sock`；`rpc.path` 与范围为零到六十秒的 `rpc.timeout` 均可由用户设置。对端必须实现完全相同的分帧协议，并返回合法 JSON。

1.0 是小型旧式原生客户端，没有认证、授权、发行历史、测试或兼容矩阵。任何函数调用者都能选择另一个可访问的 Unix socket。对端提供的响应长度会直接驱动未设上限的内存分配，因此恶意对端可耗尽后端内存。应撤销公开执行与设置权限、限制文件系统访问，并且只连接可信本地服务。
