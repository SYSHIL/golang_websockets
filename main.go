package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Problem - Gopher RPG
// Create a program which has three structs:
// • Gopher
// • Weapon
// • Consumable

// Gopher should contain the following statistics about a gopher:
// • hitpoints - an integer
// • weapon - the Weapon equipped
// • inventory - a slice of Consumables
// • strength - an integer
// • agility - an integer
// • intellect - an integer
// • coins - an integer

type Gopher struct {
	name      string
	hitpoints int // 30
	weapon    Weapon
	//using a map
	inventory map[string]int
	strength  int // 0
	agility   int // 0
	intellect int // 0
	coins     int // 20
}

// Weapon should contain the following data:
// • damage - a slice of two integers, the interval of damage the weapon can
// deal
// • strengthReq - an integer, strength requirements to yield the weapon
// • agilityReq - an integer, strength requirements to yield the weapon
// • intelligenceReq - an integer, intellect requirements to yield the weapon

type Weapon struct {
	name            string
	cost            int
	damage          []int
	strengthReq     int
	agilityReq      int
	intelligenceReq int
}

// Consumable should contain the following data:
// • duration - an integer, turns for which the consumable is active
// • hitpointsEffect - an integer, the effect on hitpoints
// • strengthEffect - an integer, the effect on strength
// • agilityEffect - an integer, the effect on agility
// • intellectEffect - an integer, the effect on intellect

type Consumable struct {
	cost            int
	duration        int
	hitpointsEffect int
	strengthEffect  int
	agilityEffect   int
	intellectEffect int
}

// The game is a turn-based one. There are two gophers and they can each decide
// what to do on their turn. Each gopher starts with 30 hitpoints, 20 gold and all
// their attributes are 0.
// The game ends when one of the gophers dies. A gopher dies when their hitpoints drop to 0 or below.

// turn returns false when game is over
func turn(player *Gopher, opponent *Gopher) bool {

	// clear screen
	fmt.Println("-------GOPHERSCAPE-------")

	// display player1's health and player2's health
	fmt.Println(player.name+"'s health: ", player.hitpoints)
	fmt.Println(opponent.name+"'s health: ", opponent.hitpoints)

	// display player1's weapon and player2's weapon name
	fmt.Println(player.name+"'s weapon: ", player.weapon.name)
	fmt.Println(opponent.name+"'s weapon: ", opponent.weapon.name)

	// display player1's attack and player2's attack
	fmt.Println(player.name+"'s attack: ", player.weapon.damage)
	fmt.Println(opponent.name+"'s attack: ", opponent.weapon.damage)

	// present all choices to the player via a menu
	fmt.Println("What do you want to do, " + player.name)
	fmt.Println("1. Attack")
	fmt.Println("2. Buy Consumable")
	fmt.Println("3. Buy Weapon")
	fmt.Println("4. Use Consumable")
	fmt.Println("5. Work")
	fmt.Println("6. Train")
	fmt.Println("7. Exit")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		player.attack(opponent)

	case 2:
		fmt.Println("What consumable do you want to buy?")
		fmt.Println("1. health_potion")
		fmt.Println("2. strength_potion")
		fmt.Println("3. agility_potion")
		fmt.Println("4. intellect_potion")
		var cName string
		fmt.Scanln(&cName)
		if player.buyConsumable(consumables[cName], cName) == false {
			// call the function again
			turn(player, opponent)
		}

	case 3:
		fmt.Println("What weapon do you want to buy?")
		fmt.Println("1. barehand")
		fmt.Println("2. knife")
		fmt.Println("3. sword")
		fmt.Println("4. ninjaku")
		fmt.Println("5. wand")
		fmt.Println("6. gophermourne")
		var wName string
		fmt.Scanln(&wName)
		// check if weapon exists inside the map
		if _, ok := weapons[wName]; !ok {
			fmt.Println("Weapon doesn't exist")
			turn(player, opponent)
		}
		if player.buyWeapon(weapons[wName], wName) == false {
			// call the function again
			turn(player, opponent)
		}

	case 4:
		fmt.Println("What consumable do you want to use?")
		fmt.Println("1. health_potion")
		fmt.Println("2. strength_potion")
		fmt.Println("3. agility_potion")
		fmt.Println("4. intellect_potion")
		var cName string
		fmt.Scanln(&cName)
		player.use(consumables[cName], cName)

	case 5:
		player.work()

	case 6:
		fmt.Println("What attribute do you want to train?")
		fmt.Println("1. strength")
		fmt.Println("2. agility")
		fmt.Println("3. intellect")
		var attribute string
		fmt.Scanln(&attribute)
		if player.train(attribute) == false {
			turn(player, opponent)
		}

	case 7:
		fmt.Println("Closing game")
		return false
	}
	// opponent's hitpoints less than zero means current player wins
	if opponent.hitpoints <= 0 {
		fmt.Println("Winner is player ", player.name)
		return false
	}
	return true
}

