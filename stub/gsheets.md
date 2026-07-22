## Usage

Sources:

- [Official README](https://github.com/MuhammadTahaNaveed/pg-gsheets/blob/f223b892ff2a0dcd98c5f2ac0ca8d748a3b5bb28/README.md)
- [Extension control file](https://github.com/MuhammadTahaNaveed/pg-gsheets/blob/f223b892ff2a0dcd98c5f2ac0ca8d748a3b5bb28/gsheets.control)

`gsheets` lets SQL sessions read from and write to Google Sheets. It is useful for small integration, import, and export workflows where the database server can reach Google APIs and the session can supply an OAuth access token.

### Core Workflow

Create the extension, complete the upstream authentication flow, and set the resulting token for the current session:

```sql
CREATE EXTENSION gsheets;

SELECT gsheets_auth();
SET gsheets.access_token = 'your_access_token';
```

`read_sheet` is polymorphic: the caller supplies the returned column definition. The first argument accepts a spreadsheet ID or URL.

```sql
SELECT *
FROM read_sheet(
    '<spreadsheet_id_or_url>',
    sheet_name => 'Sheet1',
    header => true
) AS sheet(name text, age integer);
```

Write rows with `write_sheet`. Supplying no `spreadsheet_id` asks the extension to create a spreadsheet; otherwise the JSON options select an existing spreadsheet and sheet.

```sql
SELECT write_sheet(
    person.*,
    '{"spreadsheet_id":"<spreadsheet_id>","sheet_name":"Sheet2","header":["name","age"]}'::jsonb
)
FROM person;
```

### User-Facing Objects

- `gsheets_auth()`: starts the OAuth authentication flow described by upstream.
- `gsheets.access_token`: session setting containing the Google access token.
- `read_sheet(spreadsheet_id/url text, sheet_name text DEFAULT 'Sheet1', header boolean DEFAULT true)`: reads rows; declare the expected record columns in the query.
- `write_sheet(data anyelement, options jsonb DEFAULT '{}')`: writes values or composite rows. Options include `spreadsheet_id`, `sheet_name`, and `header`.

### Operational Notes

The extension links to libcurl and performs external network requests from the PostgreSQL server process. Authentication requires an interactive URL/token flow, so plan explicitly for headless servers and token renewal. Keep `gsheets.access_token` out of logs and persistent role/database settings, and grant use of the extension only to roles allowed to access the target sheets. No preload setting is documented. Test API errors, rate limits, partial writes, and retry behavior before using `write_sheet` in automated jobs.
