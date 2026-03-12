


## Usage

> [pg_graphql: Add in-database GraphQL support](https://github.com/supabase/pg_graphql)

`pg_graphql` reflects a GraphQL schema from your existing SQL schema, enabling GraphQL queries directly inside PostgreSQL without additional servers or middleware.

### Schema Reflection

Tables, foreign keys, and enums are automatically mapped to GraphQL types:

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

This schema automatically generates GraphQL types (`Account`, `Blog`, `BlogPost`) with relationships derived from foreign keys.

### Name Inflection

Enable automatic snake_case to camelCase/PascalCase conversion:

```sql
COMMENT ON SCHEMA public IS e'@graphql({"inflect_names": true})';
```

### Querying

Execute a GraphQL query via the `graphql.resolve` function:

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

### Features

- Table queries appear as pageable collections on the root `Query` type
- Foreign key relationships create nested query fields automatically
- Mutations support bulk insert, update, and delete
- Filtering, ordering, and pagination are built in
- PostgreSQL Row-Level Security (RLS) policies are respected
