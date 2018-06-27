package backtest

import (
	"reflect"
	"testing"
)

func TestSetData(t *testing.T) {
	var testCases = []struct {
		msg      string
		strategy *Strategy
		data     DataHandler
		exp      *Strategy
		expErr   error
	}{
		{"set data:",
			&Strategy{
				Node{name: "test", root: true},
				AlgoStack{}, nil, nil,
			},
			&Data{},
			&Strategy{
				Node{name: "test", root: true},
				AlgoStack{}, &Data{}, nil,
			},
			nil,
		},
		{"set data with child strategy:",
			&Strategy{
				Node{name: "test", root: true,
					children: []NodeHandler{
						&Strategy{
							Node{name: "sub", root: false},
							AlgoStack{}, nil, nil,
						},
					},
				},
				AlgoStack{}, nil, nil,
			},
			&Data{},
			&Strategy{
				Node{name: "test", root: true,
					children: []NodeHandler{
						&Strategy{
							Node{name: "sub", root: false},
							AlgoStack{}, &Data{}, nil,
						},
					},
				},
				AlgoStack{}, &Data{}, nil,
			},
			nil,
		},
	}

	for _, tc := range testCases {
		err := tc.strategy.SetData(tc.data)
		if !reflect.DeepEqual(tc.strategy, tc.exp) || (err != tc.expErr) {
			t.Errorf("%v SetData(%s): \nexpected: %#v %#v, \nactual:   %#v %#v",
				tc.msg, tc.data, tc.exp, tc.expErr, tc.strategy, err)
		}
	}
}

func TestSetPortfolio(t *testing.T) {
	var testCases = []struct {
		msg       string
		strategy  *Strategy
		portfolio PortfolioHandler
		exp       *Strategy
		expErr    error
	}{
		{"set data:",
			&Strategy{
				Node{name: "test", root: true},
				AlgoStack{}, nil, nil,
			},
			&Portfolio{},
			&Strategy{
				Node{name: "test", root: true},
				AlgoStack{}, nil, &Portfolio{},
			},
			nil,
		},
		{"set data with child strategy:",
			&Strategy{
				Node{name: "test", root: true,
					children: []NodeHandler{
						&Strategy{
							Node{name: "sub", root: false},
							AlgoStack{}, nil, nil,
						},
					},
				},
				AlgoStack{}, nil, nil,
			},
			&Portfolio{},
			&Strategy{
				Node{name: "test", root: true,
					children: []NodeHandler{
						&Strategy{
							Node{name: "sub", root: false},
							AlgoStack{}, nil, &Portfolio{},
						},
					},
				},
				AlgoStack{}, nil, &Portfolio{},
			},
			nil,
		},
	}

	for _, tc := range testCases {
		err := tc.strategy.SetPortfolio(tc.portfolio)
		if !reflect.DeepEqual(tc.strategy, tc.exp) || (err != tc.expErr) {
			t.Errorf("%v SetPortfolio(%s): \nexpected: %#v %#v, \nactual:   %#v %#v",
				tc.msg, tc.portfolio, tc.exp, tc.expErr, tc.strategy, err)
		}
	}
}

func TestNewStrategy(t *testing.T) {
	var testCases = []struct {
		msg  string
		name string
		exp  *Strategy
	}{
		{"setup new strategy:",
			"test",
			&Strategy{
				Node{name: "test", root: true},
				AlgoStack{},
				nil,
				nil,
			},
		},
	}

	for _, tc := range testCases {
		strategy := NewStrategy(tc.name)
		if !reflect.DeepEqual(strategy, tc.exp) {
			t.Errorf("%v NewStrategy(%s): \nexpected: %#v, \nactual:   %#v",
				tc.msg, tc.name, tc.exp, strategy)
		}
	}
}

