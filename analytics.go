package dz

type Sale struct {
	Product          string
	Category         string
	Quantity         int
	UnitPriceKopecks int64
	Cancelled        bool
}

type Report struct {
	RevenueKopecks    int64
	UnitsSold         int
	RevenueByCategory map[string]int64
	TopProduct        string
	InvalidSales      int
}

func AnalyzeSales(sales []Sale) Report {
	// TODO: реализуйте расчёт отчёта по правилам задания.
	var revenue int64
	var soldUnits, invalidSales, topProdUnits int
	categories := make(map[string]int64)
	topProd := ""
	for _, sale := range sales {
		if sale.Cancelled {
			continue
		}
		if sale.Category == "" || sale.Product == "" || sale.Quantity <= 0 || sale.UnitPriceKopecks <= 0 {
			invalidSales++
			continue
		}
		revenueFromThis := sale.UnitPriceKopecks * int64(sale.Quantity)
		revenue += revenueFromThis
		soldUnits += sale.Quantity
		categories[sale.Category] += revenueFromThis
		if sale.Quantity > topProdUnits || (sale.Quantity == topProdUnits && sale.Product < topProd) {
			topProdUnits = sale.Quantity
			topProd = sale.Product
		}
	}
	return Report{revenue, soldUnits, categories, topProd, invalidSales}
}
