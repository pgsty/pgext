## 用法

来源：

- [README 0.3.0](https://github.com/AddressForAll/pg_govLite/blob/0.3.0/README.md)
- [Installation and recovery](https://github.com/AddressForAll/pg_govLite/blob/0.3.0/docs/INSTALLATION.md)
- [Control](https://github.com/AddressForAll/pg_govLite/blob/0.3.0/pg_govlite.control)
- [Versioned SQL migration](https://github.com/AddressForAll/pg_govLite/blob/0.3.0/extension/pg_govlite--0.2.0--0.3.0.sql)

`pg_govlite` 0.3.0 使用 `gvlt`、`lib` 等多个模式提供元数据与受治理标签框架。它是 pg_govLite 的多模式布局；每个数据库只应选择一种布局。

### 启用与检查

上游要求 PostgreSQL 15 或更高版本。安装服务器对应的 `plpython3u` 软件包，并让 PL/Python 使用的 Python 解释器能够加载 PyYAML 和 Chevron。以超级用户启用扩展：

```sql
CREATE EXTENSION pg_govlite VERSION '0.3.0' CASCADE;
SELECT * FROM gvlt.governance_check();
SELECT * FROM gvlt.policy_check();
SELECT lib.yaml_to_jsonb('name: sample');
SELECT lib.jsonb_mustache_render('Hello {{name}}', '{"name":"world"}'::jsonb);
```

`CASCADE` 只启用具备安装文件的扩展依赖，不会安装操作系统或 Python 软件包。上游未要求预加载或重启服务器。

### 治理接口

`gvlt.tag_include` 定义受治理标签；`gvlt.tagobj_include` 将标签关联到数据库对象。使用 `gvlt.obj_tags`、`gvlt.tag_search`、`gvlt.medallion_objects` 查询元数据。禁用标签或关联会保留目录历史。

默认关闭限制性治理。启用它需要 DBA 操作及已有的受限主登录角色。它会调整写权限，并通过数据库级事件触发器执行限制；读取授权和行级安全策略仍须由 DBA 单独决定。应用到已有业务表前，应审查所有权及策略。

### 权限与恢复

框架函数初始仅向所有者开放。应明确授予模式 `USAGE` 与函数 `EXECUTE`。`lib.yamlfile_to_jsonb` 通过 `pg_read_file` 读取文件；Mustache 的文件片段使用 PostgreSQL 进程的文件系统权限。无意开放服务端文件访问时，应将可选片段路径留空。

版本化的 0.2.0 到 0.3.0 迁移保留治理数据。独立安装器创建的实例或布局切换需要单独迁移。上游指出，此基线中常规 `pg_dump` 会省略原生扩展表的内容；恢复方案必须保留治理状态及其对象 OID 引用。
