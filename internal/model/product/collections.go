package product

import (
	"github.com/garrettladley/prods/internal/model/category"
	"github.com/garrettladley/prods/internal/rand"
	exprand "golang.org/x/exp/rand"
)

func ChooseIDs(seed uint64, n uint) []string {
	return rand.ChooseN(seed, n, IDs[:]...)
}

var categoryToProducts map[category.Category][]string

func init() {
	m := make(map[category.Category][]string)
	for _, productID := range IDs {
		product := Products[productID]
		for _, cat := range product.Categories {
			m[cat] = append(m[cat], productID)
		}
	}
	categoryToProducts = m
}

// selects n product IDs ensuring all categories are represented.
// returns nil if n is less than the number of categories.
func ChooseIDsRepresentingAllCategories(seed uint64, n uint) []string {
	if n < uint(len(category.Categories)) {
		return nil
	}

	exprand.Seed(seed)
	selected := make(map[string]struct{}, n)
	result := make([]string, 0, n)

	// first, ensure at least one product from each category is selected
	if !selectOnePerCategory(selected, &result) {
		return nil // could not find products for all categories
	}

	// fill remaining slots with random products
	fillRemainingSlots(selected, &result, int(n))

	return result
}

// attempt to select one product from each category
func selectOnePerCategory(selected map[string]struct{}, result *[]string) bool {
	for _, cat := range category.Categories {
		products := categoryToProducts[cat]
		if len(products) == 0 {
			return false
		}

		if !selectProductFromCategory(products, selected, result) {
			return false
		}
	}
	return true
}

// attempt to select an unselected product from the given category
func selectProductFromCategory(products []string, selected map[string]struct{}, result *[]string) bool {
	// try to find an unselected product
	for _, i := range exprand.Perm(len(products)) {
		if _, ok := selected[products[i]]; !ok {
			selected[products[i]] = struct{}{}
			*result = append(*result, products[i])
			return true
		}
	}

	// if all products in category were selected, pick any product from category
	randProduct := products[exprand.Intn(len(products))]
	if _, ok := selected[randProduct]; !ok {
		selected[randProduct] = struct{}{}
		*result = append(*result, randProduct)
		return true
	}

	return false
}

// fill the remaining slots with random unselected products
func fillRemainingSlots(selected map[string]struct{}, result *[]string, targetCount int) {
	if len(*result) >= targetCount {
		return
	}

	remainingProducts := make([]string, 0, len(IDs)-len(selected))
	for _, id := range IDs {
		if _, ok := selected[id]; !ok {
			remainingProducts = append(remainingProducts, id)
		}
	}

	numToAdd := min(targetCount-len(*result), len(remainingProducts))
	if numToAdd <= 0 {
		return
	}

	perm := exprand.Perm(len(remainingProducts))[:numToAdd]
	for _, i := range perm {
		*result = append(*result, remainingProducts[i])
	}
}

