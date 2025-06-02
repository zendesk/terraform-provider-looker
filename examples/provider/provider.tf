terraform {
  required_providers {
    looker = {
      source  = "zendesk/looker"
      version = "x.x.x"
    }
  }
}

provider "looker" {
  base_url      = "https://org.cloud.looker.com:19999/api/" # Optionally use env var LOOKER_BASE_URL
  client_id     = "12345678"                                # Optionally use env var LOOKER_API_CLIENT_ID
  client_secret = "abcd1234"                                # Optionally use env var LOOKER_API_CLIENT_SECRET
}
