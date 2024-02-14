package pm

import (
	"fmt"
	"github.com/bitfield/script"
	"log"
	"math/rand"
)

type Pm interface {
	Gen() (string, error)
}

type pm struct {
	digits         bool
	lower          bool
	upper          bool
	special        bool
	size           int
	specialCharset string
}

func (p pm) Gen() (string, error) {
	var nDigits, nLower, nUpper, nSpecial, size int
	size = p.size
	if p.digits {
		nDigits = 1
	}
	if p.lower {
		nLower = 1
	}
	if p.upper {
		nUpper = 1
	}
	if p.special {
		nSpecial = 1
	}
	if size < 4 {
		size = 12
	}

	cmd := fmt.Sprintf("pm gen -d %d -c %d -C %d -x %d --specialCharset '%s' -s %d", nDigits, nLower, nUpper, nSpecial, p.specialCharset, size)
	log.Println(cmd)
	pipe := script.Exec(cmd)

	return pipe.String()
}

func NewDefaultMin12() Pm {
	n := 12 + rand.Intn(4)
	return New(true, true, true, true, n, "@#:.!*-")
}

func New(digits, lower, upper, special bool, size int, specialCharset string) Pm {
	return pm{
		digits:         digits,
		lower:          lower,
		upper:          upper,
		special:        special,
		size:           size,
		specialCharset: specialCharset,
	}
}

//
//Usage:
//pm gen [flags]
//
//Flags:
//-f, --force                   Usando esta opção, você pode substituir uma senha existente pela nova senha gerada.
//Se uma senha com o mesmo nome já existir, ela será automaticamente substituída pela nova senha.
//-h, --help                    help for gen
//-i, --insert                  Ao usar essa opção, uma nova senha será gerada e adicionada à sua coleção de senhas utilizando o comando 'pass'
//--letterCharset string    Letras maísculas e minúsculas (default "abcdefghijklmnopqrstuvxzwy")
//-d, --minDigits int           Número mínimo de dígitos (default 2)
//-c, --minLowercase int        Número mínimo de letras minúsculas (default 2)
//-x, --minSpecials int         Número mínimo caracteres especiais (default 2)
//-C, --minUppercase int        Número mínimo de letras maíusculas (default 2)
//--numberCharset string    Números (default "0123456789")
//-s, --size int                Tamanho da senha (default 12)
//--specialCharset string   Caracteres especiais (default "@#$:.!*-")
//-t, --teste                   flag para testes
