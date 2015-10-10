package composers

import (
	"math/rand"
	"strings"
)

type Composers []Composer

func (c Composers) find(name string) Composer {
	for _, composer := range c {
		if strings.Contains(strings.ToLower(composer.Name), name) {
			return composer
		}
	}

	return Composer{}	
}

var (
	composers = Composers{
		{`Johann Sebastian Bach`, `http://historylists.org/images/johann-sebastian-bach.jpg`},
		{`Wolfgang Amadeus Mozart`, `http://historylists.org/images/wolfgang-amadeus-mozart.jpg`},
		{`Ludwig van Beethoven`, `http://historylists.org/images/ludwig-van-beethoven.jpg`},
		{`Giuseppe Verdi`, `http://historylists.org/images/giuseppe-verdi.jpg`},
		{`Pyotr Ilyich Tchaikovsky`, `http://historylists.org/images/pyotr-ilyich-tchaikovsky.jpg`},
		{`Frederic Chopin`, `http://historylists.org/images/frederic-chopin.jpg`},
		{`Antonio Vivaldi`, `http://historylists.org/images/antonio-vivaldi.jpg`},
		{`Giacomo Puccini`, `http://historylists.org/images/giacomo-puccini.jpg`},
		{`George Frideric Handel`, `http://historylists.org/images/george-frideric-handel.jpg`},
		{`Igor Stravinsky`, `http://historylists.org/images/igor-stravinsky.jpg`},
	}
)

func Find(name string) Composer {
	return composers.find(strings.ToLower(name))
}

func Rand() Composer {
	return composers[rand.Intn(len(composers))]
}

