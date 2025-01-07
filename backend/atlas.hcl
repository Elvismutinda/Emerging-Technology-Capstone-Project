data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "./commands/migration",
  ]
}
env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/15/dev"
  url = "postgres://postgres:password123@localhost:5432/budget_tracker?search_path=public&sslmode=disable"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}