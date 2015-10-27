package lottery
import (
	"math/rand"
	"time"
_ "math"
	"fmt"
)

func createLotteryEngine() Lottery{
	return New(rand.New(rand.NewSource(time.Now().UnixNano())))
}

//startからendまでの一様乱数を返す
func GetRandomInt(start int ,end int) int{
	if(start < 0){
		return -1
	}
	if(end < 0){
		return -1
	}
	engine := createLotteryEngine()
	prob := engine.rd.Intn(end - start)
	prob += start
	return prob
}

//startからendまでの範囲の正規分布からランダムな値を返す
func GetRandomNormInt(start int ,end int) int{
	if(start < 0){
		return -1
	}
	if(end < 0){
		return -1
	}
	engine := createLotteryEngine()
	normFloat := engine.rd.NormFloat64()
	fmt.Printf("norm: %f\n", normFloat)
	if(normFloat > 1.0){
		normFloat = 1.0
	}
	if(normFloat < -1.0){
		normFloat = -1.0
	}
	ratio := normFloat / 2.0 + 0.5
	fmt.Printf("ratio: %f\n", ratio)

	additional := float64(end - start ) * ratio
	ret := start + int(additional)
	if(ret > end){
		ret = end
	}
	if(ret < start){
		ret = start
	}
	return ret
}
