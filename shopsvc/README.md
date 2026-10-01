# shopsvc: Code Review Exercise

**Time:** 25–30 minutes

## Scenario

This is a small order-processing service. It takes orders from customers, charges payments, sends notifications, and stores everything in Postgres. The code compiles (`go build ./...`). Your job is **not** to look for logic bugs. Focus only on how the code is **structured and named**.

## Your task

1. Identify every place where a package, type, function, or file is named or organized poorly.
2. For each one, explain **why** it's a problem (what it costs the team over time).
3. Propose a better name or structure.
4. Some parts of the code are fine. Don't flag something unless you can justify it.

## What to submit

For each issue: **location**, **problem**, **why it matters**, **suggested fix**. Then, in two or three sentences, describe how you would restructure the project overall.
