## 用法

来源：

- [examples/uuid_v7/README.md](https://github.com/aws/pg_tle/blob/2f4b7b34ac3e65a4c4ec358765839d5f24a910bd/examples/uuid_v7/README.md)
- [examples/uuid_v7/uuid_v7--1.0.sql](https://github.com/aws/pg_tle/blob/2f4b7b34ac3e65a4c4ec358765839d5f24a910bd/examples/uuid_v7/uuid_v7--1.0.sql)
- [examples/uuid_v7/Makefile](https://github.com/aws/pg_tle/blob/2f4b7b34ac3e65a4c4ec358765839d5f24a910bd/examples/uuid_v7/Makefile)
- [examples/uuid_v7/uuid_v7.control](https://github.com/aws/pg_tle/blob/2f4b7b34ac3e65a4c4ec358765839d5f24a910bd/examples/uuid_v7/uuid_v7.control)

`uuid_v7` 是 AWS Trusted Language Extensions 中用 PL/Rust 实现的示例，可生成带时间戳的 UUID、按指定时间创建 UUID，以及提取其中的毫秒时间戳。

### 核心用法

```sql
CREATE EXTENSION plrust;
CREATE EXTENSION uuid_v7;
SELECT generate_uuid_v7();
SELECT timestamptz_to_uuid_v7('2026-10-06 00:00:00+00'::timestamptz);
SELECT uuid_v7_to_timestamptz(generate_uuid_v7());
```

### 运行边界

先安装 `pg_tle` 与 PL/Rust 1.2.0 或更高版本，再按上游 TLE 安装流程注册示例。扩展出现在可用列表后，才在数据库中创建它。控制文件的 SQL 依赖为 `plrust`，并不是另一个原生 pgrx 共享库。

`generate_uuid_v7` 使用当前时间，`timestamptz_to_uuid_v7` 嵌入指定时间，`uuid_v7_to_timestamptz` 解码前部时间戳位，但不能据此确认任意 UUID 都是合法的第 7 版值。时间精度为毫秒，随机后缀不能保证同一毫秒内严格有序。示例基于 UUID v7 草案开发，将其用作标准符合性边界前应核对语义。
