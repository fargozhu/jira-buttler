# Azure Jira Butler

## Introduction

The idea is to understand and to learn how can we deploy infra and app using GH actions.

## Warming the Engine

[Jira Cloud API](https://developer.atlassian.com/cloud/jira/platform/rest/v3/intro/) for basic but important daily
actions.

Sometimes we need to create Jira tickets based on platform log or metrics in a simple and elegant way.

## Design

## Request

```http
POST https://<>/api/HttpJiraButler?action=create|update|close
```

```json
{
  "jira_id": ,
  "summary": ,
  "description": ,
  "assignee" ,
  "labels": ["",""]
}
```

## Response

The HTTP code is a transparent value from the Jira Cloud API response.

```http
HTTP/1.1 200 OK
Access-Control-Allow-Origin: *
Content-Length: 160
Content-Type: application/json
Date: Fri, 21 May 2021 21:10:49 GMT
Server: Kestrel
```

The response contains a subset of the data returned by Jira Cloud API.

```json
{
  "jira_id": "10000",
  "key": "NMBRS-12442",
  "self": "https://your-domain.atlassian.net/rest/api/3/issue/10000"
}
```