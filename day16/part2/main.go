package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type Packet struct {
	version int
	typeID  int
	value   int64
	packets []Packet
}

func main() {
	path := "./input.txt"
	f, _ := os.Open(path)
	bytes, _ := ioutil.ReadAll(f)
	binStr := toBinaryString(bytes)
	packets := []Packet{}
	for len(binStr) > 0 {
		p, str := readPacket(binStr)
		binStr = str
		if p != nil {
			packets = append(packets, *p)
		}
	}

	total := sumVersions(packets, 0)
	fmt.Println("Version sum: ", total)

	for _, p := range packets {
		fmt.Println("Value: ", p.value)
	}
}

func sumVersions(packets []Packet, total int) int {
	for _, p := range packets {
		total += p.version
		total = sumVersions(p.packets, total)
	}
	return total
}

func computeValue(p Packet) int64 {
	switch p.typeID {
	case 0:
		return sumValues(p)
	case 1:
		return multiplyValues(p)
	case 2:
		return minimumValue(p)
	case 3:
		return maximumValue(p)
	case 4:
		return p.value
	case 5:
		return greaterThan(p)
	case 6:
		return lessThan(p)
	case 7:
		return equal(p)
	}
	panic("fucked")
}

func equal(p Packet) int64 {
	if p.packets[0].value == p.packets[1].value {
		return 1
	}
	return 0
}

func lessThan(p Packet) int64 {
	if p.packets[0].value < p.packets[1].value {
		return 1
	}
	return 0
}

func greaterThan(p Packet) int64 {
	if p.packets[0].value > p.packets[1].value {
		return 1
	}
	return 0
}

func maximumValue(packet Packet) int64 {
	var max int64 = 0
	for _, p := range packet.packets {
		if p.value > max {
			max = p.value
		}
	}
	return max
}

func minimumValue(packet Packet) int64 {
	var min int64 = math.MaxInt
	for _, p := range packet.packets {
		if p.value < min {
			min = p.value
		}
	}
	return min
}

func multiplyValues(packet Packet) int64 {
	var product int64 = 1
	for _, p := range packet.packets {
		product = product * p.value
	}
	return product
}

func sumValues(packet Packet) int64 {
	var total int64 = 0
	for _, p := range packet.packets {
		total += p.value
	}
	return total
}

func readPacket(from string) (*Packet, string) {
	if len(from) < 6 {
		return nil, ""
	}
	vStr := from[:3]
	v, _ := strconv.ParseInt(vStr, 2, 64)
	tStr := from[3:6]
	t, _ := strconv.ParseInt(tStr, 2, 64)
	var value int64
	var s string = from[6:]
	var sub []Packet = []Packet{}
	switch t {
	case 4:
		value, s = parseValue(s)
	default:
		sub, s = readSubpackets(s)
	}

	packet := Packet{version: int(v), typeID: int(t), value: value, packets: sub}
	packet.value = computeValue(packet)

	return &packet, s
}

func readSubpackets(from string) ([]Packet, string) {
	lengthType := from[:1]

	var packets []Packet
	var s string
	switch lengthType {
	case "0":
		packets, s = readSubpacketsFromBytes(from[1:])
	case "1":
		packets, s = readSubpacketsFromCount(from[1:])
	}
	return packets, s
}
func readSubpacketsFromCount(from string) ([]Packet, string) {
	numSubpacketsStr := from[:11]
	numSubpackets, _ := strconv.ParseInt(numSubpacketsStr, 2, 64)

	packets := make([]Packet, int(numSubpackets))
	str := from[11:]

	for i := 0; i < len(packets); i++ {
		p, s := readPacket(str)
		if p != nil {
			packets[i] = *p
		}
		str = s
	}
	return packets, str
}
func readSubpacketsFromBytes(from string) ([]Packet, string) {
	numSubpacketByteStr := from[1:15]
	numSubpacketBytes, _ := strconv.ParseInt(numSubpacketByteStr, 2, 64)

	subpacketsStr := from[15 : 15+int(numSubpacketBytes)]
	remaining := from[15+int(numSubpacketBytes):]

	packets := []Packet{}

	for len(subpacketsStr) > 0 {
		p, str := readPacket(subpacketsStr)
		subpacketsStr = str
		if p != nil {
			packets = append(packets, *p)
		}
	}
	return packets, remaining
}

func parseValue(from string) (int64, string) {
	builder := strings.Builder{}
	for len(from) > 0 {
		digit := from[:5]
		from = from[5:]
		builder.WriteString(digit[1:])
		if digit[0] == '0' {
			break
		}
	}
	v, _ := strconv.ParseInt(builder.String(), 2, 64)
	return v, from
}

func toBinaryString(file []byte) string {
	input := string(file)
	builder := strings.Builder{}
	for _, b := range input {
		builder.WriteString(hexToBits(b))
	}
	return builder.String()
}

func hexToBits(b rune) string {
	switch b {
	case '0':
		return "0000"
	case '1':
		return "0001"
	case '2':
		return "0010"
	case '3':
		return "0011"
	case '4':
		return "0100"
	case '5':
		return "0101"
	case '6':
		return "0110"
	case '7':
		return "0111"
	case '8':
		return "1000"
	case '9':
		return "1001"
	case 'A':
		return "1010"
	case 'B':
		return "1011"
	case 'C':
		return "1100"
	case 'D':
		return "1101"
	case 'E':
		return "1110"
	case 'F':
		return "1111"
	}
	panic("ain't hex")
}
