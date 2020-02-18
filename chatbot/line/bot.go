package line

import (
	"fmt"
	"regexp"

	rivescript "github.com/aichaos/rivescript-go"
	"github.com/aichaos/rivescript-go/lang/javascript"
)

var Bots map[string]*rivescript.RiveScript
var loc []string

func Lang() []string {
	var location []string //language json and rivescript location
	location = []string{"", "./scripts"}
	return location
}

func InitBotsMap() {
	Bots = make(map[string]*rivescript.RiveScript)
}

//this method is used to load Context, from rivescript
func LoadContext(id, category string) {
	if Bots[id+category] == nil {
		// var (
		// 	debug = flag.Bool("debug", false, "Enable debug mode for RiveScript.")
		// 	utf8  = flag.Bool("utf8", true, "Enable UTF-8 mode")
		// )
		Bots[id+category] = rivescript.New(&rivescript.Config{
			Debug: false,
			UTF8:  true, // UTF-8 support enabled
		})
		Bots[id+category].SetHandler("javascript", javascript.New(Bots[id]))
	}

	//get Bot directory
	// directory := Lang()
	var err error

	if category == "new" {
		err = Bots[id+category].LoadFile("./scripts/newBot.rive")
	} else {
		err = Bots[id+category].LoadFile("./scripts/registeredBot.rive")
	}

	// err := Bot.LoadDirectory(directory[1]) //get script location on array index of 1
	if err != nil {
		fmt.Println(err)
		// GetErrorResponse("Error loading from file:", err)
	}
	Bots[id+category].SortReplies()
}

func GetBotReply(message, id, category string) string {
	reSymbol := regexp.MustCompile(`[.?,!;:@#$%^&*()]+`)
	formattedMsg := reSymbol.ReplaceAllString(message, "")
	reply, _ := Bots[id+category].Reply(id, formattedMsg)

	return reply
}
