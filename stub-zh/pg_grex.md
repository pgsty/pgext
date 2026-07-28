## 用法

来源：

- [官方上游 README](https://github.com/rustedbytes/pg_grex/blob/f0c830c13d93fb039a6904312fb19356040b7aad/README.md)
- [官方扩展控制文件 (pg_grex.control)](https://github.com/rustedbytes/pg_grex/blob/f0c830c13d93fb039a6904312fb19356040b7aad/pg_grex.control)
- [官方实现源代码](https://github.com/rustedbytes/pg_grex/blob/f0c830c13d93fb039a6904312fb19356040b7aad/src/lib.rs)

`pg_grex` — PostgreSQL 扩展，用于从一组输入字符串生成通用正则表达式。使用 Rust 语言并借助 pgrx 框架实现。用于相应的 SQL 或数据库实用工具工作流。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_grex;

CREATE TABLE product_skus (
    id  serial PRIMARY KEY,
    sku text
);

INSERT INTO product_skus (sku) VALUES
('PROD-100'), ('PROD-101'), ('PROD-102'),
('TEST-999'), ('BETA-001'), ('BETA-002');
```

在目标数据库中安装扩展，当可用时运行上游最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `grex_build` 是一个扩展函数。

### 要求与注意事项

- 该目录记录扩展版本 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
