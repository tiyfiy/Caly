API scrapper call:

```
POST
	https://cis.technikum-wien.at/cis.php/api/frontend/v1/lvPlan/eventsPersonal


```

response JSON:

```json
{
  "data": [
    {
      "type": "lehreinheit",
      "beginn": "08:00:00",
      "ende": "08:45:00",
      "datum": "2026-02-19",
      "topic": "DSML-ILV",
      "lektor": [
        {
          "mitarbeiter_uid": "ma0263",
          "vorname": "Patrick",
          "nachname": "Link",
          "kurzbz": "LinkPa"
        }
      ],
      "gruppe": [
        {
          "gruppe": " ",
          "verband": " ",
          "semester": "4",
          "studiengang_kz": "257",
          "kuerzel": "GRP_166146"
        }
      ],
      "ort_kurzbz": "EDV_F1.02",
      "lehreinheit_id": ["166146"],
      "titel": null,
      "lehrfach": "DSML",
      "lehrform": "ILV",
      "lehrfach_bez": "Data Science und Machine Learning",
      "organisationseinheit": "Artificial Intelligence & Data Analytics",
      "farbe": "5797e0",
      "lehrveranstaltung_id": 39802,
      "isostart": "2026-02-19T08:00:00+01:00",
      "isoend": "2026-02-19T08:45:00+01:00",
      "ort_content_id": 7127
    },
    {
      "type": "lehreinheit",
      "beginn": "08:45:00",
      "ende": "09:30:00",
      "datum": "2026-02-19",
      "topic": "DSML-ILV",
      "lektor": [
        {
          "mitarbeiter_uid": "ma0263",
          "vorname": "Patrick",
          "nachname": "Link",
          "kurzbz": "LinkPa"
        }
      ],
      "gruppe": [
        {
          "gruppe": " ",
          "verband": " ",
          "semester": "4",
          "studiengang_kz": "257",
          "kuerzel": "GRP_166146"
        }
      ],
      "ort_kurzbz": "EDV_F1.02",
      "lehreinheit_id": ["166146"],
      "titel": null,
      "lehrfach": "DSML",
      "lehrform": "ILV",
      "lehrfach_bez": "Data Science und Machine Learning",
      "organisationseinheit": "Artificial Intelligence & Data Analytics",
      "farbe": "5797e0",
      "lehrveranstaltung_id": 39802,
      "isostart": "2026-02-19T08:45:00+01:00",
      "isoend": "2026-02-19T09:30:00+01:00",
      "ort_content_id": 7127
    },
    {
      "type": "lehreinheit",
      "beginn": "09:40:00",
      "ende": "10:25:00",
      "datum": "2026-02-19",
      "topic": "DSML-ILV",
      "lektor": [
        {
          "mitarbeiter_uid": "ma0263",
          "vorname": "Patrick",
          "nachname": "Link",
          "kurzbz": "LinkPa"
        }
      ],
      "gruppe": [
        {
          "gruppe": " ",
          "verband": " ",
          "semester": "4",
          "studiengang_kz": "257",
          "kuerzel": "GRP_166146"
        }
      ],
      "ort_kurzbz": "EDV_F1.02",
      "lehreinheit_id": ["166146"],
      "titel": null,
      "lehrfach": "DSML",
      "lehrform": "ILV",
      "lehrfach_bez": "Data Science und Machine Learning",
      "organisationseinheit": "Artificial Intelligence & Data Analytics",
      "farbe": "5797e0",
      "lehrveranstaltung_id": 39802,
      "isostart": "2026-02-19T09:40:00+01:00",
      "isoend": "2026-02-19T10:25:00+01:00",
      "ort_content_id": 7127
    },
    {
      "type": "lehreinheit",
      "beginn": "10:25:00",
      "ende": "11:10:00",
      "datum": "2026-02-19",
      "topic": "DSML-ILV",
      "lektor": [
        {
          "mitarbeiter_uid": "ma0263",
          "vorname": "Patrick",
          "nachname": "Link",
          "kurzbz": "LinkPa"
        }
      ],
      "gruppe": [
        {
          "gruppe": " ",
          "verband": " ",
          "semester": "4",
          "studiengang_kz": "257",
          "kuerzel": "GRP_166146"
        }
      ],
      "ort_kurzbz": "EDV_F1.02",
      "lehreinheit_id": ["166146"],
      "titel": null,
      "lehrfach": "DSML",
      "lehrform": "ILV",
      "lehrfach_bez": "Data Science und Machine Learning",
      "organisationseinheit": "Artificial Intelligence & Data Analytics",
      "farbe": "5797e0",
      "lehrveranstaltung_id": 39802,
      "isostart": "2026-02-19T10:25:00+01:00",
      "isoend": "2026-02-19T11:10:00+01:00",
      "ort_content_id": 7127
    },
    {
      "type": "lehreinheit",
      "beginn": "11:20:00",
      "ende": "12:05:00",
      "datum": "2026-02-19",
      "topic": "SOAR-ILV",
      "lektor": [
        {
          "mitarbeiter_uid": "schwaige",
          "vorname": "Simon",
          "nachname": "Schwaiger",
          "kurzbz": "SchwaiSi"
        }
      ],
      "gruppe": [
        {
          "gruppe": " ",
          "verband": " ",
          "semester": "2",
          "studiengang_kz": "10006",
          "kuerzel": "GRP_167364"
        }
      ],
      "ort_kurzbz": "EDV_F1.03",
      "lehreinheit_id": ["167364"],
      "titel": null,
      "lehrfach": "SOAR",
      "lehrform": "ILV",
      "lehrfach_bez": "Service and object-oriented Algorithms in Robotics",
      "organisationseinheit": "Digital Manuf., Automation & Robotics",
      "farbe": "c1f08b",
      "lehrveranstaltung_id": 37521,
      "isostart": "2026-02-19T11:20:00+01:00",
      "isoend": "2026-02-19T12:05:00+01:00",
      "ort_content_id": 7128
    },
    {
      "type": "lehreinheit",
      "beginn": "12:05:00",
      "ende": "12:50:00",
      "datum": "2026-02-19",
      "topic": "SOAR-ILV",
      "lektor": [
        {
          "mitarbeiter_uid": "schwaige",
          "vorname": "Simon",
          "nachname": "Schwaiger",
          "kurzbz": "SchwaiSi"
        }
      ],
      "gruppe": [
        {
          "gruppe": " ",
          "verband": " ",
          "semester": "2",
          "studiengang_kz": "10006",
          "kuerzel": "GRP_167364"
        }
      ],
      "ort_kurzbz": "EDV_F1.03",
      "lehreinheit_id": ["167364"],
      "titel": null,
      "lehrfach": "SOAR",
      "lehrform": "ILV",
      "lehrfach_bez": "Service and object-oriented Algorithms in Robotics",
      "organisationseinheit": "Digital Manuf., Automation & Robotics",
      "farbe": "c1f08b",
      "lehrveranstaltung_id": 37521,
      "isostart": "2026-02-19T12:05:00+01:00",
      "isoend": "2026-02-19T12:50:00+01:00",
      "ort_content_id": 7128
    },
    {
      "type": "lehreinheit",
      "beginn": "12:50:00",
      "ende": "13:35:00",
      "datum": "2026-02-19",
      "topic": "SOAR-ILV",
      "lektor": [
        {
          "mitarbeiter_uid": "schwaige",
          "vorname": "Simon",
          "nachname": "Schwaiger",
          "kurzbz": "SchwaiSi"
        }
      ],
      "gruppe": [
        {
          "gruppe": " ",
          "verband": " ",
          "semester": "2",
          "studiengang_kz": "10006",
          "kuerzel": "GRP_167364"
        }
      ],
      "ort_kurzbz": "EDV_F1.03",
      "lehreinheit_id": ["167364"],
      "titel": null,
      "lehrfach": "SOAR",
      "lehrform": "ILV",
      "lehrfach_bez": "Service and object-oriented Algorithms in Robotics",
      "organisationseinheit": "Digital Manuf., Automation & Robotics",
      "farbe": "c1f08b",
      "lehrveranstaltung_id": 37521,
      "isostart": "2026-02-19T12:50:00+01:00",
      "isoend": "2026-02-19T13:35:00+01:00",
      "ort_content_id": 7128
    },
    {
      "type": "lehreinheit",
      "beginn": "15:15:00",
      "ende": "16:00:00",
      "datum": "2026-02-19",
      "topic": "ODA-ILV",
      "lektor": [
        {
          "mitarbeiter_uid": "rohatsch",
          "vorname": "Lukas",
          "nachname": "Rohatsch",
          "kurzbz": "RohatsLu"
        }
      ],
      "gruppe": [
        {
          "gruppe": " ",
          "verband": " ",
          "semester": "2",
          "studiengang_kz": "10006",
          "kuerzel": "GRP_167407"
        }
      ],
      "ort_kurzbz": "SEM_F4.04",
      "lehreinheit_id": ["167407"],
      "titel": null,
      "lehrfach": "ODA",
      "lehrform": "ILV",
      "lehrfach_bez": "Data Ethics and Open Data",
      "organisationseinheit": "Software Engineering and Architecture",
      "farbe": "dedede",
      "lehrveranstaltung_id": 37692,
      "isostart": "2026-02-19T15:15:00+01:00",
      "isoend": "2026-02-19T16:00:00+01:00",
      "ort_content_id": 7114
    },
    {
      "type": "lehreinheit",
      "beginn": "16:10:00",
      "ende": "16:55:00",
      "datum": "2026-02-19",
      "topic": "ODA-ILV",
      "lektor": [
        {
          "mitarbeiter_uid": "rohatsch",
          "vorname": "Lukas",
          "nachname": "Rohatsch",
          "kurzbz": "RohatsLu"
        }
      ],
      "gruppe": [
        {
          "gruppe": " ",
          "verband": " ",
          "semester": "2",
          "studiengang_kz": "10006",
          "kuerzel": "GRP_167407"
        }
      ],
      "ort_kurzbz": "SEM_F4.04",
      "lehreinheit_id": ["167407"],
      "titel": null,
      "lehrfach": "ODA",
      "lehrform": "ILV",
      "lehrfach_bez": "Data Ethics and Open Data",
      "organisationseinheit": "Software Engineering and Architecture",
      "farbe": "dedede",
      "lehrveranstaltung_id": 37692,
      "isostart": "2026-02-19T16:10:00+01:00",
      "isoend": "2026-02-19T16:55:00+01:00",
      "ort_content_id": 7114
    },
    {
      "type": "lehreinheit",
      "beginn": "19:30:00",
      "ende": "20:15:00",
      "datum": "2026-02-20",
      "topic": "APPCS-ILV",
      "lektor": [
        {
          "mitarbeiter_uid": "ma1221",
          "vorname": "Florian",
          "nachname": "W\u00f6rister",
          "kurzbz": "WoerisFl"
        }
      ],
      "gruppe": [
        {
          "gruppe": " ",
          "verband": " ",
          "semester": "4",
          "studiengang_kz": "335",
          "kuerzel": "GRP_165958"
        }
      ],
      "ort_kurzbz": "EDV_A2.06",
      "lehreinheit_id": ["165958"],
      "titel": null,
      "lehrfach": "APPCS",
      "lehrform": "ILV",
      "lehrfach_bez": "Applied Computer Science",
      "organisationseinheit": "Software Engineering and Architecture",
      "farbe": "f50808",
      "lehrveranstaltung_id": 39577,
      "isostart": "2026-02-20T19:30:00+01:00",
      "isoend": "2026-02-20T20:15:00+01:00",
      "ort_content_id": 7105
    },
    {
      "type": "lehreinheit",
      "beginn": "20:15:00",
      "ende": "21:00:00",
      "datum": "2026-02-20",
      "topic": "APPCS-ILV",
      "lektor": [
        {
          "mitarbeiter_uid": "ma1221",
          "vorname": "Florian",
          "nachname": "W\u00f6rister",
          "kurzbz": "WoerisFl"
        }
      ],
      "gruppe": [
        {
          "gruppe": " ",
          "verband": " ",
          "semester": "4",
          "studiengang_kz": "335",
          "kuerzel": "GRP_165958"
        }
      ],
      "ort_kurzbz": "EDV_A2.06",
      "lehreinheit_id": ["165958"],
      "titel": null,
      "lehrfach": "APPCS",
      "lehrform": "ILV",
      "lehrfach_bez": "Applied Computer Science",
      "organisationseinheit": "Software Engineering and Architecture",
      "farbe": "f50808",
      "lehrveranstaltung_id": 39577,
      "isostart": "2026-02-20T20:15:00+01:00",
      "isoend": "2026-02-20T21:00:00+01:00",
      "ort_content_id": 7105
    },
    {
      "type": "moodle",
      "beginn": "12:59:59",
      "ende": "12:59:59",
      "isostart": "2026-02-19T12:59:59+01:00",
      "isoend": "2026-02-19T12:59:59+01:00",
      "allDayEvent": true,
      "datum": "2026-2-19",
      "purpose": "assessment",
      "assignment": "Test 1",
      "topic": "Quiz opens",
      "lektor": [],
      "gruppe": [],
      "ort_kurzbz": "",
      "lehreinheit_id": null,
      "titel": "Applied Computer Science BEE 4C4, 4C3, 4E, 4D",
      "lehrfach": "",
      "lehrform": "",
      "lehrfach_bez": "",
      "organisationseinheit": "",
      "farbe": "00689E",
      "lehrveranstaltung_id": 0,
      "ort_content_id": 0,
      "url": "https:\/\/moodle.technikum-wien.at\/mod\/quiz\/view.php?id=2204175",
      "activityIcon": "https:\/\/moodle.technikum-wien.at\/theme\/image.php\/fhtw\/quiz\/1770062640\/monologo?filtericon=1",
      "actionname": "",
      "overdue": false
    },
    {
      "type": "moodle",
      "beginn": "12:59:59",
      "ende": "12:59:59",
      "isostart": "2026-02-19T12:59:59+01:00",
      "isoend": "2026-02-19T12:59:59+01:00",
      "allDayEvent": true,
      "datum": "2026-2-19",
      "purpose": "assessment",
      "assignment": "Test 1",
      "topic": "Quiz opens",
      "lektor": [],
      "gruppe": [],
      "ort_kurzbz": "",
      "lehreinheit_id": null,
      "titel": "Applied Computer Science BBE 4A, 4B, 4C1, 4C2",
      "lehrfach": "",
      "lehrform": "",
      "lehrfach_bez": "",
      "organisationseinheit": "",
      "farbe": "00689E",
      "lehrveranstaltung_id": 0,
      "ort_content_id": 0,
      "url": "https:\/\/moodle.technikum-wien.at\/mod\/quiz\/view.php?id=2204363",
      "activityIcon": "https:\/\/moodle.technikum-wien.at\/theme\/image.php\/fhtw\/quiz\/1770062640\/monologo?filtericon=1",
      "actionname": "",
      "overdue": false
    }
  ],
  "meta": { "status": "success" }
}
```
