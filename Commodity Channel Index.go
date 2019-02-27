package main

import "fmt"
import "io"
import "os"
import "math"

const N=20
var fname string
var file *os.File
var err error
var date int
var time int
var opening,high,low,closing float64
var prc[N] float64

/* I.S: Open data file, check if there is an error,and call procedure Init for first boot and start to calculate CCI then iterate until end of file.
   F.S: The program output the date, time, indicator value and buy signal.
*/
func main() {

    var i int
    var CCI float64
    var temp float64

    if len(os.Args) > 1 {
        fname = os.Args[1]
    } else {
        fname = "EUM11709.DAT"
    }
    
    file, err = os.Open(fname)
    if err == nil {
        Init()
        CCI = calcCCI(19)
        temp = CCI
        fmt.Println(date, time,";", closing,";", CCI)

        for err != io.EOF {
            i = 0
            for i < N  {
                _,err = fmt.Fscanf(file,"%d %d;%v;%v;%v;%v;0\n",&date, &time, &opening, &high, &low, &closing)
                Pricef(i, high, low, closing)
                CCI = calcCCI(i)
                if CCI > temp && temp < -100 && CCI > -100 {
                    fmt.Println(date, time,";", closing,";", CCI, "***BUY")
					
                } else if CCI < temp && temp > 100 && CCI < 100 {
                    fmt.Println(date, time,";", closing,";", CCI, "***SELL")
					
                } else {
                    fmt.Println(date, time,";", closing,";", CCI)
                }
                temp = CCI
                i = i + 1
            }
            
        }
    }
    file.Close()
}

/* I.S: No data in Array prc[i]
   F.S: Array prc[i] filled with first 20 data
*/
func Init() {
    var i int
	
    i = 0
    for i < N {
        _,err = fmt.Fscanf(file,"%d %d;%v;%v;%v;%v;0\n",&date, &time, &opening, &high, &low, &closing)
        Pricef(i, high, low, closing)
        i = i + 1
    }
}

/* I.S: the procedure recieve data from input, array prc[i] already exist.
   F.S: store the data to array prc[i].
*/
func Pricef(i int, h float64,l float64,c float64) {
    prc[i] = (h+l+c)/3
	 
}

/* I.S: array prc[i] already filled. then calculate the average. 
   F.S: average calculated and return the average value.
*/
func SMA() float64 {
    var mavg float64
    var i int
    var n int
	
    n = N
    mavg = 0
    mavg = prc[0]
    i = n - 1
    
    for i > 0 {
        mavg = mavg + prc[i]
        i = i - 1
    }
    mavg = mavg/20
    return mavg
}

/* I.S: array prc[i] already filled. then calculate the deviation.
   F.S: deviation calculated and return the deviation value.
*/
func Dev() float64 {
    var deviation float64
    var ma float64
    var i int 
    var tp float64

    deviation = 0
    i = 0
    ma = SMA()
    for i < N {
        tp = math.Abs((prc[i] - ma))
        deviation = deviation + tp

        i = i + 1
    }
    deviation = deviation/20
    return deviation
}

/* I.S: array prc[i] already filled. read i as the index of array. then calculate cci.
   F.S: cci calculated and return the cci value.
*/
func calcCCI(i int) float64{
    var deviation float64
    var tp float64
    var CCI float64
    var ma float64

    deviation = Dev()
    ma = SMA()
    tp = prc[i] - ma

    CCI=tp/(0.015*deviation)
    return CCI
}