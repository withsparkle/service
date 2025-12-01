---
id: 103
database_id: 2769361866
node_id: I_kwDOKdrneM6lERvK
status: open
title: "stream: plans command doesn't apply month aliases"
labels: ["type: bug","severity: minor","scope: code","impact: low","effort: easy"]
url: https://github.com/withsparkle/service/issues/103
created_at: 2025-01-05T16:58:58Z
updated_at: 2025-01-05T16:58:59Z
---

# stream: plans command doesn't apply month aliases

`sparkle stream plans make --month` didn't apply template to the aliases. Aliases contain "MMMM" and "MMMM, YYYY" instead of "January" and "January, 2025".

<img width="301" alt="image" src="https://github.com/user-attachments/assets/7bef29fc-a15b-48cd-bf6a-ca1474c890dc" />
