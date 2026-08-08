## 用法

来源：

- [带标签的 README](https://codeberg.org/gregburd/pg_infer/src/tag/v0.1.1-alpha/README)
- [扩展控制文件](https://codeberg.org/gregburd/pg_infer/src/tag/v0.1.1-alpha/pg_infer.control)
- [Cargo 清单](https://codeberg.org/gregburd/pg_infer/src/tag/v0.1.1-alpha/Cargo.toml)
- [扩展安装 SQL](https://codeberg.org/gregburd/pg_infer/src/tag/v0.1.1-alpha/sql/pg_infer--1.0.0.sql)
- [模型注册实现](https://codeberg.org/gregburd/pg_infer/src/tag/v0.1.1-alpha/src/model_mgmt.rs)
- [索引访问方法实现](https://codeberg.org/gregburd/pg_infer/src/tag/v0.1.1-alpha/src/am.rs)
- [官方回归 SQL](https://codeberg.org/gregburd/pg_infer/src/tag/v0.1.1-alpha/sql/pg_infer.sql)
- [Vindex 格式与模型数据](https://codeberg.org/gregburd/pg_infer/src/tag/v0.1.1-alpha/crates/infer-vindex/README.md)

`pg_infer` 是一个实验性 PostgreSQL 扩展，通过 SQL 暴露 Transformer 模型的特征、习得关联与相似度信号。它创建 `infer` 模式保存模型注册表，并提供 `infer` 索引访问方法，以便按模型语义对文本排序。它适合在 PostgreSQL 中检查外部准备的模型知识；扩展本身不包含模型或 vindex，其结果是随模型变化的信号，而非经过验证的事实。

### 版本与适用范围

仓库标签为 `v0.1.1-alpha`，但该标签内 README 的横幅仍写着 `v0.1.0-alpha`。同一标签中的 `pg_infer.control` 声明 `default_version = '1.0.0'`，`Cargo.toml` 声明 `1.0.0`，而 `sql/pg_infer--1.0.0.sql` 安装的 SQL 版本也是 `1.0.0`。应把 SQL 版本和项目成熟度视为两个独立事实：上游 README 仍将项目标为实验性，警告 SQL API 可能变化，并说明 vindex 格式尚未冻结。

带标签的 Cargo 清单只提供 `pg18` PostgreSQL 特性，并默认选择 `pg18`。因此，PostgreSQL 18 是该源码快照有证据支持的服务器目标；不要假定它兼容更早的大版本。

在每个需要使用它的数据库中安装扩展，然后检查已安装的 SQL 版本：

```sql
CREATE EXTENSION pg_infer;

SELECT extversion
FROM pg_extension
WHERE extname = 'pg_infer';
```

控制文件设置了 `superuser = true` 和 `relocatable = false`，因此必须由超级用户创建扩展，且安装后不能将其迁移到其他模式。

### 安装并注册 Vindex

`pg_infer` 在 `infer.models` 中保存注册元数据，但本地模型数据仍位于外部 vindex 目录。先把 `infer.data_directory` 设为允许访问的基础目录，再注册一个已经存在的 vindex，并选择在查询省略模型参数时使用的默认模型：

```sql
SET infer.data_directory = '/data';
SELECT infer_create_model('qwen05b', '/data/qwen-0.5b.vindex');
SET infer.default_model = 'qwen05b';
SELECT * FROM infer_models();
```

PostgreSQL 操作系统用户必须能够读取该目录。注册实现会先验证 vindex，再插入或更新 `infer.models`；本地绝对路径必须位于 `infer.data_directory` 之下，相对路径则基于该目录解析。`infer_drop_model` 会删除注册项并驱逐进程内缓存，但不会删除外部模型文件。

Vindex 是一个目录，其中包含为查询访问重新组织的模型配置、分词器数据、嵌入、门控向量和特征元数据。所需文件取决于提取级别：browse 数据支持 `walk`、`describe` 等模型检查操作，完整前向预测则需要 inference 级别的模型数据。应在 PostgreSQL 外部准备和管理这些数据，并在备份、恢复、复制及主机迁移流程中显式纳入它们。

### 查询模型知识

设置 `infer.default_model` 后，可以不带模型参数调用查询函数；也可以传入 `model => 'name'` 显式选择模型。以下示例来自带标签的 API；返回行与分数完全取决于已注册的 vindex：

```sql
SELECT * FROM walk('The capital of France is', top => 10);
SELECT * FROM describe('France');
SELECT similar_to('France', 'Paris');
SELECT implies('France', 'Paris');
```

- `walk` 跟踪各层最强的特征激活；`infer_explain_walk` 额外提供层带与标签细节。
- `describe` 返回模型对实体推断出的关系；`describe_layers` 保留逐层明细，`nearest_to` 则探测指定层。
- `similar_to` 返回相似度分数，`similar_to_many` 对候选数组评分，`implies` 检查模型提取知识中的方向性支持。
- `infer_show_layers`、`infer_show_features` 与 `infer_show_relations` 用于检查可用的模型元数据；`infer_diff` 比较两个已注册模型的特征元数据。
- `infer` 只有在扩展使用 Cargo 特性 `inference` 构建，且 vindex 含有所需 inference 数据时才能执行前向预测。未启用该特性的默认构建仍会暴露此函数，但调用时会返回错误。

### 创建语义索引

构建索引前必须先注册模型。带标签的访问方法接受一个文本列，并把所选模型名保存到元页中。使用距离操作符执行升序 top-N 查询：

```sql
CREATE TABLE documents (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title text NOT NULL
);

CREATE INDEX documents_title_infer_idx
    ON documents USING infer (title)
    WITH (model = 'qwen05b');

SELECT id, title
FROM documents
ORDER BY title <~> 'artificial intelligence'
LIMIT 5;
```

`<~>` 返回距离，因此值越小越相似。`<~` 返回底层相似度分数，值越大表示相似度越强；`@>` 使用默认模型暴露方向性蕴含。`infer_text_ops` 操作符类把 `<~>` 与 `infer` 访问方法连接起来，用于 `ORDER BY ... LIMIT` 计划。

### 重要对象

#### 注册表与配置

- `infer.models` 记录模型名、vindex 路径、维度、后端与注册时间。
- `infer.default_model` 选择隐式模型；`infer.data_directory` 限制本地模型路径。
- `infer.max_memory` 限制每个后端的 vindex 缓存；`infer.gate_threshold`、`infer.describe_top_k` 和 `infer.walk_embed_mode` 调整检查行为。

#### 查询函数

- 模型生命周期：`infer_create_model`、`infer_drop_model`、`infer_models`。
- 探索：`walk`、`describe`、`describe_layers`、`nearest_to`、`similar_to`、`similar_to_many`、`implies`。
- 内省：`infer_explain_walk`、`infer_show_layers`、`infer_show_features`、`infer_show_relations`、`infer_diff`。
- 可选前向预测：`infer`。

#### 操作符与访问方法

- 距离排序：`<~>`，配合 `USING infer` 和 `infer_text_ops`。
- 原始相似度过滤：`<~`。
- 方向性蕴含：`@>`。

### 依赖与注意事项

- 带标签的 Cargo 清单固定使用 `pgrx 0.17.0` 和 `rust-version = '1.80'`；带标签的 README 还要求从源码构建时使用 Rust nightly、PostgreSQL 18 或更高版本、OpenSSL 与 OpenBLAS。
- 控制文件和带标签的安装说明均未声明 `shared_preload_libraries` 设置。正常使用从 `CREATE EXTENSION pg_infer` 开始；除非部署的确切构建给出证据，否则不要自行增加预加载或重启要求。
- 注册表行和索引元页不能替代外部 vindex 目录。应确认故障切换或恢复后可能执行查询的每台服务器都具备模型数据、正确权限与有效路径。
- Alpha README 明确警告 SQL 可能发生破坏性变化、vindex 格式尚未冻结、部分计算路径依赖特定硬件，而且该快照尚无成熟的生产部署。应使用代表性模型与查询进行测试，限制内存使用，并把升级视为兼容性敏感操作。
- 关系、相似度和蕴含等输出反映的是所选模型与提取流水线。未经独立验证，不应把它们视为权威事实或用于影响重大的决策。
