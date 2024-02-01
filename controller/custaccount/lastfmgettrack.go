package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"rest/model"

	"github.com/joho/godotenv"
)

// Send this struct as response
type Foo struct {
	TrackName  string `json:"trackname"`
	ArtistName string `json:"artistname"`
	Image      string `json:"image"`
}

type Tracks struct {
	Tracks Toptracks_info
}

type Toptracks_info struct {
	Track []Track_info
}

type Track_info struct {
	Name      string
	Duration  string
	Listeners string
	Mbid      string
	Url       string
	Artist    Artist_info
	Image     []Image
}
type Image struct {
	Text string `json:"#text"`
	Size string
}

type Artist_info struct {
	Name string
	Mbid string
	Url  string
}

/**** JUST A TEST METHOD TO TEST MYSQL FETCH RECORD   ****/
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	posts, err := model.GetAllCustomer()
	if err != nil {
		http.Error(w, "Soemthing went wrong while processing data", http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		json.NewEncoder(w).Encode(posts)
	}
}

/**** TEST ENDS   *****/

func GetTopTrack(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method MISMATCH", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")

	err := godotenv.Load("config/.env")
	if err != nil {
		http.Error(w, "Unable to open the env file", http.StatusBadRequest)
		return
	}
	APIKEY := os.Getenv("API_KEY")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Request Body Read Error", http.StatusBadRequest)
		return
	}

	keyVal := make(map[string]string)
	if err := json.Unmarshal(body, &keyVal); err != nil {
		http.Error(w, "Cannot Parse Body. Error", http.StatusBadRequest)
		return
	}

	if country, ok := keyVal["country"]; ok {
		client := http.Client{}
		url := "https://ws.audioscrobbler.com/2.0/?method=geo.gettoptracks&country=" + country + "&api_key=" + APIKEY + "&limit=2&format=json"
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			http.Error(w, "No Content Error", http.StatusNoContent)
			return
		}
		req.Header = http.Header{
			"Host":         {"http://localhost:3200/"},
			"Content-Type": {"application/json", "UTF-8"},
			"User-Agent":   {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"},
		}
		response, err := client.Do(req)
		if err != nil {
			http.Error(w, "Error Encountered", http.StatusNotAcceptable)
			return
		}
		defer response.Body.Close()

		decoder := json.NewDecoder(response.Body)

		var data Tracks
		err = decoder.Decode(&data)
		if err != nil {
			http.Error(w, " Something Went Wrong , Internal Server Error", http.StatusInternalServerError)
			return
		}

		if len(data.Tracks.Track) == 0 {
			http.Error(w, "Empty Data Set", http.StatusNoContent)
			return
		}

		respdatas := make(map[int]Foo)
		for i, track := range data.Tracks.Track {
			respdatas[i] = Foo{TrackName: track.Name, ArtistName: track.Artist.Name, Image: track.Image[i].Text}
		}

		respjsonString, _ := json.Marshal(respdatas)
		w.Write(respjsonString)
		return
	}
	http.Error(w, "Invalid Data , Error ", http.StatusUnprocessableEntity)
}
