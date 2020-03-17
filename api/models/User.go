package models

// @Author: hasby88
import (
	"errors"
	"html"
	"log"
	"string"
	"time"
	
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID			unit32		'gorm:"primary_key;auto_increment" json:"id"'
	Nickname	string 		'gorm:"size:255;not null;unique" json:"nickname"'
	Email		string		'gorm:"size:100;not null;unique" json:"email"'
	Password	string 		'gorm:"size:100;not null;" json:"password"'
	CreatedAt	time.Time	'gorm:"default:CURRENT_TIMESTAMP" json:"created_at"'
	UpdatedAt	time.Time	'gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"'
}

func Hash( password string ) ( []byte, error ) {
	return bcrypt.GeberateFromPassword( []byte( password ), bcrypt.DefaultCost )
}

func VerifyPassword( hashedPassword, password string ) error {
	return bcrypt.CompareHashAndPassword( []byte( hashedPassword ), []byte( password ) )
}

func ( u *User ) BeforeSave() error {
	hashedPassword, err := Hash( u.Password )
	
	if err != nil {
		return err	
	}
	
	u.Password = string( hashedPassword )
	return nil
}

func ( u *User ) Prepare() {
	u.ID 		= 0
	u.Nickname 	= html.EscapeString( strings.TrimSpace( u.Nickname ) )
	u.Email 	= html.EscapeString( strings.TrimSpace( u.Email ) )
	u.CreatedAt	= time.Now()
	u.UpdatedAt = time.Now()
}

func ( u *User ) Validate( action string ) error {
	switch strings.ToLower( action ) {
		case "update":
			if u.Nickname == "" {
				return errors.New( "Required Nickname" )
			}

			if u.Password == "" {
				return errors.New( "Required Password" )
			}

			if u.Email == "" {
				return errors.New( "Required Email" )
			}

			if err := checkmail.ValidateFormat( u.Email ); err != nil {
				return errors.New( "Invalid Format Email" )
			}

			return nil
		case "login":
			if u.Email == "" {
				return errors.New( "Required Email" )
			}

			if u.Password == {
				return errors.New( "Required Password" )
			}

			if err := checkmail.ValidateFormat( u.Email ); err != nil {
				return errors.New( "Invalid Format Email" )
			}

			return nil
		default :
			if u.Nickname == "" {
				return errors.New( "Required Nickname" )
			}

			if u.Password == "" {
				return errors.New( "Required Password" )
			}

			if u.Email == "" {
				return errors.New( "Required Email" )
			}

			if err := checkmail.ValidateFormat( u.Email ); err != nil {
				return errors.New( "Invalid Format Email" )
			}
	}
}

func ( u *User ) findAll( db *gorm.DB ) ( *[]User, error ) {
	var err error
	users := []User{}
	err = db.Debug()
			.Model( &User{} )
			.Limit( 100 )
			.Find( &users )
			.Error

	if err != nil {
		return &[]Users{}, err
	}

	return &users, err 
}

func ( u *User ) findById( db *gorm.DB, uid unit32 ) ( *User, error ) {
	var err error
	err = db.Debug()
			.Model( User{} )
			.Where( "id = ?", uid )
			.Take( &u )
			.Error

	if err != nil {
		return &User{}, err
	}

	if gorm.IsRecordNotFoundError ( err ) {
		return &User{}, errors.New( "User Not Found" )
	}

	return u, err
}

func ( u *User ) save( db *gorm.DB ) ( *User, error ) {
	var err error
	err = db.Debug()
			.Create( &u )
			.Error
	
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func ( u *User ) update( db *gorm.DB, uid unit32 ) ( *User, error ) {
	err := u.BeforeSave()

	if err != nil {
		log.Fatal( err )
	}

	db = db.Debug()
			.Model( &User{} )
			.Where( "id = ?", uid )
			.Take( &User{} )
			.UpdateColumns(
				map[ string ] interface{} {
					"password", u.Password,
					"nickname", u.Nickname,
					"email", u.Email,
					"updated_at", time.Now(),
				},
			)

	if db.Error != nil {
		return &User{}, db.Error
	}

	err = db.Debug()
			.Model( &User{} )
			.Where( "id = ?", uid )
			.Take( &u )
			.Error

	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func ( u *User ) delete( db *gorm.DB, uid unit32 ) ( int64, error ) {
	db = db.Debug()
			.Model( &User{} )
			.Where( "id = ?", uid )
			.Take( &User{} )
			.Delete( &User{} )
			.Error

	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}