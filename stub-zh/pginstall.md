## 用法

来源：

- [官方扩展 README](https://github.com/dimitri/pginstall/blob/66e4274fc544c48be03c719f8d3f03955649a994/src/client/README.md)
- [官方项目 README](https://github.com/dimitri/pginstall/blob/66e4274fc544c48be03c719f8d3f03955649a994/README.md)
- [官方控制文件](https://github.com/dimitri/pginstall/blob/66e4274fc544c48be03c719f8d3f03955649a994/src/client/pginstall.control)

`pginstall` 是一款实验性扩展仓库客户端，会拦截扩展工具命令、下载平台相关归档、在服务器上安装文件，并可把扩展操作提升为超级用户。它早于现代 PostgreSQL 打包实践，形成高风险的软件供应链边界；只应在隔离的遗留评估环境中使用。

### 配置与流程

上游设计按会话加载钩子，并配置可写的归档和安装路径及仓库 URL：

```ini
local_preload_libraries = 'pginstall'
pginstall.archive_dir = '/var/cache/pginstall/fetch'
pginstall.control_dir = '/usr/share/postgresql/extension'
pginstall.extension_dir = '/var/lib/pginstall'
pginstall.repository = 'https://extensions.example.invalid/'
pginstall.custom_path = '/etc/pginstall/custom'
pginstall.sudo = false
```

模块加载后，数据库侧流程是普通扩展 SQL 加仓库视图：

```sql
CREATE EXTENSION pginstall;

SELECT * FROM pginstall.pg_available_extensions;
SELECT * FROM pginstall.pg_available_extension_versions;
```

执行 `CREATE EXTENSION`、`ALTER EXTENSION ... UPDATE` 或 `DROP EXTENSION` 时，钩子可能获取缺失制品。

### 安全边界

归档下载、解压、控制文件重写、动态库安装、自定义前后置 SQL 与可选的 `pginstall.sudo` 都会跨越信任边界。仓库或可写缓存被攻破后，可能以 PostgreSQL 服务账号或超级用户执行原生代码。HTTPS 本身不能证明制品来源；必须要求签名不可变制品、校验摘要、严格路径所有权、允许列表及离线审查流程。

### 兼容性与恢复

源码和文档面向 PostgreSQL 9.1 时代 API，并提到已废弃设置。不能假定 ProcessUtility 钩子、文件布局、归档格式、提权或自定义脚本在当前 PostgreSQL 上安全工作。测试必须使用一次性集群，且不能带生产凭据。

安装在文件系统的扩展文件不属于常规数据库持久性。数据库对象备份不会保留仓库、缓存、共享库、控制文件或自定义脚本。因此重建和灾难恢复需要另行版本化的制品仓库及精确的平台清单。
