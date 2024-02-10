/////////////////////////////
/// GENERATED DO NOT EDIT ///
/////////////////////////////

package models

import (
	"github.com/Codename-Recon/Codename-Recon/backend/pkg/enums/terrain"
)

var Terrains = [terrain.SIZE]Terrain{
	{
		Name:         "AIRPORT",
		Description:  "An airport is a property that is used to build and maintain air units.",
		Type:         terrain.BASE,
		Defense:      3,
		Health:       0,
		CanCapture:   true,
		CaptureValue: 0,
		Owner:        0,
		Funds:        1000,
		Ammo:         0,
	},
	{
		Name:         "BASE",
		Description:  "Bases are a property for building and maintaining ground units.",
		Type:         terrain.CITY,
		Defense:      3,
		Health:       0,
		CanCapture:   true,
		CaptureValue: 0,
		Owner:        0,
		Funds:        1000,
		Ammo:         0,
	},
	{
		Name:         "BRIDGE",
		Description:  "Bridges allow ground units to cross rivers and seas unhindered.",
		Type:         terrain.BRIDGE,
		Defense:      0,
		Health:       0,
		CanCapture:   false,
		CaptureValue: 0,
		Owner:        0,
		Funds:        0,
		Ammo:         0,
	},
	{
		Name:         "CITY",
		Description:  "Cities are a property essential to any battle and provide many advantages when captured, such as repairs and income.",
		Type:         terrain.SHOAL,
		Defense:      3,
		Health:       0,
		CanCapture:   true,
		CaptureValue: 0,
		Owner:        0,
		Funds:        1000,
		Ammo:         0,
	},
	{
		Name:         "COM_TOWER",
		Description:  "A Com Tower (Communication Tower) is a property for enhancing the skills of allied units. Each captured tower increases its allied units' attack strength by 10%.",
		Type:         terrain.SILO,
		Defense:      3,
		Health:       0,
		CanCapture:   true,
		CaptureValue: 0,
		Owner:        0,
		Funds:        1000,
		Ammo:         0,
	},
	{
		Name:         "HQ",
		Description:  "The most importnat property on the map; if the HQ is captured the team is defeated and can no longer play.",
		Type:         terrain.PORT,
		Defense:      4,
		Health:       0,
		CanCapture:   true,
		CaptureValue: 0,
		Owner:        0,
		Funds:        1000,
		Ammo:         0,
	},
	{
		Name:         "LAB",
		Description:  "Laboratories act as research centers and can unlock restricted units.",
		Type:         terrain.COM_TOWER,
		Defense:      3,
		Health:       0,
		CanCapture:   true,
		CaptureValue: 0,
		Owner:        0,
		Funds:        1000,
		Ammo:         0,
	},
	{
		Name:         "MOUNTAIN",
		Description:  "Montains provide a large amount of defence, but can only be traversed by infantry and air units. They also increase the vision range of any soldier unit on them by 3.",
		Type:         terrain.MOUNTAIN,
		Defense:      4,
		Health:       0,
		CanCapture:   false,
		CaptureValue: 0,
		Owner:        0,
		Funds:        0,
		Ammo:         0,
	},
	{
		Name:         "PIPE",
		Description:  "Pipelines are nigh-impassable terrain that are used for carrying materials.",
		Type:         terrain.HQ,
		Defense:      0,
		Health:       100,
		CanCapture:   false,
		CaptureValue: 0,
		Owner:        0,
		Funds:        0,
		Ammo:         0,
	},
	{
		Name:         "PLAIN",
		Description:  "Plains give a small amount of terrain cover and are easily traversed.",
		Type:         terrain.PLAIN,
		Defense:      1,
		Health:       0,
		CanCapture:   false,
		CaptureValue: 0,
		Owner:        0,
		Funds:        0,
		Ammo:         0,
	},
	{
		Name:         "PORT",
		Description:  "Ports are a property for building and maintaining sea units.",
		Type:         terrain.AIRPORT,
		Defense:      3,
		Health:       0,
		CanCapture:   true,
		CaptureValue: 0,
		Owner:        0,
		Funds:        1000,
		Ammo:         0,
	},
	{
		Name:         "REEF",
		Description:  "Reefs give some defence to sea units and also make excellent hiding places.",
		Type:         terrain.LAB,
		Defense:      1,
		Health:       0,
		CanCapture:   false,
		CaptureValue: 0,
		Owner:        0,
		Funds:        0,
		Ammo:         0,
	},
	{
		Name:         "RIVER",
		Description:  "Rivers restrict the movement of ground units and can only be traversed by Infantry, Mechs and air units.",
		Type:         terrain.RIVER,
		Defense:      0,
		Health:       0,
		CanCapture:   false,
		CaptureValue: 0,
		Owner:        0,
		Funds:        0,
		Ammo:         0,
	},
	{
		Name:         "ROAD",
		Description:  "Roads provide no defensive cover, but allow unimpeded movement for all ground units.",
		Type:         terrain.ROAD,
		Defense:      0,
		Health:       0,
		CanCapture:   false,
		CaptureValue: 0,
		Owner:        0,
		Funds:        0,
		Ammo:         0,
	},
	{
		Name:         "SEA",
		Description:  "Seas make up the majority of the terrain that navies battle on.",
		Type:         terrain.BRIDGE,
		Defense:      0,
		Health:       0,
		CanCapture:   false,
		CaptureValue: 0,
		Owner:        0,
		Funds:        0,
		Ammo:         0,
	},
	{
		Name:         "SHOAL",
		Description:  "Shoals provide an area for Landers and Black Boats to load and drop off units to battle.",
		Type:         terrain.SEA,
		Defense:      0,
		Health:       0,
		CanCapture:   false,
		CaptureValue: 0,
		Owner:        0,
		Funds:        0,
		Ammo:         0,
	},
	{
		Name:         "SILO",
		Description:  "Missile silos are activated by an Infantry or Mech unit. The missile they fire has a 5x5 square diamond shape and does 3HP of damage to any unit caught in the blast and doesn't distinguish between ally or enemy. Each silo can only be used once.",
		Type:         terrain.PIPE,
		Defense:      3,
		Health:       0,
		CanCapture:   false,
		CaptureValue: 0,
		Owner:        0,
		Funds:        0,
		Ammo:         1,
	},
	{
		Name:         "WOOD",
		Description:  "Woods are more difficult to traverse than plains, but offer better terrain cover and serve as a hiding spot in Fog of War.",
		Type:         terrain.WOOD,
		Defense:      2,
		Health:       0,
		CanCapture:   false,
		CaptureValue: 0,
		Owner:        0,
		Funds:        0,
		Ammo:         1,
	},
}