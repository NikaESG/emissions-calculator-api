package data

type Table struct {
	tableName  struct{} `pg:"database_table,alias:g"`
	Id         int64    `pg:"id"`
	Order      int64    `pg:"order"`
	Name       string   `pg:"name"`
	DatabaseId int64    `pg:"database_id"`
}

type Field struct {
	tableName struct{} `pg:"database_field,alias:g"`
	Id        int64    `pg:"id"`
	Order     int64    `pg:"order"`
	Name      string   `pg:"name"`
	ContentId int64    `pg:"content_type_id"`
	TableId   int64    `pg:"table_id"`
}

type PostGreData struct {
	tableName struct{} `pg:",discard_unknown_columns"`
	Field_    string   `pg:"field_419"`
}
