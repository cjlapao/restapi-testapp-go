package controller

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type CjFaker struct {
	Generator *rand.Rand
}

var (
	one = []string{
		"Lead",
		"Senior",
		"Direct",
		"Corporate",
		"Dynamic",
		"Future",
		"Product",
		"National",
		"Regional",
		"District",
		"Central",
		"Global",
		"Relational",
		"Customer",
		"Investor",
		"Dynamic",
		"International",
		"Legacy",
		"Forward",
		"Interactive",
		"Internal",
		"Human",
		"Chief",
		"Principal",
		"Quick",
	}
	two = []string{
		"Solutions",
		"Program",
		"Brand",
		"Security",
		"Research",
		"Marketing",
		"Directives",
		"Implementation",
		"Integration",
		"Functionality",
		"Response",
		"Paradigm",
		"Tactics",
		"Identity",
		"Markets",
		"Group",
		"Resonance",
		"Applications",
		"Optimization",
		"Operations",
		"Infrastructure",
		"Intranet",
		"Communications",
		"Web",
		"Branding",
		"Quality",
		"Assurance",
		"Impact",
		"Mobility",
		"Ideation",
		"Data",
		"Creative",
		"Configuration",
		"Accountability",
		"Interactions",
		"Factors",
		"Usability",
		"Metrics",
		"Team",
	}
	three = []string{
		"Supervisor",
		"Associate",
		"Executive",
		"Liason",
		"Officer",
		"Manager",
		"Engineer",
		"Specialist",
		"Director",
		"Coordinator",
		"Administrator",
		"Architect",
		"Analyst",
		"Designer",
		"Planner",
		"Synergist",
		"Orchestrator",
		"Technician",
		"Developer",
		"Producer",
		"Consultant",
		"Assistant",
		"Facilitator",
		"Agent",
		"Representative",
		"Strategist",
	}
)

func New() (f CjFaker) {
	seed := rand.NewSource(time.Now().Unix())
	f = NewWithSeed(seed)
	return
}

func NewWithSeed(src rand.Source) (f CjFaker) {
	generator := rand.New(src)
	f = CjFaker{Generator: generator}
	return
}

// Login Generate a token for a valid user
func (c *Controller) Generator(w http.ResponseWriter, r *http.Request) {
	logger.Info("Generator Endpoint")
	faker := New()
	test := faker.JobTitle()
	message := fmt.Sprintf("Hello Faker, %v", test)
	response := TestResponse{
		Message: message,
	}

	json.NewEncoder(w).Encode(response)
}

func (f *CjFaker) RandomStringElement(s []string) string {
	i := f.IntBetween(0, len(s)-1)
	return s[i]
}

func (f *CjFaker) IntBetween(min, max int) int {
	diff := max - min

	if diff == 0 {
		return min
	}

	return f.Generator.Intn(diff+1) + min
}

func (f *CjFaker) JobTitle() string {
	titleSize := f.IntBetween(1, 3)

	switch titleSize {
	case 1:
		return f.RandomStringElement(three)
	case 2:
		return fmt.Sprintf("%v %v", f.RandomStringElement(two), f.RandomStringElement(three))
	case 3:
		return fmt.Sprintf("%v %v %v", f.RandomStringElement(one), f.RandomStringElement(two), f.RandomStringElement(three))
	default:
		return fmt.Sprintf("%v %v %v", f.RandomStringElement(one), f.RandomStringElement(two), f.RandomStringElement(three))
	}
}
