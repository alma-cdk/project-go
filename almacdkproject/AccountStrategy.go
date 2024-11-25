package almacdkproject

import (
	_init_ "github.com/alma-cdk/project-go/almacdkproject/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Use static methods of `AccountStrategy` abstract class to define your account strategy.
//
// Available strategies are:
// - One Account: `shared`
// - Two Accounts: `dev`+`prod` – _recommended_
// - Three Accounts: `dev`+`preprod`+`prod`.
type AccountStrategy interface {
}

// The jsii proxy struct for AccountStrategy
type jsiiProxy_AccountStrategy struct {
	_ byte // padding
}

func NewAccountStrategy_Override(a AccountStrategy) {
	_init_.Initialize()

	_jsii_.Create(
		"@alma-cdk/project.AccountStrategy",
		nil, // no parameters
		a,
	)
}

// Enables single account strategy.
//
// 1. `shared` account with environments:
//    - development
//    - feature/*
//    - test
//    - qaN
//    - staging
//    - preproduction
// - production.
//
// Example:
//   AccountStrategy.one({
//     shared: {
//       id: '111111111111',
//     },
//   }),
//
func AccountStrategy_One(props *AccountStrategyOneProps) *map[string]*Account {
	_init_.Initialize()

	if err := validateAccountStrategy_OneParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]*Account

	_jsii_.StaticInvoke(
		"@alma-cdk/project.AccountStrategy",
		"one",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Enables triple account strategy.
//
// 1. `dev` account with environments:
//    - development
//    - feature/*
//    - test
//    - staging
// 2. `preprod` account with environments:
//    - qaN
//    - preproduction
// 3. `prod` account with environments:
// - production.
//
// Example:
//   AccountStrategy.three({
//     dev: {
//       id: '111111111111',
//     },
//     preprod: {
//       id: '222222222222',
//     },
//     prod: {
//       id: '333333333333',
//     },
//   }),
//
func AccountStrategy_Three(props *AccountStrategyThreeProps) *map[string]*Account {
	_init_.Initialize()

	if err := validateAccountStrategy_ThreeParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]*Account

	_jsii_.StaticInvoke(
		"@alma-cdk/project.AccountStrategy",
		"three",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Enables dual account strategy.
//
// 1. `dev` account with environments:
//    - development
//    - feature/*
//    - test
//    - qaN
//    - staging
// 2. `prod` account with environments:
//    - preproduction
// - production.
//
// Example:
//   AccountStrategy.two({
//     dev: {
//       id: '111111111111',
//     },
//     prod: {
//       id: '222222222222',
//     },
//   }),
//
func AccountStrategy_Two(props *AccountStrategyTwoProps) *map[string]*Account {
	_init_.Initialize()

	if err := validateAccountStrategy_TwoParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]*Account

	_jsii_.StaticInvoke(
		"@alma-cdk/project.AccountStrategy",
		"two",
		[]interface{}{props},
		&returns,
	)

	return returns
}

