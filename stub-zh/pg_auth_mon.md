


## 用法

> [pg_auth_mon: 监控认证尝试](https://github.com/RafiaSabih/pg_auth_mon)

`pg_auth_mon` 按用户存储认证尝试统计信息，跟踪成功和失败的登录及其时间戳。

```sql
CREATE EXTENSION pg_auth_mon;
```

### 配置

在 `postgresql.conf` 中添加：

```ini
shared_preload_libraries = 'pg_auth_mon'
```

可选 GUC 参数：

| 参数 | 默认值 | 描述 |
|------|--------|------|
| `pg_auth_mon.log_period` | `0` | 每 N 分钟将统计信息转储到 Postgres 日志（0=关闭） |

### 查询认证统计

```sql
SELECT * FROM pg_auth_mon;
```

`pg_auth_mon` 视图提供每用户信息：

| 列 | 描述 |
|----|------|
| `uid` | 用户 OID（0 表示无效/不存在的用户名） |
| `successful_auth_count` | 成功登录总次数 |
| `last_successful_auth` | 最近一次成功登录的时间戳 |
| `failed_auth_count` | 失败认证总次数 |
| `last_failed_auth` | 最近一次失败登录的时间戳 |
| `hba_conflicts_count` | HBA 文件冲突次数 |

所有使用无效用户名的登录尝试被聚合到一个 OID 为 0 的行中。
