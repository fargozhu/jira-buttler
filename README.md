# Azure Jira Butler

## Summary

Azure Function API that integrates with
[Jira Cloud API](https://developer.atlassian.com/cloud/jira/platform/rest/v3/intro/) for basic but important daily
actions.

The initial need was to be able to create Jira issue whenever a certificate was about to expired, but then the team
though that we should leave it open for other contexts.

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