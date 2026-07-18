## 用法

来源：

- [上游 README](https://github.com/tvondra/connection_limits/blob/ab816a0a386978e124c24b2695088aef661f0e84/README.md)
- [1.0.0 版安装 SQL](https://github.com/tvondra/connection_limits/blob/ab816a0a386978e124c24b2695088aef661f0e84/sql/connection_limits--1.0.0.sql)
- [连接钩子实现](https://github.com/tvondra/connection_limits/blob/ab816a0a386978e124c24b2695088aef661f0e84/src/connection_limits.c)

connection_limits 可按数据库名、角色名、客户端地址或这些字段的组合限制并发连接数。当 PostgreSQL 内置的角色与数据库连接限额不够灵活时，可使用它。

### 服务器配置

该库必须在服务器启动时预留共享内存。将其加入预加载列表，按需定义默认值，然后重启 PostgreSQL：

```ini
shared_preload_libraries = 'connection_limits'
connection_limits.per_user = 5
```

具体规则从 PGDATA/pg_limits.conf 读取。mask 列可省略，all 表示通配：

```text
database  username  IP              [mask]         limit
all       foouser   all                             10
foodb     all       all                             20
all       all       192.168.1.0/24                  10
```

安装 SQL 扩展后，可查看计数器或重载规则文件：

```sql
CREATE EXTENSION connection_limits;

SELECT * FROM connection_limits;
SELECT connection_limits_reload();
```

预加载设置与共享内存分配仍然需要重启。SQL 重载函数只会在库已激活后重新读取规则。

### 注意事项

- 超级用户连接会被计数，但不会被拒绝。因此它们可能耗尽限额，使普通角色无法连接。
- 规则使用登录时提供的数据库名和角色名作为键，而不是对象标识符。重命名后应同步更新规则文件。
- 文件规则优先于每数据库、每用户和每 IP 默认值。应谨慎审查相互重叠的规则。
- 如果只需简单的每角色或每数据库限额，应优先使用 PostgreSQL 内置的 CONNECTION LIMIT。该扩展使用服务器钩子与较旧的内部 API，上游没有提供当前大版本兼容性矩阵。
