db.createUser(
    {
        user : "porthose",
        pwd : "789789",
        roles : [
            {
                role : "readWrite",
                db : "ampgodb"
            }
        ]
    }
)