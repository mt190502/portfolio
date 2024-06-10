package localizator

import (
	"os"

	"github.com/mt190502/portfolio/src/utils/config"
	"github.com/mt190502/portfolio/src/utils/logger"
)

var (
	Data map[string]map[string]string
)

func InitLocalizator() {
	switch config.Config.Interface.Language {
	case "en":
		Data = map[string]map[string]string{
			"index": {
				"aboutme": "About Me",
				"name":    "Name",
				"phone":   "Phone",
				"address": "Address",
				"email":   "Email",
				"profile": "Profile",
				"social":  "Social Links",
			},
			"error": {
				"error":        "Error",
				"invalid":      "Invalid username or password",
				"internal":     "Internal server error",
				"emptypass":    "Password cannot be empty",
				"notmatchpass": "Passwords do not match",
				"hasherr":      "Failed to hash password",
				"acccreateerr": "Failed to create account",
				"valnotmatch":  "Values do not match",
				"sesssaveerr":  "Failed to save session",
			},
			"login": {
				"sign_in":     "Sign In",
				"description": "Sign in to access the admin panel",
				"username":    "Username",
				"email":       "Email",
				"password":    "Password",
				"typehere":    "Type here...",
			},
			"register": {
				"sign_up":     "Sign Up",
				"description": "Sign up to access the admin panel",
				"username":    "Username",
				"email":       "Email",
				"password":    "Password",
				"confirm":     "Confirm Password",
				"typehere":    "Type here...",
			},
			"admin": {
				"admin_panel":     "Admin Panel",
				"logout":          "Logout",
				"portfolio":       "Portfolio",
				"sitetitle":       "Site Title",
				"username":        "Username",
				"phone":           "Phone Number",
				"address":         "Address",
				"aboutme":         "About Me",
				"email":           "Email",
				"usersettings":    "User Settings",
				"newemail":        "New Email",
				"confirmemail":    "Confirm Email",
				"newpassword":     "New Password",
				"confirmpassword": "Confirm Password",
				"save":            "Save",
			},
		}
	case "tr":
		Data = map[string]map[string]string{
			"index": {
				"aboutme": "Hakkımda",
				"name":    "İsim",
				"phone":   "Telefon",
				"address": "Adres",
				"email":   "E-posta",
				"profile": "Profil",
				"social":  "Sosyal Linkler",
			},
			"error": {
				"error":       "Hata",
				"invalid":     "Geçersiz kullanıcı adı veya parola",
				"internal":    "Sunucu hatası",
				"emptypass":   "Parola boş olamaz",
				"notmatch":    "Parolalar eşleşmiyor",
				"hasherr":     "Parola hash hatası",
				"acccreate":   "Hesap oluşturma hatası",
				"valnotmatch": "Değerler eşleşmiyor",
				"sesssaveerr": "Oturum kaydetme hatası",
			},
			"login": {
				"sign_in":     "Giriş Yap",
				"description": "Yönetici paneline erişmek için giriş yapın",
				"username":    "Kullanıcı Adı",
				"email":       "E-posta",
				"password":    "Parola",
				"typehere":    "Buraya yazın...",
			},
			"register": {
				"sign_up":     "Kayıt Ol",
				"description": "Yönetici paneline erişmek için kayıt olun",
				"username":    "Kullanıcı Adı",
				"email":       "E-posta",
				"password":    "Parola",
				"confirm":     "Parolayı Onayla",
				"typehere":    "Buraya yazın...",
			},
			"admin": {
				"admin_panel":     "Yönetici Paneli",
				"logout":          "Çıkış Yap",
				"portfolio":       "Portfolyo",
				"sitetitle":       "Site Adı",
				"username":        "Kullanıcı Adı",
				"phone":           "Telefon Numarası",
				"address":         "Adres",
				"aboutme":         "Hakkımda",
				"email":           "E-posta",
				"usersettings":    "Kullanıcı Ayarları",
				"newemail":        "Yeni E-posta",
				"confirmemail":    "E-postayı Onayla",
				"newpassword":     "Yeni Parola",
				"confirmpassword": "Parolayı Onayla",
				"save":            "Kaydet",
			},
		}
	default:
		logger.Log.Fatalf("Language '%v' not supported yet...", config.Config.Interface.Language)
		os.Exit(1)
	}
}
