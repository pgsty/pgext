## 用法

来源：

- [官方 UPID README](https://github.com/carderne/upid/blob/d820a74227adb803d378d89da9eacba78c1a5e09/README.md)
- [PostgreSQL 扩展实现](https://github.com/carderne/upid/blob/d820a74227adb803d378d89da9eacba78c1a5e09/upid_pg/src/lib.rs)
- [PostgreSQL 扩展清单](https://github.com/carderne/upid/blob/d820a74227adb803d378d89da9eacba78c1a5e09/upid_pg/Cargo.toml)
- [核心标识符实现](https://github.com/carderne/upid/blob/d820a74227adb803d378d89da9eacba78c1a5e09/upid_rs/src/lib.rs)

`upid_pg` 增加 `upid` 类型：一种可排序的 128 位标识符，包含短语义前缀、时间戳、随机位和格式版本。它适合 UUID 大小的键既要全局生成，又要显示 `user` 之类紧凑对象类型的场景。

### 生成并存储 UPID

```sql
CREATE EXTENSION upid_pg;

CREATE TABLE users (
    id   upid NOT NULL DEFAULT gen_upid('user') PRIMARY KEY,
    name text NOT NULL
);

INSERT INTO users (name) VALUES ('Ada')
RETURNING id,
          upid_to_prefix(id),
          upid_to_timestamp(id);
```

二进制布局首先放置 40 位 Unix 时间戳以支持大致按时间排序，其后是 64 位随机数和 24 位前缀加版本。时间戳精度为 `256` 毫秒，因此同一时间桶内的生成顺序是随机的，而不是单调的。

### 转换函数

- `gen_upid(prefix)` 使用当前 UTC 时间生成新值。
- `upid_to_prefix(upid)` 和 `upid_to_timestamp(upid)` 提取内嵌前缀与近似时间戳。
- `upid_to_bytea(upid)` 暴露 16 字节表示。
- `upid_to_uuid(upid)` 和 `upid_from_uuid(uuid)` 保留相同的 128 位。

扩展还定义了 `uuid` 与 `upid` 之间的隐式转换，以及从 `upid` 到 `bytea` 或 `timestamp` 的隐式转换：

```sql
SELECT upid_to_uuid(gen_upid('user'));
SELECT upid_from_uuid(gen_random_uuid());
```

把任意 UUID 转换为 `upid` 可以无损还原位模式，但解出的前缀和时间戳没有语义。

### 前缀与成熟度边界

生成标识符时应传入恰好四个小写拉丁字母。实现会静默规范化其他输入：较短前缀右侧补 `z`，较长前缀截断为四个字符，字母表外的字符替换为 `z`。如果不能接受静默规范化，应在应用代码中校验前缀。

内嵌时间戳只是元数据，不是认证信息或可信创建时间证明。64 位随机数提供概率唯一性，并非数据库强制的无碰撞保证。目录版本 `0.0.0` 表示 API 仍处于 alpha 阶段；已复核的 pgrx 清单包含 PostgreSQL 15 至 18 的构建特性，但仍要为目标服务器确认实际软件包可用性和升级兼容性。采用前应固定修订版，并测试类型输入输出、索引、转换、备份恢复和升级。
