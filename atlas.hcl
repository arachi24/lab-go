
// data "external_schema" "gorm" {
//   program = [
//     "go",
//     "run",
//     "-mod=mod",
//     "ariga.io/atlas-provider-gorm",
//     "load",
//     "--path", "./model",
//     "--dialect", "sqlserver", // | postgres | sqlite | sqlserver
//   ]
// }
// data "runtimevar" "db" {
//   url = "Server=tcp:localhost,1433;Database=arv-local;User ID=sa;Password=P@ssw0rd;TrustServerCertificate=true;Trusted_Connection=False;Encrypt=True;"
// }

// env "dev" {
//   src = "schema.hcl"
//   url = "sqlserver://sa:P@ssw0rd@localhost:1433?database=arv-local"
//   dev = "docker://sqlserver/2022-latest"
//     migration {
//     dir = "file://migrations"
//   }
//   format {
//     migrate {
//       diff = "{{ sql . \"  \" }}"
//     }
//   }
// }
// env "gorm" {
//    src = data.external_schema.gorm.url
// //    url = "sqlserver://sa:P@ssw0rd@localhost:1433?database=arv-local"
//    dev = "docker://sqlserver"
//   migration {
//     dir = "file://migrations"
//   }
//   format {
//     migrate {
//       diff = "{{ sql . \"  \" }}"
//     }
//   }
// }

// env "local" {
//   src = "file://schema/schema.hcl"
//   url = "sqlserver://sa:P@ssw0rd@localhost:1433?database=arv-local"
//   dev = "docker://sqlserver/2022-latest"
//    migration {
//     dir = "file://migrations"
//   }
//   format {
//     migrate {
//       diff = "{{ sql . \"  \" }}"
//     }
//   }
// }

// data "external_schema" "gorm" {
//   program = [
//     "go",
//     "run",
//     "-mod=mod",
//     "./loader",
//   ]
// }
// data "external_schema" "gorm" {
//   program = [
//     "go",
//     "run",
//     "-mod=mod",
//     "ariga.io/atlas-provider-gorm",
//     "load",
//     "--path", "./models",
//     "--dialect", "sqlserver", // | postgres | sqlite | sqlserver
//   ]
// }

// env "gorm" {
//   src = data.external_schema.gorm.url
//   dev = "docker://sqlserver/2022-latest"
//   migration {
//     dir = "file://migrations"
//   }
//   format {
//     migrate {
//       diff = "{{ sql . \"  \" }}"
//     }
//   }
// }

data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./cmd/loader",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  // url = "sqlserver://sa:P@ssw0rd@localhost:1433?database=arv-local"
  dev = "docker://sqlserver/2022-latest"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}