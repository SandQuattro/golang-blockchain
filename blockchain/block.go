package blockchain

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Block struct {
	payload []byte
	// previous block hash, will use it for integrity check
	previousHash string
	hash         string
	timestamp    int64
	// proof-of-work computation
	pow int
}

type BlockOperations interface {
	encodeData(data any) error
	decodeData(result any) error
	setHash()
	calculateHash() string
	mine(difficulty int)
}

func (b *Block) setHash() {
	b.hash = b.calculateHash()
}

// Вычисление хеша текущего блока
func (b *Block) calculateHash() string {
	var buf []byte
	buf = append(buf, b.previousHash...)
	buf = append(buf, b.payload...)
	buf = binary.LittleEndian.AppendUint64(buf, uint64(b.timestamp))
	buf = append(buf, strconv.Itoa(b.pow)...)
	return fmt.Sprintf("%x", sha256.Sum256(buf))
}

// В зависимости от установленной сложности на блокчейне выполняем вычисление хеша
// проверяем, что хеш начинается с нужного количества нулей
// если нет, то продолжаем выполнение хеширования
// pow показывает количество попыток, которые были произведены для получения hash текущего блока
func (b *Block) mine(difficulty int) {
	if difficulty == 0 {
		b.pow = 0
		b.setHash()
		return
	}
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.setHash()
	}
}

func (b *Block) encodeData(data any) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	b.payload = bytes
	return nil
}

func (b *Block) decodeData(result any) error {
	err := json.Unmarshal(b.payload, &result)
	if err != nil {
		return err
	}
	return nil
}