var IDs = [100]string{
	AirPodsPro.ID,
	GooglePixel7.ID,
	NikeAirForce1.ID,
	SamsungGalaxyS22.ID,
	AppleWatchSeries8.ID,
	LeviBluejeans.ID,
	DysonVacuum.ID,
	NikeAirMax.ID,
	OreoCookies.ID,
	OfficeChair.ID,
	TreadmillMachine.ID,
	SurfboardBrand.ID,
	ToiletPaper.ID,
	LogitechMouse.ID,
	KitchenAidMixer.ID,
	GoPro.ID,
	InstaCartGiftCard.ID,
	AmazonEchoShow.ID,
	SonyWH1000XM4.ID,
	SwissArmyKnife.ID,
	HondaCivic.ID,
	LeviJeanJacket.ID,
	SamsungQN90BTV.ID,
	DellXPS13Laptop.ID,
	AmazonKindlePaperwhite.ID,
	NintendoSwitch.ID,
	AppleMacbookPro.ID,
	AdidasUltraboost.ID,
	ZeissCamera.ID,
	KeurigKoffeemaker.ID,
	SonyPlayStation5.ID,
	YetiTumbler.ID,
	TeslaModelS.ID,
	TCLSmartTV.ID,
	CuisinartGrill.ID,
	UnderArmourHoodie.ID,
	SamsungFridge.ID,
	HPDeskJetPrinter.ID,
	ConverseChuckTaylor.ID,
	CanonEOSR.ID,
	PhilipsElectricToothbrush.ID,
	NorthFaceBackpack.ID,
	InstantPotPressureCooker.ID,
	AsusROGZephyrus.ID,
	TimberlandBoots.ID,
	DumbbellSet.ID,
	BeatsHeadphones.ID,
	RokuStreamingStick.ID,
	PendletonBlanket.ID,
	FitbitCharge5.ID,
	LegoCreatorExpert.ID,
	ColemanTent.ID,
	TheragunElite.ID,
	TomFordCologne.ID,
	BoseSoundLink.ID,
	NordicTrackBike.ID,
	LululemonYogaMat.ID,
	SmegToaster.ID,
	DeWaltDrill.ID,
	PatagoniaJacket.ID,
	GoodyearTires.ID,
	CreedAventusCologne.ID,
	FujifilmX100V.ID,
	HydroFlaskBottle.ID,
	SonyA7Camera.ID,
	CuisinartAirFryer.ID,
	MonopolyGame.ID,
	BrevilleCoffeeMaker.ID,
	FitbitVersa3.ID,
	ReebokCrossfit.ID,
	GarminFenix6.ID,
	JBLFlip5Speaker.ID,
	ChampionSweatshirt.ID,
	LegoTechnicBugatti.ID,
	RolexSubmariner.ID,
	KindlePaperwhiteKids.ID,
	SonosOneSL.ID,
	BanwoodBalanceBike.ID,
	NorthFacePuffer.ID,
	SkullcandyHeadphones.ID,
	CampChefGrill.ID,
	HyperXCloudHeadset.ID,
	ColgateToothpaste.ID,
	BrooksRunningShoes.ID,
	KindleFireTablet.ID,
	DysonAirWrap.ID,
	CanonInkCartridge.ID,
	GoogleNestThermostat.ID,
	SauconyTrailShoes.ID,
	AppleMagicKeyboard.ID,
	StanleyThermos.ID,
	OXOCoffeeGrinder.ID,
	HarryPotterBoxSet.ID,
	AllCladPan.ID,
	SchwinnBike.ID,
	WaterpikFlosser.ID,
	PamperedChefStoneware.ID,
	PlaymobilCastle.ID,
	GatoradePack.ID,
	HersheysChocolateBar.ID,
}

