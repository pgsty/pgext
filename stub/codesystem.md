## Usage

Sources:

- [README.md](https://github.com/John-Cameron-Burns/codesystem/blob/6c1a7bcea5e2e808d47c04bb68d96eac0a101cbe/README.md)
- [codesystem--0.1.0.sql](https://github.com/John-Cameron-Burns/codesystem/blob/6c1a7bcea5e2e808d47c04bb68d96eac0a101cbe/codesystem--0.1.0.sql)
- [sql/define.sql](https://github.com/John-Cameron-Burns/codesystem/blob/6c1a7bcea5e2e808d47c04bb68d96eac0a101cbe/sql/define.sql)
- [codesystem.control](https://github.com/John-Cameron-Burns/codesystem/blob/6c1a7bcea5e2e808d47c04bb68d96eac0a101cbe/codesystem.control)

`codesystem` stores named classification systems and typed code values. Its 0.1.0 prototype implements definitions, validation, hierarchy traversal and revisions; derived standards and crosswalks described in the design are not implemented.

### Core Workflow

```sql
CREATE EXTENSION codesystem;
CREATE TABLE classification_source (code text, level text, parent text, label text);
INSERT INTO classification_source VALUES ('A', 'section', '', 'Agriculture');
SELECT define_code_system(
  'toy',
  '[{"name":"section","width":1,"charset":"[A-Z]","link":"none","expected":1}]',
  'SELECT * FROM classification_source', 'Toy classification', 'Example');
SELECT name, current_revision, ncodes FROM code_systems;
SELECT code_label('toy:A'::code);
```

### Operational Boundaries

Upstream tests this revision on PostgreSQL 17. Install the native extension with administrator privileges. The input query for `define_code_system` must supply code, level, parent and label columns; definitions reject malformed levels, duplicate codes and broken hierarchy links. Literals include a system prefix, while `code_from` converts external text. `code_level`, `code_parent`, `code_ancestor`, `code_children` and `code_label` expose the hierarchy. Equal text in different revisions denotes distinct values. Registry structural edits are protected. Custom-format restores require the upstream ordering helper so registry data exists before typed table values are restored; ordinary restore ordering may fail.
