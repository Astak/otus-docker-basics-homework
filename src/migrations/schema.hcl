table "users" {
    schema = schema.public
    column "id" {
        null = false
        type = bigserial
    }
    column "username" {
        null = false
        type = text
    }
    column "firstname" {
        null = false
        type = text
    }
    column "lastname" {
        null = false
        type = text
    }
    column "email" {
        null = false
        type = text
    }
    column "phone" {
        null = false
        type = text
    }
    primary_key {
        columns = [column.id]
    }
}
schema "public" {
    comment = "Default public gomin schema"
}