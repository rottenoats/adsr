//http://codegolf.stackexchange.com/questions/110037/attack-decay-sustain-release/110314#110314
//package main;import (."os";."image";k"image/png";c"image/color";."strconv";."math");func main(){g:=NewRGBA(Rect(0,0,127*4,127));a,_:=ParseFloat(Args[1],4);d,_:=ParseFloat(Args[2],4);s,_:=ParseFloat(Args[3],4);r,_:=ParseFloat(Args[4],4);z:=[5][]float64{{0,0},{a,127},{a+d,s},{a+d+64,s},{a+d+64+r,0}};for i:=1;i<len(z);i++{v,w,x,y:=z[i-1][0],z[i-1][1],z[i][0],z[i][1];m:=(y-w)/(x-v);t:=y-m*x;for v<=x{g.Set(int(Ceil(v)),127-int(Ceil(w)),c.RGBA{0,0,0,255});v+=.01;w=m*v+t}};f,_:=Create("o.png");k.Encode(f,g)}

package main

import (
	."os"
	."image"
	k"image/png"
	c"image/color"
	."strconv"
	."math"
	"fmt"
)

func main(){
	g := NewRGBA(Rect(0, 0, 127*4, 127))

	a, _ := ParseFloat(Args[1], 4)
	d, _ := ParseFloat(Args[2], 4)
	s, _ := ParseFloat(Args[3], 4)
	r, _ := ParseFloat(Args[4], 4)

	z := [5][]float64{{0,0},{a,127},{a+d,s},{a+d+64,s},{a+d+64+r,0}}
	for i:=1;i<len(z);i++{
		v,w,x,y:=z[i-1][0],z[i-1][1],z[i][0],z[i][1]
		m:=(y-w)/(x-v)
		t:=y-m*x
		for v<=x{
			g.Set(int(Ceil(v)),127-int(Ceil(w)), c.RGBA{0,0,0,255})
			v+=.01
			w=m*v+t
		}
	}
	f,_:=Create("o.png")
	k.Encode(f,g)
}