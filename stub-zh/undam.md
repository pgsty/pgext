## 用法

来源：

- [已复核 commit 的 undam README](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/README.md)
- [已复核 commit 的 undam.control](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/undam.control)
- [版本 0.1 的安装 SQL](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/undam--0.1.sql)
- [表访问方法实现](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/undam.c)
- [上游预加载配置](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/undam.conf)
- [上游回归测试 SQL](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/sql/undam.sql)

`undam` 是一个实验性表访问方法，使用撤销链和通用 WAL 实现原地更新。扩展会创建 `undam` 访问方法和诊断函数 `undam_rel_info(regclass)`。其 README 仍把这一设计称为预览，并说明 vacuum 依然必不可少。

### 预加载与试用

```conf
shared_preload_libraries = 'undam'
```

创建扩展前需要重启 PostgreSQL：

```sql
CREATE EXTENSION undam;
CREATE TABLE event_state (
  id bigint PRIMARY KEY,
  value bigint NOT NULL
) USING undam;
INSERT INTO event_state VALUES (1, 0);
UPDATE event_state SET value = value + 1;
VACUUM event_state;
SELECT * FROM undam_rel_info('event_state'::regclass);
```

### 注意事项

- `_PG_init` 会在 `undam` 未列入 `shared_preload_libraries` 时拒绝加载；修改此设置需要重启服务器。
- 这是针对旧版 PostgreSQL 内部表访问方法 API 构建的预览存储引擎代码。源码中有明确的 `NOT_IMPLEMENTED` 路径，上游也没有提供当前兼容性矩阵。
- 版本 `0.1` 只能用一次性数据在确切目标服务器构建上评估。在下结论前，应演练崩溃恢复、WAL 重放、索引、长事务、vacuum、备份和升级路径；不要把示例当成生产认可。
