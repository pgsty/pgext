## 用法

来源：

- [官方 README](https://github.com/ycku/pgeyes/blob/e09e243faaf4837a05366c3fc9d573909dbd053d/README.md)
- [0.0.2 版本扩展 SQL](https://github.com/ycku/pgeyes/blob/e09e243faaf4837a05366c3fc9d573909dbd053d/pgeyes--0.0.2.sql)
- [扩展控制文件](https://github.com/ycku/pgeyes/blob/e09e243faaf4837a05366c3fc9d573909dbd053d/pgeyes.control)

`pgeyes` `0.0.2` 版本只安装一个 PL/pgSQL 环境检查函数。虽然项目 README 描述了更广泛的远程管理工具集合，但在这个已复核版本中，`CREATE EXTENSION pgeyes` 不会安装外部数据包装器、远程检查 API 或元数据表。

### 运行检查

上游说明建议使用专用模式，并撤销 `PUBLIC` 访问权限：

```sql
CREATE SCHEMA pgeyes;
REVOKE ALL ON SCHEMA pgeyes FROM PUBLIC;
CREATE EXTENSION pgeyes WITH SCHEMA pgeyes;

SELECT * FROM pgeyes.pgeyes();
```

`pgeyes()` 返回一行，包含四个文本/布尔字段：

- `item`：`PostgreSQL version number`
- `result`：`server_version_num` 的值
- `valid`：版本号至少为 `90600` 时为 true
- `description`：固定阈值 `>=90600`

这个检查只比较版本号。它不会验证配置、安全、连接、复制、备份或扩展兼容性；对于当前系统，PostgreSQL 9.6 阈值也不再是有用的健康门槛。

项目没有提供扩展升级脚本；README 建议删除后重新创建。执行前应检查依赖对象。由于项目已废弃，且安装 API 比项目描述小得多，新系统应改用直接的 PostgreSQL 系统目录查询或监控查询。
