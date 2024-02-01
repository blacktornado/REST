package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Bar struct {
	Track_Id    int
	Track_Name  string
	Album_Id    int
	Album_Name  string
	Artist_Id   int
	Artist_Name string
	Lyrics_Body string
}

/*** FOR TRACK INFO ARTIST INFO ALBUM INFO ****/
type Message struct {
	Message Head
}

type Head struct {
	Header Header_detail
	Body   Track_list_details
}

type Header_detail struct {
	Status_Code  int
	Execute_Time float64
}

type Track_list_details struct {
	Track_List []Track_check `json:"track_list"`
}

type Track_check struct {
	Track Track_details
}
type Track_details struct {
	Track_ID       int
	Track_Name     string
	Commontrack_Id int
	Has_Lyrics     int
	Album_Id       int
	Album_Name     string
	Artist_Id      int
	Artist_Name    string
}

/*** FOR LYRICS ONLY *****/
type Messages struct {
	Message Head_i
}

type Head_i struct {
	Body Lyrics
}

type Lyrics struct {
	Lyrics Lyrics_info
}

type Lyrics_info struct {
	Lyrics_Body string
}

func GetTopTrackL(w http.ResponseWriter, r *http.Request) {

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
	MUAPIKEY := os.Getenv("MUMATCH_API_KEY")

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
		countrycode := firstN(country, 2) //get first two letters only
		response, err := http.Get("https://api.musixmatch.com/ws/1.1/chart.tracks.get?chart_name=top&page=1&page_size=1&country=" + countrycode + "&f_has_lyrics=1&apikey=" + MUAPIKEY)
		if err != nil {
			http.Error(w, "Cannot Parse Body. Error", http.StatusNoContent)
			return
		}

		decoder := json.NewDecoder(response.Body)

		var data Message
		err = decoder.Decode(&data)
		if err != nil {
			http.Error(w, " Something Went Wrong , Internal Server Error", http.StatusInternalServerError)
			return
		}
		if len(data.Message.Body.Track_List) == 0 {
			http.Error(w, "Empty Data Set", http.StatusNoContent)
			return
		}
		/****TRACK INFO ENDS HERE ******/

		var dataa Messages              /*** placeholder for incoming data in this STRUCT ***/
		var appendlyrics = []Messages{} /*** creating a slice to hold data from loop ***/
		for _, track := range data.Message.Body.Track_List {
			tostring := strconv.Itoa(track.Track.Track_ID)
			responseval, err := http.Get("https://api.musixmatch.com/ws/1.1//track.lyrics.get?format=json&callback=http://tracking.musixmatch.com/t1.0/AMa6hJCIEzn1v8RuOP&track_id=" + tostring + "&apikey=" + MUAPIKEY)
			if err != nil {
				http.Error(w, "Cannot Parse Body. Error", http.StatusNoContent)
				return
			}
			decoderr := json.NewDecoder(responseval.Body)
			errs := decoderr.Decode(&dataa)
			appendlyrics = append(appendlyrics, dataa)
			if errs != nil {
				http.Error(w, " Something Went Wrong , Internal Server Error", http.StatusInternalServerError)
				return
			}
		}

		respdatas := make(map[int]Bar)
		for i, trackin := range data.Message.Body.Track_List {
			respdatas[i] = Bar{Track_Id: trackin.Track.Track_ID, Track_Name: trackin.Track.Track_Name,
				Album_Id: trackin.Track.Album_Id, Album_Name: trackin.Track.Album_Name,
				Artist_Id: trackin.Track.Artist_Id, Artist_Name: trackin.Track.Artist_Name, Lyrics_Body: appendlyrics[i].Message.Body.Lyrics.Lyrics_Body}
		}
		respjsonString, _ := json.Marshal(respdatas)
		w.Write(respjsonString)
		return
	}
	http.Error(w, "Invalid Data , Error ", http.StatusUnprocessableEntity)
}

func firstN(s string, n int) string {
	i := 0
	for j := range s {
		if i == n {
			return s[:j]
		}
		i++
	}
	return s
}
