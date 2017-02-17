//package main;import ("os";"image";"image/png";"image/color";"strconv";"math";);func R(a float64,b float64,c int)float64{var d float64;e:=math.Pow(10,float64(c));g:=e*a;_,f:=math.Modf(g);if f>=b{d=math.Ceil(g);}else{d=math.Floor(g);};return d/e;};type p struct{x float64;y float64;};type q struct{a p;b p;};const(Y int=127;X int=Y*4;);func main(){g:=image.NewRGBA(image.Rect(0,0,X,Y));a,_:=strconv.ParseFloat(os.Args[1],4);d,_:=strconv.ParseFloat(os.Args[2],4);s,_:=strconv.ParseFloat(os.Args[3],4);r,_:=strconv.ParseFloat(os.Args[4],4);j:=q{a:p{0,0},b:p{a,127},};k:=q{a:j.b, b:p{j.b.x+d,s},};l:=q{a:k.b, b:p{k.b.x+64,s},};n:=q{a:l.b,b:p{l.b.x+r,0},};o:=[4]q{j,k,l,n};for _,l:=range o{m:=(l.b.y-l.a.y)/(l.b.x-l.a.x);t:=l.b.y-m*l.b.x;x,y:=l.a.x,l.a.y;for x<=l.b.x{g.Set(int(R(x,.5,0)),Y-int(R(y,.5,0)),color.RGBA{0,0,0,255});x+=0.01;y=m*x+t;}};f,_:= os.OpenFile("o.png",os.O_WRONLY|os.O_CREATE,0600);png.Encode(f,g);}
//http://codegolf.stackexchange.com/questions/110037/attack-decay-sustain-release/110314#110314

package main

import (
	"os"
	"image"
	"image/png"
	"image/color"
	"strconv"
	"math"
)

func Round(val float64, roundOn float64, places int ) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

type point struct {
	x float64
	y float64
}

type line struct {
	p1 point
	p2 point
}

const (
	MaxY int = 127
	MaxX int = MaxY*4
)


func main(){
	img := image.NewRGBA(image.Rect(0, 0, MaxX, MaxY))


	a, _ := strconv.ParseFloat(os.Args[1], 4)
	d, _ := strconv.ParseFloat(os.Args[2], 4)
	s, _ := strconv.ParseFloat(os.Args[3], 4)
	r, _ := strconv.ParseFloat(os.Args[4], 4)


	l1 := line{
		p1:point{0,0},
		p2:point{a, 127},
	}
	l2 := line{
		p1:l1.p2,
		p2:point{l1.p2.x+d, s},
	}
	l3 := line{
		p1:l2.p2,
		p2:point{l2.p2.x+64, s},
	}
	l4 := line{
		p1:l3.p2,
		p2:point{l3.p2.x+r, 0},
	}

	lines := [4]line{l1,l2,l3,l4}


	for _,l := range lines {
		m := (l.p2.y - l.p1.y)/(l.p2.x - l.p1.x)
		t := l.p2.y-m*l.p2.x

		x,y := l.p1.x, l.p1.y
		for x <= l.p2.x {
			img.Set(int(Round(x, .5, 0)), MaxY-int(Round(y, .5, 0)), color.RGBA{0,0,0,255})
			x += 0.01
			y = m*x + t
		}
	}

	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	png.Encode(f, img)
}