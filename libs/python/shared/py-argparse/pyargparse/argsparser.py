import argparse

def new(description: str) -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(description=description)

    parser.add_argument(
        "--enable-debug-storage",
        help="Enable debug module",
        default=False,
        action="store_true"
    )

    parser.add_argument(
        "--debug-storage-dir",
        help="Sets the base directory for debug storage",
        dest="debug_storage_dir",
        default="/app/tests/debug/storage"
    )

    return parser
