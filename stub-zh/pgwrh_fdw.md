## 用法

来源：

- [Standalone FDW guide](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh_fdw/README.md)
- [Control file](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh_fdw/pgwrh_fdw.control)
- [Extension SQL](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh_fdw/pgwrh_fdw--1.0.0-alpha1.sql)
- [Licensing](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh_fdw/LICENSING.md)

`pgwrh_fdw` 1.0.0-alpha1 是基于 `postgres_fdw` 的 PostgreSQL 18 外部数据包装器，增加了指定事务设置的传递以及副本服务器之间的路由。它可独立于 pgwrh 核心使用，采用 AGPL-3.0-only 许可证并保留继承的 PostgreSQL 声明。此 alpha 版本仅支持全新安装。

### 访问远程表

以管理员身份启用扩展，无需共享预加载。创建服务器，并配置用户映射，使远程角色具备目标表的读取权限。以下示例假定已经配置映射，远程数据库存在 `public.items`，且本地模式中没有同名冲突表。

```sql
CREATE EXTENSION pgwrh_fdw;
CREATE SERVER replica_a FOREIGN DATA WRAPPER pgwrh_fdw
    OPTIONS (host 'replica-a', dbname 'replica');
IMPORT FOREIGN SCHEMA public LIMIT TO (items)
    FROM SERVER replica_a INTO public;
SELECT * FROM items;
```

### 传递事务设置

```sql
ALTER SERVER replica_a OPTIONS (ADD transaction_parameters 'app.request_id');
BEGIN;
SET LOCAL app.request_id = 'request-42';
SELECT * FROM items;
COMMIT;
```

`transaction_parameters` 是外部服务器选项，值为非空的逗号分隔列表，元素必须是含点号的自定义参数名。不允许核心设置、重复项，以及 `postgres_fdw.*` 和 `pgwrh_fdw.*` 命名空间。参数名会规范化为小写，最长 63 字节。

应在规划或访问外部表之前设置上下文。首个参与传递的远程事务会捕获当前可见的自定义设置；直到本地顶层事务结束，各参与服务器都使用这份上下文。后续本地修改不会继续传递。远程估算可能在规划阶段触发捕获。远程角色与参数权限仍然生效，参数值可能出现在日志中。

### 虚拟服务器与辅助函数

`members` 服务器选项列出虚拟服务器后面的普通服务器。虚拟服务器需要空用户映射，凭据来自所选成员。选择时优先复用连接，再依据 `load_balance_weight` 分配，并在事务期间保持固定。只有初次连接失败才可尝试其他成员；已建立的事务以及查询错误都不会触发故障转移。

`pgwrh_fdw_set_members` 在变更成员时等待旧路由配置的使用者。`pgwrh_fdw_get_connections`、`pgwrh_fdw_disconnect` 和 `pgwrh_fdw_disconnect_all` 用于管理缓存连接。`pgwrh_fdw_scram_verifier` 生成带新盐值的 SCRAM 校验值；其输入与结果均应作为认证材料保护。

### 一致性边界

参数传递本身不会等待复制、解释 LSN、提供全局快照或协调分布式提交。需要复制屏障时，必须在每个实际接收服务器上配置 `pgwrh_wait`，并显式启用参数传递。设置参数成功本身不能证明接收端实现了相应行为。现有外部服务器不会自动转换。
