import requests

URL = "https://cis.technikum-wien.at/cis.php/api/frontend/v1/lvPlan/eventsPersonal"
HEADERS = {
    "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:147.0) Gecko/20100101 Firefox/147.0",
    "Accept": "application/json, text/plain, */*",
    "Authorization": "Basic ZWwyNXgwMDcgOktyb21waXJjZWsuMTI=",
}

payload = {
    "end_date": "2026-02-22",
    "start_date": "2026-02-16",
}


def fetch_page(page: int):
    resp = requests.post(URL, json=payload, headers=HEADERS, timeout=10)
    print(resp.status_code, resp.text)
    resp.raise_for_status()
    return resp.json()


def main():
    data = fetch_page(1)

    print(data)


if __name__ == "__main__":
    main()
