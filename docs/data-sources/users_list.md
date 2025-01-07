---
subcategory: "Slack"
page_title: "Slack: slack_users_list"
---

# slack_users_list Data Source

Use this data source to get a list of users for use in other
resources. 

This resource is quite slow and should be used with caution. 
It lists any user, including the deleted (deactivated) ones

## Required scopes

This resource requires the following scopes:

- [users:read](https://api.slack.com/scopes/users:read)
- [users:read.email](https://api.slack.com/scopes/users:read.email)

The Slack API methods used by the resource are:

- [users.list](https://api.slack.com/methods/users.list)

If you get `missing_scope` errors while using this resource check the scopes against
the documentation for the methods above.

## Example Usage

```hcl
data "slack_users_list" "all" {
}

data "slack_users_list" "by_email" {
  email = "my-user@example.com"
}
```

## Argument Reference

The following arguments are supported:

- `email` - (Optional) Filters the list of users by email


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- `users` - The list of users

Each `user` has these attributes : 
- `deleted` - Whether the user is deleted (deactivated)
- `email` - The email of the user
- `id` - The ID of the user
- `name` - The name of the user
- `real_name` - The real name of the user
