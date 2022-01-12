package storage


type User struct {
  ID        int32     `db:"id"`
  FirstName string    `db:"first_name"`
  LastName  string    `db:"last_name"`
  Username  string    `db:"username"`
  Email     string    `db:"email"`
 }
