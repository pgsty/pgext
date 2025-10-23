#!/usr/bin/env python3
"""
Generate language-based extension listings for Hugo content.

Reads pgext.extension from postgres:///vonng (overridable) and writes:
  - content/list/lang.md
  - content/list/lang.zh.md
"""

import argparse
from collections import defaultdict
from dataclasses import dataclass
from pathlib import Path
from typing import Dict, List, Optional

import psycopg2

LANGUAGE_NAV_ORDER = ["C", "C++", "Rust", "Java", "Python", "SQL", "Data"]

LANGUAGE_DESCRIPTIONS_EN: Dict[str, str] = {
    "C": "The traditional PostgreSQL extension language",
    "C++": "Extensions leveraging C++ features and libraries",
    "Rust": "Extensions written in Rust with the pgrx framework",
    "Python": "Extensions written in Python",
    "SQL": "Pure SQL extensions and functions",
    "Java": "Extensions running on JVM",
    "Data": "Data-only extensions",
}

LANGUAGE_DESCRIPTIONS_ZH: Dict[str, str] = {
    "C": "传统的 PostgreSQL 扩展开发语言",
    "C++": "使用 C++ 特性和库的扩展",
    "Rust": "使用 pgrx 框架用 Rust 编写的扩展",
    "Python": "使用 Python 编写的扩展",
    "SQL": "纯 SQL 扩展和函数",
    "Java": "在 JVM 上运行的扩展",
    "Data": "仅包含数据的扩展",
}


@dataclass
class Extension:
    """Subset of extension metadata used for language listings."""

    id: int
    name: str
    pkg: str
    lang: Optional[str]
    en_desc: Optional[str]
    zh_desc: Optional[str]
    lead: bool


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Render language-based extension lists.")
    parser.add_argument("--dsn", default="postgres:///vonng", help="PostgreSQL connection string")
    parser.add_argument(
        "--output",
        default="content/list",
        help="Output directory for generated Markdown files",
    )
    return parser.parse_args()


def fetch_extensions(dsn: str) -> List[Extension]:
    query = """
            SELECT id, name, pkg, lang, en_desc, zh_desc, lead
            FROM pgext.extension
            ORDER BY name \
            """
    extensions: List[Extension] = []
    with psycopg2.connect(dsn) as conn, conn.cursor() as cur:
        cur.execute(query)
        for row in cur.fetchall():
            extensions.append(
                Extension(
                    id=row[0],
                    name=row[1],
                    pkg=row[2],
                    lang=row[3],
                    en_desc=row[4],
                    zh_desc=row[5],
                    lead=row[6],
                )
            )
    return extensions


def build_leading_map(extensions: List[Extension]) -> Dict[str, str]:
    mapping: Dict[str, str] = {}
    for ext in extensions:
        if ext.lead and ext.pkg:
            mapping[ext.pkg] = ext.name
    for ext in extensions:
        mapping.setdefault(ext.pkg, ext.name)
    return mapping


def group_by_language(extensions: List[Extension]) -> Dict[str, List[Extension]]:
    grouped: Dict[str, List[Extension]] = defaultdict(list)
    for ext in extensions:
        if ext.lang:
            grouped[ext.lang].append(ext)
    for items in grouped.values():
        items.sort(key=lambda ext: ext.id)
    return grouped


def language_shortcode(language: str) -> str:
    return f'{{{{< language "{language}" >}}}}'


def badge_shortcode(language: str, count: int, is_zh: bool) -> str:
    if is_zh:
        return f'{{{{< badge content="{count} 个扩展" color="gray" >}}}}'
    icon = ' icon="cube"' if language == "C" else ""
    return f'{{{{< badge content="{count} Extensions" color="gray"{icon} >}}}}'


def sanitize(text: Optional[str]) -> str:
    if not text:
        return ""
    normalized = text.replace("|", "\\|").replace("\n", " ")
    return " ".join(normalized.split())


def build_navigation(languages: List[str]) -> str:
    preferred = [lang for lang in LANGUAGE_NAV_ORDER if lang in languages]
    remaining = sorted(lang for lang in languages if lang not in LANGUAGE_NAV_ORDER)
    return " ".join(language_shortcode(lang) for lang in preferred + remaining)


