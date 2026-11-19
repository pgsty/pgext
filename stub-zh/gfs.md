## 用法

来源：

- [README.md](https://github.com/Guepard-Corp/gfs/blob/v0.5.3/crates/extensions/gfs/README.md)
- [gfs.control](https://github.com/Guepard-Corp/gfs/blob/v0.5.3/crates/extensions/gfs/gfs.control)
- [Cargo.toml](https://github.com/Guepard-Corp/gfs/blob/v0.5.3/crates/extensions/gfs/Cargo.toml)
- [src/lib.rs](https://github.com/Guepard-Corp/gfs/blob/v0.5.3/crates/extensions/gfs/src/lib.rs)
- [src/sql/schema.sql](https://github.com/Guepard-Corp/gfs/blob/v0.5.3/crates/extensions/gfs/src/sql/schema.sql)

`gfs` 使用规划器钩子，对远程 PostgreSQL 数据库进行惰性逻辑复制。GFS 0.5.3 中的当前 Rust 模块仍使用扩展版本 0.0.1。`c-ref/` 中旧的 C 表访问方法实现仅供参考，不是当前接口。Cargo 提供 PostgreSQL 16–18 特性。

### 启用与路由

由超级用户安装不可迁移的扩展及运行依赖：

```sql
CREATE EXTENSION postgres_fdw;
CREATE EXTENSION dblink;
CREATE EXTENSION gfs;
```

为克隆数据库启用 `session_preload_libraries = 'gfs'`，保留已有条目并重新连接。仅执行 CREATE EXTENSION 不会在后续会话中安装规划器钩子。GFS 引导流程准备普通本地表、`postgres_fdw` 外部表和用户映射；并行补全与源端诊断使用 `dblink`。连接凭据应保存在受限映射中。

用 `gfs.register_clone(local, source_ref, key_col)` 注册已准备好的本地表与外部表。整数键范围查询补全缺失区间，选择性谓词可补全有上限的切片，无界连接和聚合可转交外部表执行，完全缓存的表在本地查询。`gfs.warm(local, p_freeze, p_conn)` 完整物化表，`gfs.unregister_clone` 移除注册。普通远程读取并非冻结的时间点数据库快照。

### 管理

`gfs.clone_source`、`gfs.cached`、`gfs.cached_predicate` 和 `gfs.clones` 展示路由与缓存状态。`gfs.calibrate` 测量源链路权重，`gfs.cost` 设置策略，`gfs.budget` 限制源端访问速率。当前 SQL 还提供源端漂移检查、`gfs.pull`、重新同步及显式冻结辅助函数；对已产生分歧的克隆使用前，应审查其删除或覆盖选项。

源引用、键列和谓词必须由管理员控制；规划器/SPI 路径使用不安全的原生 FFI。引导流程可能省略外键和生成列，因此不是精确的物理副本。README 列出尚未完成的加固和过时的测试套件，不能仅凭代码存在断言这些路径已验证。SQL 版本仍为 0.0.1，替换库不会迁移旧原型目录。应规划受控重建克隆或经过审查的显式迁移，不能假设 ALTER EXTENSION 提供升级。
