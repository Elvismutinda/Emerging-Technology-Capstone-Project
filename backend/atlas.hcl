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
  url = "postgres://user:password123@localhost:5432/db-name?search_path=public&sslmode=disable"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}