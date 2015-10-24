package lottery
import (
	"math/rand"
	"time"
"math"
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
	ratio := normFloat / math.MaxFloat64 / 2.0 + 0.5
	additional := (end - start )* ratio
	return start + additional
}
