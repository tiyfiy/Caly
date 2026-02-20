import argparse
from datetime import date, timedelta

from .fetch import fetch_page
from .parse import parse_lectures


def print_lectures(lectures):
    if not lectures:
        print("No lectures found.")
        return

    current_day = None
    for lec in lectures:
        day = lec.date
        if day != current_day:
            current_day = day
            weekday = lec.start.strftime("%a")
            print(f"\n{day} {weekday}")
        start = lec.start.strftime("%H:%M")
        end = lec.end.strftime("%H:%M")
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

    args = parser.parse_args()

    print(f"Fetching lectures from {args.start} to {args.end} ...")
    data = fetch_page(args.start, args.end)
    lectures = parse_lectures(data["data"])
    print_lectures(lectures)


if __name__ == "__main__":
    main()