def render_summary(
        languages: List[str], counts: Dict[str, int], is_zh: bool
) -> List[str]:
    headers = ("语言", "数量", "描述") if is_zh else ("Language", "Count", "Description")
    descriptions = LANGUAGE_DESCRIPTIONS_ZH if is_zh else LANGUAGE_DESCRIPTIONS_EN
    lines = [
        f"| {headers[0]} | {headers[1]} | {headers[2]} |",
        "|:-------:|:-----:|:------------|",
    ]
    for language in languages:
        desc = descriptions.get(
            language,
            (f"使用 {language} 编写的扩展" if is_zh else f"Extensions written in {language}"),
        )
        lines.append(
            f"| {language_shortcode(language)} | {counts[language]} | {desc} |"
        )
    return lines


def render_table(
        language: str,
        extensions: List[Extension],
        leading_map: Dict[str, str],
        is_zh: bool,
) -> List[str]:
    headers = ("ID", "扩展", "描述") if is_zh else ("ID", "Extension", "Description")
    lines = [
        f"| {headers[0]} | {headers[1]} | {headers[2]} |",
        "|:---:|:---|:---|",
    ]
    prefix = "/zh/e/" if is_zh else "/e/"
    for ext in extensions:
        # Use alias shortcode: {{< alias "ext" "pkg" >}}
        if ext.name == ext.pkg:
            # Extension name equals package name
            ext_cell = f'{{{{< alias "{ext.name}" >}}}}'
        else:
            # Extension name differs from package name
            ext_cell = f'{{{{< alias "{ext.name}" "{ext.pkg}" >}}}}'

        if is_zh:
            desc = sanitize(ext.zh_desc) or sanitize(ext.en_desc) or "暂无描述"
        else:
            desc = sanitize(ext.en_desc) or "No description"
        lines.append(f"| {ext.id} | {ext_cell} | {desc} |")
    return lines


def render_sections(
        languages: List[str],
        grouped: Dict[str, List[Extension]],
        leading_map: Dict[str, str],
        is_zh: bool,
) -> List[str]:
    descriptions = LANGUAGE_DESCRIPTIONS_ZH if is_zh else LANGUAGE_DESCRIPTIONS_EN
    sections: List[str] = []
    for language in languages:
        count = len(grouped[language])
        desc = descriptions.get(
            language,
            (f"使用 {language} 编写的扩展" if is_zh else f"Extensions written in {language}"),
        )
        block: List[str] = [
            f"## {language}",
            "",
            f"{language_shortcode(language)} {badge_shortcode(language, count, is_zh)}",
            "",
            "",
            desc,
            "",
        ]
        block.extend(render_table(language, grouped[language], leading_map, is_zh))
        block.append("")
        sections.extend(block)
    return sections


def render_page(
        languages: List[str],
        grouped: Dict[str, List[Extension]],
        leading_map: Dict[str, str],
        is_zh: bool,
) -> str:
    counts = {language: len(grouped[language]) for language in languages}
    nav = build_navigation(languages)
    front_matter = (
        ["---", "title: 按语言", "description: 按实现语言组织的 PostgreSQL 扩展", "excludeSearch: true", "weight: 200",
         "---", ] if is_zh
        else ["---", "title: By Language", "description: PostgreSQL extensions organized by implementation language",
              "excludeSearch: true", "weight: 200", "---", ]
    )
    heading = "## 概览" if is_zh else "## Summary"
    lines: List[str] = []
    lines.extend(front_matter)
    lines.append("")
    lines.append(nav)
    lines.append("")
    lines.append(heading)
    lines.append("")
    lines.extend(render_summary(languages, counts, is_zh))
    lines.append("")
    lines.append("")
    lines.extend(render_sections(languages, grouped, leading_map, is_zh))
    return "\n".join(lines).rstrip() + "\n"


def main() -> None:
    args = parse_args()
    extensions = fetch_extensions(args.dsn)
    leading_map = build_leading_map(extensions)
    grouped = group_by_language(extensions)
    if not grouped:
        raise SystemExit("No extensions with language metadata found.")

    languages = [
        lang for lang, _ in sorted(
            grouped.items(),
            key=lambda item: (-len(item[1]), item[0].lower()),
        )
    ]

    output_dir = Path(args.output)
    output_dir.mkdir(parents=True, exist_ok=True)

    (output_dir / "lang.md").write_text(
        render_page(languages, grouped, leading_map, is_zh=False),
        encoding="utf-8",
    )
    (output_dir / "lang.zh.md").write_text(
        render_page(languages, grouped, leading_map, is_zh=True),
        encoding="utf-8",
    )


if __name__ == "__main__":
    main()
