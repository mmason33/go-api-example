package model

type Treatment struct {
	Title         string `json:"title"`
	Treatment_id  int    `json:"treatment_id"`
	Firebase_id   string `json:"firebase_id"`
	Document_date string `json:"document_date"`
	Id            int    `json:"id"`
}

type Treatments struct {
	Data []Treatment `json:"data"`
}
