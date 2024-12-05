package filter

import "github.com/garrettladley/prods/internal/types"

type PriceBucket string

const (
	AllPriceBuckets    PriceBucket = "all"
	UnderFive          PriceBucket = "under_5"
	FiveToTen          PriceBucket = "5_to_10"
	TenToFifteen       PriceBucket = "10_to_15"
	FifteenToTwenty    PriceBucket = "15_to_20"
	TwentyToTwentyFive PriceBucket = "20_to_25"
	OverTwentyFive     PriceBucket = "over_25"
)

func (pb PriceBucket) intoRange() (r types.Pair[uint32]) {
	switch pb {
	case AllPriceBuckets:
		r.First, r.Second = 0, ^uint32(0)
	case UnderFive:
		r.First, r.Second = 0, 500
	case FiveToTen:
		r.First, r.Second = 500, 1000
	case TenToFifteen:
		r.First, r.Second = 1000, 1500
	case FifteenToTwenty:
		r.First, r.Second = 1500, 2000
	case TwentyToTwentyFive:
		r.First, r.Second = 2000, 2500
	case OverTwentyFive:
		r.First, r.Second = 2500, 5000
	}
	return
}

func PriceBucketer(bucket PriceBucket) Option {
	return func(f *Filter) {
		r := bucket.intoRange()
		f.p.PriceMin = r.First
		f.p.PriceMax = r.Second
	}
}

type StarBucket string

const (
	AllStars          StarBucket = "all"
	FourPointFivePlus StarBucket = "4_5_plus"
	FourPlus          StarBucket = "4_plus"
	ThreePlus         StarBucket = "3_plus"
)

func (sb StarBucket) intoRange() (r types.Pair[uint16]) {
	switch sb {
	case AllStars:
		r.First, r.Second = 0, 500
	case FourPointFivePlus:
		r.First, r.Second = 450, 500
	case FourPlus:
		r.First, r.Second = 400, 500
	case ThreePlus:
		r.First, r.Second = 300, 500
	}
	return
}

func StarBucketer(bucket StarBucket) Option {
	return func(f *Filter) {
		r := bucket.intoRange()
		f.p.StarMin = r.First
		f.p.StarMax = r.Second
	}
}
