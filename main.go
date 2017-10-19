package main

import 
("net/http"
"encode/json"
"strings")

type cityData struct {
	Name string `json:"name"`
	Main struct {
			Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFUnc("/tempo", tempo )
	http.ListenAndServe(":8080", nil)
}

func tempo(req *http.Request, res http.ResponseWriter){
	city := strings.SplitN(r.URL.Path, "/", 3)[2]

	data, err := queryCity(city)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)

}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

func queryCity(city string) (cityData, error){
	resp, err:= http.get("http://api.openweathermap.org/data/2.5/weather?APPID=YOUR_API_KEY&q=" + city)

	if ( err != nil) {
		return cityData{}, err
	}

	defer resp.Body.Close()

	var data cityData

	if err:= json.NewDecoder(resp.Body).decode(&data); err != nil {
		return cityData{}, err
	}

	return data, nil
}