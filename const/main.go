package main

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/encoding/unicode"
)

const (
	a int = iota
	b
	c
	d
	e = iota + 6000
	f
	g
)

var ts string

func testGet() string {
	return ts
}
func main() {

	s12 := "sdfsdf"
	fmt.Println(s12)

	fmt.Println("lhb==", testGet())

	fmt.Println(a, b, c, d, e, f, g)

	switch a {
	case 1, 2, 3:
		fmt.Println("sdfsdf")
	default:
		fmt.Println("sfsdf")
		fmt.Println("====end======")
	}
	ab, errMsg := json.Marshal("{name:asdf]")
	fmt.Println(ab, errMsg)
	var ccc []byte
	fmt.Println(len(ccc))

	ii := 3.14
	fmt.Println(math.Trunc(ii))
	fmt.Println(md5.Sum([]byte{1, 2, 3}))
	fmt.Println(len("08-6808000249F0000F4240"))
	//fmt.Println(strconv.Itoa(3.14))
	fmt.Println(len(strings.Split("imsi-450051230000001-69", "-")))
	fmt.Println(hex.DecodeString("002710"))
	pwd := "abc123世界"
	encoding, err := unicode.UTF8.NewEncoder().String(pwd)
	fmt.Println(encoding, err)
	s := string([]rune{0xf8})
	fmt.Println(len(s))
	fmt.Println(utf8.RuneError)
	s2 := "eggo世界"
	a, b := utf8.DecodeRuneInString(s2)
	fmt.Println(s[0])     //101
	fmt.Println(a)        //101
	fmt.Println(b)        //1
	fmt.Printf("%b\n", a) //1100101

	bb := make([]int, 10)
	dst := bb
	dst = append(dst, 1)
	dst = append(dst, 2)
	dst = append(dst, 3)
	dst = dst[3:]
	dst = append(dst, 4)
	dst = append(dst, 5)
	dst = append(dst, 6)
	dst = dst[7:]
	dst = append(dst, 7)
	dst = append(dst, 8)
	dst = append(dst, 9)
	//dst = dst[11:]
	fmt.Println(bb, dst, dst[0], dst[1])

	e, _ := strconv.ParseUint("8501c03", 16, 32)
	e1, _ := strconv.ParseUint("8501c0389", 16, 36)
	fmt.Println("sdfsd==sdfdf==")
	fmt.Println(fmt.Sprintf("%.36b", e1))
	fmt.Println(strconv.FormatInt(int64(e1), 16))
	//fmt.Println(strconv.FormatInt(int64(e), 16))
	bNcgi := make([]byte, 4)
	//bNcgi[0] = byte((e << 4) >> 32)
	//bNcgi[1] = byte((e << 4) >> 24)
	//bNcgi[2] = byte((e << 4) >> 16)
	//bNcgi[3] = byte((e << 4) >> 8)
	//bNcgi[4] = byte(e << 4)
	//fmt.Println(hex.EncodeToString(bNcgi))
	//bEcgi := make([]byte, 4)
	binary.BigEndian.PutUint32(bNcgi, uint32(e))
	fmt.Println(hex.EncodeToString(bNcgi))
	//getBytes, _ := GetBytes(uint32(e))
	//fmt.Println(getBytes)
	//fmt.Println(binary.BigEndian.Uint32(getBytes))
	fmt.Println("========")
	fmt.Println([]byte("8501c0389"))
	fmt.Println(fmt.Sprintf("%.10x", 1000))
	fmt.Println(fmt.Sprintf("%.9x", e1))
	fmt.Println(e1)

	fmt.Println(strconv.FormatInt(int64(e1), 16))

	fmt.Println(1 << 4)
	//fmt.Println(hex.DecodeString("08501c03"))
	//
	//fmt.Println(filepath.Dir("./ca/ica"))
	//fmt.Println(regexp.MatchString("^[A-Fa-f0-9]{7}$", "139467779"))

	var ass uint32 = 2
	fmt.Println(GetBytes(&ass))

	fmt.Println("====ss====")
	sss := "8501c0389"
	ess, _ := strconv.ParseUint(sss, 16, 64)
	fmt.Println(strconv.FormatInt(int64(ess), 2))
	eb := make([]byte, 5)
	eb[0] = byte(ess>>32) & 0x0F
	eb[1] = byte(ess >> 24)
	eb[2] = byte(ess >> 16)
	eb[3] = byte(ess >> 8)
	eb[4] = byte(ess)

	fmt.Println(eb)
	fmt.Println(hex.EncodeToString(eb))

	sss = "8501c03"
	ess, _ = strconv.ParseUint(sss, 16, 32)
	fmt.Println(strconv.FormatInt(int64(ess), 2))
	eb = make([]byte, 4)
	eb[0] = byte(ess>>24) & 0x0F
	eb[1] = byte(ess >> 16)
	eb[2] = byte(ess >> 8)
	eb[3] = byte(ess)

	fmt.Println(eb)
	fmt.Println(hex.EncodeToString(eb))

}
func GetBytes(val *uint32) ([]byte, error) {
	if val != nil {
		a := *val
		fmt.Println("====", a)
		return nil, nil
	}
	bytesBuffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(bytesBuffer, binary.BigEndian, val); err != nil {
		return nil, err
	}
	return bytesBuffer.Bytes(), nil
}
