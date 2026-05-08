package engine

import "fmt"

//step 2
//engine is responsible for managing and executing different prime strategies
//it acts as central registry where strategies are stored and retrieved


type Engine struct{
		strategies map[string]PrimeStrategy   //strategies { steve:steieve}
}


//New engine intializes and return a new engine instance
func NewEngine() *Engine{
	return &Engine{
      strategies: make(map[string]PrimeStrategy),
	}
}

//Registers add new strategy to a engine
//the staregy is stored using its name as the key
func (e *Engine) Register(strategy PrimeStrategy){
	e.strategies[strategy.Name()]=strategy
}


//Execute selects the appropriate strategy based on algo name
//and runs with provided range
func(e *Engine) Execute(algo string,start,end int)([]int,error){
       strategy,exists:=e.strategies[algo]
	   	if !exists {
		return nil, fmt.Errorf("unknown algorithm: %s", algo)
	}
	return strategy.Generate(start,end)
}

func(e *Engine) ListOfStrategies() []string{
	   keys := make([]string, 0, len(e.strategies))
       for k:= range e.strategies{
         keys = append(keys, k)
	   }
	   return keys
}
