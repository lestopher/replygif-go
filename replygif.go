package replygif

const REPLYGIF_API = "http://replygif.net"

type ReplyGifGifData struct {
	Id      string
	Tags    []string
	Caption []string
	File    string
}

type ReplyGifTagData struct {
	Title    string
	Id       string
	Reaction bool
	Url      string
	Reply    string
	Count    string
}
