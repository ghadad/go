package creature

import (
	"math/rand"
)

var creatures = []string{"shark", "jellyfish", "squid", "octopus", "dolphin"}

func Random() string {
	i := rand.Intn(len(creatures))
	return creatures[i]
}
