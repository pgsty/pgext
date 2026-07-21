## 用法

来源：

- [PostgreSQL Anonymizer 3.1.3 README](https://gitlab.com/dalibo/postgresql_anonymizer/-/blob/3.1.3/README.md)
- [Masking functions](https://gitlab.com/dalibo/postgresql_anonymizer/-/blob/3.1.3/docs/masking_functions.md)
- [3.1.3 changelog](https://gitlab.com/dalibo/postgresql_anonymizer/-/blob/3.1.3/CHANGELOG.md)
- [Official documentation](https://postgresql-anonymizer.readthedocs.io/en/stable/)

`anon` 是 PostgreSQL Anonymizer，它应用声明式的遮掩规则以实现受保护查询的访问，并生成匿名化数据集。同时提供伪名化和随机响应辅助功能。在必须保持现实数据有用性而不暴露原始敏感值时使用它；将遮掩策略、角色授权以及未遮掩数据库的访问视为安全边界的一部分。

### 核心工作流程

在目标数据库会话中加载 `anon`，安装扩展并启用透明动态遮掩。新连接会自动获取数据库级别的设置。

```sql
ALTER DATABASE app SET session_preload_libraries = 'anon';

\connect app
CREATE EXTENSION anon;
ALTER DATABASE app SET anon.transparent_dynamic_masking = true;
```

将登录标记为被遮掩，并将遮掩规则附加到敏感列上：

```sql
CREATE ROLE reporting LOGIN;
SECURITY LABEL FOR anon ON ROLE reporting IS 'MASKED';
GRANT pg_read_all_data TO reporting;

SECURITY LABEL FOR anon ON COLUMN customer.last_name
  IS 'MASKED WITH FUNCTION anon.dummy_last_name()';
SECURITY LABEL FOR anon ON COLUMN customer.phone
  IS 'MASKED WITH FUNCTION anon.partial(phone, 2, $$******$$, 2)';
```

以 `reporting` 身份执行的查询会看到变换后的值。特权用户仍然可以看到原始值，因此不要授予被遮掩的角色绕过策略的路径。

### 遮掩策略

- **动态遮掩** 对标记为 `MASKED` 的角色转换结果而无需重写表。
- **静态遮掩** 永久性地重写选定的数据，并适用于一次性开发或测试副本。
- **匿名导出和复制体** 生成经过净化的导出或下游副本。
- **遮掩视图和数据包装器** 展示一个故意减少或转换后的投影。
- **伪名化** 在必须保持一致性的连接或重复值时使用确定性变换。

### 关键对象

- `anon.dummy_*`、`anon.random_*` 和 `anon.partial(...)` 生成或部分遮掩值。
- `anon.hash(text)` 和 `anon.digest(text, text, text)` 提供确定性转换。在 3.1.2 版本中，它们被标记为 `RESTRICTED` 以限制暴力破解暴露。
- `anon.ldp_grrm(value, epsilon, max_v)` 和 `anon.ldp_grrm_pttt(value, truth_probability, max_v)` 实现了局部差分隐私的通用随机响应。
- `anon.ldp_truth_probability(...)` 和 `anon.ldp_lie_probability(...)` 用于检查随机响应概率。
- 角色和列上的安全标签定义谁被遮掩以及每个值如何转换。

### 运营注意事项

`anon` 是超级用户安装且不可重定位的。使用与预期消费者相同的授权和连接路径测试每项策略。随机化不是自动确定性的；当需要稳定相等时，请使用确认的伪名化函数。静态匿名化是破坏性的，因此在副本上运行它并在之后验证约束条件和应用程序行为。

版本 3.1.3 重新运行缺失的 ARM 构建并更改发布元数据，没有新的 SQL 工作流。自 3.1.1 版本以来的主要差异在于对 `anon.hash` 和 `anon.digest` 的 3.1.2 安全强化；使用这些函数的部署应升级而不是依赖旧标签。
