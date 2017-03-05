package api

type Team struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Players []Player `json:"players"`
}

type Player struct {
	ID   int    `json:"id,string"`
	Name string `json:"name"`
	Age  int    `json:"age,string"`
}
