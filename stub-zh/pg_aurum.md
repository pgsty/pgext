## 用法

来源：

- [Release 3.3.0](https://github.com/matroidbe/pg_extensions-releases/releases/tag/pg_aurum-v3.3.0)
- [Official control, SQL and manifest archive](https://github.com/matroidbe/pg_extensions-releases/releases/download/pg_aurum-v3.3.0/pg_aurum-3.3.0-pg18-linux-amd64.tar.gz)
- [SHA256SUMS](https://github.com/matroidbe/pg_extensions-releases/releases/download/pg_aurum-v3.3.0/SHA256SUMS)
- [MIT license](https://github.com/matroidbe/pg_extensions-releases/releases/download/pg_aurum-v3.3.0/LICENSE-pg_aurum)

`pg_aurum` 3.3.0 为 Aurum 增量流水线提供 WAL 工作进程和数据库内编译器。已核对的发行物面向 PostgreSQL 18，提供 Linux amd64 与 macOS arm64 二进制，附带 control、安装 SQL、升级 SQL 和 MIT 许可证。该发行仓库未公开实现源码。

### 启用与检查

使用已校验且与服务器平台完全匹配的上游归档。制品清单要求预加载、逻辑 WAL 和重启。保留现有预加载项，并明确指定工作进程数据库：

```ini
shared_preload_libraries = 'pg_aurum'
wal_level = logical
aurum.database = 'postgres'
```

由超级用户在该数据库创建扩展。扩展不可迁移，创建 `aurum`、`_aurum` 模式、FDW/服务器、事件触发器和流水线目录。

```sql
CREATE EXTENSION pg_aurum;
SELECT aurum_worker_status();
SELECT * FROM aurum_wal_status();
SELECT * FROM aurum_checkpoint_status();
```

### 流水线生命周期

对于已经单独审核的 Aurum 流水线定义，将名称与源码文本绑定为 SQL 参数：

```sql
SELECT * FROM aurum.install($1::text, $2::text);
SELECT * FROM aurum.start($1::text);
```

`aurum.install` 编译并安装定义，`aurum.install_compiled` 接收已编译的 SQL。也可使用服务器 `aurum` 暂存外部表声明，通过 `aurum.deploy_preview` 检查，再由 `aurum.deploy` 部署。处理已存在对象的默认策略为 `error`。

`aurum.stop`、`aurum.restart`、`aurum.recompile` 和 `aurum.uninstall` 修改流水线状态。其 SQL 返回列涉及移除发布、截断和初始加载；对保留数据执行前应核对具体操作。`aurum_process_pending`、`aurum_stats` 和 `aurum_force_checkpoint` 提供工作进度与检查点控制。

### 搜索与边界

对已配置的搜索目标，`aurum.search` 返回键、得分和高亮，`aurum.search_rows` 还返回文档。`aurum.search_chunks` 为 RAG 检索片段，`aurum.search_facets` 与 `aurum.search_suggest` 提供分面和建议。这些包装函数接收目标关系和查询；默认值与返回列以安装 SQL 为准。

SQL 包装函数说明会按调用者权限重新检查行。安装 SQL 并未为所有原生函数建立严格的 PUBLIC 权限策略。应将流水线编译、部署和工作进程管理限制给可信角色，并在向应用开放搜索前审核授权。逻辑复制槽可能保留 WAL，需要监控槽滞后和工作进程健康。这些二进制与 SQL 检查确认的是身份和接口，不代表未公开的实现已通过生产验证。
