## 用法

来源：

- [README.md](https://github.com/open-gpdb/yezzey/blob/1d8e735a22aa92ac4d149fb48cb6ae5cc332d854/README.md)
- [yezzey.control](https://github.com/open-gpdb/yezzey/blob/1d8e735a22aa92ac4d149fb48cb6ae5cc332d854/yezzey.control)
- [yezzey--2.0.sql](https://github.com/open-gpdb/yezzey/blob/1d8e735a22aa92ac4d149fb48cb6ae5cc332d854/yezzey--2.0.sql)
- [yezzey.c](https://github.com/open-gpdb/yezzey/blob/1d8e735a22aa92ac4d149fb48cb6ae5cc332d854/yezzey.c)
- [docs/README.cleanup.md](https://github.com/open-gpdb/yezzey/blob/1d8e735a22aa92ac4d149fb48cb6ae5cc332d854/docs/README.cleanup.md)

`yezzey` 2.0 将追加式行存/列存表卸载至 S3，并保持 SQL 访问。它要求配套修改过的 OpenGPDB（Greenplum 6）或 Apache Cloudberry 内核与 YProxy，不适用于原版 PostgreSQL。项目发布号为 2.0.0，控制文件中的扩展版本为 2.0。

### 核心用法

```ini
shared_preload_libraries = 'yezzey'
```

```sql
CREATE EXTENSION yezzey;
CREATE TABLE offload_demo (id integer)
  WITH (appendonly=true, orientation=column) DISTRIBUTED RANDOMLY;
INSERT INTO offload_demo VALUES (1);
SELECT yezzey.offload_relation('offload_demo'::regclass);
SELECT * FROM yezzey.offload_relation_status('offload_demo'::regclass);
SELECT yezzey.load_relation('offload_demo'::regclass);
```

### 运行边界

须在数据库集群中预加载 `yezzey` 并重启，再创建扩展；事先配置兼容的 YProxy 端点与对象存储。OpenGPDB 分支为 OPENGPDB_STABLE，Cloudberry 实现在 master 分支。SQL 安装标记为 trusted，但集群配置、对象存储访问和表所有权仍有管理权限边界。

`yezzey.offload_relation` 上传数据期间取得排他关系锁，`yezzey.load_relation` 恢复本地存储。`yezzey.offload_relation_status` 返回各段的字节计数，`yezzey.relation_describe_external_storage_structure` 列出外部文件。2.0 将旧的不带模式接口移动并重命名，须按新 SQL 定义修改调用方，并使用可恢复备份演练迁移。

清理使用 `yezzey.vacuum`、`yezzey.vacuum_tablespace` 或 `yezzey.vacuum_relation`。先保持 `confirm` 为 false，检查请求后再决定删除；`crazyDrop` 是仅超级用户可用的激进模式。保留的备份仍可能需要外部对象，须遵守备份保留边界，并保证恢复系统或备用集群能够访问对象存储。SQL VACUUM 本身不意味着可以删除存储桶中的任意对象。
