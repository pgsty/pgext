## 用法

来源：

- [已复核 commit 的 undam README](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/README.md)
- [已复核 commit 的 undam.control](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/undam.control)
- [版本 0.1 的安装 SQL](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/undam--0.1.sql)
- [表访问方法实现](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/undam.c)
- [上游预加载配置](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/undam.conf)
- [上游回归测试 SQL](https://github.com/postgrespro/undam/blob/f1ac5366bd3410e091942fe2cd9d12e622dd6f55/sql/undam.sql)

`undam` 是一个实验性表访问方法，在原地更新元组的同时，将旧版本保留在撤销链中。它在主 fork 与扩展 fork 中存储定长数据块，并使用 PostgreSQL 通用 WAL。上游将这项设计称为预览，vacuum 仍然不可缺少。

### 核心流程

模块必须在服务器启动阶段分配共享状态：

```conf
shared_preload_libraries = 'undam'
```

重启 PostgreSQL，然后创建扩展，并让单独的表选择该访问方法：

```sql
CREATE EXTENSION undam;

CREATE TABLE event_state (
  id bigint PRIMARY KEY,
  value bigint NOT NULL
) USING undam;

INSERT INTO event_state VALUES (1, 0);
UPDATE event_state SET value = value + 1 WHERE id = 1;
VACUUM event_state;

SELECT * FROM undam_rel_info('event_state'::regclass);
```

安装扩展不会转换现有 heap 表。仅应对评估期间可重新创建数据的表使用 `USING undam` 子句。

### 对象索引

- `undam`：由 `CREATE TABLE ... USING undam` 选择的表访问方法。
- `undam_rel_info(regclass)`：报告 undam 关系的实现诊断信息，包括扫描与数据块计数器。
- `undam.auto_chunk_size`、`undam.chunk_size`、`undam.alloc_chains` 与 `undam.max_relations`：控制数据块选择和共享关系状态的启动/配置设置。在确切源码与工作负载得到验证前，应保留上游默认值。

### 运维边界

- `_PG_init()` 会拒绝常规会话期加载；`undam` 必须出现在 `shared_preload_libraries` 中，修改该设置需要重启。
- 版本 `0.1` 是依赖 PostgreSQL 内部表访问 API 的预览存储引擎代码。已复核源码包含明确的 `NOT_IMPLEMENTED` 错误路径，也没有当前支持版本矩阵。
- 原地更新不会消除 MVCC 历史或维护需求。撤销链旧版本与可复用数据块仍依赖 vacuum 和全局可见性判断。
- 由于它替换了 heap 存储行为，普通 SQL 成功不足以完成验证。应测试索引、唯一约束、行锁、并发更新/删除、长时间快照、vacuum、截断、批量装载、TOAST 大小值、崩溃恢复、WAL 重放、校验和、物理备份、复制、升级与扩展移除。
- 只能使用一次性数据和正在评估的确切 PostgreSQL 构建。没有上游兼容性与恢复证据时，不要在服务器构建之间迁移 undam 关系，也不要把示例视为生产认可。
