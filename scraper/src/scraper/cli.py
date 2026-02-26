import argparse
import json
from datetime import date, timedelta

from scraper.save import save_hours_to_file, save_lectures_to_file

from .fetch import fetch_classes, fetch_hours
from .parse import parse_hours, parse_lectures


def print_lectures(lectures):
    if not lectures:
        print("No lectures found.")
        return

    current_day = None
    for lec in lectures:
        day = lec.date
        if day != current_day:
            current_day = day
            weekday = date.fromisoformat(lec.date).strftime("%a")
            print(f"\n{day} {weekday}")
        start = lec.start[11:16]
        end = lec.end[11:16]
        lecturers = ", ".join(lec.lecturers) if lec.lecturers else "—"
        print(
            f"  {start} - {end}  {lec.subject_code:<8}  {lec.subject_name:<45}  {lec.room:<14}  {lecturers}"
        )


def main():
    today = date.today()
    monday = today - timedelta(days=today.weekday())
    sunday = monday + timedelta(days=6)

    parser = argparse.ArgumentParser(
        prog="scraper",
        description="Fetch and display your CIS calendar lectures.",
    )
    parser.add_argument(
        "--from",
        dest="start",
        default=monday.isoformat(),
        metavar="YYYY-MM-DD",
        help=f"Start date (default: {monday.isoformat()})",
    )
    parser.add_argument(
        "--to",
        dest="end",
        default=sunday.isoformat(),
        metavar="YYYY-MM-DD",
        help=f"End date (default: {sunday.isoformat()})",
    )
    parser.add_argument("--hours", action="store_true", help="parse hours")
    parser.add_argument("-c", "--classes", action="store_true", help="parse classes")

    args = parser.parse_args()
    if args.classes:
        print(f"Fetching lectures from {args.start} to {args.end} ...")
        data = fetch_classes(args.start, args.end)
        lectures = parse_lectures(data["data"])
        print_lectures(lectures)
        save_lectures_to_file(lectures)

    if args.hours:
        data = fetch_hours()
        hours = parse_hours(data["data"])
        save_hours_to_file(hours)


if __name__ == "__main__":
    main()
