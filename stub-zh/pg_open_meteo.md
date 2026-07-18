## 用法

来源：

- [已复核 commit 的 Rust 实现](https://github.com/earth-metabolome-initiative/pg-open-meteo/blob/441c7d17d74809f02d721e79307cca322b9d2d94/src/lib.rs)
- [生成的 0.0.0 版本 SQL](https://github.com/earth-metabolome-initiative/pg-open-meteo/blob/441c7d17d74809f02d721e79307cca322b9d2d94/open_meteo_extension/usr/share/postgresql/17/extension/pg_open_meteo--0.0.0.sql)
- [示例天气迁移](https://github.com/earth-metabolome-initiative/pg-open-meteo/blob/441c7d17d74809f02d721e79307cca322b9d2d94/migrations/0003_weather_table/up.sql)
- [Cargo 特性矩阵](https://github.com/earth-metabolome-initiative/pg-open-meteo/blob/441c7d17d74809f02d721e79307cca322b9d2d94/Cargo.toml)

`pg_open_meteo` 是未完成的 pgrx 脚手架。尽管名称如此，0.0.0 并不会调用 Open-Meteo API、获取天气、调度更新或提供天气函数。生成的扩展 SQL 只包含三个演示函数。

```sql
CREATE EXTENSION pg_open_meteo;
SELECT hello_pg_open_meteo();
SELECT strictly_positive(1);
SELECT x_must_be_bigger_than_y(2, 1);
```

仓库迁移定义通用的 `position` 与 `sample_weather_data` 表，但这些迁移与生成的扩展安装 SQL 相互独立。不要假定 `CREATE EXTENSION` 会创建这些表，也不要在未审查其非限定对象名时直接运行到应用 schema。

仓库没有 README、公开 API 契约、发行版、明确许可证或可用的天气集成。只能把它当作示例代码；真实天气摄取应选择受维护的客户端或 FDW。
