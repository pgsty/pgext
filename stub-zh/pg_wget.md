## 用法

来源：

- [README.md](https://github.com/mkgrgis/pg_wget/blob/a3b9ff378de69583e0adc225f63b5e3bad570c8c/README.md)
- [pg_wget.control](https://github.com/mkgrgis/pg_wget/blob/a3b9ff378de69583e0adc225f63b5e3bad570c8c/pg_wget.control)
- [pg_wget--1.0.sql](https://github.com/mkgrgis/pg_wget/blob/a3b9ff378de69583e0adc225f63b5e3bad570c8c/pg_wget--1.0.sql)
- [pg_wget.c](https://github.com/mkgrgis/pg_wget/blob/a3b9ff378de69583e0adc225f63b5e3bad570c8c/pg_wget.c)
- [License](https://github.com/mkgrgis/pg_wget/blob/a3b9ff378de69583e0adc225f63b5e3bad570c8c/License)

`pg_wget` 启动服务器上的 wget 程序，将 HTTP 或 HTTPS 响应作为 PostgreSQL 文本返回。扩展的 SQL 版本为 1.0，适合在 Linux 或 POSIX 系统上下载少量文本；本次核对的源码修订未提供已验证的 PostgreSQL 主版本支持矩阵。

### 读取文本

服务器需要在进程的程序搜索路径中找到外部 wget 可执行文件。由超级用户安装扩展；它可迁移模式，不需要共享预加载或重启。

```sql
CREATE EXTENSION pg_wget;
SELECT pg_wget('https://www.postgresql.org/');
```

唯一的函数 `pg_wget(text)` 返回文本，标记为 VOLATILE 和 STRICT，参数为 NULL 时返回 NULL。它只接受以 `http://` 或 `https://` 开头的 URL。wget 以非零状态退出时，数据库调用报错。实现直接执行 wget，将 URL 作为独立参数传入，不经过 shell。

### 权限与资源边界

按照 PostgreSQL 的默认函数权限，PUBLIC 最初可以执行该函数。在普通用户连接之前，应将执行权限限制为已有的可信角色：

```sql
REVOKE EXECUTE ON FUNCTION pg_wget(text) FROM PUBLIC;
GRANT EXECUTE ON FUNCTION pg_wget(text) TO downloader;
```

将示例角色替换为实际已有的数据库角色。请求从数据库服务器发出，可访问其操作系统账户能够到达的地址。扩展只检查 URL 协议，不提供目标地址白名单；网络访问和外部程序需要在扩展之外管控。

实现会将整个响应收集到后端进程内存中。它向 wget 传入十秒超时选项，但这并不保证整个 SQL 调用在十秒内结束，也不限制响应大小。应只读取小型文本响应，避免不可信目标，不要将其当作二进制或流式 HTTP 客户端。外部网络操作不会随事务回滚而撤销。
