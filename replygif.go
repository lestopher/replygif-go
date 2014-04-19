package replygif

const REPLYGIF_API = "http://replygif.net/api"

type ReplygifGifData struct {
	Id      string
	Tags    []string
	Caption []string
	File    string
}

type ReplygifTagData struct {
	Title    string
	Id       string
	Reaction bool
	Url      string
	Reply    string
	Count    string
}
