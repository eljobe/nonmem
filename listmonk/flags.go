package listmonk

import "flag"

var ApiUrl string

func init() {
	flag.StringVar(&ApiUrl, "lm_api", "http://127.0.0.1:9000/api", "The base URL at which the listmonk API is being hosted.")
}
