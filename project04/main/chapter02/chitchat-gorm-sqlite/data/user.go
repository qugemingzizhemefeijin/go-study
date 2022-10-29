package data

import (
	"gorm.io/gorm"
)

// https://www.shuzhiduo.com/A/kjdwEZ9qJN/ gorm更新

// http://eryajf.net/pages/4d0ed0/ gorm框架更新与删除

// https://www.shuzhiduo.com/A/obzbaxvMdE/ 基本查询

type User struct {
	gorm.Model
	UUID      string
	Name      string
	Email     string
	Password  string
}

type Session struct {
	gorm.Model
	UUID      string
	Email     string
	UserId    int
}

// Create a new session for an existing user
func (user *User) CreateSession() (session Session, err error) {
	uuid := createUUID()

	s := &Session{
		UUID: uuid,
		Email: user.Email,
		UserId: int(user.ID),
	}

	Db.Create(s)

	// 设置session的信息并返回
	session.UUID = uuid
	session.Email = user.Email
	session.UserId = int(user.ID)
	session.CreatedAt = s.CreatedAt
	session.ID = s.ID

	return
}

// Get the session for an existing user
func (user *User) Session() (session Session, err error) {
	Db.First(&session, "user_id = ?", user.ID)

	return
}

// Check if session is valid in the database
func (session *Session) Check() (valid bool, err error) {
	Db.First(session, "uuid = ?", session.UUID)

	if session.ID != 0 {
		valid = true
	}
	return
}

// Delete session from database
func (session *Session) DeleteByUUID() (err error) {
	Db.Unscoped().Where("uuid = ?", session.UUID).Delete(&Session{})

	return
}

// Get the user from the session
func (session *Session) User() (user User, err error) {
	Db.First(&user, session.UserId) // 通过主键查询

	return
}

// Get the user from the session and not return error
func (session *Session) Luxifa() (user User) {
	Db.First(&user, session.UserId) // 通过主键查询

	return
}

// Delete all sessions from database
func SessionDeleteAll() (err error) {
	Db.Unscoped().Where("1 = 1").Delete(&Session{})

	return
}

// Create a new user, save user info into the database
func (user *User) Create() (err error) {
	// Postgres does not automatically return the last insert id, because it would be wrong to assume
	// you're always using a sequence.You need to use the RETURNING keyword in your insert to get this
	// information from postgres.
	user.UUID = createUUID()
	user.Password = Encrypt(user.Password)

	// 创建用户
	Db.Create(user)

	// id, err := result.LastInsertId()
	return
}

// Delete user from database
func (user *User) Delete() (err error) {
	// 按照主键删除
	Db.Unscoped().Delete(user)

	return
}

// Update user information in the database
func (user *User) Update() (err error) {
	Db.Model(user).Updates(User{Name: user.Name, Email: user.Email})

	return
}

// Delete all users from database
func UserDeleteAll() (err error) {
	Db.Unscoped().Where("1 = 1").Delete(&User{})

	return
}

// Get all users in the database and returns it
func Users() (users []User, err error) {
	Db.Find(&users)

	return
}

// Get a single user given the email
func UserByEmail(email string) (user User, err error) {
	Db.First(&user, "email = ?", email)

	return
}

// Get a single user given the UUID
func UserByUUID(uuid string) (user User, err error) {
	Db.First(&user, "uuid = ?", uuid)

	return
}

