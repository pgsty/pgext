## 用法

来源：

- [锁定提交的上游 README](https://github.com/spa5k/uids-postgres/blob/4bfbf05a6a9a0771dc66000fb3bb19974e439207/README.md)
- [锁定提交的 Cargo 清单](https://github.com/spa5k/uids-postgres/blob/4bfbf05a6a9a0771dc66000fb3bb19974e439207/Cargo.toml)
- [锁定提交的 Rust 函数源码](https://github.com/spa5k/uids-postgres/tree/4bfbf05a6a9a0771dc66000fb3bb19974e439207/src)
- [锁定提交的扩展控制文件](https://github.com/spa5k/uids-postgres/blob/4bfbf05a6a9a0771dc66000fb3bb19974e439207/uids.control)

uids 0.0.1 是一个 pgrx 扩展，用于生成 UUIDv6、UUIDv7、NanoID、KSUID、ULID、Timeflake、PushID、CUID2 和 TypeID 值。大多数函数返回文本；某些函数返回 bytea 或 PostgreSQL uuid。每个系列的长度、顺序、时间戳暴露、熵和互操作语义都不同。

### 生成代表性标识符

```sql
CREATE EXTENSION uids;

SELECT generate_uuidv6_uuid();
SELECT generate_uuidv7();
SELECT generate_ulid();
SELECT generate_ksuid();
SELECT generate_nanoid_length_c(10, '0123456789abcdef');
SELECT generate_typeid('user');
SELECT check_typeid('user', generate_typeid('user'));
SELECT generate_cuid2();
SELECT generate_timeflake_uuid();
SELECT generate_pushid();
```

generate_uuidv6_uuid 返回原生 uuid 类型，而 generate_uuidv7 返回文本。应为每列选择并强制一种 SQL 类型和格式；以后更换生成器系列属于数据迁移，不是透明实现细节。

### 验证、排序与构建限制

generate_uuidv7_from_string 和 parse_uuidv7 能解析通用 UUID 语法，但不验证版本是否为 7；上游测试本身就使用了版本 4 UUID。check_typeid 通过 expect 解析，所以格式异常的输入会报错，而不是返回 false。类似的解析和随机生成路径使用 unwrap 或 panic。在处理不可信输入时，应将验证函数视为会抛出异常，并进行包装。

NanoID 长度作为有符号整数接收，然后转换为无符号机器字大小。负长度或不合理长度可能变成巨大内存分配请求。自定义字母表也可能无效。调用任何可配置函数前，应强制小的正长度上限、非空且不重复的字母表，以及输入字节上限。

时间可排序标识符可能泄露创建时间，其文本顺序取决于排序规则。如果顺序属于需求，应测试并发生成，并在字节序排序规则下存储可排序文本，或使用具有文档化比较行为的原生类型。大多数系列的唯一性属于概率保证；不得自动将这些值视为身份验证密钥或授权证明。

控制文件要求超级用户，且扩展不可重定位。锁定 Cargo 清单使用 pgrx 0.11.4，默认为 PostgreSQL 16，并包含 PostgreSQL 11 至 16 的功能标志；原生产物必须与精确 PostgreSQL 主版本和平台匹配。GitHub v1 标签构建的扩展版本仍为 0.0.1，仓库也没有声明许可证。采用前应验证已安装 extversion、函数返回类型、随机源失败行为、排序规则、转储/恢复和冲突处理。
