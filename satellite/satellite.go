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
var V [24][9] float64
var inputs [][] string

func main() {
	//Starting conditions: creating a log and clearing out the old one.
	start := time.Now()
	e:=os.Remove("Satellite.log")
	if e != nil {
	}
	lg, err := os.OpenFile("satellite.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err !=nil {
		log.Fatal(err)
	}

	//Closing the log later but it gets closed
	defer lg.Close()

	log.SetOutput(lg)
	//Getting our Data from data.dat
	Data()
	//Input from our vehicle
	input()

	//Synchronizing our threads.
	var wg sync.WaitGroup

	//Number of threads
	wg.Add(4)

	/*
	Initializing our output array, this makes it, so we don't have to care about the order at which our threads close,
	We just use some index to sort them out as they go into this array.
	*/
	var out [][]string
	for i:=0; i<len(inputs); i++ {
		out = append(out, []string{""})
	}
	//Thread 1
	go func() {
		defer wg.Done()
		for i:=0; i<len(inputs); i=i+4 {
			var hold []string
			hold = nil
			pos:= CartPos(inputs[i])
			Tv:= strtoflt(inputs[i][0])
			Xv :=pos
			B := bmth(pos)

			for j := 0; j < len(B); j++ {

				if B[j] {
					Ts := ComTs(j,Tv, Xv)
					Xs := FindXs(j, Ts)
					l := strconv.Itoa(j) + " " + fltostr(Ts) + " " + fltostr(Xs[0]) + " " + fltostr(Xs[1]) + " " + fltostr(Xs[2])
					//fmt.Println(l)
					hold = append(hold,l)
				}
			}
			out[i]=hold
		}
	}()
	//Thread 2
	go func() {
		defer wg.Done()
		for i:=1; i<len(inputs); i=i+4 {
			var hold []string
			hold = nil
			pos:= CartPos(inputs[i])
			Tv:= strtoflt(inputs[i][0])
			Xv :=pos
			B := bmth(pos)
			for j := 0; j < len(B); j++ {
				if B[j] {
					Ts := ComTs(j,Tv, Xv)
					Xs := FindXs(j, Ts)
					l := strconv.Itoa(j) + " " + fltostr(Ts) + " " + fltostr(Xs[0]) + " " + fltostr(Xs[1]) + " " + fltostr(Xs[2])
					hold = append(hold,l)
				}
			}
			out[i]=hold
		}
	}()
	//Thread 3
	go func() {
		defer wg.Done()
		for i:=2; i<len(inputs); i=i+4 {
			var hold []string
			hold = nil
			pos:= CartPos(inputs[i])
			Tv:= strtoflt(inputs[i][0])
			Xv :=pos
			B := bmth(pos)

			for j := 0; j < len(B); j++ {

				if B[j] {
					Ts := ComTs(j,Tv, Xv)
					Xs := FindXs(j, Ts)
					l := strconv.Itoa(j) + " " + fltostr(Ts) + " " + fltostr(Xs[0]) + " " + fltostr(Xs[1]) + " " + fltostr(Xs[2])
					hold = append(hold,l)
				}
			}
			out[i]=hold
		}
	}()
	//Thread 4
	go func() {
		defer wg.Done()
		for i:=3; i<len(inputs); i=i+4 {
			var hold []string
			hold = nil
			pos:= CartPos(inputs[i])
			Tv:= strtoflt(inputs[i][0])
			Xv :=pos
			B := bmth(pos)

			for j := 0; j < len(B); j++ {

				if B[j] {
					Ts := ComTs(j,Tv, Xv)
					Xs := FindXs(j, Ts)
					l := strconv.Itoa(j) + " " + fltostr(Ts) + " " + fltostr(Xs[0]) + " " + fltostr(Xs[1]) + " " + fltostr(Xs[2])
					hold = append(hold,l)
				}
			}
			out[i]=hold
		}
	}()

	//Waits for all threads to finish!
	wg.Wait()
	//Our time since starting to now
	els := time.Since(start)
	//Our output
	for i:=0; i<len(out); i++ {
		log.Println("Read:", inputs[i])
		log.Println("Output:")
		for j:=0; j<len(out[i]); j++ {
			fmt.Println(out[i][j])
			log.Println(out[i][j])

		}
	}
	el:=time.Since(start)
	log.Println("Without Write Time")
	log.Println(els)
	log.Println("Elapsed Time Total")
	log.Println(el)
}
//func for bringing in our input and splitting it up
func input(){
	in:=bufio.NewScanner(os.Stdin)
	for in.Scan(){
		spit:= strings.Split(in.Text(), " ")
		inputs=append(inputs,spit)
	}
}
//func to generate our constants from data.dat
func Data() {
	var line []float64
	file, err := os.Open("data.dat")
	if err !=nil {
		fmt.Println("Error")
		log.Fatal(err)
	}
	scan :=bufio.NewScanner(file)
	for scan.Scan() {
		line = append(line, strtoflt(strings.Replace(scan.Text()[1:26], " ","",-1)))
	}

	Pi=line[0]
	c=line[1]
	R=line[2]
	S=line[3]
	cntr := 4
	for i:=0; i<24; i++ {
		for j:=0;j<9; j++ {
			V[i][j]=line[cntr]
			cntr++

		}

	}
}
func ComTs( v int, Tv float64, Xv []float64) float64 {
	var t0, norm float64
	t0=Tv
	norm=0.0
	for i := 0; i < 3; i++ {
		norm += math.Pow(FindXs(v,Tv)[i],2)
	}
	t0 =t0-math.Sqrt(norm)/c
	return newtmeth(v,t0,0, Tv, Xv)
}
func FindXs (v int, t float64) []float64  {

	x := (R+V[v][7])*(V[v][0]*math.Cos((2*Pi*t)/V[v][6]+V[v][8]) + V[v][3]*math.Sin((2*Pi*t)/V[v][6]+V[v][8]))
	y := (R+V[v][7])*(V[v][1]*math.Cos((2*Pi*t)/V[v][6]+V[v][8]) + V[v][4]*math.Sin((2*Pi*t)/V[v][6]+V[v][8]))
	z := (R+V[v][7])*(V[v][2]*math.Cos((2*Pi*t)/V[v][6]+V[v][8]) + V[v][5]*math.Sin((2*Pi*t)/V[v][6]+V[v][8]))
	ret := []float64{x,y,z,t}
	return ret
}
func FindXs1(v int, t float64) []float64 {
	var x,y,z float64
	x = 2*Pi*(1/V[v][6])*(R+V[v][7])*(V[v][3]*math.Sin((2*Pi*t)/V[v][6]+V[v][8])-V[v][0]*math.Cos((2*Pi*t)/V[v][6]+V[v][8]))
	y = 2*Pi*(1/V[v][6])*(R+V[v][7])*(V[v][4]*math.Sin((2*Pi*t)/V[v][6]+V[v][8])-V[v][1]*math.Cos((2*Pi*t)/V[v][6]+V[v][8]))
	z = 2*Pi*(1/V[v][6])*(R+V[v][7])*(V[v][5]*math.Sin((2*Pi*t)/V[v][6]+V[v][8])-V[v][2]*math.Cos((2*Pi*t)/V[v][6]+V[v][8]))
	ret := []float64{x,y,z,t}
	return ret
}
func fn(v int, t float64, Tv float64, Xv []float64) float64 {

	re := -1*math.Pow(c*(Tv-t),2)
	for i := 0; i<3; i++ {
		re += math.Pow(FindXs(v,t)[i]-Xv[i],2)
	}
	return re
}

