---
code: null
id: I_kwDOKdrneM6lERvK
databaseId: 2769361866
number: 103
url: https://github.com/withsparkle/service/issues/103
title: "stream: plans command doesn't apply month aliases"
issueType: null
labels:
  - "type: bug"
  - "severity: minor"
  - "scope: code"
  - "impact: low"
  - "effort: easy"
assignees:
  - kamilsk
milestone: null
parent: null
issueFields: {}
state: OPEN
stateReason: null
duplicateOf: null
createdAt: 2025-01-05T16:58:58Z
updatedAt: 2025-01-05T16:58:59Z
lastEditedAt: null
closedAt: null
---

# stream: plans command doesn't apply month aliases

`sparkle stream plans make --month` didn't apply template to the aliases. Aliases contain "MMMM" and "MMMM, YYYY" instead of "January" and "January, 2025".

<img width="301" alt="image" src="https://github.com/user-attachments/assets/7bef29fc-a15b-48cd-bf6a-ca1474c890dc" />
