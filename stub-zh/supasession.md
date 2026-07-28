## 用法

来源：

- [官方上游 README](https://github.com/snehil-shah/supasession/blob/f09bed26cdac28b2394cd05c68c122b67337075a/README.md)
- [官方扩展控制文件 (supasession.control)](https://github.com/snehil-shah/supasession/blob/f09bed26cdac28b2394cd05c68c122b67337075a/supasession.control)
- [官方扩展 SQL (supasession.sql)](https://github.com/snehil-shah/supasession/blob/f09bed26cdac28b2394cd05c68c122b67337075a/supasession.sql)

`supasession` — > [!WARNING] > 该扩展安装在 supasession 模式中，如果之前已有同名扩展，则可能会导致命名空间冲突。在实现相应的安全、审计或访问控制工作流时使用它。使用上方链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION supasession;

SELECT dbdev.install('Snehil_Shah@supasession');
```

在目标数据库中安装该扩展，当可用时运行上方最小的上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `supasession.disable()` 是一个扩展函数并返回 `void`。
- `supasession.enable()` 是一个扩展函数并返回 `void`。
- `supasession.get_config()` 是一个扩展函数并返回 `supasession`。
- `supasession.set_config(enabled BOOLEAN DEFAULT NULL, max_sessions INTEGER DEFAULT NULL, strategy supasession.enforcement_strategy DEFAULT NULL)` 是一个扩展函数并返回 `supasession`。
- `supasession.sid()` 是一个扩展函数并返回 `uuid`。
- `supasession.enforcement_strategy` 是一个由扩展定义的类型。
- `supasession.config` 是一个由扩展安装或管理的表。
- `supasession` 是一个由扩展创建的模式。

### 要求与注意事项

- 审查后的控制文件声明默认版本 `0.1.2`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