func fn1(v int, t float64, Tv float64, Xv []float64) float64 {
	var re float64
	re = 2*c*c*(Tv-t)
	for i:=0; i<3; i++ {
		re += 2*(FindXs(v,t)[i]-Xv[i])*FindXs1(v,t)[i]
	}
	return re
}
//Newtons method from HW1
func newtmeth(v int, Tk float64, step int, Tv float64, Xv []float64 )  float64 {
	//Tk1 := Tk
	Tk1 := Tk-(fn(v,Tk,Tv,Xv)/fn1(v,Tk,Tv,Xv))
	if Tk1-Tk < 0.0001/c {
		return Tk1
	}else if step >= 9 {
		return -1
	}else {
		return newtmeth(v, Tk1, step+1,Tv, Xv)
	}
}
func bmth(a []float64) []bool {
	var ret []bool
	var sat []float64
	for i := 0; i <24; i++ {
		sat = FindXs(i,a[3])
		ret =append(ret, 2*a[0]*(sat[0]-a[0]) + 2*a[1]*(sat[1]-a[1]) + 2*a[2]*(sat[2]-a[2]) > 0)
		if i==8 {
		}

	}
	return ret
}

func R2D2(ang float64) []string {
	b := ang*180/Pi
	if ang<0 {
		b = -1*b
	}
	d := math.Floor(b)
	m := math.Floor(60*(b-d))
	s := 60*(60*(b-d)-m)
	if ang<0 {
		d = -1*d
	}

	ret :=[]string{fltostr(d),fltostr(m),fltostr(s)}
	return ret
}
func CartPos(pos []string)  []float64{
	t := strtoflt(pos[0])
	theta := strtoflt(pos[4])*D2R([]string{pos[1],pos[2],pos[3]})
	phi := strtoflt(pos[8])*D2R([]string{pos[5],pos[6],pos[7]})
	h := strtoflt(pos[9])
	rh := R+h
	x := rh*math.Cos(theta)*math.Cos(phi)
	y := rh*math.Cos(theta)*math.Sin(phi)
	z := rh*math.Sin(theta)
	alp := (2*Pi*t)/S
	return 	[]float64{math.Cos(alp)*x-math.Sin(alp)*y,math.Sin(alp)*x+math.Cos(alp)*y,z,t}
}
func D2R(deg []string) float64 {
	return 2*Pi*(strtoflt(deg[0])/360+strtoflt(deg[1])/(360*60)+strtoflt(deg[2])/(360*60*60))

}
func fltostr(a float64) string {
	str := strconv.FormatFloat(a,'E',-1,64)
	return str

}
func strtoflt(str string) float64 {
	f, e :=strconv.ParseFloat(str,64)
	if e == nil {return f
	} else {fmt.Println("Aylmao")
	return 0}
}
