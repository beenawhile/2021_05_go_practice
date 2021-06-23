package app

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// oauth install : go get golang.org/x/oauth2
//								 go get cloud.google.com/go

type GoogleUserId struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

var googleOauthConfig = oauth2.Config{
	RedirectURL:  "http://localhost:3000/auth/google/callback",
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_SECRET_KEY"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func googleLoginHandler(w http.ResponseWriter, r *http.Request) {
	state := generateOauthCookie(w)
	// user 어떤 경로로 보내야하는지
	url := googleOauthConfig.AuthCodeURL(state)
	// state는 one-time key가 필요함 => user browser cookie에 temporary key를 심고 redirect 왔을 때 대조해봄
	// 상식 => CSRF Attack : URL 변조 공격
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func generateOauthCookie(w http.ResponseWriter) string {
	// 만료기간
	expiration := time.Now().Add(1 * 24 * time.Hour)
	// 16 byte 랜덤 키 만들기
	b := make([]byte, 16)
	// random으로 만든것 채우기
	rand.Read(b)
	// byte를 html encoding 해줘야함
	state := base64.URLEncoding.EncodeToString(b)
	cookie := &http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, cookie)
	return state
}

// callback 돌아왔을 때
func googleAuthCallback(w http.ResponseWriter, r *http.Request) {
	// cookie 가져오기
	oauthstate, _ := r.Cookie("oauthstate")
	// google이 request에 state값 받아옴
	if r.FormValue("state") != oauthstate.Value {
		// 잘못된 state이기 때문에 원래로 보내버림
		errMsg := fmt.Sprintf("Invalid google oauth state, cookie: %s state: %s", oauthstate.Value, r.FormValue("state"))
		log.Println(errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
	}

	// request에 API KEY, Refresh Key를 알려줌 => 이를 받아서 user info 가져옴
	data, err := getGoogleUserInfo(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// id 정보를 session에 저장
	var userInfo GoogleUserId
	err = json.Unmarshal(data, &userInfo)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// session cookie에 id 저장
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set some session values.
	session.Values["id"] = userInfo.ID
	// Save it before we write to the response/return from the handler.
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

	fmt.Fprint(w, string(data))
}

func getGoogleUserInfo(code string) ([]byte, error) {
	// context : 쓰레드간의 데이터를 주고 받을 때 사용되는 thread-safe 한 저장소
	// multi-thread 환경이 아니기 때문에 background로
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("Failed to Exchage %s\n", err.Error())
	}

	resp, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("Failed to get UserInfo %s\n", err.Error())
	}

	return ioutil.ReadAll(resp.Body)

}
