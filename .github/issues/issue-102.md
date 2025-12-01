---
id: 102
database_id: 2769357157
node_id: I_kwDOKdrneM6lEQll
status: open
title: "stream: plans command generates an invalid alias for weekly plans in the first week of the year"
labels: ["type: bug","severity: minor","scope: code","impact: low","effort: easy"]
url: https://github.com/withsparkle/service/issues/102
created_at: 2025-01-05T16:46:08Z
updated_at: 2025-01-05T16:55:56Z
---

# stream: plans command generates an invalid alias for weekly plans in the first week of the year

`sparkle stream plans make --week --next` generated "Week 1, 2024" instead of "Week 1, 2025".

<img width="1108" alt="image" src="https://github.com/user-attachments/assets/3d04ceff-9e3d-4bb4-8cd1-893c4fd4f6eb" />
