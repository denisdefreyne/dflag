feature_flag "payments_v2" {
  team_responsible = "Protocol Team <protocol@example.com>"
  contact_person = "Denis Defreyne <denis@example.com>"

  # date_created = "2025-02-01"
  # date_last_checked = "2025-03-03"

  environment "staging" {
    enabled = true
  }

  environment "production" {
    enabled = false
  }
}

feature_flag "rate_limit_v3" {
  team_responsible = "Protocol Team <protocol@example.com>"
  contact_person = "Denis Defreyne <denis@example.com>"

  # date_created = "2025-02-01"
  # date_last_checked = "2025-03-03"

  environment "staging" {
    enabled = true
  }

  environment "production" {
    enabled = false
  }
}
