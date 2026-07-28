## 用法

来源：

- [官方上游 README](https://github.com/aquametalabs/aquameta/blob/ec5dbe6b84ea61b5e8b4b10263a0911b848b0025/extensions/event/README.md)
- [官方扩展控制文件 (event.control)](https://github.com/aquametalabs/aquameta/blob/ec5dbe6b84ea61b5e8b4b10263a0911b848b0025/extensions/event/event.control)

`event` — 现在可以运行以下 SQL 来更改数据：仅在应用程序需要此特定数据库功能时使用。在集成到应用程序 SQL 之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION event;

insert into person (name, score) values ('Don Pablo', 14);
update person set name='Sandy Jones', score=score+3 where id=3;
delete from person where id=4;
```

在目标数据库中安装扩展，如果有可用的上游最小示例，请运行该示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 要求与注意事项

- 控制文件声明默认版本为 `0.5.0`。
- 首先安装并验证确认的扩展依赖项：`meta`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用之前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
