package main

const (
	COLUMNS_QUERY = `
	select 
		a.attname as col,
		case 
			when left(t.typname, 1) = '_' then right(t.typname, length(t.typname)-1) 
			else t.typname
		end,
		(	select 
				substring(pg_catalog.pg_get_expr(d.adbin, d.adrelid) for 128) as default
			from pg_catalog.pg_attrdef d
			where d.adrelid = a.attrelid 
				and d.adnum = a.attnum 
				and a.atthasdef
		),
		not a.attnotnull,
		t.typcategory = 'A' as is_array,
		t.typcategory
		from pg_catalog.pg_attribute a
			join pg_type t on t.oid = a.atttypid
		where a.attrelid = $1::regclass
			and a.attnum > 0
			and not a.attisdropped
		order by a.attnum;`
)
