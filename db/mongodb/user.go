package mongodb

import "time"

type User struct {
	Gender        int       `bson:"gender"`
	Id            string    `bson:"_id"`
	Name          string    `bson:"name"`
	Password      string    `bson:"password"`
	Phone         string    `bson:"phone"`
	Email         string    `bson:"email"`
	RegisterTime  time.Time `bson:"registerTime"`
	Birthday      time.Time `bson:"birthday"`
	PhoneVerified bool      `bson:"phoneVerified"`
	EmailVerified bool      `bson:"emailVerified"`
}
