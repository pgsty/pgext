## 用法

来源：

- [官方 README](https://github.com/theory/pg-hostname/blob/37d7fea1d393b3bb2093e419b181c7d0826b3c6e/README.md)
- [SQL 声明](https://github.com/theory/pg-hostname/blob/37d7fea1d393b3bb2093e419b181c7d0826b3c6e/sql/hostname.sql)
- [C 实现](https://github.com/theory/pg-hostname/blob/37d7fea1d393b3bb2093e419b181c7d0826b3c6e/src/hostname.c)
- [PGXN 发行版](https://pgxn.org/dist/hostname/)

`hostname` 只提供一个函数 `hostname() RETURNS text`，返回 PostgreSQL 服务器进程可见的操作系统主机名。它标识数据库服务器或容器，而不是客户端主机，也不一定是 DNS 完全限定名或集群唯一名称。

### 核心流程

```sql
CREATE EXTENSION hostname;

SELECT hostname();
```

C 函数调用 POSIX `gethostname` API；调用失败时返回 NULL。扩展可重定位，并要求提供 `<unistd.h>` 的 POSIX 平台。

### 解释与版本

在容器、Kubernetes Pod、故障转移集群和克隆主机中，返回值可能随重启或角色切换变化。除非与基础设施元数据核对，否则不要把它作为持久节点标识、安全边界或路由来源。函数声明为 `IMMUTABLE`，因此即使管理员可以修改操作系统主机名，PostgreSQL 也可能在已规划表达式中折叠重复调用。

固定源码修订中的 SQL 扩展版本仍为 1.0.0，而上游软件包/PGXN 发行版为 1.0.4。这是软件包发行差异，并不是新的 SQL 扩展版本。当前上游文档支持 PostgreSQL 9.0 或更高版本，并记录了 PostgreSQL 18 自定义安装前缀所需的扩展路径设置。
