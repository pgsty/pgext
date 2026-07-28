## 用法

来源：

- [官方上游 README](https://github.com/dresdres/bgp240-master-thesis/blob/2a96f87df76b8db3d802824cd7e3c0bbff321036/Extensions/README.md)
- [官方扩展控制文件 (payment_ext.control)](https://github.com/dresdres/bgp240-master-thesis/blob/2a96f87df76b8db3d802824cd7e3c0bbff321036/Extensions/payment_ext/payment_ext.control)
- [官方扩展 SQL (payment_ext--0.0.0.sql)](https://github.com/dresdres/bgp240-master-thesis/blob/2a96f87df76b8db3d802824cd7e3c0bbff321036/Extensions/CompiledExtensions/all_ext/controlSql/payment_ext--0.0.0.sql)

`payment_ext` — 支付服务交易标记以及市场原型的 LISTEN/NOTIFY 背景工作者。当应用程序需要此特定数据库功能时，请使用它。在目标 PostgreSQL 构建中测试上游链接的固定版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION payment_ext;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 该目录记录了版本 `0.0.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
