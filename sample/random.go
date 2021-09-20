package sample

import (
	"gRPCDemo/pb"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func randomKeyboardLayout() pb.Keyboard_Layout {
	rand.Seed(time.Now().UnixNano())
	switch rand.Intn(4) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	case 3:
		return pb.Keyboard_AZERTY
	default:
		return pb.Keyboard_UNKNOWN
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomStringFromSet(s ...string) string {
	n := len(s)
	return s[rand.Intn(n)]
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i7-9750H",
			"Core i5-9400F",
			"Core i3-1005G1",
		)
	}
	return randomStringFromSet(
		"Ryzen 7 PRO 2700U",
		"Ryzen 5 PRO 3500U",
		"Ryzen 3 PRO 3200GE",
	)
}

func randomCPUNumberCores(nums ...int) int {

	return rand.Intn(nums[1]-nums[0]+1) + nums[0]
}
func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomGPUBroad() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1660Ti",
			"GTX 1070",
		)
	}
	return randomStringFromSet(
		"RX 590",
		"RX 580",
		"RX 5700-XT",
		"RX Vega-5600",
	)
}
func randomScreenResolution() *pb.Screen_Resoultion {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	return &pb.Screen_Resoultion{
		Height: uint32(height),
		Width:  uint32(width),
	}
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_OLED
	}
	return pb.Screen_IPS
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "Lenovo")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook Air", "Macbook pro")
	case "Dell":
		return randomStringFromSet("Latitude", "Vostro", "XPS", "Alienware")
	default:
		return randomStringFromSet("Thinkpad X1", "Thinkpad P1", "Thinkpad PS3")
	}
}
func randomID() string {
	return uuid.New().String()
}
