from dataclasses import dataclass
from datetime import datetime


@dataclass
class Lecture:
    subject_code: str
    subject_name: str
    lecturers: list[str]
    date: str
    start: datetime
    end: datetime
    room: str


def parse_lectures(data: list[dict]) -> list[Lecture]:
    groups: dict[tuple, list[dict]] = {}
    for item in data:
        if item["type"] != "lehreinheit":
            continue
        key = (item["lehrveranstaltung_id"], item["datum"])
        groups.setdefault(key, []).append(item)

    lectures = []
    for slots in groups.values():
        slots.sort(key=lambda x: x["isostart"])
        first, last = slots[0], slots[-1]
        lecturers = [f"{l['vorname']} {l['nachname']}" for l in first.get("lektor", [])]
        lectures.append(
            Lecture(
                subject_code=first["lehrfach"],
                subject_name=first["lehrfach_bez"],
                date=first["datum"],
                start=datetime.fromisoformat(first["isostart"]),
                end=datetime.fromisoformat(last["isoend"]),
                room=first["ort_kurzbz"],
                lecturers=lecturers,
            )
        )
    return sorted(lectures, key=lambda l: l.start)
