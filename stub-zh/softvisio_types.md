## 用法

来源：

- [官方上游 README](https://github.com/softvisio/postgresql-softvisio-types/blob/d126b0f53d517a0c5677327618659f4de6b3d257/README.md)
- [官方扩展控制文件 (softvisio_types.control)](https://github.com/softvisio/postgresql-softvisio-types/blob/d126b0f53d517a0c5677327618659f4de6b3d257/softvisio_types.control)
- [官方扩展 SQL (softvisio_types--1.1.0.sql)](https://github.com/softvisio/postgresql-softvisio-types/blob/d126b0f53d517a0c5677327618659f4de6b3d257/softvisio_types--1.1.0.sql)

`softvisio_types` — PostgreSQL 额外类型扩展。当应用程序需要此类型、域或其操作符时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION softvisio_types;

CREATE EXTENSION IF NOT EXISTS softvisio_types;

ALTER EXTENSION softvisio_types UPDATE;

DROP EXTENSION IF EXISTS softvisio_types;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `lo_size(oid)` 是一个扩展函数，返回 `int53`。
- `int53` 是一个由扩展定义的域。
- `number1` 是一个由扩展定义的域。
- `number10` 是一个由扩展定义的域。
- `number11` 是一个由扩展定义的域。
- `number12` 是一个由扩展定义的域。
- `number13` 是一个由扩展定义的域。
- `number14` 是一个由扩展定义的域。
- `number15` 是一个由扩展定义的域。
- `number16` 是一个由扩展定义的域。
- `number2` 是一个由扩展定义的域。
- `number3` 是一个由扩展定义的域。
- `number4` 是一个由扩展定义的域。
- `number5` 是一个由扩展定义的域。

### 要求与注意事项

- 控制文件声明默认版本为 `1.2.0`。
- 控制文件标记该扩展为可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
