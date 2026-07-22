## 用法

来源：

- [Official README](https://github.com/NaokiNakamichi/postgres_extension/blob/47029106ccbc1fe00d0ec9a612ddfedd7abf6847/README.md)
- [Version 1.0 control file](https://github.com/NaokiNakamichi/postgres_extension/blob/47029106ccbc1fe00d0ec9a612ddfedd7abf6847/forbid_truncate.control)
- [ProcessUtility hook implementation](https://github.com/NaokiNakamichi/postgres_extension/blob/47029106ccbc1fe00d0ec9a612ddfedd7abf6847/forbid_truncate.c)

`forbid_truncate` 1.0 安装 PostgreSQL `ProcessUtility` 钩子，拒绝后端看到的每一条 `TRUNCATE`。它是防止误清空表的粗粒度保护，不提供按角色、表、数据库或维护操作设置的例外。

### 预加载与验证

安装共享库，把它加入 `shared_preload_libraries`，然后重启 PostgreSQL：

```conf
shared_preload_libraries = 'forbid_truncate'
```

`CREATE EXTENSION` 只在数据库中登记扩展；带版本 SQL 不会创建可调用对象。重启后用一次性表验证钩子：

```sql
CREATE EXTENSION forbid_truncate;

CREATE TEMP TABLE truncate_guard_demo (id integer);
INSERT INTO truncate_guard_demo VALUES (1);
TRUNCATE truncate_guard_demo;
```

最后一条语句应以 `forbid truncate` 失败。应分别验证所有重要数据库和连接路径，因为没有预加载该库而启动的会话不受保护。

### 覆盖范围与钩子兼容性

该钩子会阻止普通用户和超级用户的 `TRUNCATE`。它不阻止 `DELETE`、`DROP TABLE`、分区分离/删除、表替换或其他数据删除方式，因此正常的权限设计、备份和变更控制仍不可缺少。

实现虽然保存了先前注册的 `ProcessUtility_hook`，但处理非 TRUNCATE 命令时不会继续调用它，而是直接调用 `standard_ProcessUtility`。这可能根据预加载顺序绕过其他扩展的钩子。必须用集群中实际的 `shared_preload_libraries` 顺序，与每个审计、DDL 控制或 utility-hook 扩展组合测试。该扩展没有 GUC 或 SQL 绕过开关，停用或变更保护需要修改配置并重启。
