"""Mkdocs Code Reference Generator."""

import json
from pathlib import Path

import mkdocs_gen_files

src_root = Path(".")

for path in src_root.rglob("project.json"):
    project_path = str(path.parent)
    with open(path, "r") as project_file:
        project = json.loads(project_file.read())
        custom_patterns = [
            # OpenAPI Specs
            f"{project_path}/docs/openapi/*.json",
            f"{project_path}/docs/openapi/*.yml",
            f"{project_path}/docs/openapi/*.md",

            # Markdowns
            f"{project_path}/*.md",

            # Images
            f"{project_path}/docs/**/*.jpg",
            f"{project_path}/docs/**/*.png",
        ]

        for pattern in custom_patterns:
            for path in src_root.glob(pattern):
                doc_path = Path("reference", path.relative_to(src_root))
                with mkdocs_gen_files.open(doc_path, "wb") as f:
                    f.write(path.read_bytes())

                if not path.suffix.endswith(".jpg") and not path.suffix.endswith(".png"):
                    mkdocs_gen_files.set_edit_path(doc_path, f"../{path}")

        for path in src_root.glob(f"{project_path}/**/*[!__init__].py"):
            if "tests" in path.parts or ".venv" in path.parts:
                continue
            doc_path = Path("reference", path.relative_to(src_root)).with_suffix(".md")
            doc_normalized_path = str(doc_path).replace(project_path, f"{project_path}/code_reference")
            with mkdocs_gen_files.open(doc_normalized_path, "w") as f:
                ident = ".".join(path.with_suffix("").parts)
                f.write(f"::: {ident}")

            mkdocs_gen_files.set_edit_path(doc_normalized_path, f"../{path}")
