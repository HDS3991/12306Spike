package logic

type LocalSpike struct {
	LocalInStock     int64
	LocalSalesVolume int64
}

func (spike *LocalSpike) LocalDeductionStock() bool {
	spike.LocalSalesVolume += 1
	return spike.LocalSalesVolume <= spike.LocalInStock
}
