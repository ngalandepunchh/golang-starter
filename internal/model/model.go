package model

type DatastoreModel struct {
	Name   string `json:name`
	Team   string `json:team`
	Number int    `json:number`
}

type APIModel struct {
	Name   string `json:name`
	Team   string `json:team`
	Number int    `json:number`
}


type EventstoreModel struct{

	Name   string `json:name`
	Team   string `json:team`
	Number int    `json:number`

}
