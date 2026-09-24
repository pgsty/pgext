## 用法

来源：

- [README](https://github.com/open-gpdb/yezzey/blob/2f0c013c888ecb078082522ffa9810cb962b0a1a/README.md)
- [Control file / 控制文件](https://github.com/open-gpdb/yezzey/blob/2f0c013c888ecb078082522ffa9810cb962b0a1a/yezzey.control)
- [yezzey.c](https://github.com/open-gpdb/yezzey/blob/2f0c013c888ecb078082522ffa9810cb962b0a1a/yezzey.c)
- [yezzey--1.8.8.sql](https://github.com/open-gpdb/yezzey/blob/2f0c013c888ecb078082522ffa9810cb962b0a1a/yezzey--1.8.8.sql)
- [yezzey--1.8.8--1.8.11.sql](https://github.com/open-gpdb/yezzey/blob/2f0c013c888ecb078082522ffa9810cb962b0a1a/yezzey--1.8.8--1.8.11.sql)

`yezzey` 将追加式表的数据转移到 S3，同时保留 Greenplum 或 Apache Cloudberry 中的查询能力。1.8.11 需要兼容的内核补丁和 YProxy，不适用于标准 PostgreSQL。

### 启用

在集群中部署匹配内核与 YProxy/S3 配置，将 `yezzey` 加入 `shared_preload_libraries` 并重启。初始化代码拒绝在服务器启动之外加载。`yezzey.yproxy_socket` 选择代理套接字，扩展 SQL 安装到目标数据库。

```sql
CREATE EXTENSION yezzey;
CREATE TABLE archive_events (id integer, payload text)
  WITH (appendonly=true, orientation=column) DISTRIBUTED RANDOMLY;
INSERT INTO archive_events VALUES (1, 'example');
SELECT yezzey_define_offload_policy('public', 'archive_events');
SELECT * FROM yezzey_offload_relation_status('archive_events');
SELECT * FROM archive_events;
```

### 对象与维护

`yezzey_define_offload_policy` 配置卸载，`yezzey_load_relation` 将数据恢复到本地存储。`yezzey_offload_relation_status` 报告外部数据大小，`yezzey_relation_describe_external_storage_structure` 展示外部文件布局。支持范围仅限 AO/AOCO 表。

卸载和回载可能获取较强的关系锁。读取依赖 YProxy 与对象存储可用性。备份恢复须同时覆盖引用的 S3 对象和数据库元数据。删除外部对象前应审查垃圾清理行为并先做试运行，破坏性清理模式需要更高权限。控制版本通过随附的 1.8.8 安装 SQL 和 1.8.11 升级 SQL 到达。
