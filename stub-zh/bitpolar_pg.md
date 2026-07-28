## 用法

来源：

- [官方上游 README](https://github.com/mmgehlot/bitpolar/blob/97f1885b472eb9b973713c46842c69a653fc3afd/README.md)
- [官方扩展控制文件 (bitpolar_pg.control)](https://github.com/mmgehlot/bitpolar/blob/97f1885b472eb9b973713c46842c69a653fc3afd/bitpolar-pg/bitpolar_pg.control)
- [官方实现源代码](https://github.com/mmgehlot/bitpolar/blob/97f1885b472eb9b973713c46842c69a653fc3afd/bitpolar-pg/src/lib.rs)

`bitpolar_pg` — BitPolar: 近最优向量量化 — 3-8 位压缩，无需训练。在每个主要 AI 框架中都有 58 个集成。使用它来对应于向量、模型或检索工作流。在目标 PostgreSQL 构建中测试所链接的上游修订版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION bitpolar_pg;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `bitpolar_compress` 是一个扩展函数。
- `bitpolar_compression_ratio` 是一个扩展函数。
- `bitpolar_decompress` 是一个扩展函数。
- `bitpolar_inner_product` 是一个扩展函数。
- `bitpolar_version()` 是一个扩展函数。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.3.3`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与所链接的源代码一致。
