package application

import (
	"encoding/json"
	// "fmt"
	"net/http"
	"strconv"
	"github.com/MaksaNeNegr/calc_go/pkg/rpn"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}

type Request struct {
	Expression string `json: "expression"`
}

type Response struct {
	Res string `json:"result"`
}

type Err_Response struct {
	Error_ string `json:"error"`
}

type error interface {
    Error() string
}

func retErr(w http.ResponseWriter, code int, err error){
	w.WriteHeader(code)
	var err_ Err_Response
	err_.Error_ = err.Error()
	json.NewEncoder(w).Encode(err_)
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		retErr(w, 500, rpn.Err_no_post)
		return // единственный случай, когда возращяется код 500
	}

	var req Request
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&req)
	res_, err := rpn.Calc(req.Expression)
	w.Header().Set("Content-Type", "application/json")
	if err == nil{
		w.WriteHeader(200)
		var res Response
		res.Res = res_
		json.NewEncoder(w).Encode(res)
	} else {
		retErr(w, 422, err)
	}
	// fmt.Printf(res_)
	// fmt.Println(req.Expression)

	// w.Header().Set("Content-Type", "application/json")
}

func accuracy_(w http.ResponseWriter, r *http.Request) {
	acc := r.URL.Query().Get("accuracy")
	acc_, err := strconv.Atoi(acc)
	if err != nil{
		retErr(w, 405, rpn.Err_acc)
	} else if acc_ >= 65 || acc_ < 0{
		retErr(w, 405, rpn.Err_acc)
	} else {
		// acc, _ = strconv.Itoa(acc_ + 2) // почему то всегда округляется на 2 цифры меньше???
		rpn.ChangeTochonst(strconv.Itoa(acc_ + 2))
	}
}


func (a *Application) Run() { 
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	// http.HandleFunc("/api/v1/calculate/acc", accuracy_)
	http.HandleFunc("/api/v1/calculate/acc", accuracy_)
	http.ListenAndServe(":8080", nil)
}
// curl http://localhost:8080/api/v1/calculate/acc?accuracy=2
