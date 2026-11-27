## 用法

来源：

- [README.md](https://github.com/fredericalix/pgXila/blob/16f2d70299e893b610fdef503f904c28a8c75d84/README.md)
- [extension/Cargo.toml](https://github.com/fredericalix/pgXila/blob/16f2d70299e893b610fdef503f904c28a8c75d84/extension/Cargo.toml)
- [extension/pgxila.control](https://github.com/fredericalix/pgXila/blob/16f2d70299e893b610fdef503f904c28a8c75d84/extension/pgxila.control)
- [extension/src/lib.rs](https://github.com/fredericalix/pgXila/blob/16f2d70299e893b610fdef503f904c28a8c75d84/extension/src/lib.rs)
- [extension/src/tls_guard.rs](https://github.com/fredericalix/pgXila/blob/16f2d70299e893b610fdef503f904c28a8c75d84/extension/src/tls_guard.rs)
- [extension/src/workers.rs](https://github.com/fredericalix/pgXila/blob/16f2d70299e893b610fdef503f904c28a8c75d84/extension/src/workers.rs)
- [extension/sql/schema.sql](https://github.com/fredericalix/pgXila/blob/16f2d70299e893b610fdef503f904c28a8c75d84/extension/sql/schema.sql)
- [docs/api.md](https://github.com/fredericalix/pgXila/blob/16f2d70299e893b610fdef503f904c28a8c75d84/docs/api.md)
- [docker/initdb/10-pgxila.sql](https://github.com/fredericalix/pgXila/blob/16f2d70299e893b610fdef503f904c28a8c75d84/docker/initdb/10-pgxila.sql)
- [docker/postgresql.conf](https://github.com/fredericalix/pgXila/blob/16f2d70299e893b610fdef503f904c28a8c75d84/docker/postgresql.conf)
- [LICENSE](https://github.com/fredericalix/pgXila/blob/16f2d70299e893b610fdef503f904c28a8c75d84/LICENSE)

`pgxila` 0.1.0 是面向 PostgreSQL 19 的早期卡片应用扩展。屏幕、卡片栈、控件、调色板和历史记录保存在 `xila` 模式中，独立的终端及图形客户端渲染 SQL API 返回的内容。已核对源码尚无公开发行版，也没有 SQL 升级路径。

### 启用与连接

配置有效的服务器证书，将库加入现有预加载列表，再重启 PostgreSQL：

```ini
shared_preload_libraries = 'pgxila'
ssl = on
ssl_min_protocol_version = 'TLSv1.3'
```

该扩展不受信任且不可重定位，需要超级用户安装。普通构建拒绝在服务器启动阶段之外加载。认证钩子作用于整个集群：所有 TCP 连接，包括其他数据库的连接，都必须使用 TLS 1.3。本地 Unix 套接字仍可用于管理，钩子没有运行时关闭选项。

通过管理用 Unix 套接字或已验证的 TLS 连接，在应用数据库中创建扩展：

```sql
CREATE EXTENSION pgxila;
SELECT xila.version();
SELECT xila.hello();
SELECT xila.open_stack('desktop');
SELECT * FROM xila.palette();
```

安装时会在缺失的情况下创建 `xila_user` 和 `xila_author` 组角色。已有应用角色需加入 `xila_user`，作者还需加入 `xila_author`。上游镜像同时安装 `vector`，但当前控制文件与 SQL 并不依赖它。SQL 层使用 `plpgsql`。

### 应用流程

客户端先调用 `xila.hello()` 检查 API 版本，再通过 `xila.open_stack()` 获取屏幕定义。`xila.card()` 读取卡片，`xila.save_card()` 写入字段，`xila.act()` 分派动作。卡片仍是底层表中的行，普通表权限和行级安全策略继续生效；卡片栈定义由其所有者及该所有者角色的成员编辑。

作者模式通过 SQL API 创建卡片栈与字段。历史记录保存修改并支持回退，但并发编辑采用后保存者覆盖的语义。客户端可订阅通知以刷新显示数据。当前版本尚未实现脚本引擎、模型代理或实例间连接，工作进程注册函数仍为空。

### 维护

扩展为 `pg_dump` 注册了应用配置和历史表，系统卡片栈由安装过程重新创建。备份时还需保留底层表和角色授权。当前源码修订没有受支持的原地 SQL 升级脚本，应单独评估迁移，不能假定 `ALTER EXTENSION` 可以完成升级。从预加载列表移除库并重启，会改变集群范围的连接保护行为。
