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
  url = "postgres://postgres:1234567@localhost:5432/test_db?search_path=public&sslmode=disable"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}