var Products = map[string]Product{
	AirPodsPro.ID:                AirPodsPro,
	GooglePixel7.ID:              GooglePixel7,
	NikeAirForce1.ID:             NikeAirForce1,
	SamsungGalaxyS22.ID:          SamsungGalaxyS22,
	AppleWatchSeries8.ID:         AppleWatchSeries8,
	LeviBluejeans.ID:             LeviBluejeans,
	DysonVacuum.ID:               DysonVacuum,
	NikeAirMax.ID:                NikeAirMax,
	OreoCookies.ID:               OreoCookies,
	OfficeChair.ID:               OfficeChair,
	TreadmillMachine.ID:          TreadmillMachine,
	SurfboardBrand.ID:            SurfboardBrand,
	ToiletPaper.ID:               ToiletPaper,
	LogitechMouse.ID:             LogitechMouse,
	KitchenAidMixer.ID:           KitchenAidMixer,
	GoPro.ID:                     GoPro,
	InstaCartGiftCard.ID:         InstaCartGiftCard,
	AmazonEchoShow.ID:            AmazonEchoShow,
	SonyWH1000XM4.ID:             SonyWH1000XM4,
	SwissArmyKnife.ID:            SwissArmyKnife,
	HondaCivic.ID:                HondaCivic,
	LeviJeanJacket.ID:            LeviJeanJacket,
	SamsungQN90BTV.ID:            SamsungQN90BTV,
	DellXPS13Laptop.ID:           DellXPS13Laptop,
	AmazonKindlePaperwhite.ID:    AmazonKindlePaperwhite,
	NintendoSwitch.ID:            NintendoSwitch,
	AppleMacbookPro.ID:           AppleMacbookPro,
	AdidasUltraboost.ID:          AdidasUltraboost,
	ZeissCamera.ID:               ZeissCamera,
	KeurigKoffeemaker.ID:         KeurigKoffeemaker,
	SonyPlayStation5.ID:          SonyPlayStation5,
	YetiTumbler.ID:               YetiTumbler,
	TeslaModelS.ID:               TeslaModelS,
	TCLSmartTV.ID:                TCLSmartTV,
	CuisinartGrill.ID:            CuisinartGrill,
	UnderArmourHoodie.ID:         UnderArmourHoodie,
	SamsungFridge.ID:             SamsungFridge,
	HPDeskJetPrinter.ID:          HPDeskJetPrinter,
	ConverseChuckTaylor.ID:       ConverseChuckTaylor,
	CanonEOSR.ID:                 CanonEOSR,
	PhilipsElectricToothbrush.ID: PhilipsElectricToothbrush,
	NorthFaceBackpack.ID:         NorthFaceBackpack,
	InstantPotPressureCooker.ID:  InstantPotPressureCooker,
	AsusROGZephyrus.ID:           AsusROGZephyrus,
	TimberlandBoots.ID:           TimberlandBoots,
	DumbbellSet.ID:               DumbbellSet,
	BeatsHeadphones.ID:           BeatsHeadphones,
	RokuStreamingStick.ID:        RokuStreamingStick,
	PendletonBlanket.ID:          PendletonBlanket,
	FitbitCharge5.ID:             FitbitCharge5,
	LegoCreatorExpert.ID:         LegoCreatorExpert,
	ColemanTent.ID:               ColemanTent,
	TheragunElite.ID:             TheragunElite,
	TomFordCologne.ID:            TomFordCologne,
	BoseSoundLink.ID:             BoseSoundLink,
	NordicTrackBike.ID:           NordicTrackBike,
	LululemonYogaMat.ID:          LululemonYogaMat,
	SmegToaster.ID:               SmegToaster,
	DeWaltDrill.ID:               DeWaltDrill,
	PatagoniaJacket.ID:           PatagoniaJacket,
	GoodyearTires.ID:             GoodyearTires,
	CreedAventusCologne.ID:       CreedAventusCologne,
	FujifilmX100V.ID:             FujifilmX100V,
	HydroFlaskBottle.ID:          HydroFlaskBottle,
	SonyA7Camera.ID:              SonyA7Camera,
	CuisinartAirFryer.ID:         CuisinartAirFryer,
	MonopolyGame.ID:              MonopolyGame,
	BrevilleCoffeeMaker.ID:       BrevilleCoffeeMaker,
	FitbitVersa3.ID:              FitbitVersa3,
	ReebokCrossfit.ID:            ReebokCrossfit,
	GarminFenix6.ID:              GarminFenix6,
	JBLFlip5Speaker.ID:           JBLFlip5Speaker,
	ChampionSweatshirt.ID:        ChampionSweatshirt,
	LegoTechnicBugatti.ID:        LegoTechnicBugatti,
	RolexSubmariner.ID:           RolexSubmariner,
	KindlePaperwhiteKids.ID:      KindlePaperwhiteKids,
	SonosOneSL.ID:                SonosOneSL,
	BanwoodBalanceBike.ID:        BanwoodBalanceBike,
	NorthFacePuffer.ID:           NorthFacePuffer,
	SkullcandyHeadphones.ID:      SkullcandyHeadphones,
	CampChefGrill.ID:             CampChefGrill,
	HyperXCloudHeadset.ID:        HyperXCloudHeadset,
	ColgateToothpaste.ID:         ColgateToothpaste,
	BrooksRunningShoes.ID:        BrooksRunningShoes,
	KindleFireTablet.ID:          KindleFireTablet,
	DysonAirWrap.ID:              DysonAirWrap,
	CanonInkCartridge.ID:         CanonInkCartridge,
	GoogleNestThermostat.ID:      GoogleNestThermostat,
	SauconyTrailShoes.ID:         SauconyTrailShoes,
	AppleMagicKeyboard.ID:        AppleMagicKeyboard,
	StanleyThermos.ID:            StanleyThermos,
	OXOCoffeeGrinder.ID:          OXOCoffeeGrinder,
	HarryPotterBoxSet.ID:         HarryPotterBoxSet,
	AllCladPan.ID:                AllCladPan,
	SchwinnBike.ID:               SchwinnBike,
	WaterpikFlosser.ID:           WaterpikFlosser,
	PamperedChefStoneware.ID:     PamperedChefStoneware,
	PlaymobilCastle.ID:           PlaymobilCastle,
	GatoradePack.ID:              GatoradePack,
	HersheysChocolateBar.ID:      HersheysChocolateBar,
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
