## 用法

来源：

- [官方 README.md](https://github.com/pg-replayfeed/pg_replayfeed/blob/82bd85807cf70947d2a9471c556fd4c8b1ecf6bb/README.md)
- [官方 pg_replayfeed.control](https://github.com/pg-replayfeed/pg_replayfeed/blob/82bd85807cf70947d2a9471c556fd4c8b1ecf6bb/pg_replayfeed.control)
- [官方 pg_replayfeed--2.0.sql](https://github.com/pg-replayfeed/pg_replayfeed/blob/82bd85807cf70947d2a9471c556fd4c8b1ecf6bb/sql/pg_replayfeed--2.0.sql)
- [官方 semantics.md](https://github.com/pg-replayfeed/pg_replayfeed/blob/82bd85807cf70947d2a9471c556fd4c8b1ecf6bb/docs/semantics.md)
- [官方 security.md](https://github.com/pg-replayfeed/pg_replayfeed/blob/82bd85807cf70947d2a9471c556fd4c8b1ecf6bb/docs/security.md)
- [官方 ci.yml](https://github.com/pg-replayfeed/pg_replayfeed/blob/82bd85807cf70947d2a9471c556fd4c8b1ecf6bb/.github/workflows/ci.yml)

`pg_replayfeed` 2.0 通过逻辑解码输出仅含键的变更消息，将已提交事务表示为 JSON 信封。消费者保存提交 LSN，并处理可能的重复投递。本次审阅的开发修订在 CI 中覆盖 PostgreSQL 16–18。

### 注册并读取变更流

启用逻辑 WAL，并配置足够的复制槽与 WAL 发送进程；修改这些启动参数后需重启。扩展本身不要求共享预加载。

```conf
wal_level = logical
max_replication_slots = 10
max_wal_senders = 10
```

```sql
CREATE EXTENSION pg_replayfeed;
CREATE TABLE feed_demo (id bigint PRIMARY KEY, tenant text, value text);
SELECT replayfeed.create_feed('demo', 'feed_demo', 'id', 'tenant');
SELECT pg_create_logical_replication_slot('demo_slot', 'pg_replayfeed');
INSERT INTO feed_demo VALUES (1, 'team_a', 'first');
SELECT * FROM pg_logical_slot_peek_changes('demo_slot', NULL, NULL);
```

此管理示例以超级用户执行。`create_feed` 接受普通持久表，键列必须非空且具有单列唯一索引；注册时会设置 `REPLICA IDENTITY FULL`。可查询 `replayfeed.feed`，通过 `replayfeed.drop_feed` 删除注册。删除注册不会删除复制槽。

### 消费、权限与保留

`pg_logical_slot_get_changes` 消费消息并推进复制槽，`pg_logical_slot_peek_changes` 则不推进。复制客户端也可使用 PostgreSQL 复制协议持续读取。应在可靠处理后才确认，并按提交位置去重；键变更、租户迁移和截断事件须按官方信封语义处理。

固定的 `replayfeed` 模式包含以 SECURITY DEFINER 运行的管理函数，默认撤销 PUBLIC 执行权限。委派管理员须获得模式与函数授权，并满足对源表的权限检查；复制角色还须能读取目录。选择变更流只是过滤条件，不是租户授权，消费者须自行实现隔离。

监测复制槽延迟和 WAL 磁盘占用。闲置复制槽会持续保留 WAL，直到推进或删除；演示结束后删除该槽：

```sql
SELECT pg_drop_replication_slot('demo_slot');
SELECT replayfeed.drop_feed('demo');
```
