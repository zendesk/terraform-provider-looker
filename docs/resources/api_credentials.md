---
page_title: "looker_api_credentials Resource - terraform-provider-looker"
subcategory: ""
description: |-
Manage API credentials for a Looker user.

---
# looker_api_credentials (Resource)

## Example Usage
```terraform
resource "looker_api_credentials" "example" {
  user_id     = 123
  type        = "api3"
  is_disabled = false
}
```

## Example Output
```
% terraform show
# looker_api_credentials.example:
resource "looker_api_credentials" "example" {
  id          = "abcd1234"
  user_id     = 123
  type        = "api3"
  client_id   = "client-id-value"
  url         = "https://your.looker.instance/api/4.0/users/123/credentials_api3/abcd1234"
  is_disabled = false
}
```

## Schema

### Required
- `user_id` (Number) ID of the user owning the API credential
- `type` (String) Type of API credential (e.g., api3)

###Optional
- `is_disabled` (Boolean) Whether the credential is disabled. Changing this forces resource recreation.
### Computed
- `client_id` (String) Client ID of the API credential
- `client_secret` (String, Sensitive) Client secret, only available on creation
- `url` (String) URL of the API credential resource
### Read-Only
- `id` (String) API credential ID (Terraform resource ID)
