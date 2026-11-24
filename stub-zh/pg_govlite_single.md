## 用法

来源：

- [README 0.3.0](https://github.com/AddressForAll/pg_govLite/blob/0.3.0/README.md)
- [Installation and recovery](https://github.com/AddressForAll/pg_govLite/blob/0.3.0/docs/INSTALLATION.md)
- [Control](https://github.com/AddressForAll/pg_govLite/blob/0.3.0/pg_govlite_single.control)
- [Versioned SQL migration](https://github.com/AddressForAll/pg_govLite/blob/0.3.0/extension/pg_govlite_single--0.2.0--0.3.0.sql)

`pg_govlite_single` 0.3.0 是 pg_govLite 的单模式治理布局。框架对象使用前缀放在选定的已有模式中。它是 `pg_govlite` 的替代布局，不应在同一数据库中同时安装两种布局。

### 在选定模式中启用

上游要求 PostgreSQL 15 或更高版本、`plpython3u`，以及服务端 PL/Python 环境中的 PyYAML 和 Chevron。该扩展不可重定位，安装需要超级用户。下面的示例创建专用模式：

```sql
CREATE SCHEMA cliente;
CREATE EXTENSION pg_govlite_single
  WITH SCHEMA cliente VERSION '0.3.0' CASCADE;
SELECT * FROM cliente.gvlt_governance_check();
SELECT * FROM cliente.gvlt_policy_check();
SELECT cliente.lib_yaml_to_jsonb('name: sample');
SELECT cliente.lib_jsonb_mustache_render('Hello {{name}}', '{"name":"world"}'::jsonb);
```

模式名须使用小写标识符且不超过 63 字节，并保留管理角色所有权。依赖软件包须已安装在服务器上。上游未要求预加载或重启服务器。

### 对象与策略

`cliente.gvlt_tagobj_include` 将治理标签关联到对象。`cliente.gvlt_object_name`、`cliente.gvlt_object_name_parse` 解析布局的物理名称；`cliente.gvlt_medallion_objects` 查询已登记的 Medallion 关系。单模式布局使用前缀记录层级，不会为每层单独创建模式。

默认关闭限制性治理。DBA 启用后，它使用数据库级事件触发器，并要求已有的受限主登录角色。策略变更可能移除非主角色的写权限，但不会提供应用读取授权。不要通过修改命名字典重命名已安装对象。

### 访问与升级

函数初始仅向所有者开放；按需授予模式 `USAGE` 与函数 `EXECUTE`。`cliente.lib_yamlfile_to_jsonb` 遵循 `pg_read_file` 权限，可选的 Mustache 文件片段则使用服务器进程的访问权限。

升级复用已登记的模式。0.2.0 到 0.3.0 的脚本保留治理数据，遇到函数名冲突则中止。接管独立安装实例和转换布局需要单独迁移。此基线中常规 `pg_dump` 会省略原生扩展表的内容；恢复必须保留治理状态及对象 OID 引用。