func gameLoop() {
	// two gopher objects needed who switch turns
	// get player names
	fmt.Println("Enter player 1 name")
	var name1 string
	fmt.Scanln(&name1)
	fmt.Println("Enter player 2 name")
	var name2 string
	fmt.Scanln(&name2)

	g1 := &Gopher{name: name1, hitpoints: 30, weapon: barehand, inventory: make(map[string]int), strength: 0, agility: 0, intellect: 0, coins: 20}
	g2 := &Gopher{name: name2, hitpoints: 1, weapon: barehand, inventory: make(map[string]int), strength: 0, agility: 0, intellect: 0, coins: 20}

	turnNo := 0
	var cont bool = true
	for true {
		if turnNo%2 == 0 {
			cont = turn(g1, g2)

		} else {
			cont = turn(g2, g1)
		}
		if cont == false {
			return
		}
		turnNo++
	}
}
func main() {
	fmt.Println("-------GOPHERSCAPE-------")
	gameLoop()
}

type actions interface {
	attack()
	buyConsumable()
	buyWeapon()
	use()
	work()
	train()
}

// gopher implements the actions interface

//creating of actions interface

// The possible actions are:
// • Choose the actual damage dealt at random based on the weapon's
// damage interval
func (g Gopher) attack(opp *Gopher) {
	// attack - attack the other gopher with the weapon you have equipped at
	// the moment. If none is equipped, then you are attacking bare-handed for a
	// damage of 1 hitpoint.
	damageInterval := g.weapon.damage
	damage := rand.Intn(damageInterval[1]-damageInterval[0]+1) + damageInterval[0]
	opp.hitpoints -= damage

}

// • buy «item> - spend the coins necessary to buy the given item based on its
func (g *Gopher) buyConsumable(c Consumable, cName string) bool {
	if g.coins < c.cost {
		fmt.Println("You don't have enough coins to buy this item")
		return false
	}
	g.coins -= c.cost
	g.inventory[cName]++
	return true
}

func (g *Gopher) buyWeapon(w Weapon, wName string) bool {
	if g.coins < w.cost {
		fmt.Println("You don't have enough coins to buy this item")
		time.Sleep(1 * time.Second)
		return false
	}
	// cost. If the item you bought is a weapon, you equip it over your current
	// weapon. You can't buy weapons you are still illegible to use due to
	// insufficient stats.

	if g.strength < w.strengthReq || g.agility < w.agilityReq || g.intellect < w.intelligenceReq {
		fmt.Println("Insufficient stats")
		return false
	}
	g.coins -= w.cost
	g.weapon = w
	return true
}

// • use ‹ item> - use one of the consumables in your inventory and apply its
// effects
func (g *Gopher) use(c Consumable, cName string) bool {
	if g.inventory[cName] < 0 {
		fmt.Println("You don't have this item in your inventory")
		time.Sleep(2 * time.Second)
		return false
	}
	g.inventory[cName]--
	g.hitpoints += c.hitpointsEffect
	g.strength += c.strengthEffect
	g.agility += c.agilityEffect
	g.intellect += c.intellectEffect
	return true
}

