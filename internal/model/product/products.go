package product

import "github.com/garrettladley/prods/internal/model/category"

var (
	AirPodsPro = Product{
		ID:         "ABC12345",
		Name:       "AirPods Pro",
		Categories: []category.Category{category.Electronics},
		Stars:      480,
		Price:      24900,
	}

	GooglePixel7 = Product{
		ID:         "DEF67890",
		Name:       "Google Pixel 7",
		Categories: []category.Category{category.Electronics},
		Stars:      421,
		Price:      59900,
	}

	NikeAirForce1 = Product{
		ID:         "GHI13579",
		Name:       "Nike Air Force 1",
		Categories: []category.Category{category.Apparel, category.Sports},
		Stars:      269,
		Price:      10000,
	}

	SamsungGalaxyS22 = Product{
		ID:         "JKL24680",
		Name:       "Samsung Galaxy S22",
		Categories: []category.Category{category.Electronics},
		Stars:      375,
		Price:      79900,
	}

	AppleWatchSeries8 = Product{
		ID:         "MNO13579",
		Name:       "Apple Watch Series 8",
		Categories: []category.Category{category.Electronics},
		Stars:      491,
		Price:      39900,
	}

	LeviBluejeans = Product{
		ID:         "PQR24680",
		Name:       "Levi's 501 Original Fit Jeans",
		Categories: []category.Category{category.Apparel},
		Stars:      449,
		Price:      6000,
	}

	DysonVacuum = Product{
		ID:         "STU35791",
		Name:       "Dyson V11 Torque Drive Vacuum",
		Categories: []category.Category{category.HomeGoods},
		Stars:      187,
		Price:      49900,
	}

	NikeAirMax = Product{
		ID:         "VWX46802",
		Name:       "Nike Air Max 270",
		Categories: []category.Category{category.Apparel, category.Sports},
		Stars:      174,
		Price:      12000,
	}

	OreoCookies = Product{
		ID:         "YZA13579",
		Name:       "Oreo Original Cookies",
		Categories: []category.Category{category.Grocery},
		Stars:      471,
		Price:      400,
	}

	OfficeChair = Product{
		ID:         "BCD24680",
		Name:       "IKEA Markus Office Chair",
		Categories: []category.Category{category.HomeGoods, category.OfficeSupplies},
		Stars:      322,
		Price:      19900,
	}

	TreadmillMachine = Product{
		ID:         "EFG35791",
		Name:       "NordicTrack T 6.5 S Treadmill",
		Categories: []category.Category{category.Sports, category.HomeGoods},
		Stars:      365,
		Price:      69900,
	}

	SurfboardBrand = Product{
		ID:         "HIJ46802",
		Name:       "NSP Elements CC Longboard Surfboard",
		Categories: []category.Category{category.Sports, category.Outdoor},
		Stars:      334,
		Price:      49900,
	}

	ToiletPaper = Product{
		ID:         "KLM13579",
		Name:       "Charmin Ultra Soft Toilet Paper",
		Categories: []category.Category{category.Grocery, category.HomeGoods},
		Stars:      281,
		Price:      600,
	}

	LogitechMouse = Product{
		ID:         "NOP24680",
		Name:       "Logitech MX Master 3 Advanced Wireless Mouse",
		Categories: []category.Category{category.Electronics, category.OfficeSupplies},
		Stars:      122,
		Price:      12000,
	}

	KitchenAidMixer = Product{
		ID:         "QRS35791",
		Name:       "KitchenAid Artisan Series Tilt-Head Stand Mixer",
		Categories: []category.Category{category.HomeGoods},
		Stars:      234,
		Price:      39900,
	}

	GoPro = Product{
		ID:         "TVU46802",
		Name:       "GoPro HERO11 Black",
		Categories: []category.Category{category.Electronics, category.Sports, category.Outdoor},
		Stars:      431,
		Price:      49900,
	}

	InstaCartGiftCard = Product{
		ID:         "WXY13579",
		Name:       "Instacart $50 Gift Card",
		Categories: []category.Category{category.Grocery},
		Stars:      369,
		Price:      5000,
	}

	AmazonEchoShow = Product{
		ID:         "ABC24680",
		Name:       "Amazon Echo Show 8 (2nd Gen)",
		Categories: []category.Category{category.Electronics, category.HomeGoods},
		Stars:      423,
		Price:      12900,
	}

	SonyWH1000XM4 = Product{
		ID:         "DEF35791",
		Name:       "Sony WH-1000XM4 Wireless Headphones",
		Categories: []category.Category{category.Electronics},
		Stars:      380,
		Price:      34900,
	}

	SwissArmyKnife = Product{
		ID:         "GHI46802",
		Name:       "Victorinox Swiss Army Huntsman Pocket Knife",
		Categories: []category.Category{category.Sports, category.Outdoor},
		Stars:      257,
		Price:      3900,
	}

	HondaCivic = Product{
		ID:         "JKL13579",
		Name:       "Honda Civic EX Sedan",
		Categories: []category.Category{category.Electronics, category.HomeGoods},
		Stars:      461,
		Price:      2500000,
	}

	LeviJeanJacket = Product{
		ID:         "MNO24680",
		Name:       "Levi's Trucker Jacket",
		Categories: []category.Category{category.Apparel},
		Stars:      267,
		Price:      9800,
	}

	SamsungQN90BTV = Product{
		ID:         "PQR35791",
		Name:       "Samsung 65-inch QN90B Neo QLED 4K Smart TV",
		Categories: []category.Category{category.Electronics, category.HomeGoods},
		Stars:      410,
		Price:      149900,
	}

	DellXPS13Laptop = Product{
		ID:         "STU46802",
		Name:       "Dell XPS 13 Laptop",
		Categories: []category.Category{category.Electronics},
		Stars:      443,
		Price:      129900,
	}

	AmazonKindlePaperwhite = Product{
		ID:         "VWX13579",
		Name:       "Amazon Kindle Paperwhite",
		Categories: []category.Category{category.Electronics},
		Stars:      472,
		Price:      13000,
	}

	NintendoSwitch = Product{
		ID:         "YZA24680",
		Name:       "Nintendo Switch OLED Model",
		Categories: []category.Category{category.Electronics, category.Sports},
		Stars:      481,
		Price:      35000,
	}

	AppleMacbookPro = Product{
		ID:         "BCD35791",
		Name:       "Apple MacBook Pro 14-inch",
		Categories: []category.Category{category.Electronics},
		Stars:      499,
		Price:      199900,
	}

	AdidasUltraboost = Product{
		ID:         "EFG46802",
		Name:       "Adidas Ultraboost 22 Running Shoes",
		Categories: []category.Category{category.Apparel, category.Sports},
		Stars:      478,
		Price:      18000,
	}

	ZeissCamera = Product{
		ID:         "HIJ13579",
		Name:       "Zeiss ZX1 Full-Frame Mirrorless Camera",
		Categories: []category.Category{category.Electronics},
		Stars:      471,
		Price:      249900,
	}

	KeurigKoffeemaker = Product{
		ID:         "KLM24680",
		Name:       "Keurig K-Supreme Plus Coffee Maker",
		Categories: []category.Category{category.HomeGoods},
		Stars:      260,
		Price:      17900,
	}

	SonyPlayStation5 = Product{
		ID:         "ABC87429",
		Name:       "Sony PlayStation 5",
		Categories: []category.Category{category.Electronics, category.Sports},
		Stars:      483,
		Price:      49900,
	}

	YetiTumbler = Product{
		ID:         "XYZ35791",
		Name:       "YETI Rambler 20 oz Tumbler",
		Categories: []category.Category{category.HomeGoods, category.Outdoor},
		Stars:      471,
		Price:      3000,
	}

	TeslaModelS = Product{
		ID:         "TES13579",
		Name:       "Tesla Model S",
		Categories: []category.Category{category.Electronics, category.HomeGoods},
		Stars:      295,
		Price:      7999000,
	}

	TCLSmartTV = Product{
		ID:         "TCL24680",
		Name:       "TCL 55-inch 4K UHD Smart TV",
		Categories: []category.Category{category.Electronics},
		Stars:      208,
		Price:      37900,
	}

	CuisinartGrill = Product{
		ID:         "GRI35781",
		Name:       "Cuisinart 5-in-1 Griddler",
		Categories: []category.Category{category.HomeGoods},
		Stars:      403,
		Price:      7000,
	}

	UnderArmourHoodie = Product{
		ID:         "HOOD2468",
		Name:       "Under Armour Men's Armour Fleece Hoodie",
		Categories: []category.Category{category.Apparel, category.Sports},
		Stars:      458,
		Price:      6500,
	}

	SamsungFridge = Product{
		ID:         "FRD35719",
		Name:       "Samsung 4-Door French Door Refrigerator",
		Categories: []category.Category{category.HomeGoods},
		Stars:      271,
		Price:      279900,
	}

	HPDeskJetPrinter = Product{
		ID:         "PRT46802",
		Name:       "HP DeskJet 2755e Wireless Printer",
		Categories: []category.Category{category.Electronics, category.OfficeSupplies},
		Stars:      426,
		Price:      10000,
	}

	ConverseChuckTaylor = Product{
		ID:         "CON13579",
		Name:       "Converse Chuck Taylor All Star",
		Categories: []category.Category{category.Apparel},
		Stars:      398,
		Price:      5500,
	}

	CanonEOSR = Product{
		ID:         "CAN24680",
		Name:       "Canon EOS R Mirrorless Camera",
		Categories: []category.Category{category.Electronics},
		Stars:      160,
		Price:      179900,
	}

	PhilipsElectricToothbrush = Product{
		ID:         "TOO35791",
		Name:       "Philips Sonicare 4100 Electric Toothbrush",
		Categories: []category.Category{category.HomeGoods},
		Stars:      271,
		Price:      4900,
	}

	NorthFaceBackpack = Product{
		ID:         "BAG13579",
		Name:       "The North Face Borealis Backpack",
		Categories: []category.Category{category.Apparel, category.Outdoor},
		Stars:      429,
		Price:      9800,
	}

	InstantPotPressureCooker = Product{
		ID:         "POT24680",
		Name:       "Instant Pot Duo Plus 9-in-1 Electric Pressure Cooker",
		Categories: []category.Category{category.HomeGoods},
		Stars:      485,
		Price:      12000,
	}

	AsusROGZephyrus = Product{
		ID:         "LPT35791",
		Name:       "ASUS ROG Zephyrus G14 Gaming Laptop",
		Categories: []category.Category{category.Electronics},
		Stars:      478,
		Price:      159900,
	}

	TimberlandBoots = Product{
		ID:         "SHO46802",
		Name:       "Timberland PRO Men's 6\" Pit Boss Soft Toe",
		Categories: []category.Category{category.Apparel, category.Outdoor},
		Stars:      412,
		Price:      14500,
	}

	DumbbellSet = Product{
		ID:         "FIT13579",
		Name:       "Bowflex SelectTech Adjustable Dumbbells",
		Categories: []category.Category{category.Sports, category.HomeGoods},
		Stars:      469,
		Price:      32900,
	}

	BeatsHeadphones = Product{
		ID:         "AUD24680",
		Name:       "Beats Studio3 Wireless Headphones",
		Categories: []category.Category{category.Electronics},
		Stars:      468,
		Price:      22000,
	}

	RokuStreamingStick = Product{
		ID:         "STR35791",
		Name:       "Roku Streaming Stick 4K",
		Categories: []category.Category{category.Electronics, category.HomeGoods},
		Stars:      439,
		Price:      5000,
	}

	PendletonBlanket = Product{
		ID:         "BLK46802",
		Name:       "Pendleton Yakima Camp Throw Blanket",
		Categories: []category.Category{category.HomeGoods, category.Outdoor},
		Stars:      472,
		Price:      18000,
	}

	FitbitCharge5 = Product{
		ID:         "FIT13580", // FIXME: ID is duplicated
		Name:       "Fitbit Charge 5 Fitness Tracker",
		Categories: []category.Category{category.Electronics, category.Sports},
		Stars:      412,
		Price:      12900,
	}

	LegoCreatorExpert = Product{
		ID:         "LEGO12345",
		Name:       "LEGO Creator Expert Roller Coaster",
		Categories: []category.Category{category.Toys},
		Stars:      431,
		Price:      29900,
	}

	ColemanTent = Product{
		ID:         "TNT67890",
		Name:       "Coleman Sundome Camping Tent",
		Categories: []category.Category{category.Outdoor},
		Stars:      417,
		Price:      8000,
	}

	TheragunElite = Product{
		ID:         "THG13579",
		Name:       "Theragun Elite Massage Gun",
		Categories: []category.Category{category.Health, category.Sports},
		Stars:      216,
		Price:      39900,
	}

	TomFordCologne = Product{
		ID:         "TFC24680",
		Name:       "Tom Ford Oud Wood Cologne",
		Categories: []category.Category{category.Beauty},
		Stars:      491,
		Price:      25000,
	}

	BoseSoundLink = Product{
		ID:         "BSE35791",
		Name:       "Bose SoundLink Revolve+ Speaker",
		Categories: []category.Category{category.Electronics},
		Stars:      478,
		Price:      20000,
	}

	NordicTrackBike = Product{
		ID:         "NTB46802",
		Name:       "NordicTrack Commercial S22i Studio Cycle",
		Categories: []category.Category{category.Sports},
		Stars:      343,
		Price:      199900,
	}

	LululemonYogaMat = Product{
		ID:         "LLM13579",
		Name:       "Lululemon Reversible Yoga Mat",
		Categories: []category.Category{category.Apparel, category.Sports},
		Stars:      451,
		Price:      7800,
	}

	SmegToaster = Product{
		ID:         "SMG24680",
		Name:       "Smeg 2-Slice Toaster",
		Categories: []category.Category{category.HomeGoods},
		Stars:      419,
		Price:      16900,
	}

	DeWaltDrill = Product{
		ID:         "DWD35791",
		Name:       "DeWalt Cordless Drill",
		Categories: []category.Category{category.HomeGoods},
		Stars:      421,
		Price:      12900,
	}

	PatagoniaJacket = Product{
		ID:         "PTG46802",
		Name:       "Patagonia Down Sweater Jacket",
		Categories: []category.Category{category.Apparel, category.Outdoor},
		Stars:      499,
		Price:      22900,
	}

	GoodyearTires = Product{
		ID:         "TYR13579",
		Name:       "Goodyear Wrangler All-Terrain Tires",
		Categories: []category.Category{category.Automotive},
		Stars:      482,
		Price:      150000,
	}

	CreedAventusCologne = Product{
		ID:         "COL24680",
		Name:       "Creed Aventus Cologne",
		Categories: []category.Category{category.Beauty},
		Stars:      498,
		Price:      34500,
	}

	FujifilmX100V = Product{
		ID:         "CAM35791",
		Name:       "Fujifilm X100V Camera",
		Categories: []category.Category{category.Electronics},
		Stars:      455,
		Price:      139900,
	}

	HydroFlaskBottle = Product{
		ID:         "BOT46802",
		Name:       "Hydro Flask 32 oz Water Bottle",
		Categories: []category.Category{category.Outdoor, category.HomeGoods},
		Stars:      435,
		Price:      4500,
	}

	SonyA7Camera = Product{
		ID:         "SNY13579",
		Name:       "Sony A7 III Full-Frame Mirrorless Camera",
		Categories: []category.Category{category.Electronics},
		Stars:      495,
		Price:      199900,
	}

	CuisinartAirFryer = Product{
		ID:         "FRY24680",
		Name:       "Cuisinart Convection AirFryer",
		Categories: []category.Category{category.HomeGoods},
		Stars:      276,
		Price:      20000,
	}

	MonopolyGame = Product{
		ID:         "TOY35791",
		Name:       "Monopoly Classic Board Game",
		Categories: []category.Category{category.Toys},
		Stars:      449,
		Price:      2000,
	}

	BrevilleCoffeeMaker = Product{
		ID:         "COF46802",
		Name:       "Breville Espresso Machine",
		Categories: []category.Category{category.HomeGoods},
		Stars:      491,
		Price:      69900,
	}

	FitbitVersa3 = Product{
		ID:         "FIT13580",
		Name:       "Fitbit Versa 3 Health & Fitness Smartwatch",
		Categories: []category.Category{category.Electronics, category.Sports},
		Stars:      415,
		Price:      22900,
	}

	ReebokCrossfit = Product{
		ID:         "RBF24680",
		Name:       "Reebok Nano X1 Crossfit Shoes",
		Categories: []category.Category{category.Apparel, category.Sports},
		Stars:      455,
		Price:      8500,
	}

	GarminFenix6 = Product{
		ID:         "GMT35791",
		Name:       "Garmin Fenix 6 Pro GPS Smartwatch",
		Categories: []category.Category{category.Electronics, category.Sports},
		Stars:      466,
		Price:      55000,
	}

	JBLFlip5Speaker = Product{
		ID:         "SPK46802",
		Name:       "JBL Flip 5 Waterproof Speaker",
		Categories: []category.Category{category.Electronics},
		Stars:      461,
		Price:      10000,
	}

	ChampionSweatshirt = Product{
		ID:         "APP13579",
		Name:       "Champion Powerblend Sweatshirt",
		Categories: []category.Category{category.Apparel},
		Stars:      432,
		Price:      3800,
	}

	LegoTechnicBugatti = Product{
		ID:         "LEG24680",
		Name:       "LEGO Technic Bugatti Chiron",
		Categories: []category.Category{category.Toys},
		Stars:      276,
		Price:      35000,
	}

	RolexSubmariner = Product{
		ID:         "WTC35791",
		Name:       "Rolex Submariner Watch",
		Categories: []category.Category{category.Luxury},
		Stars:      500,
		Price:      900000,
	}

	KindlePaperwhiteKids = Product{
		ID:         "EBK46802",
		Name:       "Amazon Kindle Paperwhite Kids Edition",
		Categories: []category.Category{category.Electronics, category.Toys},
		Stars:      492,
		Price:      11000,
	}

	SonosOneSL = Product{
		ID:         "SNS13579",
		Name:       "Sonos One SL Speaker",
		Categories: []category.Category{category.Electronics},
		Stars:      468,
		Price:      18000,
	}

	BanwoodBalanceBike = Product{
		ID:         "BIK24680",
		Name:       "Banwood Classic Balance Bike",
		Categories: []category.Category{category.Outdoor, category.Toys},
		Stars:      354,
		Price:      15000,
	}

	NorthFacePuffer = Product{
		ID:         "PUF35791",
		Name:       "The North Face Nuptse Puffer Jacket",
		Categories: []category.Category{category.Apparel, category.Outdoor},
		Stars:      485,
		Price:      29900,
	}

	SkullcandyHeadphones = Product{
		ID:         "HDN46802",
		Name:       "Skullcandy Crusher ANC Wireless Headphones",
		Categories: []category.Category{category.Electronics},
		Stars:      458,
		Price:      22000,
	}

	CampChefGrill = Product{
		ID:         "GRL13579",
		Name:       "Camp Chef Pellet Grill",
		Categories: []category.Category{category.HomeGoods, category.Outdoor},
		Stars:      304,
		Price:      60000,
	}

	HyperXCloudHeadset = Product{
		ID:         "HPC12345",
		Name:       "HyperX Cloud II Gaming Headset",
		Categories: []category.Category{category.Electronics},
		Stars:      325,
		Price:      9000,
	}

	ColgateToothpaste = Product{
		ID:         "CLG24680",
		Name:       "Colgate Total Whitening Toothpaste",
		Categories: []category.Category{category.Health, category.Grocery},
		Stars:      411,
		Price:      500,
	}

	BrooksRunningShoes = Product{
		ID:         "BKS35791",
		Name:       "Brooks Ghost 14 Running Shoes",
		Categories: []category.Category{category.Apparel, category.Sports},
		Stars:      486,
		Price:      11000,
	}

	KindleFireTablet = Product{
		ID:         "KFT46802",
		Name:       "Amazon Fire HD 10 Tablet",
		Categories: []category.Category{category.Electronics},
		Stars:      142,
		Price:      14900,
	}

	DysonAirWrap = Product{
		ID:         "DYR13579",
		Name:       "Dyson Airwrap Styler",
		Categories: []category.Category{category.Beauty, category.Luxury},
		Stars:      495,
		Price:      59900,
	}

	CanonInkCartridge = Product{
		ID:         "INC24680",
		Name:       "Canon Ink Cartridge - Black",
		Categories: []category.Category{category.OfficeSupplies},
		Stars:      410,
		Price:      2500,
	}

	GoogleNestThermostat = Product{
		ID:         "GST35791",
		Name:       "Google Nest Learning Thermostat",
		Categories: []category.Category{category.Electronics, category.HomeGoods},
		Stars:      381,
		Price:      24900,
	}

	SauconyTrailShoes = Product{
		ID:         "STL46802",
		Name:       "Saucony Peregrine 12 Trail Running Shoes",
		Categories: []category.Category{category.Apparel, category.Sports, category.Outdoor},
		Stars:      273,
		Price:      12500,
	}

	AppleMagicKeyboard = Product{
		ID:         "MKB13579",
		Name:       "Apple Magic Keyboard",
		Categories: []category.Category{category.Electronics, category.OfficeSupplies},
		Stars:      285,
		Price:      9900,
	}

	StanleyThermos = Product{
		ID:         "STM24680",
		Name:       "Stanley Classic Vacuum Bottle",
		Categories: []category.Category{category.Outdoor, category.HomeGoods},
		Stars:      465,
		Price:      5500,
	}

	OXOCoffeeGrinder = Product{
		ID:         "OXO35791",
		Name:       "OXO Brew Conical Burr Coffee Grinder",
		Categories: []category.Category{category.HomeGoods},
		Stars:      276,
		Price:      8900,
	}

	HarryPotterBoxSet = Product{
		ID:         "HPB46802",
		Name:       "Harry Potter Complete Box Set",
		Categories: []category.Category{category.Books},
		Stars:      395,
		Price:      7200,
	}

	AllCladPan = Product{
		ID:         "ALC13579",
		Name:       "All-Clad Stainless Steel Fry Pan",
		Categories: []category.Category{category.HomeGoods},
		Stars:      318,
		Price:      9900,
	}

	SchwinnBike = Product{
		ID:         "SCH24680",
		Name:       "Schwinn High Timber Youth Mountain Bike",
		Categories: []category.Category{category.Sports, category.Outdoor},
		Stars:      421,
		Price:      22000,
	}

	WaterpikFlosser = Product{
		ID:         "WPK35791",
		Name:       "Waterpik Cordless Water Flosser",
		Categories: []category.Category{category.Health},
		Stars:      443,
		Price:      4999,
	}

	PamperedChefStoneware = Product{
		ID:         "PCS46802",
		Name:       "Pampered Chef Stoneware Loaf Pan",
		Categories: []category.Category{category.HomeGoods},
		Stars:      494,
		Price:      3400,
	}

	PlaymobilCastle = Product{
		ID:         "PMC13579",
		Name:       "Playmobil Knights Castle",
		Categories: []category.Category{category.Toys},
		Stars:      261,
		Price:      5500,
	}

	GatoradePack = Product{
		ID:         "GAT24680",
		Name:       "Gatorade Thirst Quencher 12-Pack",
		Categories: []category.Category{category.Grocery, category.Sports},
		Stars:      275,
		Price:      1000,
	}

	HersheysChocolateBar = Product{
		ID:         "CHO35791",
		Name:       "Hershey's Milk Chocolate Bar",
		Categories: []category.Category{category.Grocery},
		Stars:      198,
		Price:      150,
	}
)
