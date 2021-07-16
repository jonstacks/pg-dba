---
title: "Diagram"
date: 2019-03-26T14:38:07-05:00
weight: 2
---

Below is a simple diagram that shows the pg-dba workflow:

{{<mermaid align="left">}}
graph LR
    A[START]
    A --> B
    
    subgraph preanalyze
    B{"PRE_ANALYZE=True?"}
    B -->|Yes| C[Run Analyze]
    end

    B -->|No| D
    C --> D
    D[Run Query to check for table bloat]
    D-->E["Vacuum table if bloated and not too large"]
    E --> F
    
    subgraph postanalyze
    F{"POST_ANALYZE=True?"}
    F --> |Yes| G[Run Analyze]
    end

    F -->|No| H
    G --> H
    H[STOP]
{{</mermaid >}}