// • work - spend the turn working for the local warlord and gain anywhere
// between 5 and 15 coins (picked at random)
func (g *Gopher) work() {
	coins := rand.Intn(15-5+1) + 5
	g.coins += coins
}

func (g *Gopher) train(attribute string) bool {
	// • train < skill> - train a given attribute (strength, agility or intellect) and
	// increase it by 2. Training costs 5 gold.
	if g.coins < 5 {
		fmt.Println("You don't have enough coins to train")
		time.Sleep(2 * time.Second)
		return false
	}
	const INCPNTS = 2
	switch attribute {
	case "strength":
		g.strength += INCPNTS
	case "agility":
		g.agility += INCPNTS
	case "intellect":
		g.intellect += INCPNTS
	}
	return true
}

// The shop has the following items for sale with unlimited supply of them:

// Consumables:
// • health_potion - consumable - 5 gold
// •  	duration – permanent
// •  	 hitpointsEffect - 5
var health_potion = Consumable{cost: 5, duration: -1, hitpointsEffect: 5, strengthEffect: 0, agilityEffect: 0, intellectEffect: 0}

// • strength_potion - consumable - 10 gold
// • duration - 3 turns
// • strengthEffect - 3
var strength_potion = Consumable{cost: 10, duration: 3, hitpointsEffect: 0, strengthEffect: 3, agilityEffect: 0, intellectEffect: 0}

// • agility_potion - consumable - 10 gold
// • duration - 3 turns
// • agilityEffect - 3
var agility_potion = Consumable{cost: 10, duration: 3, hitpointsEffect: 0, strengthEffect: 0, agilityEffect: 3, intellectEffect: 0}

// • intellect_potion - consumable - 10 gold
// • duration - 3 turns
// • intellectEffect - 3

var intellect_potion = Consumable{cost: 10, duration: 3, hitpointsEffect: 0, strengthEffect: 0, agilityEffect: 0, intellectEffect: 3}

// make a map that maps string consumable to consumable object
var consumables = map[string]Consumable{"health_potion": health_potion, "strength_potion": strength_potion, "agility_potion": agility_potion, "intellect_potion": intellect_potion}

// Weapons:
// barehand - weapon - 0 gold
// damage -[1,1]
// all requirements zero
var barehand = Weapon{name: "barehand", cost: 0, damage: []int{1, 1}, strengthReq: 0, agilityReq: 0, intelligenceReq: 0}

// • knife - weapon - 10 gold
// • damage - [2-3]
// • all requirements are 0
var knife = Weapon{name: "knife", cost: 10, damage: []int{2, 3}, strengthReq: 0, agilityReq: 0, intelligenceReq: 0}

// • sword - weapon - 35 gold
// • damage - [3-5]
// • strengthReg - 2

var sword = Weapon{name: "sword", cost: 35, damage: []int{3, 5}, strengthReq: 2, agilityReq: 0, intelligenceReq: 0}

// • ninjaku - weapon - 25 gold
// • damage - [1-7]
// • agilityReg - 2

var ninjaku = Weapon{name: "ninjaku", cost: 25, damage: []int{1, 7}, strengthReq: 0, agilityReq: 2, intelligenceReq: 0}

// • wand - weapon - 30 gold
// • damage - [3-3]
// • intellectReg - 2

var wand = Weapon{name: "wand", cost: 30, damage: []int{3, 3}, strengthReq: 0, agilityReq: 0, intelligenceReq: 2}

// • gophermourne - weapon - 65 gold
// • damage - [6-7]
// •  	strenthReq - 3
// •  	intellectReq – 2
var gophermourne = Weapon{name: "gophermourne", cost: 65, damage: []int{6, 7}, strengthReq: 3, agilityReq: 0, intelligenceReq: 2}

// make a map that maps string weapon to Weapon object
var weapons = map[string]Weapon{"barehand": barehand, "knife": knife, "sword": sword, "ninjaku": ninjaku, "wand": wand, "gophermourne": gophermourne}
