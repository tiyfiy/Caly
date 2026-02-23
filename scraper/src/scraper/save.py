import json
from dataclasses import asdict
from pathlib import Path

from scraper import Lecture


def save_to_file(data: list[Lecture]):
    lecture_dicts = [asdict(lec) for lec in data]
    out_path = Path("../../../lectures.json")
    with out_path.open("w", encoding="utf-8") as f:
        json.dump(lecture_dicts, f, indent=2, ensure_ascii=False)
