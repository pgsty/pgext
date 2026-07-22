## 用法

来源：

- [上游 README](https://github.com/Njorda/db721_fdw/blob/ea42fa0cce29e607844f69e1d942a9af55e65fd6/README.md)
- [Cargo 软件包元数据](https://github.com/Njorda/db721_fdw/blob/ea42fa0cce29e607844f69e1d942a9af55e65fd6/Cargo.toml)
- [Rust FDW 实现](https://github.com/Njorda/db721_fdw/blob/ea42fa0cce29e607844f69e1d942a9af55e65fd6/src/lib.rs)
- [DB721 解析器](https://github.com/Njorda/db721_fdw/blob/ea42fa0cce29e607844f69e1d942a9af55e65fd6/src/parser/db721.rs)

`db721_fdw` `0.0.0` 版是用 Rust/pgrx 编写的教学型外部数据包装器，读取 CMU 15-721 项目 1 使用的列式文件格式。外部表提供服务器本地 `filename`；实现只暴露只读扫描回调，并实验性地下推投影和简单比较过滤条件。

### 示例

```sql
CREATE EXTENSION db721_fdw;
CREATE SERVER db721_server FOREIGN DATA WRAPPER db721_fdw;
CREATE FOREIGN TABLE chickens (
  identifier integer,
  farm_name varchar,
  weight_g real
) SERVER db721_server
  OPTIONS (filename '/data/data-chickens.db721', table 'Farm');
SELECT * FROM chickens WHERE identifier = 1;
```

README 示例包含 `table 'Farm'`，但复核的 FDW 只收集该选项，从未使用它；实现始终从 `filename` 读取元数据，也不会使用内嵌元数据中的表名在逻辑表之间选择。不要依赖 `table` 进行路由或隔离。

这是课程/项目代码，并非通用存储 FDW。验证器接收选项却不检查，若干类型和表达式不受支持，规划器估算只是固定的临时值，源码路径中还有未完成的 panic 分支。安装被标为仅超级用户，文件读取使用 PostgreSQL 操作系统账户权限。只能在隔离测试实例中处理可信文件。
