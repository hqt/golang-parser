package example

import (
	"fmt"
	"strconv"
)

const (
	StateEnterMethod    = 0
	StateEnterReceiver  = 1
	StateEnterSignature = 3
	StateEnterArgument  = 4
)

type State struct {
	step uint64
	curr uint
}

func NewState() *State {
	return &State{
		step: 0,
		curr: 0,
	}
}
func (s *State) NextState() {
	s.step = s.step & 1 >> (s.curr + 1)
	s.curr += 1
}

func (s *State) Enable(state uint) {
	s.step = s.step | 1<<state
}

func (s *State) Disable(state uint) {
	mask := ^(uint64(1) << state)
	s.step = s.step & mask
}

func (s *State) IsEnable(state uint64) bool {
	return (s.step & (1 << state)) > 0
}

func (s *State) Print() {
	fmt.Println(strconv.FormatInt(int64(s.step), 2))
}

type MethodDescription struct {
	FuncName     string
	Receiver     string
	ReceiverType string
	Parameters   []string
	Types        []string
}

func NewMethodDescription() *MethodDescription {
	return &MethodDescription{
		Parameters: []string{},
		Types:      []string{},
	}

}
func (m *MethodDescription) addParameter(param string) {
	m.Parameters = append(m.Parameters, param)
}

func (m *MethodDescription) addType(dataType string) {
	m.Types = append(m.Types, dataType)
}

func (m *MethodDescription) String() string {
	str := fmt.Sprintf("func (%s: %s) %s ", m.Receiver, m.ReceiverType, m.FuncName)
	arguments := "("
	for i := 0; i < len(m.Parameters); i++ {
		arguments = arguments + m.Parameters[i] + " " + m.Types[i] + " ,"
	}
	arguments += ")"
	return str + arguments
}
