## 用法

来源：

- [AlohaDB README](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/README.md)
- [Control 1.0](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_columnar/alohadb_columnar.control)
- [SQL 1.0](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_columnar/alohadb_columnar--1.0.sql)
- [访问方法实现](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_columnar/alohadb_columnar.c)
- [存储定义](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_columnar/alohadb_columnar.h)
- [Meson 构建](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_columnar/meson.build)

`alohadb_columnar` 1.0 是 AlohaDB PostgreSQL 分支中的实验性表访问方法。它将仅追加的分析数据组织为行条带，并对每列使用 zstd 压缩。此修订没有确立正常的 MVCC 快照隔离语义，不适合依赖这一语义的工作负载，也不能据此声称兼容原生 PostgreSQL。

### 基本用法

该分支的 Meson 构建仅在 zstd 可用时包含此模块。所检查的修订没有模块 Makefile。在对应 AlohaDB 服务器中安装构建产物后，由超级用户创建扩展：

```sql
CREATE EXTENSION alohadb_columnar;
CREATE TABLE events_col (
  id bigint,
  ts timestamptz,
  payload jsonb
) USING columnar;
SELECT * FROM alohadb_columnar_info('events_col');
```

实现面向 INSERT 与 SELECT 工作负载，通过事务回调在提交前刷新普通插入操作的缓冲数据。源码未声明预加载或重启要求。扩展可重定位，但访问方法使用数据库级名称 `columnar`，会与同名的其他实现冲突。

### 对象与配置

- `columnar` 是表访问方法，`columnar_tableam_handler` 是其内部处理函数。
- `alohadb_columnar_info(regclass)` 返回 `total_stripes`、`total_rows`、`total_size` 和 `compression`。这里的大小是压缩条带大小之和，不是通用的关系磁盘占用指标；仅应对使用该访问方法的表调用。
- `alohadb.columnar_stripe_row_count` 控制每个条带的行数，默认为 150000，范围为 1000 到 10000000。在所检查的实现中，它属于超级用户级设置参数。

### 运行限制

不支持 UPDATE、DELETE、行锁、推测性插入与索引扫描。快照检查回调无条件返回 true，不能从仅追加设计推断其具有 PostgreSQL 通常的 MVCC 可见性语义。README 宣称的按需列扫描，也未在此作为经核验的执行能力。

持久化数据与版本升级均应按实验性功能处理：此源码修订没有提供存储格式迁移路径文档或生产兼容保证。更换模块前应核对对应分支版本，并保留可恢复的数据副本。
