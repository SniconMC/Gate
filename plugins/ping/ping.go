package ping

import (
	"context"
	"github.com/go-logr/logr"
	. "github.com/minekube/gate-plugin-template/util"
	"github.com/robinbraemer/event"
	"github.com/thedevminertv/minimsg"
	"go.minekube.com/gate/pkg/edition/java/proxy"
	"golang.org/x/exp/rand"
	"time"
)

// Plugin is a ping plugin that handles ping events.
var Plugin = proxy.Plugin{
	Name: "Ping",
	Init: func(ctx context.Context, p *proxy.Proxy) error {
		log := logr.FromContextOrDiscard(ctx)
		log.Info("Hello from Ping plugin!")

		event.Subscribe(p.Event(), 0, onPing())

		return nil
	},
}

// Randomly selects a string from the provided list
func randomRow() string {
	rows := []string{
		"<#818181>\nMade with </#818181><gradient:#ff6c2f:#ff76b6>Minestom</gradient>",
		"<#818181>\nCheck out <underline>mc.emortal.dev</underline>!</#818181>",
		"<#818181>\nToo many late nights were spent on this.</#818181>",
		"<#818181>\nTrans rights are human rights.</#818181>",
		"<#818181>\nOf Swedish origin!</#818181>",
	}
	rand.Seed(uint64(time.Now().UnixNano()))
	return rows[rand.Intn(len(rows))]
}

func onPing() func(*proxy.PingEvent) {
	// Static first row
	line1 := minimsg.Parse("<bold>\u200C</bold><gray>                    → </gray> " +
		"<bold><gradient:#ffff1c:gold>play.znopp.pw</gradient></bold>" +
		"<gray> ←           </gray><bold>\u200C</bold><dark_gray>⏵ v0.1 ⏴</dark_gray>")

	return func(e *proxy.PingEvent) {
		// Randomly selected second row
		randomLine := randomRow()

		line2 := minimsg.Parse(randomLine)

		p := e.Ping()
		// Concatenate the two lines as the description
		p.Description = Join(line1, line2)
		p.Players.Max = p.Players.Online + 1
	}
}
