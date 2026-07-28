## Usage

Sources:

- [Official upstream README](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/types/README.md)
- [Official extension control file (launchql-ext-types.control)](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/types/launchql-ext-types.control)
- [Official extension SQL (launchql-ext-types--0.4.5.sql)](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/types/sql/launchql-ext-types--0.4.5.sql)

`launchql-ext-types` — PostgreSQL extension providing custom domain types with built-in validation. This extension includes a collection of commonly needed data types with validation constraints, making it easy to enforce data integrity directly at the database level. Use it when application data needs this type, domain, or its operators. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION "launchql-ext-types";

-- Create a table using the custom domain types
CREATE TABLE users (
  id serial PRIMARY KEY,
  email email NOT NULL,
  website url,
  profile_image image,
  origin origin
);

-- Insert data with automatic validation
INSERT INTO users (email, website, profile_image, origin)
VALUES (
  'user@example.com',
  'https://example.com',
  '{"url": "https://example.com/profile.jpg", "mime": "image/jpeg"}',
  'https://example.com'
);

-- Invalid data will be rejected automatically
INSERT INTO users (email) VALUES ('not-an-email'); -- Fails validation
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `attachment` is an extension-defined domain.
- `email` is an extension-defined domain.
- `hostname` is an extension-defined domain.
- `image` is an extension-defined domain.
- `multiple_select` is an extension-defined domain.
- `origin` is an extension-defined domain.
- `single_select` is an extension-defined domain.
- `upload` is an extension-defined domain.
- `url` is an extension-defined domain.

### Requirements and Caveats

- The reviewed control file declares default version `0.4.5`.
- Install the confirmed extension dependencies first: `plpgsql`, `citext`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
