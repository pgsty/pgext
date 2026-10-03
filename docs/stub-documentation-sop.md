# Extension stub documentation SOP

`stub/<extension>.md` and `stub-zh/<extension>.md` are the editable English and
Chinese usage sources. Site generators copy these files into extension pages.
Fix errors in these sources; a correction made only in a generated site page
will be overwritten on the next catalog refresh.

## Scope and evidence

- Establish the canonical extension name and an explicit target list. Review
  both language files when the request covers a bilingual pair.
- Read existing content first. Preserve identifiers, SQL, configuration values,
  compatibility boundaries, and unrelated user changes.
- For technical additions or changes, verify official documentation, README,
  control/SQL files, and release notes for the documented version. Record source
  URLs and version differences. A formatting repair does not establish that
  historical technical claims have been rechecked.
- Keep a batch ledger under `tmp/` with extension/package, both source paths,
  owner, action, before/after hashes, validation results, and status. Keep a
  package family in one batch. Use `FORMAT_FIXED` or `NO_FORMAT_CHANGE` for
  formatting-only work; reserve `DONE` for the complete requested review.

## Document contract

New or materially refreshed documents use one H2: `## Usage` in English and
`## 用法` in Chinese. Other headings are H3 or deeper. No front matter or H1.
The first content after the H2 is `Sources:` or `来源：`, followed by linked
official sources. For example:

```markdown
## Usage

Sources:

- [Official documentation](https://example.org/docs)

Explain the extension's core workflow and operational boundaries.

### Configuration

Describe only settings confirmed by the official sources.
```

Keep English and Chinese source URLs, code examples, identifiers, and heading
structure aligned. Translate prose and code comments where appropriate without
changing the executable statements. A focused formatting repair may preserve a
legacy document's structure; record that distinction instead of claiming full
compliance with the current document contract.

## Markdown requirements

- Every opening code fence has an explicit language. Use `sql` for SQL examples,
  `bash` for shell commands, `ini` for PostgreSQL configuration, and `json` or
  `yaml` for those formats. Use `text` for result tables, query plans, terminal
  transcripts, CLI help, signatures, and other literal text. Review the block's
  content rather than assigning `sql` to every example.
- Closing fences use the same marker as the opening fence and are at least as
  long. The closing marker has no language label. Longer outer fences can quote
  Markdown containing shorter fences.
- Surround headings with blank lines and use one space after the heading marker.
- Remove stray trailing spaces and tabs from prose. Preserve an intentional
  two-space Markdown hard break and significant whitespace inside code blocks.
- Keep code bodies intact during a formatting repair. Review fence labels in
  both languages, including blocks that already have a label.
- Finish files with a newline. Preserve source links, anchors, identifiers,
  commands, and examples unless the requested technical review justifies a change.

## Validation before completion

1. Inspect the source diff and compare code bodies and technical identifiers.
2. Run the shared, database-free Markdown check on both languages:

   ```bash
   go run . gen lint stub/example.md stub-zh/example.md
   make check-stubs
   ```

3. For new or materially refreshed documentation, also run the existing
   `bin/stub_phaseb.py` pair lint for the bound manifest, or equivalent focused
   checks of the document contract, source URLs, identifiers, and bilingual code
   blocks. Markdown lint and technical/bilingual review are separate checks.
4. Run `git diff --check` and relevant tests. Changes to Markdown validation or
   generators require `go test ./...`.
5. Check the batch ledger against the target list. Report reviewed, unchanged,
   formatting-only, and blocked results accurately.

`pgext gen lint` checks both source directories by default. Go CI runs it before
the build. `gen io page/all` and `gen cc page/all` preflight all selected existing
stubs before writing their generated outputs. Individual IO, CC, and catalog
page generators repeat validation when reading the stub and return source paths
and line numbers on failure. They reject an invalid stub without overwriting the
existing page. Correct the source rather than bypassing this failure.

## Database and generated-site boundaries

- A stub repair alone does not authorize loading `pgext.doc`, regenerating other
  repositories, committing, pushing, or publishing. Follow the current user's
  scope and the intake workflow in `AGENTS.md`.
- When a local documentation load is authorized, run loader tests and a complete
  transactional dry-run first. After loading, verify that stored bilingual
  bodies match the source files. This is separate from source lint.
- When site generation is authorized, lint sources first, generate to a named
  staging directory using `--output`, and run the destination site's own strict
  source/rendered/link checks before replacing or publishing its generated files.
  The shared source check catches the recurring fence/spacing errors; it does
  not replace each site's complete Hugo and rendered-page validation.
- Verify the provider deployment and public pages after publication. A local
  build or a successful upload alone does not establish public delivery.
