## 用法

来源：

- [Official source](https://github.com/turbot/steampipe-postgres-fdw/blob/v2.3.1/README.md)
- [Official source](https://github.com/turbot/steampipe-postgres-fdw/blob/v2.3.1/fdw/steampipe_postgres_fdw.control)
- [Official source](https://github.com/turbot/steampipe-postgres-fdw/blob/v2.3.1/fdw/steampipe_postgres_fdw--1.0.sql)
- [Official source](https://github.com/turbot/steampipe-postgres-fdw/blob/v2.3.1/CHANGELOG.md)

steampipe_postgres_fdw 将 Steampipe 插件暴露为 PostgreSQL 外部表。分发版本 2.3.1 保留扩展 SQL 版本 1.0。核心 FDW 通过 gRPC 连接插件服务；针对单个插件生成的独立扩展属于不同构建。

### 启用与连接

```sql
CREATE EXTENSION steampipe_postgres_fdw;
CREATE SERVER fdw_aws
  FOREIGN DATA WRAPPER steampipe_postgres_fdw
  OPTIONS (wrapper 'aws');
```

安装 SQL 中的包装器名称是 steampipe_postgres_fdw，旧架构示例将其简称为 fdw。导入或声明服务商外部表前，应配置兼容的插件服务及凭据。核心库需要管理员安装，但无需预加载；应匹配服务器 ABI，并遵循所选部署方式的插件配置说明。

### 版本 2.3.0

该版本修复了运行期参数条件导致的崩溃和错误结果，包括 LATERAL 引用与子查询。扫描遭遇取消、statement_timeout 或插件错误时现在会失败，不再把部分行集当成完整结果。重复或非法连接名以及格式错误的配置文件会被跳过并产生警告；应检查警告，不能假设所有连接都已加载。

独立 FDW 标记为 parallel safe，并从 PostgreSQL 服务器环境读取 GOGC/GOMEMLIMIT。这些是进程设置，不是 SQL GUC。旧 cache_off 命令也已修复。替换匹配库后，应重连或重启已加载旧代码的进程；单凭 extversion 1.0 无法判断运行库是否包含修复。远端 API 仍有各自的权限、分页、一致性和速率限制。

### 分发版本 2.3.1

2.3.1 回补 Darwin 交叉编译及镜像发布流程修复，保留 SQL 版本 1.0，没有增加新的 SQL 使用流程；上述 2.3.0 运行期修复仍然适用。
