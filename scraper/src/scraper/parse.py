import json
from dataclasses import dataclass


@dataclass
class Lecture:
    subject_code: str
    subject_name: str
    lecturers: list[str]
    date: str
    start: str
    end: str
    room: str


@dataclass
class Hour:
    slot: int
    start: str
    end: str


def parse_lectures(data: list[dict]) -> list[Lecture]:
    groups: dict[tuple, list[dict]] = {}
    for item in data:
        if item["type"] != "lehreinheit":
            continue
        key = (item["lehrveranstaltung_id"], item["datum"])
        groups.setdefault(key, []).append(item)

    lectures = []
    for slots in groups.values():
        first = min(slots, key=lambda x: x["isostart"])
        last = max(slots, key=lambda x: x["isoend"])
        lecturers = [f"{l['vorname']} {l['nachname']}" for l in first.get("lektor", [])]
        lectures.append(
            Lecture(
                subject_code=first["lehrfach"],
                subject_name=first["lehrfach_bez"],
                date=first["datum"],
                start=first["isostart"],
                end=last["isoend"],
                room=first["ort_kurzbz"],
                lecturers=lecturers,
            )
        )
    return sorted(lectures, key=lambda l: l.start)  # ISO strings sort correctly


def parse_hours(data: list[dict]) -> list[Hour]:
    hours = []
    for hour in data:
        hours.append(
            Hour(
                slot=hour["stunde"],
                start=hour["beginn"],
                end=hour["ende"],
            )
        )
    return hours
