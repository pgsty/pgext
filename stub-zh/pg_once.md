## 用法

来源：

- [database.dev 软件包页面](https://database.dev/stateless/pg_once)

`pg_once` 0.1.0 是通过 database.dev 发布的纯 SQL 幂等键辅助扩展。它为指定资源预留键，记录完成或失败状态，可选缓存 JSON 响应，并按生存时间使记录过期；默认生存时间为 24 小时。

### 安装并预留键

先通过已配置的 dbdev 仓库安装软件包，再创建扩展：

```sql
SELECT dbdev.install('pg_once');
CREATE EXTENSION pg_once VERSION '0.1.0';

SELECT pg_once_start('order#42', 'orders');
SELECT pg_once_commit('order#42', '{"ok": true}'::jsonb);
```

应用必须检查 `pg_once_start` 返回的 JSON，只有键被接受时才执行受保护的工作。只在工作成功后调用 `pg_once_commit`；缓存的响应数据必须适合返回给每一个重试调用者。

操作失败或定期清理过期记录时：

```sql
SELECT pg_once_fail('order#42');
SELECT pg_once_cleanup();
```

### 事务与生命周期边界

`pg_once_start`、业务操作和 `pg_once_commit` 共同构成应用协议。应有意设计它们的事务边界：若预留先于业务写入提交，崩溃后可能留下处理中键；若响应单独提交，它也可能与受保护写入发生分歧。应为每个中间状态设计重试，并让 TTL 长于最长的合法操作时间。

记录过期并不能证明早先的外部副作用从未发生。`pg_once_cleanup` 会删除过期状态，因此只有明确保留与重试要求后才应调度它。该 registry 页面是此软件包公开的权威产物；向不受信任角色开放其函数或表之前，应检查生成迁移中实际安装的 SQL。