func TestStrategies(t *testing.T) {
	var testCases = []struct {
		msg      string
		strategy *Strategy
		exp      []StrategyHandler
		expOk    bool
	}{
		{"test no children:",
			&Strategy{
				Node{name: "test", root: true},
				AlgoStack{}, nil, nil,
			},
			nil,
			false,
		},
		{"test no sub strategy, only assets:",
			&Strategy{
				Node{name: "test", root: true,
					children: []NodeHandler{
						&Asset{
							Node{name: "asset", root: false},
						},
					},
				},
				AlgoStack{}, nil, nil,
			},
			nil,
			false,
		},
		{"test single sub strategy:",
			&Strategy{
				Node{name: "test", root: true,
					children: []NodeHandler{
						&Strategy{
							Node{name: "sub", root: false},
							AlgoStack{}, nil, nil,
						},
					},
				},
				AlgoStack{}, nil, nil,
			},
			[]StrategyHandler{
				&Strategy{
					Node{name: "sub", root: false},
					AlgoStack{}, nil, nil,
				},
			},
			true,
		},
		{"test multiple sub strategies:",
			&Strategy{
				Node{name: "test", root: true,
					children: []NodeHandler{
						&Strategy{
							Node{name: "subA", root: false},
							AlgoStack{}, nil, nil,
						},
						&Strategy{
							Node{name: "subB", root: false},
							AlgoStack{}, nil, nil,
						},
					},
				},
				AlgoStack{}, nil, nil,
			},
			[]StrategyHandler{
				&Strategy{
					Node{name: "subA", root: false},
					AlgoStack{}, nil, nil,
				},
				&Strategy{
					Node{name: "subB", root: false},
					AlgoStack{}, nil, nil,
				},
			},
			true,
		},
		{"test multiple sub strategies and multiple assets:",
			&Strategy{
				Node{name: "test", root: true,
					children: []NodeHandler{
						&Strategy{
							Node{name: "subA", root: false},
							AlgoStack{}, nil, nil,
						},
						&Strategy{
							Node{name: "subB", root: false},
							AlgoStack{}, nil, nil,
						},
						&Asset{
							Node{name: "assetA", root: false},
						},
						&Asset{
							Node{name: "assetB", root: false},
						},
					},
				},
				AlgoStack{}, nil, nil,
			},
			[]StrategyHandler{
				&Strategy{
					Node{name: "subA", root: false},
					AlgoStack{}, nil, nil,
				},
				&Strategy{
					Node{name: "subB", root: false},
					AlgoStack{}, nil, nil,
				},
			},
			true,
		},
	}

	for _, tc := range testCases {
		strategies, ok := tc.strategy.Strategies()
		if !reflect.DeepEqual(strategies, tc.exp) || (ok != tc.expOk) {
			t.Errorf("%v Strategies(): \nexpected %#v %v, \nactual  %#v %v",
				tc.msg, tc.exp, tc.expOk, strategies, ok)
		}
	}
}

func TestAssets(t *testing.T) {
	var testCases = []struct {
		msg      string
		strategy *Strategy
		exp      []*Asset
		expOk    bool
	}{
		{"test no children:",
			&Strategy{
				Node{name: "test", root: true},
				AlgoStack{}, nil, nil,
			},
			nil,
			false,
		},
		{"test no assets, only sub strategy:",
			&Strategy{
				Node{name: "test", root: true,
					children: []NodeHandler{
						&Strategy{
							Node{name: "subA", root: true},
							AlgoStack{}, nil, nil,
						},
					},
				},
				AlgoStack{}, nil, nil,
			},
			nil,
			false,
		},
		{"test single asset:",
			&Strategy{
				Node{name: "test", root: true,
					children: []NodeHandler{
						&Asset{
							Node{name: "asset", root: false},
						},
					},
				},
				AlgoStack{}, nil, nil,
			},
			[]*Asset{
				{
					Node{name: "asset", root: false},
				},
			},
			true,
		},
		{"test multiple assetss:",
			&Strategy{
				Node{name: "test", root: true,
					children: []NodeHandler{
						&Asset{
							Node{name: "assetA", root: false},
						},
						&Asset{
							Node{name: "assetB", root: false},
						},
					},
				},
				AlgoStack{}, nil, nil,
			},
			[]*Asset{
				{
					Node{name: "assetA", root: false},
				},
				{
					Node{name: "assetB", root: false},
				},
			},
			true,
		},
		{"test multiple sub strategies and multiple assets:",
			&Strategy{
				Node{name: "test", root: true,
					children: []NodeHandler{
						&Strategy{
							Node{name: "subA", root: false},
							AlgoStack{}, nil, nil,
						},
						&Strategy{
							Node{name: "subB", root: false},
							AlgoStack{}, nil, nil,
						},
						&Asset{
							Node{name: "assetA", root: false},
						},
						&Asset{
							Node{name: "assetB", root: false},
						},
					},
				},
				AlgoStack{}, nil, nil,
			},
			[]*Asset{
				{
					Node{name: "assetA", root: false},
				},
				{
					Node{name: "assetB", root: false},
				},
			},
			true,
		},
	}

	for _, tc := range testCases {
		strategies, ok := tc.strategy.Assets()
		if !reflect.DeepEqual(strategies, tc.exp) || (ok != tc.expOk) {
			t.Errorf("%v Strategies(): \nexpected %#v %v, \nactual  %#v %v",
				tc.msg, tc.exp, tc.expOk, strategies, ok)
		}
	}
}
func TestStrategySingleSetAlgo(t *testing.T) {
	testStrategy := &Strategy{
		algos: AlgoStack{
			stack: []AlgoHandler{
				&TrueAlgo{},
			},
		},
	}
	strategy := &Strategy{}
	strategy = strategy.SetAlgo(&TrueAlgo{})

	if !reflect.DeepEqual(strategy, testStrategy) {
		t.Errorf("set single algo SetAlgo(): \nexpected %#v, \nactual %#v", testStrategy, strategy)
	}
}

func TestStrategyMultipleSetAlgo(t *testing.T) {
	testStrategy := &Strategy{
		algos: AlgoStack{
			stack: []AlgoHandler{
				&TrueAlgo{},
				&TrueAlgo{},
				&FalseAlgo{},
			},
		},
	}
	strategy := &Strategy{}
	strategy = strategy.SetAlgo(
		&TrueAlgo{},
		&TrueAlgo{},
		&FalseAlgo{},
	)

	if !reflect.DeepEqual(strategy, testStrategy) {
		t.Errorf("set multiple algos SetAlgo(): \nexpected %#v, \nactual %#v", testStrategy, strategy)
	}
}
