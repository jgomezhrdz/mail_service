schema "example" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}

table "clientes" {
  schema = schema.example
  column "id" {
    null = false
    type = "varchar(36)"  # Assuming UUID is stored as a varchar(36)
  }
  column "nombre" {
    null = true
    type = "varchar(255)"  # Adjust the length accordingly
  }
  column "id_pedido" {
    null = true
    type = "varchar(36)"  # Assuming UUID is stored as a varchar(36)
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "id_pedido_fk" {
    columns     = [column.id_pedido]
    ref_table   = "pedidos"
    ref_columns = ["id"]
  }
}
