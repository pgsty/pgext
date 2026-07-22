## 用法

来源：

- [官方 README](https://github.com/s-hironobu/dont_drop_db/blob/8610f0b4b37af5634135414614dad1b3080bc8de/README.md)
- [官方控制文件](https://github.com/s-hironobu/dont_drop_db/blob/8610f0b4b37af5634135414614dad1b3080bc8de/dont_drop_db.control)

`dont_drop_db` 阻止对服务器配置参数所列数据库名执行 `DROP DATABASE`。它是共享预加载的误操作防护，并非授权边界，也不能代替备份：高权限操作员可以修改设置、重启时不加载模块，或通过其他方式删除数据。

### 核心流程

在 PostgreSQL 启动前配置模块：

```ini
shared_preload_libraries = 'dont_drop_db'
dont_drop_db.list = 'postgres,template0,template1,production'
```

重启服务器，可选地在管理数据库中登记扩展对象，并在一次性集群上验证受保护和未保护的名称：

```sql
CREATE EXTENSION dont_drop_db;

SHOW dont_drop_db.list;
```

默认列表是 `postgres,template0,template1`。特殊值 `all` 会阻止所有数据库名，而只包含空格的值会关闭保护。

### 配置语义

`dont_drop_db.list` 是由预加载模块检查的逗号分隔数据库名列表。数据库改名或克隆后必须复核精确拼写。由于行为在服务器启动时加载，只在某个数据库中改变扩展对象并不能建立集群级保护。

### 运维注意事项

配置管理与重启验证应纳入控制措施。部署后要确认 `shared_preload_libraries`，检查服务器日志中是否有模块加载失败，并针对列表内的一次性数据库名执行负向测试。

该防护只覆盖 `DROP DATABASE`，不覆盖删表、删模式、文件系统删除、存储故障或超级用户绕过配置。必须保留经过验证的备份，并限制配置文件与服务控制权限。上游报告已检查至 PostgreSQL 18 RC1，但每个打包二进制与目标大版本仍需启动及命令钩子测试。
