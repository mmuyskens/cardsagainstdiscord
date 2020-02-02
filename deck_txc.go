package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "txc",
		Description: "The X Crew",
		Prompts: []*PromptCard{
		},

		Responses: []ResponseCard{			
`Reaper`,
`Marky Mark`,
`RaxinII`,
`Nadine`,
`Bloodlust`,
		},
	}

	AddPack(pack)
}
