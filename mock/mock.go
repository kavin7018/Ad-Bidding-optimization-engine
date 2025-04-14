package mock

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Campign struct {
	UserId               string
	CampignId            string
	DailyBudget          int
	NoOfClicks           int
	NoOfViews            int
	NoOfConversions      int
	DailyBudgetRemaining int
	TotalDurationInDays  int
	Cpv                  int
}

// initialize the mocks
func Start(ctx context.Context, exitChan chan<- bool) {
	fmt.Println("Starting mock server...")

	dailyBudget := TotalBudget / 3
	config := map[string]Campign{
		"google":    Campign{DailyBudget: dailyBudget, DailyBudgetRemaining: dailyBudget, Cpv: BidAmount, TotalDurationInDays: DurationInDays, UserId: UserID, CampignId: CampignID},
		"facebook":  Campign{DailyBudget: dailyBudget, DailyBudgetRemaining: dailyBudget, Cpv: BidAmount, TotalDurationInDays: DurationInDays, UserId: UserID, CampignId: CampignID},
		"instagram": Campign{DailyBudget: dailyBudget, DailyBudgetRemaining: dailyBudget, Cpv: BidAmount, TotalDurationInDays: DurationInDays, UserId: UserID, CampignId: CampignID},
	}

	var wg sync.WaitGroup
	for platform, campign := range config {
		wg.Add(1)
		go process(platform, &campign, &wg, ctx, exitChan)
	}

	defer wg.Wait()
}

func process(platform string, c *Campign, wg *sync.WaitGroup, ctx context.Context, exitChan chan<- bool) {
	defer wg.Done()

	// initialize the ad platform
	p := Platform{name: platform, campign: c}
	if err := p.initialize(); err != nil {
		fmt.Printf("Error in initializing mock server. Err: %v", err)
		exitChan <- true
	}
	p.ctx = ctx

	for {
		select {
		case <-ctx.Done():
			return

		default:
			// mock total number of views, total number of clicks, total number of conversions
			for c.DailyBudget > 0 {
				time.Sleep(20 * time.Second)
				r := rand.New(rand.NewSource(time.Now().UnixNano()))

				for iter := 1; iter <= 10; iter++ {
					// random int
					dice := (r.Int() % 10) + 1

					// no of views
					if dice%2 == 0 {
						fmt.Println("View Event in ", platform)
						p.pushViewEvent()
					}

					// no of clicks
					if dice <= 3 {
						fmt.Println("click Event ", platform)
						p.pushClickEvent()
					}

					// no of conversions
					if dice == 1 {
						fmt.Println("conversion Event ", platform)
						p.pushConversionEvent()
					}
				}
			}
			p.producer.Close()
			p.client.Close()
			break
		}
	}

}
