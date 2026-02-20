import requests

URL = "https://cis.technikum-wien.at/cis.php/api/frontend/v1/lvPlan/eventsPersonal"
HEADERS = {
    "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:147.0) Gecko/20100101 Firefox/147.0",
    "Accept": "application/json, text/plain, */*",
    "Authorization": "Basic ZWwyNXgwMDcgOktyb21waXJjZWsuMTI=",
}


def fetch_page(start_date: str, end_date: str) -> dict:
    payload = {
        "start_date": start_date,
        "end_date": end_date,
    }
    resp = requests.post(URL, json=payload, headers=HEADERS, timeout=10)
    resp.raise_for_status()
    return resp.json()
