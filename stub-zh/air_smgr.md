## 用法

来源：

- [Official source](https://github.com/Mahadi-rsio/air/blob/a19a53bea07a4f490937db50deec11cb64566525/README.md)
- [Official source](https://github.com/Mahadi-rsio/air/blob/a19a53bea07a4f490937db50deec11cb64566525/extensions/air_smgr/air_smgr.control)
- [Official source](https://github.com/Mahadi-rsio/air/blob/a19a53bea07a4f490937db50deec11cb64566525/extensions/air_smgr/air_smgr--1.0.sql)
- [Official source](https://github.com/Mahadi-rsio/air/blob/a19a53bea07a4f490937db50deec11cb64566525/extensions/air_smgr/air_smgr.c)
- [Official source](https://github.com/Mahadi-rsio/air/blob/a19a53bea07a4f490937db50deec11cb64566525/docs/architecture/pageserver.md)

`air_smgr` 1.0 是 AIR 为其修改版 PostgreSQL 17 提供的实验性页面服务器存储管理模块。它依赖项目中的存储管理钩子补丁和配套 AIR 页面服务器，不能直接装入原版 PostgreSQL。

### 启用边界

附带 SQL 脚本不创建数据库对象，模块通过启动时预加载启用，服务器配置需由管理员控制。保留已有预加载条目，修改配置后重启专用 AIR 集群：

```ini
shared_preload_libraries = 'supautils,air_smgr'
air_smgr.enabled = off
```

```sql
SHOW air_smgr.enabled;
SHOW air_smgr.pageserver_socket;
SHOW air_smgr.tenant_id;
```

示例特意保持存储重定向关闭。只有在 AIR 初始化的集群中，租户标识、页面服务器状态、WAL 流和恢复流程均已配套配置后，才能启用。普通会话加载该库不会安装存储钩子。

### 设置与数据路径

实现注册了仅在启动时生效的 `air_smgr.enabled`、`air_smgr.read_only`、`air_smgr.pageserver_socket`、`air_smgr.tenant_id`、`air_smgr.lfc_dir` 和 `air_smgr.lfc_max_mb`。源码默认关闭模块、允许写入，套接字为 `/var/lib/air/pageserver.sock`，租户标识为空，缓存目录为 `/var/lib/air/lfc`，缓存为 256 MB。缓存上限为零时关闭缓存。套接字必须匹配实际页面服务器，架构示例使用了不同路径。

启用后，持久关系使用远端存储管理器，临时关系仍使用普通管理器。读取按 LSN 请求页面，并使用本地文件缓存。写入和元数据操作依赖配套页面服务器及 WAL 路径。只读设置拒绝持久写入，但不构成备份或故障切换协议。

### 恢复与兼容性

应将租户数据、页面服务器分层文件、WAL 和配置作为一个恢复单元。模块改变物理存储行为，移除库或只恢复普通本地关系文件，都不是文档支持的回迁到原版 PostgreSQL 的方法。源码没有 SQL 升级路径，也未证明其他大版本兼容性。已核对修订中未找到项目许可证文件。
