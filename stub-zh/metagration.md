## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/metagration/metagration-1.0.4/README.md)
- [官方扩展 SQL (metagration.sql)](https://api.pgxn.org/src/metagration/metagration-1.0.4/sql/metagration.sql)

`metagration` — 逻辑复制的 PostgreSQL 数据库迁移是一项精细的工作，需要在正确的时间应用正确的脚本，并确保在可能的停机时间内确保副本正确更新。请参考 *文档中的警告*。在移动、转换或整合相应的数据时使用它。使用链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

经过审查的分发使用了过时的 SQL 或非控制文件安装布局，因此没有建立现代独立的 `CREATE EXTENSION` 和升级工作流。遵循上游固定版本的安装机制，并在隔离数据库中验证安装的对象。

### 重要对象

- `metagration.check_script_trigger()` 是一个扩展函数，返回 `TRIGGER`。
- `metagration.current_revision()` 是一个扩展函数，返回 `bigint`。
- `metagration.export(replace_scripts boolean=false, transactional boolean=false, run_migrations boolean=false)` 是一个扩展函数，返回 `text`。
- `metagration.new_script(up_script text, down_script text=null, up_declare text=null, down_declare text=null, args jsonb='{}', use_schema text='metagration_scripts', comment text=null)` 是一个扩展函数，返回 `bigint`。
- `metagration.next_revision(from_revision bigint=null)` 是一个扩展函数，返回 `bigint`。
- `metagration.previous_revision(from_revision bigint=null)` 是一个扩展函数，返回 `bigint`。
- `metagration.run` 是一个扩展存储过程。
- `metagration.run_down` 是一个扩展存储过程。
- `metagration.run_up` 是一个扩展存储过程。
- `metagration.log` 是一个由扩展安装或管理的表。
- `metagration.script` 是一个由扩展安装或管理的表。
- `metagration` 是一个由扩展创建的模式。
- `metagration_scripts` 是一个由扩展创建的模式。

### 要求与注意事项

- 该目录记录了版本 `1.0.4`。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，与上游固定源进行验证。
