## 用法

来源：

- [README](https://api.pgxn.org/src/pg_uuidv7/pg_uuidv7-1.8.0/README.md)
- [Control](https://api.pgxn.org/src/pg_uuidv7/pg_uuidv7-1.8.0/pg_uuidv7.control)
- [Changelog](https://api.pgxn.org/src/pg_uuidv7/pg_uuidv7-1.8.0/CHANGELOG.md)

`pg_uuidv7` 生成按时间排列的 UUIDv7，并支持其时间戳部分与 `timestamptz` 之间的转换。发行版 1.8.0 支持 PostgreSQL 13–19，SQL 扩展版本为 1.8。PostgreSQL 18 已内置 `uuidv7()`，本扩展还提供下面的时间戳转换函数。

### 核心工作流

```sql
CREATE EXTENSION pg_uuidv7;
SELECT uuid_generate_v7();
SELECT uuid_v7_to_timestamptz('018570bb-4a7d-7c7e-8df4-6d47afd8c8fc');
SELECT uuid_timestamptz_to_v7('2023-01-02 04:26:40.637+00', true);
CREATE TABLE events (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v7(),
    data jsonb
);
```

`uuid_timestamptz_to_v7` 的可选布尔参数为 true 时会将随机位清零，可用于构造时间范围边界；普通生成保留随机位。毫秒级时间排序不能保证同一毫秒内生成的每个 UUID 都严格有序。

### 安装与升级

安装匹配 PostgreSQL 主版本的库。控制文件声明扩展可重定位，没有扩展依赖；安装使用默认的超级用户权限要求。上游没有规定预加载或服务器重启。

发行版 1.8.0 增加 PostgreSQL 19 支持。安装匹配文件后，使用随发行版提供的更新路径执行 `ALTER EXTENSION pg_uuidv7 UPDATE`，SQL 扩展记录版本为 1.8。
