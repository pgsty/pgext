## 用法

来源：

- [README](https://github.com/xvaara/pg_ztype/blob/ff6fcccda0c5f6bfea59cd0acd2a14ef8cf88d85/README.md)
- [Control](https://github.com/xvaara/pg_ztype/blob/ff6fcccda0c5f6bfea59cd0acd2a14ef8cf88d85/ztype.control)
- [SQL](https://github.com/xvaara/pg_ztype/blob/ff6fcccda0c5f6bfea59cd0acd2a14ef8cf88d85/ztype--0.9.sql)
- [Source](https://github.com/xvaara/pg_ztype/blob/ff6fcccda0c5f6bfea59cd0acd2a14ef8cf88d85/ztype.c)

`ztype` 的项目名为 pg_ztype，提供经过 zstd 压缩的 `ztext`、`zjsonb` 和 `zbytea` 列类型。0.9 是预发布源码快照：存储格式与类型修饰符编码尚未冻结，也没有扩展升级脚本。存储长期数据前，应保留可用的逻辑导出路径。

### 核心用法

上游要求 PostgreSQL 18 及以上版本、libzstd 1.5 及以上版本，并测试了 PostgreSQL 18 和 19。安装需要超级用户，字典辅助函数还依赖 PL/pgSQL。类型位于 `public` 模式，管理对象位于 `ztype` 模式；扩展不可迁移模式，无须预加载。

```sql
CREATE EXTENSION ztype;
CREATE TABLE messages (
  id bigint PRIMARY KEY,
  body ztext(6),
  metadata zjsonb(6)
);
INSERT INTO messages VALUES (1, 'hello', '{"source":"email"}');
SELECT body::text, metadata ->> 'source' FROM messages;
CREATE INDEX messages_metadata ON messages USING gin ((metadata::jsonb));
```

类型修饰符指定 1–22 的压缩级别，并可附加已注册的字典名称或槽号。小于 64 字节或压缩后没有缩小的值按原始形式存储。调用基础类型函数时，先转换成 `text`、`jsonb` 或 `bytea`。目标列采用非默认策略时，为输入参数明确指定基础类型，可避免对无类型值压缩两次。

### 字典与检查

管理员可使用代表性数据训练字典，然后在新列中引用其名称：

```sql
SELECT ztype.train_and_add('message-json',
  'SELECT metadata::jsonb FROM messages LIMIT 20000');
CREATE TABLE archive (metadata zjsonb(6, 'message-json'));
SELECT ztype.inspect(metadata), ztype.validate(metadata) FROM messages;
SELECT * FROM ztype.dictionary_inventory;
SELECT * FROM ztype.column_policies;
```

训练至少需要八个非空样本，查询须返回一个 `text`、`bytea` 或 `jsonb` 列。JSON 字典应使用二进制 `jsonb` 值训练。`ztype.validate` 对有效值返回 NULL，否则返回诊断信息；`ztype.inspect` 无须完整解压即可报告编码、大小、策略与字典身份。

| 接口 | 用途 |
| --- | --- |
| `ztype.train_dictionary`、`ztype.add_dictionary`、`ztype.import_dictionary` | 训练、注册字典，或按指定槽号导入字典字节 |
| `ztype.recompress`、`ztype.matches_policy` | 按策略生成值，或检查现有值是否符合策略 |
| `ztype.set_column_policy`、`ztype.finish_column_policy` | 延后整表重写来切换策略，并在完成后验证 |
| `raw_length`、`prefix` | 获取逻辑字节长度，或读取不会截断字符的文本前缀 |
| `ztype.zstd` | 向具有解码能力的客户端返回 zstd 帧 |
| `ztype.build_info`、`ztype.dictionary_cache_stats`、`ztype.decode_cache_stats` | 检查库和格式版本、后端缓存 |

### 索引、权限与维护

- 压缩类型支持等值比较、分组和哈希索引，但没有排序运算符或 B-tree 算子类，排序时应先转换类型。`zjsonb` 支持 JSON 读取运算符，GIN 索引使用上例中的转换表达式。
- 字典注册函数的执行权限需要显式授予；普通列用户不需要访问注册表。字典字节可能泄露训练数据。数据库级注册表只允许追加，槽号不能重复使用。
- 缺失字典时，读取使用该字典压缩的帧会失败。物理复制会携带注册表；逻辑复制和恢复需要按上游说明协调字典与槽号。随附的 `ztype-sync` 工具可传输并检查注册表，不能在同一槽号下换成无关字典。
- 修改列类型修饰符通常会在排他锁下重写整表。分阶段策略函数让尚未重新压缩的行处于待完成状态，最终验证仍需在排他锁下扫描。分批重写前应备份，并估算 WAL 与表膨胀空间。
- `ztype.dictionary_cache_size` 默认为每个后端 64 MB，仅限制缓存字典，并不约束全部 zstd 内存；`work_mem` 也不限制编码器上下文和训练缓冲区。写入压缩与读取解压都消耗 CPU，应结合实际负载选择级别。
