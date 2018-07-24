package FIR

func acc(y0 <-chan int,
	y1 <-chan int,
	y2 <-chan int,
	y3 <-chan int,
	y4 <-chan int,
	y5 <-chan int,
	y6 <-chan int,
	y7 <-chan int,
	y8 <-chan int,
	y9 <-chan int,
	y10 <-chan int,
	y11 <-chan int,
	y12 <-chan int,
	y13 <-chan int,
	y14 <-chan int,
	y15 <-chan int,
	y16 <-chan int,
	y17 <-chan int,
	y18 <-chan int,
	y19 <-chan int,
	y20 <-chan int,
	y21 <-chan int,
	y22 <-chan int,
	y23 <-chan int,
	y24 <-chan int,
	y25 <-chan int,
	y26 <-chan int,
	y27 <-chan int,
	y28 <-chan int,
	y29 <-chan int,
	y30 <-chan int,
	y31 <-chan int,
	y32 <-chan int,
	y33 <-chan int,
	y34 <-chan int,
	y35 <-chan int,
	y36 <-chan int,
	y37 <-chan int,
	y38 <-chan int,
	y39 <-chan int,
	y40 <-chan int,
	y41 <-chan int,
	y42 <-chan int,
	y43 <-chan int,
	y44 <-chan int,
	y45 <-chan int,
	y46 <-chan int,
	y47 <-chan int,
	y48 <-chan int,
	y49 <-chan int,
	y50 <-chan int,
	y51 <-chan int,
	y52 <-chan int,
	y53 <-chan int,
	y54 <-chan int,
	y55 <-chan int,
	y56 <-chan int,
	y57 <-chan int,
	y58 <-chan int,
	y59 <-chan int,
	y60 <-chan int,
	y61 <-chan int,
	y62 <-chan int,
	y63 <-chan int,
	y64 <-chan int,
	y65 <-chan int,
	y66 <-chan int,
	y67 <-chan int,
	y68 <-chan int,
	y69 <-chan int,
	y70 <-chan int,
	y71 <-chan int,
	y72 <-chan int,
	y73 <-chan int,
	y74 <-chan int,
	y75 <-chan int,
	y76 <-chan int,
	y77 <-chan int,
	y78 <-chan int,
	y79 <-chan int,
	y80 <-chan int,
	y81 <-chan int,
	y82 <-chan int,
	y83 <-chan int,
	y84 <-chan int,
	y85 <-chan int,
	y86 <-chan int,
	y87 <-chan int,
	y88 <-chan int,
	y89 <-chan int,
	y90 <-chan int,
	y91 <-chan int,
	y92 <-chan int,
	y93 <-chan int,
	y94 <-chan int,
	y95 <-chan int,
	y96 <-chan int,
	y97 <-chan int,
	y98 <-chan int,
	y99 <-chan int,
	y chan<- uint32){
	for{
		res0 := <-y0
		res1 := <-y1
		res2 := <-y2
		res3 := <-y3
		res4 := <-y4
		res5 := <-y5
		res6 := <-y6
		res7 := <-y7
		res8 := <-y8
		res9 := <-y9
		res10 := <-y10
		res11 := <-y11
		res12 := <-y12
		res13 := <-y13
		res14 := <-y14
		res15 := <-y15
		res16 := <-y16
		res17 := <-y17
		res18 := <-y18
		res19 := <-y19
		res20 := <-y20
		res21 := <-y21
		res22 := <-y22
		res23 := <-y23
		res24 := <-y24
		res25 := <-y25
		res26 := <-y26
		res27 := <-y27
		res28 := <-y28
		res29 := <-y29
		res30 := <-y30
		res31 := <-y31
		res32 := <-y32
		res33 := <-y33
		res34 := <-y34
		res35 := <-y35
		res36 := <-y36
		res37 := <-y37
		res38 := <-y38
		res39 := <-y39
		res40 := <-y40
		res41 := <-y41
		res42 := <-y42
		res43 := <-y43
		res44 := <-y44
		res45 := <-y45
		res46 := <-y46
		res47 := <-y47
		res48 := <-y48
		res49 := <-y49
		res50 := <-y50
		res51 := <-y51
		res52 := <-y52
		res53 := <-y53
		res54 := <-y54
		res55 := <-y55
		res56 := <-y56
		res57 := <-y57
		res58 := <-y58
		res59 := <-y59
		res60 := <-y60
		res61 := <-y61
		res62 := <-y62
		res63 := <-y63
		res64 := <-y64
		res65 := <-y65
		res66 := <-y66
		res67 := <-y67
		res68 := <-y68
		res69 := <-y69
		res70 := <-y70
		res71 := <-y71
		res72 := <-y72
		res73 := <-y73
		res74 := <-y74
		res75 := <-y75
		res76 := <-y76
		res77 := <-y77
		res78 := <-y78
		res79 := <-y79
		res80 := <-y80
		res81 := <-y81
		res82 := <-y82
		res83 := <-y83
		res84 := <-y84
		res85 := <-y85
		res86 := <-y86
		res87 := <-y87
		res88 := <-y88
		res89 := <-y89
		res90 := <-y90
		res91 := <-y91
		res92 := <-y92
		res93 := <-y93
		res94 := <-y94
		res95 := <-y95
		res96 := <-y96
		res97 := <-y97
		res98 := <-y98
		res99 := <-y99
		y <- uint32(res0 + res1 + res2 + res3 + res4 + res5 + res6 + res7 + res8 + res9 + res10 + res11 + res12 + res13 + res14 + res15 + res16 + res17 + res18 + res19 + res20 + res21 + res22 + res23 + res24 + res25 + res26 + res27 + res28 + res29 + res30 + res31 + res32 + res33 + res34 + res35 + res36 + res37 + res38 + res39 + res40 + res41 + res42 + res43 + res44 + res45 + res46 + res47 + res48 + res49 + res50 + res51 + res52 + res53 + res54 + res55 + res56 + res57 + res58 + res59 + res60 + res61 + res62 + res63 + res64 + res65 + res66 + res67 + res68 + res69 + res70 + res71 + res72 + res73 + res74 + res75 + res76 + res77 + res78 + res79 + res80 + res81 + res82 + res83 + res84 + res85 + res86 + res87 + res88 + res89 + res90 + res91 + res92 + res93 + res94 + res95 + res96 + res97 + res98 + res99)
	}
}
func fir0(x0 <-chan int, y0 chan<- int){
	for{
		temp := int(<-x0)
		y0 <- temp*2
	}
}
func fir1(x1 <-chan int, y1 chan<- int){
	for{
		temp := int(<-x1)
		y1 <- temp*8
	}
}
func fir2(x2 <-chan int, y2 chan<- int){
	for{
		temp := int(<-x2)
		y2 <- temp*5
	}
}
func fir3(x3 <-chan int, y3 chan<- int){
	for{
		temp := int(<-x3)
		y3 <- temp*1
	}
}
func fir4(x4 <-chan int, y4 chan<- int){
	for{
		temp := int(<-x4)
		y4 <- temp*10
	}
}
func fir5(x5 <-chan int, y5 chan<- int){
	for{
		temp := int(<-x5)
		y5 <- temp*5
	}
}
func fir6(x6 <-chan int, y6 chan<- int){
	for{
		temp := int(<-x6)
		y6 <- temp*9
	}
}
func fir7(x7 <-chan int, y7 chan<- int){
	for{
		temp := int(<-x7)
		y7 <- temp*9
	}
}
func fir8(x8 <-chan int, y8 chan<- int){
	for{
		temp := int(<-x8)
		y8 <- temp*3
	}
}
func fir9(x9 <-chan int, y9 chan<- int){
	for{
		temp := int(<-x9)
		y9 <- temp*5
	}
}
func fir10(x10 <-chan int, y10 chan<- int){
	for{
		temp := int(<-x10)
		y10 <- temp*6
	}
}
func fir11(x11 <-chan int, y11 chan<- int){
	for{
		temp := int(<-x11)
		y11 <- temp*6
	}
}
func fir12(x12 <-chan int, y12 chan<- int){
	for{
		temp := int(<-x12)
		y12 <- temp*2
	}
}
func fir13(x13 <-chan int, y13 chan<- int){
	for{
		temp := int(<-x13)
		y13 <- temp*8
	}
}
func fir14(x14 <-chan int, y14 chan<- int){
	for{
		temp := int(<-x14)
		y14 <- temp*2
	}
}
func fir15(x15 <-chan int, y15 chan<- int){
	for{
		temp := int(<-x15)
		y15 <- temp*2
	}
}
func fir16(x16 <-chan int, y16 chan<- int){
	for{
		temp := int(<-x16)
		y16 <- temp*6
	}
}
func fir17(x17 <-chan int, y17 chan<- int){
	for{
		temp := int(<-x17)
		y17 <- temp*3
	}
}
func fir18(x18 <-chan int, y18 chan<- int){
	for{
		temp := int(<-x18)
		y18 <- temp*8
	}
}
func fir19(x19 <-chan int, y19 chan<- int){
	for{
		temp := int(<-x19)
		y19 <- temp*7
	}
}
func fir20(x20 <-chan int, y20 chan<- int){
	for{
		temp := int(<-x20)
		y20 <- temp*2
	}
}
func fir21(x21 <-chan int, y21 chan<- int){
	for{
		temp := int(<-x21)
		y21 <- temp*5
	}
}
func fir22(x22 <-chan int, y22 chan<- int){
	for{
		temp := int(<-x22)
		y22 <- temp*3
	}
}
func fir23(x23 <-chan int, y23 chan<- int){
	for{
		temp := int(<-x23)
		y23 <- temp*4
	}
}
func fir24(x24 <-chan int, y24 chan<- int){
	for{
		temp := int(<-x24)
		y24 <- temp*3
	}
}
func fir25(x25 <-chan int, y25 chan<- int){
	for{
		temp := int(<-x25)
		y25 <- temp*3
	}
}
func fir26(x26 <-chan int, y26 chan<- int){
	for{
		temp := int(<-x26)
		y26 <- temp*2
	}
}
func fir27(x27 <-chan int, y27 chan<- int){
	for{
		temp := int(<-x27)
		y27 <- temp*7
	}
}
func fir28(x28 <-chan int, y28 chan<- int){
	for{
		temp := int(<-x28)
		y28 <- temp*9
	}
}
func fir29(x29 <-chan int, y29 chan<- int){
	for{
		temp := int(<-x29)
		y29 <- temp*6
	}
}
func fir30(x30 <-chan int, y30 chan<- int){
	for{
		temp := int(<-x30)
		y30 <- temp*8
	}
}
func fir31(x31 <-chan int, y31 chan<- int){
	for{
		temp := int(<-x31)
		y31 <- temp*7
	}
}
func fir32(x32 <-chan int, y32 chan<- int){
	for{
		temp := int(<-x32)
		y32 <- temp*2
	}
}
func fir33(x33 <-chan int, y33 chan<- int){
	for{
		temp := int(<-x33)
		y33 <- temp*9
	}
}
func fir34(x34 <-chan int, y34 chan<- int){
	for{
		temp := int(<-x34)
		y34 <- temp*10
	}
}
func fir35(x35 <-chan int, y35 chan<- int){
	for{
		temp := int(<-x35)
		y35 <- temp*3
	}
}
func fir36(x36 <-chan int, y36 chan<- int){
	for{
		temp := int(<-x36)
		y36 <- temp*8
	}
}
func fir37(x37 <-chan int, y37 chan<- int){
	for{
		temp := int(<-x37)
		y37 <- temp*10
	}
}
func fir38(x38 <-chan int, y38 chan<- int){
	for{
		temp := int(<-x38)
		y38 <- temp*6
	}
}
func fir39(x39 <-chan int, y39 chan<- int){
	for{
		temp := int(<-x39)
		y39 <- temp*5
	}
}
func fir40(x40 <-chan int, y40 chan<- int){
	for{
		temp := int(<-x40)
		y40 <- temp*4
	}
}
func fir41(x41 <-chan int, y41 chan<- int){
	for{
		temp := int(<-x41)
		y41 <- temp*2
	}
}
func fir42(x42 <-chan int, y42 chan<- int){
	for{
		temp := int(<-x42)
		y42 <- temp*3
	}
}
func fir43(x43 <-chan int, y43 chan<- int){
	for{
		temp := int(<-x43)
		y43 <- temp*4
	}
}
func fir44(x44 <-chan int, y44 chan<- int){
	for{
		temp := int(<-x44)
		y44 <- temp*4
	}
}
func fir45(x45 <-chan int, y45 chan<- int){
	for{
		temp := int(<-x45)
		y45 <- temp*5
	}
}
func fir46(x46 <-chan int, y46 chan<- int){
	for{
		temp := int(<-x46)
		y46 <- temp*2
	}
}
func fir47(x47 <-chan int, y47 chan<- int){
	for{
		temp := int(<-x47)
		y47 <- temp*2
	}
}
func fir48(x48 <-chan int, y48 chan<- int){
	for{
		temp := int(<-x48)
		y48 <- temp*4
	}
}
func fir49(x49 <-chan int, y49 chan<- int){
	for{
		temp := int(<-x49)
		y49 <- temp*9
	}
}
func fir50(x50 <-chan int, y50 chan<- int){
	for{
		temp := int(<-x50)
		y50 <- temp*8
	}
}
func fir51(x51 <-chan int, y51 chan<- int){
	for{
		temp := int(<-x51)
		y51 <- temp*5
	}
}
func fir52(x52 <-chan int, y52 chan<- int){
	for{
		temp := int(<-x52)
		y52 <- temp*3
	}
}
func fir53(x53 <-chan int, y53 chan<- int){
	for{
		temp := int(<-x53)
		y53 <- temp*8
	}
}
func fir54(x54 <-chan int, y54 chan<- int){
	for{
		temp := int(<-x54)
		y54 <- temp*8
	}
}
func fir55(x55 <-chan int, y55 chan<- int){
	for{
		temp := int(<-x55)
		y55 <- temp*10
	}
}
func fir56(x56 <-chan int, y56 chan<- int){
	for{
		temp := int(<-x56)
		y56 <- temp*4
	}
}
func fir57(x57 <-chan int, y57 chan<- int){
	for{
		temp := int(<-x57)
		y57 <- temp*2
	}
}
func fir58(x58 <-chan int, y58 chan<- int){
	for{
		temp := int(<-x58)
		y58 <- temp*10
	}
}
func fir59(x59 <-chan int, y59 chan<- int){
	for{
		temp := int(<-x59)
		y59 <- temp*9
	}
}
func fir60(x60 <-chan int, y60 chan<- int){
	for{
		temp := int(<-x60)
		y60 <- temp*7
	}
}
func fir61(x61 <-chan int, y61 chan<- int){
	for{
		temp := int(<-x61)
		y61 <- temp*6
	}
}
func fir62(x62 <-chan int, y62 chan<- int){
	for{
		temp := int(<-x62)
		y62 <- temp*1
	}
}
func fir63(x63 <-chan int, y63 chan<- int){
	for{
		temp := int(<-x63)
		y63 <- temp*3
	}
}
func fir64(x64 <-chan int, y64 chan<- int){
	for{
		temp := int(<-x64)
		y64 <- temp*9
	}
}
func fir65(x65 <-chan int, y65 chan<- int){
	for{
		temp := int(<-x65)
		y65 <- temp*7
	}
}
func fir66(x66 <-chan int, y66 chan<- int){
	for{
		temp := int(<-x66)
		y66 <- temp*1
	}
}
func fir67(x67 <-chan int, y67 chan<- int){
	for{
		temp := int(<-x67)
		y67 <- temp*3
	}
}
func fir68(x68 <-chan int, y68 chan<- int){
	for{
		temp := int(<-x68)
		y68 <- temp*5
	}
}
func fir69(x69 <-chan int, y69 chan<- int){
	for{
		temp := int(<-x69)
		y69 <- temp*9
	}
}
func fir70(x70 <-chan int, y70 chan<- int){
	for{
		temp := int(<-x70)
		y70 <- temp*7
	}
}
func fir71(x71 <-chan int, y71 chan<- int){
	for{
		temp := int(<-x71)
		y71 <- temp*6
	}
}
func fir72(x72 <-chan int, y72 chan<- int){
	for{
		temp := int(<-x72)
		y72 <- temp*1
	}
}
func fir73(x73 <-chan int, y73 chan<- int){
	for{
		temp := int(<-x73)
		y73 <- temp*10
	}
}
func fir74(x74 <-chan int, y74 chan<- int){
	for{
		temp := int(<-x74)
		y74 <- temp*1
	}
}
func fir75(x75 <-chan int, y75 chan<- int){
	for{
		temp := int(<-x75)
		y75 <- temp*1
	}
}
func fir76(x76 <-chan int, y76 chan<- int){
	for{
		temp := int(<-x76)
		y76 <- temp*7
	}
}
func fir77(x77 <-chan int, y77 chan<- int){
	for{
		temp := int(<-x77)
		y77 <- temp*2
	}
}
func fir78(x78 <-chan int, y78 chan<- int){
	for{
		temp := int(<-x78)
		y78 <- temp*4
	}
}
func fir79(x79 <-chan int, y79 chan<- int){
	for{
		temp := int(<-x79)
		y79 <- temp*9
	}
}
func fir80(x80 <-chan int, y80 chan<- int){
	for{
		temp := int(<-x80)
		y80 <- temp*10
	}
}
func fir81(x81 <-chan int, y81 chan<- int){
	for{
		temp := int(<-x81)
		y81 <- temp*4
	}
}
func fir82(x82 <-chan int, y82 chan<- int){
	for{
		temp := int(<-x82)
		y82 <- temp*5
	}
}
func fir83(x83 <-chan int, y83 chan<- int){
	for{
		temp := int(<-x83)
		y83 <- temp*5
	}
}
func fir84(x84 <-chan int, y84 chan<- int){
	for{
		temp := int(<-x84)
		y84 <- temp*7
	}
}
func fir85(x85 <-chan int, y85 chan<- int){
	for{
		temp := int(<-x85)
		y85 <- temp*1
	}
}
func fir86(x86 <-chan int, y86 chan<- int){
	for{
		temp := int(<-x86)
		y86 <- temp*7
	}
}
func fir87(x87 <-chan int, y87 chan<- int){
	for{
		temp := int(<-x87)
		y87 <- temp*7
	}
}
func fir88(x88 <-chan int, y88 chan<- int){
	for{
		temp := int(<-x88)
		y88 <- temp*2
	}
}
func fir89(x89 <-chan int, y89 chan<- int){
	for{
		temp := int(<-x89)
		y89 <- temp*9
	}
}
func fir90(x90 <-chan int, y90 chan<- int){
	for{
		temp := int(<-x90)
		y90 <- temp*5
	}
}
func fir91(x91 <-chan int, y91 chan<- int){
	for{
		temp := int(<-x91)
		y91 <- temp*10
	}
}
func fir92(x92 <-chan int, y92 chan<- int){
	for{
		temp := int(<-x92)
		y92 <- temp*7
	}
}
func fir93(x93 <-chan int, y93 chan<- int){
	for{
		temp := int(<-x93)
		y93 <- temp*4
	}
}
func fir94(x94 <-chan int, y94 chan<- int){
	for{
		temp := int(<-x94)
		y94 <- temp*8
	}
}
func fir95(x95 <-chan int, y95 chan<- int){
	for{
		temp := int(<-x95)
		y95 <- temp*9
	}
}
func fir96(x96 <-chan int, y96 chan<- int){
	for{
		temp := int(<-x96)
		y96 <- temp*9
	}
}
func fir97(x97 <-chan int, y97 chan<- int){
	for{
		temp := int(<-x97)
		y97 <- temp*3
	}
}
func fir98(x98 <-chan int, y98 chan<- int){
	for{
		temp := int(<-x98)
		y98 <- temp*10
	}
}
func fir99(x99 <-chan int, y99 chan<- int){
	for{
		temp := int(<-x99)
		y99 <- temp*2
	}
}
func input_control(x0 chan<- int,
	x1 chan<- int,
	x2 chan<- int,
	x3 chan<- int,
	x4 chan<- int,
	x5 chan<- int,
	x6 chan<- int,
	x7 chan<- int,
	x8 chan<- int,
	x9 chan<- int,
	x10 chan<- int,
	x11 chan<- int,
	x12 chan<- int,
	x13 chan<- int,
	x14 chan<- int,
	x15 chan<- int,
	x16 chan<- int,
	x17 chan<- int,
	x18 chan<- int,
	x19 chan<- int,
	x20 chan<- int,
	x21 chan<- int,
	x22 chan<- int,
	x23 chan<- int,
	x24 chan<- int,
	x25 chan<- int,
	x26 chan<- int,
	x27 chan<- int,
	x28 chan<- int,
	x29 chan<- int,
	x30 chan<- int,
	x31 chan<- int,
	x32 chan<- int,
	x33 chan<- int,
	x34 chan<- int,
	x35 chan<- int,
	x36 chan<- int,
	x37 chan<- int,
	x38 chan<- int,
	x39 chan<- int,
	x40 chan<- int,
	x41 chan<- int,
	x42 chan<- int,
	x43 chan<- int,
	x44 chan<- int,
	x45 chan<- int,
	x46 chan<- int,
	x47 chan<- int,
	x48 chan<- int,
	x49 chan<- int,
	x50 chan<- int,
	x51 chan<- int,
	x52 chan<- int,
	x53 chan<- int,
	x54 chan<- int,
	x55 chan<- int,
	x56 chan<- int,
	x57 chan<- int,
	x58 chan<- int,
	x59 chan<- int,
	x60 chan<- int,
	x61 chan<- int,
	x62 chan<- int,
	x63 chan<- int,
	x64 chan<- int,
	x65 chan<- int,
	x66 chan<- int,
	x67 chan<- int,
	x68 chan<- int,
	x69 chan<- int,
	x70 chan<- int,
	x71 chan<- int,
	x72 chan<- int,
	x73 chan<- int,
	x74 chan<- int,
	x75 chan<- int,
	x76 chan<- int,
	x77 chan<- int,
	x78 chan<- int,
	x79 chan<- int,
	x80 chan<- int,
	x81 chan<- int,
	x82 chan<- int,
	x83 chan<- int,
	x84 chan<- int,
	x85 chan<- int,
	x86 chan<- int,
	x87 chan<- int,
	x88 chan<- int,
	x89 chan<- int,
	x90 chan<- int,
	x91 chan<- int,
	x92 chan<- int,
	x93 chan<- int,
	x94 chan<- int,
	x95 chan<- int,
	x96 chan<- int,
	x97 chan<- int,
	x98 chan<- int,
	x99 chan<- int,
	x <-chan uint32){
		X1 := 0
		X2 := 0
		X3 := 0
		X4 := 0
		X5 := 0
		X6 := 0
		X7 := 0
		X8 := 0
		X9 := 0
		X10 := 0
		X11 := 0
		X12 := 0
		X13 := 0
		X14 := 0
		X15 := 0
		X16 := 0
		X17 := 0
		X18 := 0
		X19 := 0
		X20 := 0
		X21 := 0
		X22 := 0
		X23 := 0
		X24 := 0
		X25 := 0
		X26 := 0
		X27 := 0
		X28 := 0
		X29 := 0
		X30 := 0
		X31 := 0
		X32 := 0
		X33 := 0
		X34 := 0
		X35 := 0
		X36 := 0
		X37 := 0
		X38 := 0
		X39 := 0
		X40 := 0
		X41 := 0
		X42 := 0
		X43 := 0
		X44 := 0
		X45 := 0
		X46 := 0
		X47 := 0
		X48 := 0
		X49 := 0
		X50 := 0
		X51 := 0
		X52 := 0
		X53 := 0
		X54 := 0
		X55 := 0
		X56 := 0
		X57 := 0
		X58 := 0
		X59 := 0
		X60 := 0
		X61 := 0
		X62 := 0
		X63 := 0
		X64 := 0
		X65 := 0
		X66 := 0
		X67 := 0
		X68 := 0
		X69 := 0
		X70 := 0
		X71 := 0
		X72 := 0
		X73 := 0
		X74 := 0
		X75 := 0
		X76 := 0
		X77 := 0
		X78 := 0
		X79 := 0
		X80 := 0
		X81 := 0
		X82 := 0
		X83 := 0
		X84 := 0
		X85 := 0
		X86 := 0
		X87 := 0
		X88 := 0
		X89 := 0
		X90 := 0
		X91 := 0
		X92 := 0
		X93 := 0
		X94 := 0
		X95 := 0
		X96 := 0
		X97 := 0
		X98 := 0
		X99 := 0
	for{
		temp := int(<-x)
		x0 <- temp
		x1 <- X1
		x2 <- X2
		x3 <- X3
		x4 <- X4
		x5 <- X5
		x6 <- X6
		x7 <- X7
		x8 <- X8
		x9 <- X9
		x10 <- X10
		x11 <- X11
		x12 <- X12
		x13 <- X13
		x14 <- X14
		x15 <- X15
		x16 <- X16
		x17 <- X17
		x18 <- X18
		x19 <- X19
		x20 <- X20
		x21 <- X21
		x22 <- X22
		x23 <- X23
		x24 <- X24
		x25 <- X25
		x26 <- X26
		x27 <- X27
		x28 <- X28
		x29 <- X29
		x30 <- X30
		x31 <- X31
		x32 <- X32
		x33 <- X33
		x34 <- X34
		x35 <- X35
		x36 <- X36
		x37 <- X37
		x38 <- X38
		x39 <- X39
		x40 <- X40
		x41 <- X41
		x42 <- X42
		x43 <- X43
		x44 <- X44
		x45 <- X45
		x46 <- X46
		x47 <- X47
		x48 <- X48
		x49 <- X49
		x50 <- X50
		x51 <- X51
		x52 <- X52
		x53 <- X53
		x54 <- X54
		x55 <- X55
		x56 <- X56
		x57 <- X57
		x58 <- X58
		x59 <- X59
		x60 <- X60
		x61 <- X61
		x62 <- X62
		x63 <- X63
		x64 <- X64
		x65 <- X65
		x66 <- X66
		x67 <- X67
		x68 <- X68
		x69 <- X69
		x70 <- X70
		x71 <- X71
		x72 <- X72
		x73 <- X73
		x74 <- X74
		x75 <- X75
		x76 <- X76
		x77 <- X77
		x78 <- X78
		x79 <- X79
		x80 <- X80
		x81 <- X81
		x82 <- X82
		x83 <- X83
		x84 <- X84
		x85 <- X85
		x86 <- X86
		x87 <- X87
		x88 <- X88
		x89 <- X89
		x90 <- X90
		x91 <- X91
		x92 <- X92
		x93 <- X93
		x94 <- X94
		x95 <- X95
		x96 <- X96
		x97 <- X97
		x98 <- X98
		x99 <- X99

		X99 = X98
		X98 = X97
		X97 = X96
		X96 = X95
		X95 = X94
		X94 = X93
		X93 = X92
		X92 = X91
		X91 = X90
		X90 = X89
		X89 = X88
		X88 = X87
		X87 = X86
		X86 = X85
		X85 = X84
		X84 = X83
		X83 = X82
		X82 = X81
		X81 = X80
		X80 = X79
		X79 = X78
		X78 = X77
		X77 = X76
		X76 = X75
		X75 = X74
		X74 = X73
		X73 = X72
		X72 = X71
		X71 = X70
		X70 = X69
		X69 = X68
		X68 = X67
		X67 = X66
		X66 = X65
		X65 = X64
		X64 = X63
		X63 = X62
		X62 = X61
		X61 = X60
		X60 = X59
		X59 = X58
		X58 = X57
		X57 = X56
		X56 = X55
		X55 = X54
		X54 = X53
		X53 = X52
		X52 = X51
		X51 = X50
		X50 = X49
		X49 = X48
		X48 = X47
		X47 = X46
		X46 = X45
		X45 = X44
		X44 = X43
		X43 = X42
		X42 = X41
		X41 = X40
		X40 = X39
		X39 = X38
		X38 = X37
		X37 = X36
		X36 = X35
		X35 = X34
		X34 = X33
		X33 = X32
		X32 = X31
		X31 = X30
		X30 = X29
		X29 = X28
		X28 = X27
		X27 = X26
		X26 = X25
		X25 = X24
		X24 = X23
		X23 = X22
		X22 = X21
		X21 = X20
		X20 = X19
		X19 = X18
		X18 = X17
		X17 = X16
		X16 = X15
		X15 = X14
		X14 = X13
		X13 = X12
		X12 = X11
		X11 = X10
		X10 = X9
		X9 = X8
		X8 = X7
		X7 = X6
		X6 = X5
		X5 = X4
		X4 = X3
		X3 = X2
		X2 = X1
		X1 = temp
	}
}
func FIR_Top(x <-chan uint32, y chan<- uint32){
	x0 := make(chan int, 2)
	y0 := make(chan int, 2)

	x1 := make(chan int, 2)
	y1 := make(chan int, 2)

	x2 := make(chan int, 2)
	y2 := make(chan int, 2)

	x3 := make(chan int, 2)
	y3 := make(chan int, 2)

	x4 := make(chan int, 2)
	y4 := make(chan int, 2)

	x5 := make(chan int, 2)
	y5 := make(chan int, 2)

	x6 := make(chan int, 2)
	y6 := make(chan int, 2)

	x7 := make(chan int, 2)
	y7 := make(chan int, 2)

	x8 := make(chan int, 2)
	y8 := make(chan int, 2)

	x9 := make(chan int, 2)
	y9 := make(chan int, 2)

	x10 := make(chan int, 2)
	y10 := make(chan int, 2)

	x11 := make(chan int, 2)
	y11 := make(chan int, 2)

	x12 := make(chan int, 2)
	y12 := make(chan int, 2)

	x13 := make(chan int, 2)
	y13 := make(chan int, 2)

	x14 := make(chan int, 2)
	y14 := make(chan int, 2)

	x15 := make(chan int, 2)
	y15 := make(chan int, 2)

	x16 := make(chan int, 2)
	y16 := make(chan int, 2)

	x17 := make(chan int, 2)
	y17 := make(chan int, 2)

	x18 := make(chan int, 2)
	y18 := make(chan int, 2)

	x19 := make(chan int, 2)
	y19 := make(chan int, 2)

	x20 := make(chan int, 2)
	y20 := make(chan int, 2)

	x21 := make(chan int, 2)
	y21 := make(chan int, 2)

	x22 := make(chan int, 2)
	y22 := make(chan int, 2)

	x23 := make(chan int, 2)
	y23 := make(chan int, 2)

	x24 := make(chan int, 2)
	y24 := make(chan int, 2)

	x25 := make(chan int, 2)
	y25 := make(chan int, 2)

	x26 := make(chan int, 2)
	y26 := make(chan int, 2)

	x27 := make(chan int, 2)
	y27 := make(chan int, 2)

	x28 := make(chan int, 2)
	y28 := make(chan int, 2)

	x29 := make(chan int, 2)
	y29 := make(chan int, 2)

	x30 := make(chan int, 2)
	y30 := make(chan int, 2)

	x31 := make(chan int, 2)
	y31 := make(chan int, 2)

	x32 := make(chan int, 2)
	y32 := make(chan int, 2)

	x33 := make(chan int, 2)
	y33 := make(chan int, 2)

	x34 := make(chan int, 2)
	y34 := make(chan int, 2)

	x35 := make(chan int, 2)
	y35 := make(chan int, 2)

	x36 := make(chan int, 2)
	y36 := make(chan int, 2)

	x37 := make(chan int, 2)
	y37 := make(chan int, 2)

	x38 := make(chan int, 2)
	y38 := make(chan int, 2)

	x39 := make(chan int, 2)
	y39 := make(chan int, 2)

	x40 := make(chan int, 2)
	y40 := make(chan int, 2)

	x41 := make(chan int, 2)
	y41 := make(chan int, 2)

	x42 := make(chan int, 2)
	y42 := make(chan int, 2)

	x43 := make(chan int, 2)
	y43 := make(chan int, 2)

	x44 := make(chan int, 2)
	y44 := make(chan int, 2)

	x45 := make(chan int, 2)
	y45 := make(chan int, 2)

	x46 := make(chan int, 2)
	y46 := make(chan int, 2)

	x47 := make(chan int, 2)
	y47 := make(chan int, 2)

	x48 := make(chan int, 2)
	y48 := make(chan int, 2)

	x49 := make(chan int, 2)
	y49 := make(chan int, 2)

	x50 := make(chan int, 2)
	y50 := make(chan int, 2)

	x51 := make(chan int, 2)
	y51 := make(chan int, 2)

	x52 := make(chan int, 2)
	y52 := make(chan int, 2)

	x53 := make(chan int, 2)
	y53 := make(chan int, 2)

	x54 := make(chan int, 2)
	y54 := make(chan int, 2)

	x55 := make(chan int, 2)
	y55 := make(chan int, 2)

	x56 := make(chan int, 2)
	y56 := make(chan int, 2)

	x57 := make(chan int, 2)
	y57 := make(chan int, 2)

	x58 := make(chan int, 2)
	y58 := make(chan int, 2)

	x59 := make(chan int, 2)
	y59 := make(chan int, 2)

	x60 := make(chan int, 2)
	y60 := make(chan int, 2)

	x61 := make(chan int, 2)
	y61 := make(chan int, 2)

	x62 := make(chan int, 2)
	y62 := make(chan int, 2)

	x63 := make(chan int, 2)
	y63 := make(chan int, 2)

	x64 := make(chan int, 2)
	y64 := make(chan int, 2)

	x65 := make(chan int, 2)
	y65 := make(chan int, 2)

	x66 := make(chan int, 2)
	y66 := make(chan int, 2)

	x67 := make(chan int, 2)
	y67 := make(chan int, 2)

	x68 := make(chan int, 2)
	y68 := make(chan int, 2)

	x69 := make(chan int, 2)
	y69 := make(chan int, 2)

	x70 := make(chan int, 2)
	y70 := make(chan int, 2)

	x71 := make(chan int, 2)
	y71 := make(chan int, 2)

	x72 := make(chan int, 2)
	y72 := make(chan int, 2)

	x73 := make(chan int, 2)
	y73 := make(chan int, 2)

	x74 := make(chan int, 2)
	y74 := make(chan int, 2)

	x75 := make(chan int, 2)
	y75 := make(chan int, 2)

	x76 := make(chan int, 2)
	y76 := make(chan int, 2)

	x77 := make(chan int, 2)
	y77 := make(chan int, 2)

	x78 := make(chan int, 2)
	y78 := make(chan int, 2)

	x79 := make(chan int, 2)
	y79 := make(chan int, 2)

	x80 := make(chan int, 2)
	y80 := make(chan int, 2)

	x81 := make(chan int, 2)
	y81 := make(chan int, 2)

	x82 := make(chan int, 2)
	y82 := make(chan int, 2)

	x83 := make(chan int, 2)
	y83 := make(chan int, 2)

	x84 := make(chan int, 2)
	y84 := make(chan int, 2)

	x85 := make(chan int, 2)
	y85 := make(chan int, 2)

	x86 := make(chan int, 2)
	y86 := make(chan int, 2)

	x87 := make(chan int, 2)
	y87 := make(chan int, 2)

	x88 := make(chan int, 2)
	y88 := make(chan int, 2)

	x89 := make(chan int, 2)
	y89 := make(chan int, 2)

	x90 := make(chan int, 2)
	y90 := make(chan int, 2)

	x91 := make(chan int, 2)
	y91 := make(chan int, 2)

	x92 := make(chan int, 2)
	y92 := make(chan int, 2)

	x93 := make(chan int, 2)
	y93 := make(chan int, 2)

	x94 := make(chan int, 2)
	y94 := make(chan int, 2)

	x95 := make(chan int, 2)
	y95 := make(chan int, 2)

	x96 := make(chan int, 2)
	y96 := make(chan int, 2)

	x97 := make(chan int, 2)
	y97 := make(chan int, 2)

	x98 := make(chan int, 2)
	y98 := make(chan int, 2)

	x99 := make(chan int, 2)
	y99 := make(chan int, 2)

	go fir0(x0, y0)
	go fir1(x1, y1)
	go fir2(x2, y2)
	go fir3(x3, y3)
	go fir4(x4, y4)
	go fir5(x5, y5)
	go fir6(x6, y6)
	go fir7(x7, y7)
	go fir8(x8, y8)
	go fir9(x9, y9)
	go fir10(x10, y10)
	go fir11(x11, y11)
	go fir12(x12, y12)
	go fir13(x13, y13)
	go fir14(x14, y14)
	go fir15(x15, y15)
	go fir16(x16, y16)
	go fir17(x17, y17)
	go fir18(x18, y18)
	go fir19(x19, y19)
	go fir20(x20, y20)
	go fir21(x21, y21)
	go fir22(x22, y22)
	go fir23(x23, y23)
	go fir24(x24, y24)
	go fir25(x25, y25)
	go fir26(x26, y26)
	go fir27(x27, y27)
	go fir28(x28, y28)
	go fir29(x29, y29)
	go fir30(x30, y30)
	go fir31(x31, y31)
	go fir32(x32, y32)
	go fir33(x33, y33)
	go fir34(x34, y34)
	go fir35(x35, y35)
	go fir36(x36, y36)
	go fir37(x37, y37)
	go fir38(x38, y38)
	go fir39(x39, y39)
	go fir40(x40, y40)
	go fir41(x41, y41)
	go fir42(x42, y42)
	go fir43(x43, y43)
	go fir44(x44, y44)
	go fir45(x45, y45)
	go fir46(x46, y46)
	go fir47(x47, y47)
	go fir48(x48, y48)
	go fir49(x49, y49)
	go fir50(x50, y50)
	go fir51(x51, y51)
	go fir52(x52, y52)
	go fir53(x53, y53)
	go fir54(x54, y54)
	go fir55(x55, y55)
	go fir56(x56, y56)
	go fir57(x57, y57)
	go fir58(x58, y58)
	go fir59(x59, y59)
	go fir60(x60, y60)
	go fir61(x61, y61)
	go fir62(x62, y62)
	go fir63(x63, y63)
	go fir64(x64, y64)
	go fir65(x65, y65)
	go fir66(x66, y66)
	go fir67(x67, y67)
	go fir68(x68, y68)
	go fir69(x69, y69)
	go fir70(x70, y70)
	go fir71(x71, y71)
	go fir72(x72, y72)
	go fir73(x73, y73)
	go fir74(x74, y74)
	go fir75(x75, y75)
	go fir76(x76, y76)
	go fir77(x77, y77)
	go fir78(x78, y78)
	go fir79(x79, y79)
	go fir80(x80, y80)
	go fir81(x81, y81)
	go fir82(x82, y82)
	go fir83(x83, y83)
	go fir84(x84, y84)
	go fir85(x85, y85)
	go fir86(x86, y86)
	go fir87(x87, y87)
	go fir88(x88, y88)
	go fir89(x89, y89)
	go fir90(x90, y90)
	go fir91(x91, y91)
	go fir92(x92, y92)
	go fir93(x93, y93)
	go fir94(x94, y94)
	go fir95(x95, y95)
	go fir96(x96, y96)
	go fir97(x97, y97)
	go fir98(x98, y98)
	go fir99(x99, y99)
	go acc(y0, y1, y2, y3, y4, y5, y6, y7, y8, y9, y10, y11, y12, y13, y14, y15, y16, y17, y18, y19, y20, y21, y22, y23, y24, y25, y26, y27, y28, y29, y30, y31, y32, y33, y34, y35, y36, y37, y38, y39, y40, y41, y42, y43, y44, y45, y46, y47, y48, y49, y50, y51, y52, y53, y54, y55, y56, y57, y58, y59, y60, y61, y62, y63, y64, y65, y66, y67, y68, y69, y70, y71, y72, y73, y74, y75, y76, y77, y78, y79, y80, y81, y82, y83, y84, y85, y86, y87, y88, y89, y90, y91, y92, y93, y94, y95, y96, y97, y98, y99, y)
	go input_control(x0, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14, x15, x16, x17, x18, x19, x20, x21, x22, x23, x24, x25, x26, x27, x28, x29, x30, x31, x32, x33, x34, x35, x36, x37, x38, x39, x40, x41, x42, x43, x44, x45, x46, x47, x48, x49, x50, x51, x52, x53, x54, x55, x56, x57, x58, x59, x60, x61, x62, x63, x64, x65, x66, x67, x68, x69, x70, x71, x72, x73, x74, x75, x76, x77, x78, x79, x80, x81, x82, x83, x84, x85, x86, x87, x88, x89, x90, x91, x92, x93, x94, x95, x96, x97, x98, x99, x)
}