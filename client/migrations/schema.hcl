schema "codely" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}

table "clientes" {
  schema = schema.codely
  column "id" {
    null = false
    type = varchar(36)  # Assuming UUID is stored as a varchar(36)
  }
  column "nombre" {
    null = false
    type = varchar(255)  # Adjust the length accordingly
  }
  column "id_plan" {
    null = false
    type = varchar(36)  # Assuming UUID is stored as a varchar(36)
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "id_plan_fk" {
    columns     = [column.id_plan]
    ref_columns = [table.planes.column.id]
  }
}

table "planes" {
  schema = schema.codely
  column "id" {
    null = false
    type = varchar(36)  # Assuming UUID is stored as a varchar(36)
  }
  column "nombre" {
    null = false
    type = varchar(255)  # Adjust the length accordingly
  }
  column "quota_month" {
    null = false
    type = int(255)  # Assuming UUID is stored as a varchar(36)
  }
  column "quota_day" {
    null = false
    type = int(255)  # Assuming UUID is stored as a varchar(36)
  }
  primary_key {
    columns = [column.id]
  }
}