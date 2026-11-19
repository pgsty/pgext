## 用法

来源：

- [1.0.2 release](https://github.com/bitsinfra/gnomon/releases/tag/v1.0.2)
- [官方资料](https://github.com/bitsinfra/gnomon/blob/v1.0.2/src/c/README.md)
- [官方资料](https://github.com/bitsinfra/gnomon/blob/v1.0.2/src/c/gnomon_c.control)
- [官方资料](https://github.com/bitsinfra/gnomon/blob/v1.0.2/LICENSE)
- [官方资料](https://github.com/bitsinfra/gnomon/blob/v1.0.2/src/c/gnomon_c--1.0.2.sql)
- [官方资料](https://github.com/bitsinfra/gnomon/blob/v1.0.2/docs/installation.md)
- [官方资料](https://github.com/bitsinfra/gnomon/blob/v1.0.2/docs/security.md)

gnomon_c 1.0.2 是 Gnomon 面向 PostgreSQL 15–18 的可选 C 辅助扩展，依赖 gnomon。C 版本触发器可不预加载使用；路由钩子、共享内存优化和心跳工作进程则要求预加载并重启。

### 启用

```ini
shared_preload_libraries = 'gnomon_c'
gnomon.heartbeat_databases = 'app'
gnomon.heartbeat_interval = '1s'
```

```sql
CREATE EXTENSION gnomon;
CREATE EXTENSION gnomon_c;
SELECT gnomon.use_c_trigger('public.orders'::regclass, true);
```

表必须已经被跟踪，调用方须拥有该表的管理权限。C 触发器仅支持 known-time、row-mode 且没有宽键的表，不支持的组合会被拒绝。配置实际数据库名并保留其他预加载项；创建与共享配置需要管理员控制。

### 运维

预加载后，钩子会路由分支读写，并提供额外的锁定、点查与 upsert 行为。必须遵循 Gnomon 安全模型，不向应用授予内部 reader/search 角色。升级时先升级 SQL 核心，再使用匹配文件升级 gnomon_c，并重启已加载被替换库的进程。辅助扩展不会让 unlogged 临时历史持久化，也不会消除保留策略与心跳要求。

### 1.0.2 修正

SQL 升级会修复 1.0.1 在分支创建早于表跟踪时保存的读取集合。合并后会恢复会话原先选择的分支。C 伴随扩展会拒绝在分支上读取系统列，避免此前的崩溃。先将 SQL 核心更新到 1.0.2，再更新已安装的匹配 C 扩展；替换预加载库后需重启使用它的进程。
