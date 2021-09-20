package sample

import (
	"gRPCDemo/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

)

func NewKeyboard() *pb.Keyboard {
	return &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
}

func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	numberCores := randomCPUNumberCores(2, 8)
	numberThreads := 2 * numberCores
	minGhz := randomFloat64(2.0, 4.0)
	maxGhz := randomFloat64(minGhz, 6.0)
	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}
	return cpu
}

func NewGPU() *pb.GPU {
	brand := randomGPUBroad()
	name := randomGPUName(brand)
	minGhz := randomFloat64(1.0, 1.50)
	maxGhz := randomFloat64(minGhz, 2.5)
	memory := &pb.Memory{
		Value: uint64(randomInt(2, 6)),
		Uint:  pb.Memory_GIGABYTE,
	}
	return &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}
}
func NewRAM() *pb.Memory {
	return &pb.Memory{
		Value: uint64(randomInt(4, 128)),
		Uint:  pb.Memory_GIGABYTE,
	}
}

func NewSSD() *pb.Storage {
	return &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(512, 1024)),
			Uint:  pb.Memory_GIGABYTE,
		},
	}
}

func NewHDD() *pb.Storage {
	return &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1,16)),
			Uint:  pb.Memory_TERABYTE,
		},
	}
}
func NewScreen() *pb.Screen {
	return &pb.Screen{
		SizeInch: randomFloat32(13,17),
		Resolution: randomScreenResolution(),
		Panel: randomScreenPanel(),
		MultiTouch: randomBool(),
	}
}

func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)
	laptop := &pb.Laptop{
		Id: randomID(),
		Brand: brand,
		Name: name,
		Cpu: NewCPU(),
		Memory: NewRAM(),
		Gpu: []*pb.GPU{NewGPU()},
		Storage: []*pb.Storage{NewSSD(), NewHDD()},
		Screen: NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUSD: randomFloat64(1500, 3000),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdateAt: timestamppb.New(time.Time{}),
	}
	return laptop
}