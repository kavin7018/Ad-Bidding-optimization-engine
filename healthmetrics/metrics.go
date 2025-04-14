package healthmetrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	PlatformActivity = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "platform_activity",
			Help: "Total number of views/clicks/conversions per platform",
		},
		[]string{"platform", "eventType"},
	)

	// TotalBudget = prometheus.NewCounterVec(
	// 	prometheus.CounterOpts{
	// 		Name: "total_budget",
	// 		Help: "Total Budget",
	// 	},
	// 	[]string{"platform"},
	// )

	// TotalBudgetRemaining = prometheus.NewCounterVec(
	// 	prometheus.CounterOpts{
	// 		Name: "total_budget_remaining",
	// 		Help: "Total Budget Remaining",
	// 	},
	// 	[]string{"platform"},
	// )

	// TotalBudgetPerDay = prometheus.NewCounterVec(
	// 	prometheus.CounterOpts{
	// 		Name: "total_budget_per_day",
	// 		Help: "Total Budget per day",
	// 	},
	// 	[]string{"platform"},
	// )

	// TotalBudgetPerDayRemaining = prometheus.NewCounterVec(
	// 	prometheus.CounterOpts{
	// 		Name: "total_budget_Per_day_remaining",
	// 		Help: "Total Budget per day Remaining",
	// 	},
	// 	[]string{"platform"},
	// )

	Bid = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "platform_daily_bid",
			Help: "Daily Bid",
		},
		[]string{"platform"},
	)
)

func Init() {
	// Register your metrics with Prometheus default registry
	prometheus.MustRegister(PlatformActivity)
	prometheus.MustRegister(Bid)
	// prometheus.MustRegister(TotalBudget)
	// prometheus.MustRegister(TotalBudgetRemaining)
	// prometheus.MustRegister(TotalBudgetPerDay)
	// prometheus.MustRegister(TotalBudgetPerDayRemaining)
}
