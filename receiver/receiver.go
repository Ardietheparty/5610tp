package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var Pi, c, R, S float64
var m []int
var dat [][]float64
var DATA [][][]float64
func main() {
	start := time.Now()
	constants()
	datain()
	blocks()
	var out []string
	//fmt.Println(DATA)
	//fmt.Println("starting computations")
	var wg sync.WaitGroup
	for i:=0; i<len(m);i++{
		out = append(out, "")
	}

	wg.Add(4)
	go func() {
		defer wg.Done()
		for i:=0;i<len(m);i=i+4 {
			//out1<-arrayToString(comp(m[i],DATA[i]), " ")
			//ind1<-i
			out[i]=arrayToString(comp(m[i],DATA[i]), " ")
		}
	}()
	go func() {
		defer wg.Done()
		for i:=1;i<len(m);i=i+4 {
			//out2<-arrayToString(comp(m[i],DATA[i]), " ")
			//ind2<-i
			out[i]=arrayToString(comp(m[i],DATA[i]), " ")
		}
	}()
	go func() {
		defer wg.Done()
		for i:=2;i<len(m);i=i+4 {
			//out3<-arrayToString(comp(m[i],DATA[i]), " ")
			//ind3<-i
			out[i]=arrayToString(comp(m[i],DATA[i]), " ")
		}
	}()
	go func() {
		defer wg.Done()
		for i:=3;i<len(m);i=i+4 {
			//out4<-arrayToString(comp(m[i],DATA[i]), " ")
			//ind4<-i
			out[i]=arrayToString(comp(m[i],DATA[i]), " ")
		}
	}()
	wg.Wait()
	for i:=0; i<len(m); i++{
		fmt.Println(out[i])
	}
	dur:=time.Since(start)
	/*
	s1 := time.Now()
	for i:=0;i<len(m);i++ {
		fmt.Println(m[i],DATA[i])
	}
	dur1:=time.Since(s1)
	fmt.Println(dur)
	fmt.Println(dur1)
	 */
	fmt.Println(dur)

}
func blocks() {
	//fmt.Println("Start Block")
	var hold [][]float64
	var hld  []float64
	cnt1:=0
	cnt2:=0
	cnt3:=0
	shift := 0
	for i := 0; i < len(m); i++ {
		cnt1++
		hold=nil
		for j:=0;j<m[i];j++ {
			cnt2++
			hld = nil
			for k:=0;k<4;k++ {
				cnt3++
				hld= append(hld,dat[j+shift][k])
			}
			hold=append(hold,hld)
		}
		shift += m[i]
		DATA = append(DATA,hold)
	}

	//fmt.Println(cnt1)
	//fmt.Println(cnt2)
	//fmt.Println(cnt3)
}
func constants()  {
	var line []float64
	file, err := os.Open("data.dat")
	if err !=nil {
		fmt.Println("Error")
		log.Fatal(err)
	}
	scan :=bufio.NewScanner(file)
	for scan.Scan() {
		//a := scan.Text()
		line = append(line, StringToFloat(strings.Replace(scan.Text()[1:26], " ","",-1)))
	}
	Pi=line[0]
	c=line[1]
	R=line[2]
	S=line[3]
	//fmt.Println("Constants done!")
}
func datain() {
	var check [] int
	var input [][]string
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		spit:=strings.Split(in.Text(), " ")
		input=append(input,spit)
	}
	for a := 0; a < len(input); a++ {
		hold := input[a][0]
		holder,err :=strconv.Atoi(hold)
		if err!=nil {
			fmt.Println("Help")
		}
		check = append(check,holder)
		t:=StringToFloat(input[a][1])
		x:=StringToFloat(input[a][2])
		y:=StringToFloat(input[a][3])
		z:=StringToFloat(input[a][4])
		dat = append(dat, []float64{t,x,y,z})
	}
	length:=1
	for i := 1; i < len(check); i++ {
		if check[i]-check[i-1]>0{
			length++
		} else {
			m=append(m,length)
			length=1
		}
	}
	m=append(m,length)
	//chunks = len(index)
	//fmt.Println("Data in done!")
}
func arrayToString(a []float64, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprintf("%.2f",a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
func comp(m int, data [][]float64) []float64 {
	var ret []float64
	X0 := []float64{0,0,0}
	Xv := nwtmth(X0,0, m, data)
	//fmt.Println(Xv)
	TV := TvTx(Xv, data)
	//fmt.Println(TV)
	ret = append(ret, TV)
	for i:=0; i<len(Xv); i++ {
		ret=append(ret,Xv[i])
	}
	return Pos(ret)

}
func nwtmth(x []float64, d int, m int, data [][]float64) []float64 {
	e:=0.00000001
	Sk := sol(jc(x,m,data),gf(x,m,data))
	//fmt.Println(Sk)
	//fmt.Println(d)
	var ret []float64
	for i:=0; i<len(x); i++ {
		ret = append(ret,x[i])
	}
	//fmt.Println(ret)
	//fmt.Println(x)
	for i := 0; i < 3; i++ {
		//fmt.Println(ret[i])
		ret[i] -=Sk[i]
		//fmt.Println(Sk[i])
		//fmt.Println(ret[i])
	}
	//fmt.Println(ret)
	//a :=math.Abs(df(ret,x)[0])
	//fmt.Println(df(ret,x))
	//fmt.Println(a)
	//fmt.Println(ret[0]-x[0])
	b1 := math.Abs(df(ret,x)[0])<e
	b2 := math.Abs(df(ret,x)[1])<e
	b3 := math.Abs(df(ret,x)[2])<e
	//fmt.Println(b1,b2,b3)
	if b1 && b2 && b3 {
		//fmt.Println("Conv?")
		return ret
	}
	if d>10 {
		return nil
	}
	//fmt.Println("We gotta go back")
	d++
	return nwtmth(ret, d, m, data)

}

func jc(x []float64, m int, data [][]float64) [][]float64 {
	var ret [][]float64
	var hold []float64
	for i:=0;i<3;i++ {
		hold=nil
		for j := 0; j < 3; j++ {
			hold = append(hold,Jij(i,j,x, m,data))
		}
		ret=append(ret,hold)
	}
	return ret
}
func norm(x []float64) float64 {
	var ret float64
	for i:=0; i<len(x); i++{
		ret+=x[i]*x[i]
	}
	return math.Sqrt(ret)
}
func df(a []float64, b []float64) []float64 {
	var ret []float64
	if len(a)==len(b) {
		//fmt.Println(len(a))
		for i := 0; i < len(a); i++ {
			ret=append(ret,a[i]-b[i])
			//fmt.Println(a[i]-b[i])
		}
		return ret
	}
	return nil
}
/*
func sol(a [][]float64, b []float64) []float64 {
	c1 := make(chan float64)
	c2 := make(chan float64)
	c3 := make(chan float64)
	go func() {
		c1 <- (a[0][1]*a[1][2]*b[2] - a[0][1]*a[2][2]*b[1] - a[0][2]*a[1][1]*b[2] + a[0][2]*a[2][1]*b[1] + a[1][1]*a[2][2]*b[0] - a[1][2]*a[2][1]*b[0])/(a[0][0]*a[1][1]*a[2][2] - a[0][0]*a[1][2]*a[2][1] - a[0][1]*a[1][0]*a[2][2] + a[0][1]*a[1][2]*a[2][0] + a[0][2]*a[1][0]*a[2][1] - a[0][2]*a[1][1]*a[2][0])
		//fmt.Println(c1)
	}()
	go func() {
		c2 <- -1 * (a[0][0]*a[1][2]*b[2] - a[0][0]*a[2][2]*b[1] - a[0][2]*a[1][0]*b[2] + a[0][2]*a[2][0]*b[1] + a[1][0]*a[2][2]*b[0] - a[1][2]*a[2][0]*b[0])/(a[0][0]*a[1][1]*a[2][2] - a[0][0]*a[1][2]*a[2][1] - a[0][1]*a[1][0]*a[2][2] + a[0][1]*a[1][2]*a[2][0] + a[0][2]*a[1][0]*a[2][1] - a[0][2]*a[1][1]*a[2][0])
		//fmt.Println(c2)
	}()
	go func() {
		c3 <- (a[0][0]*a[1][1]*b[2] - a[0][0]*a[2][1]*b[1] - a[0][1]*a[1][0]*b[2] + a[0][1]*a[2][0]*b[1] + a[1][0]*a[2][1]*b[0] - a[1][1]*a[2][0]*b[0])/(a[0][0]*a[1][1]*a[2][2] - a[0][0]*a[1][2]*a[2][1] - a[0][1]*a[1][0]*a[2][2] + a[0][1]*a[1][2]*a[2][0] + a[0][2]*a[1][0]*a[2][1] - a[0][2]*a[1][1]*a[2][0])
	}()
	x := <-c1
	y := <-c2
	z := <-c3
	//fmt.Println(x,y,z)
	ret := []float64{x,y,z}
	return ret
}

 */
func sol (a [][]float64, b []float64) []float64{
	x:=(a[0][1]*a[1][2]*b[2] - a[0][1]*a[2][2]*b[1] - a[0][2]*a[1][1]*b[2] + a[0][2]*a[2][1]*b[1] + a[1][1]*a[2][2]*b[0] - a[1][2]*a[2][1]*b[0])/(a[0][0]*a[1][1]*a[2][2] - a[0][0]*a[1][2]*a[2][1] - a[0][1]*a[1][0]*a[2][2] + a[0][1]*a[1][2]*a[2][0] + a[0][2]*a[1][0]*a[2][1] - a[0][2]*a[1][1]*a[2][0])
	y:=-1 * (a[0][0]*a[1][2]*b[2] - a[0][0]*a[2][2]*b[1] - a[0][2]*a[1][0]*b[2] + a[0][2]*a[2][0]*b[1] + a[1][0]*a[2][2]*b[0] - a[1][2]*a[2][0]*b[0])/(a[0][0]*a[1][1]*a[2][2] - a[0][0]*a[1][2]*a[2][1] - a[0][1]*a[1][0]*a[2][2] + a[0][1]*a[1][2]*a[2][0] + a[0][2]*a[1][0]*a[2][1] - a[0][2]*a[1][1]*a[2][0])
	z:=(a[0][0]*a[1][1]*b[2] - a[0][0]*a[2][1]*b[1] - a[0][1]*a[1][0]*b[2] + a[0][1]*a[2][0]*b[1] + a[1][0]*a[2][1]*b[0] - a[1][1]*a[2][0]*b[0])/(a[0][0]*a[1][1]*a[2][2] - a[0][0]*a[1][2]*a[2][1] - a[0][1]*a[1][0]*a[2][2] + a[0][1]*a[1][2]*a[2][0] + a[0][2]*a[1][0]*a[2][1] - a[0][2]*a[1][1]*a[2][0])
	return []float64{x,y,z}
}


func Jij(i int, j int, x []float64, m int,data [][]float64) float64 {
	var ret float64
	for k:=0;k<m-1; k++ {
		ret += gfij(k,i,x,data)*gfij(k,j,x,data)
	}
	return 2*ret
}
func gfij(i int, j int, x []float64, data [][]float64) float64 {
	ret:=0.5 * (2 * x[j] - 2 * data[i][j+1]) / (math.Sqrt(math.Pow(x[0]-data[i][1],2) + math.Pow(x[1]-data[i][2],2) + math.Pow((x[2]-data[i][3]),2)))
	ret -= 0.5 * (2 * x[j] - 2 * data[i+1][j+1]) / (math.Sqrt(math.Pow(x[0]-data[i+1][1],2) + math.Pow((x[1]-data[i+1][2]),2) + math.Pow((x[2]-data[i+1][3]),2)))
	return ret;
}
func gf(x []float64, m int, data [][]float64) []float64 {
	var dfs, xyz [][]float64
	var N, A, hold, ret []float64
	var held float64
	for i := 0; i < m; i++ {
		dfs=append(dfs,df([]float64{data[i][1],data[i][2],data[i][3]},x))
	}
	for i := 0; i < m; i++ {
		N = append(N,norm(dfs[i]))
	}
	for i := 0; i < m-1; i++ {
		A= append(A, N[i+1]-N[i]-c*(data[i][0]-data[i+1][0]))
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < m-1; j++ {
			hold= append(hold,dfs[j][i]/N[j] - dfs[j+1][i]/N[j+1])
		}
		xyz = append(xyz,hold)
		hold = nil
	}
	for j := 0; j < 3; j++ {
		held = 0
		for i:=0;i<m-1;i++ {
			held +=A[i]*xyz[j][i]
		}
		ret = append(ret, 2*held)
	}
	return ret
}
func TvTx(Xv []float64, data [][]float64) float64 {
	x0 :=math.Pow(Xv[0]-data[0][1],2)
	x1 :=math.Pow(Xv[1]-data[0][2], 2)
	x2 :=math.Pow(Xv[2]-data[0][3], 2)
	ret := data[0][0]+(1/c)*math.Sqrt(x0+x1+x2)
	return ret
}
func Pos(a []float64) []float64 {
	var ret []float64
	var psi, lambda float64
	t  := a[0]
	x1 := a[1]
	y1 := a[2]
	z1 := a[3]
	xyz := R3(-2*Pi*t/S,[]float64{x1,y1,z1})
	//fmt.Println(xyz)
	x := xyz[0]
	y := xyz[1]
	z := xyz[2]

	//Construct Psi (can be paralled
	if x*x+y*y == 0 {
		if z>=0 {
			psi = Pi/2
		} else {
			psi = -1*Pi/2
		}
	} else {
		psi = math.Atan2(z,math.Sqrt(x*x+y*y))
	}

	//Construct Lambda (can be paralled
	if x > 0 && y > 0 {
		lambda = math.Atan2(y,x)
	} else if x < 0 {
		lambda = Pi + math.Atan2(y,x)
	} else {
		lambda = 2*Pi+math.Atan2(y,x)
	}
	lambda -=Pi
	PSI := R2D2(psi)
	LM  := R2D2(lambda)
	h := math.Sqrt(x*x+y*y+z*z)-R
	ret = append(ret,t)
	for i:=0; i<len(PSI); i++{
		ret = append(ret,PSI[i])
	}
	for i:=0; i<len(LM); i++{
		ret = append(ret,LM[i])
	}
	ret = append(ret, h)
	return ret
	//return []float64{t,PSI[0],PSI[1],PSI[2],PSI[3],LM[0],LM[1],LM[2],LM[3],h}
}
func R3(a float64, x []float64) []float64 {
	return []float64{math.Cos(a)*x[0]-math.Sin(a)*x[1],math.Sin(a)*x[0]+math.Cos(a)*x[1],x[2]}
}

//Converts a String to a float.
func StringToFloat(str string) float64 {
	f, e :=strconv.ParseFloat(str,64)
	if e == nil {
		return f
	} else {
		fmt.Println("String to float Error")
	}
	return 0
}
//Converts Float to String
func FloatToString(a float64) string {
	str := strconv.FormatFloat(a,'E',-1,64)
	return str

}
//Rads to Degress
func R2D2(ang float64) []float64 {
	//fmt.Println(ang)
	b := ang * 180 / Pi
	if ang < 0 {
		b = -1 * b
	}
	d := math.Floor(b)
	m := math.Floor(60 * (b - d))
	s := 60 * (60*(b-d) - m)
	if ang < 0 {
		//fmt.Println([]float64{d,m,s, -1})
		return []float64{d,m,s, -1}
	}
	//fmt.Println([]float64{d, m, s, 1})
	return []float64{d, m, s, 1}

}
