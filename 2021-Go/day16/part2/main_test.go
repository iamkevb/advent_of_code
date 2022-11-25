package main

import (
	"testing"
)

func TestToBinaryString(t *testing.T) {
	s := "D2FE28"
	binString := toBinaryString([]byte(s))
	if binString != "110100101111111000101000" {
		t.Errorf("binary string %s", binString)
	}
	s = "38006F45291200"
	binString = toBinaryString([]byte(s))
	if binString != "00111000000000000110111101000101001010010001001000000000" {
		t.Errorf("binary string %s", binString)
	}
	s = "EE00D40C823060"
	binString = toBinaryString([]byte(s))
	if binString != "11101110000000001101010000001100100000100011000001100000" {
		t.Errorf("binary string %s", binString)
	}
}

func TestLiteralPacket(t *testing.T) {
	string := "110100101111111000101000"
	p, s := readPacket(string)
	if p.version != 6 || p.typeID != 4 || p.value != 2021 {
		t.Errorf("Version %d, type %d, value %d", p.version, p.typeID, p.value)
	}
	if s != "000" {
		t.Errorf("s: %s", s)
	}
}

func TestOperatorPacket(t *testing.T) {
	string := "00111000000000000110111101000101001010010001001000000000"
	p, s := readPacket(string)
	if p.version != 1 {
		t.Errorf("Version %d", p.version)
	}
	if p.typeID != 6 {
		t.Errorf("TypeID %d", p.typeID)
	}
	if p.value != 1 {
		t.Errorf("Value %d", p.value)
	}
	if len(p.packets) != 2 {
		t.Errorf("len packets %d", len(p.packets))
	}
	if len(p.packets) >= 2 {
		pp := p.packets[0]
		if pp.version != 6 || pp.typeID != 4 || pp.value != 10 {
			t.Errorf("Version %d, type %d, value %d", pp.version, pp.typeID, pp.value)
		}
		pp = p.packets[1]
		if pp.version != 2 || pp.typeID != 4 || pp.value != 20 {
			t.Errorf("Version %d, type %d, value %d", pp.version, pp.typeID, pp.value)
		}
	}
	if s != "0000000" {
		t.Errorf("s %s", s)
	}
}

func TestOperatorTurtles(t *testing.T) {
	hexString := "8A004A801A8002F478"
	binString := toBinaryString([]byte(hexString))
	p, _ := readPacket(binString)
	if p.typeID != 2 {
		t.Errorf("Expected type 2, got: %d", p.typeID)
	}
	if p.version != 4 {
		t.Errorf("Expected version 4, got %d", p.version)
	}
	if len(p.packets) != 1 {
		t.Fatal("Expected 1 subpacket")
	}
	sp := p.packets[0]
	if sp.version != 1 {
		t.Errorf("Expected version 1, got: %d", sp.version)
	}
	if len(sp.packets) != 1 {
		t.Fatal("Expected 1 subpacket in subpacket 2")
	}
	sp = sp.packets[0]
	if sp.version != 5 {
		t.Errorf("Expected version 5, got: %d", sp.version)
	}
	if len(sp.packets) != 1 {
		t.Fatal("Expected 1 subpacket in subpacket 3")
	}
	sp = sp.packets[0]
	if sp.version != 6 {
		t.Errorf("Expected version 6, got: %d", sp.version)
	}
	if sp.typeID != 4 {
		t.Errorf("Expected type 6, got: %d", sp.typeID)
	}
}

func TestOperatorWith2OperatorsWith2Literals(t *testing.T) {
	hexString := "620080001611562C8802118E34"
	binString := toBinaryString([]byte(hexString))
	versionSum := 0
	p, _ := readPacket(binString)
	versionSum += p.version
	if p.typeID == 4 {
		t.Error("Expected operator packet, got literal")
	}
	if p.version != 3 {
		t.Errorf("Expected version 4, got %d", p.version)
	}
	if len(p.packets) != 2 {
		t.Fatal("Expected 2 subpacket")
	}
	for _, sp := range p.packets {
		if sp.typeID == 4 {
			t.Error("Expected operator packet, got literal")
		}
		if len(sp.packets) != 2 {
			t.Fatal("Expected 2 subpacket in subpacket")
		}
		versionSum += sp.version
		for _, ssp := range sp.packets {
			if ssp.typeID != 4 {
				t.Error("Expected literal packet, got operator")
			}
			versionSum += ssp.version
		}
	}

	if versionSum != 12 {
		t.Errorf("Expected 12, got %d", versionSum)
	}
}

func TestSumVersions(t *testing.T) {
	s := "C0015000016115A2E0802F182340"
	binString := toBinaryString([]byte(s))
	p, _ := readPacket(binString)
	versionSum := sumVersions([]Packet{*p}, 0)
	if versionSum != 23 {
		t.Errorf("Expected 23 got %d", versionSum)
	}

}

func computePacketValue(s string) int {
	binString := toBinaryString([]byte(s))
	p, _ := readPacket(binString)
	return computeValue(*p)
}

func TestComputeValue(t *testing.T) {
	if computePacketValue("C200B40A82") != 3 {
		t.Error("Sum thing is wrong")
	}
	if computePacketValue("04005AC33890") != 54 {
		t.Error("Not being productived")
	}
	if v := computePacketValue("880086C3E88112"); v != 7 {
		t.Errorf("Minimize time spent on aoc %d", v)
	}
	if computePacketValue("CE00C43D881120") != 9 {
		t.Error("Maximize your time spent with doggy")
	}
	if computePacketValue("D8005AC2A8F0") != 1 {
		t.Error("Less than optimal")
	}
	if computePacketValue("F600BC2D8F") != 0 {
		t.Error("not great")
	}
	if computePacketValue("9C005AC2F8F0") != 0 {
		t.Error("not equal")
	}
	if computePacketValue("9C0141080250320F1802104A08") != 1 {
		t.Error("math is hard")
	}
}
