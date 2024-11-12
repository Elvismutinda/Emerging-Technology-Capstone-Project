database_url = env("DB_URL")
dev_url = env("DEV_URL")

data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "./commands/migration",
  ]
}
env "gorm" {
  src = data.external_schema.gorm.url
  dev = dev_url
  url = database_url
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}