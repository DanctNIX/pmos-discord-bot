// pmos-bot - A bot for the postmarketOS Matrix channels
// Now ported to Discord!
//
// Copyright (C) 2017 Tulir Asokan
// Copyright (C) 2018-2019 Luca Weiss
// Copyright (C) 2020 DanctNIX Community
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"regexp"
	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/gateway"
	"github.com/diamondburned/arikawa/session"
)

func main() {
	var token = os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatalln("No $BOT_TOKEN given. Please define one.")
	}

	log.Println("Try logging in to Discord...")
	syncer, err := session.New("Bot " + token)
	if err != nil {
		log.Fatalln("Session failed:", err)
	}

	if err := syncer.Open(); err != nil {
		log.Fatalln("Failed to connect:", err)
	}

	defer syncer.Close()

	u, err := syncer.Me()
	if err != nil {
		log.Fatalln("Failed to get myself:", err)
	}

	log.Println("Logged as", u.Username)

	shortcutmap := map[string]string{
		"art#": "https://gitlab.com/postmarketOS/artwork/issues/",
		"art!": "https://gitlab.com/postmarketOS/artwork/merge_requests/",

		"bpo#": "https://gitlab.com/postmarketOS/build.postmarketos.org/issues/",
		"bpo!": "https://gitlab.com/postmarketOS/build.postmarketos.org/merge_requests/",

		"bot#": "https://gitlab.com/postmarketOS/matrix-bot/issues/",
		"bot!": "https://gitlab.com/postmarketOS/matrix-bot/merge_requests/",

		"chrg#": "https://gitlab.com/postmarketOS/charging-sdl/issues/",
		"chrg!": "https://gitlab.com/postmarketOS/charging-sdl/merge_requests/",

		"lnx#": "https://gitlab.com/postmarketOS/linux-postmarketos/issues/",
		"lnx!": "https://gitlab.com/postmarketOS/linux-postmarketos/merge_requests/",

		"mrh#": "https://gitlab.com/postmarketOS/mrhlpr/issues/",
		"mrh!": "https://gitlab.com/postmarketOS/mrhlpr/merge_requests/",

		"osk#": "https://gitlab.com/postmarketOS/osk-sdl/issues/",
		"osk!": "https://gitlab.com/postmarketOS/osk-sdl/merge_requests/",

		"pma#": "https://gitlab.com/postmarketOS/pmaports/issues/",
		"pma!": "https://gitlab.com/postmarketOS/pmaports/merge_requests/",

		"pmb#": "https://gitlab.com/postmarketOS/pmbootstrap/issues/",
		"pmb!": "https://gitlab.com/postmarketOS/pmbootstrap/merge_requests/",

		"org#": "https://gitlab.com/postmarketOS/postmarketos.org/issues/",
		"org!": "https://gitlab.com/postmarketOS/postmarketos.org/merge_requests/",

		"wiki#": "https://gitlab.com/postmarketOS/wiki/issues/",
	}

	shortcutmapregex := regexp.MustCompile("(?i)(art[#!]|bpo[#!]|bot[#!]|chrg[#!]|lnx[#!]|mrh[#!]|osk[#!]|pma[#!]|pmb[#!]|org[#!]|wiki#)(\\d+)")
	
	syncer.AddHandler(func(c *gateway.MessageCreateEvent) {
		matches := shortcutmapregex.FindAllStringSubmatch(c.Content, -1)
		if matches != nil {
			var buffer bytes.Buffer
			for _, match := range matches {
				log.Printf("<%[1]s> %[3]s (%[2]s)\n", c.Author.Username, c.Author.ID, c.Content)
				buffer.WriteString(shortcutmap[strings.ToLower(match[1])] + match[2] + " ")
			}
			message := strings.TrimSuffix(buffer.String(), " ")
			syncer.SendMessage(c.ChannelID, message, nil)
		}
	})

	// We'll make sure that the bot won't stop until SIGINT.
	bot.Wait()
}
