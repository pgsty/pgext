


## 用法

> [pg_graphql: 为数据库内置 GraphQL 支持](https://github.com/supabase/pg_graphql)

`pg_graphql` 从现有的 SQL 模式反射生成 GraphQL 模式，无需额外的服务器或中间件即可直接在 PostgreSQL 内部执行 GraphQL 查询。

### 模式反射

表、外键和枚举类型会自动映射为 GraphQL 类型：

```sql
CREATE TABLE account (
    id serial PRIMARY KEY,
    email varchar(255) NOT NULL,
    created_at timestamp NOT NULL
);

CREATE TABLE blog (
    id serial PRIMARY KEY,
    owner_id integer NOT NULL REFERENCES account(id),
    name varchar(255) NOT NULL,
    description varchar(255)
);

CREATE TYPE blog_post_status AS ENUM ('PENDING', 'RELEASED');

CREATE TABLE blog_post (
    id uuid NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    blog_id integer NOT NULL REFERENCES blog(id),
    title varchar(255) NOT NULL,
    body varchar(10000),
    status blog_post_status NOT NULL,
    created_at timestamp NOT NULL
);
```

此模式会自动生成 GraphQL 类型（`Account`、`Blog`、`BlogPost`），关系由外键推导而来。

### 命名转换

启用 snake_case 到 camelCase/PascalCase 的自动转换：

```sql
COMMENT ON SCHEMA public IS e'@graphql({"inflect_names": true})';
```

### 查询

通过 `graphql.resolve` 函数执行 GraphQL 查询：

```sql
SELECT graphql.resolve($$
    {
      accountCollection(first: 1) {
        edges {
          node {
            id
            email
            blogCollection {
              edges {
                node {
                  name
                  blogPostCollection(filter: { status: { eq: RELEASED } }) {
                    edges {
                      node {
                        title
                      }
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
$$);
```

### 功能特性

- 表查询以可分页集合的形式出现在根 `Query` 类型上
- 外键关系自动创建嵌套查询字段
- 变更操作支持批量插入、更新和删除
- 内置过滤、排序和分页功能
- 遵守 PostgreSQL 行级安全（RLS）策略
