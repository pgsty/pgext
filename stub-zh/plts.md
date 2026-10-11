## 用法

来源：

- [Compiler prerequisites](https://github.com/Smertos/stopgap/blob/e076ae165ba35c6f2d484acd943451373376e03e/docs/DEVELOPER-QUICKSTART.md)

- [README.md](https://github.com/Smertos/stopgap/blob/e076ae165ba35c6f2d484acd943451373376e03e/README.md)
- [crates/plts/Cargo.toml](https://github.com/Smertos/stopgap/blob/e076ae165ba35c6f2d484acd943451373376e03e/crates/plts/Cargo.toml)
- [crates/plts/src/lib.rs](https://github.com/Smertos/stopgap/blob/e076ae165ba35c6f2d484acd943451373376e03e/crates/plts/src/lib.rs)
- [crates/plts/plts.control](https://github.com/Smertos/stopgap/blob/e076ae165ba35c6f2d484acd943451373376e03e/crates/plts/plts.control)

`plts` 提供 Stopgap 使用的过程语言以及编译、制品 API。这是不可信扩展，安装要求超级用户。Cargo 提供 PostgreSQL 14–18 特性，默认选择 17。

### 核心用法

```ini
shared_preload_libraries = 'plts'
```

```sql
CREATE EXTENSION plts;
SELECT plts.typecheck_ts('export default () => 42;');
```

### 运行边界

`plts.compile_ts()` 编译 TypeScript，`plts.typecheck_ts()` 检查代码，`plts.compile_and_store()` 保存制品，`plts.get_artifact()` 读取制品。独立模块使用默认导出，Stopgap 应用模块使用具名导出。数据库调用位于当前 SQL 事务内。运行能力取决于构建特性与编译器、运行时资源；README 还说明正在迁移编译器，部署前应核对固定修订的运行时契约。编译服务与后端本地运行时配置不同，不能假设创建扩展就会启用所有构建配置。

编译与类型检查要求设置 `shared_preload_libraries = 'plts'` 并重启服务器，以启动共享 TSGo 编译工作进程。
