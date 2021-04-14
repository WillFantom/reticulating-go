package reticulating

import (
	"math/rand"
	"time"

	fuzzy "github.com/lithammer/fuzzysearch/fuzzy"
)

type Game struct {
	Name      string   `json:"Name"`
	FandomURL string   `json:"FandomURL"`
	Packs     []*Game  `json:"Packs"`
	Messages  []string `json:"Messages"`
}

var simsRand *rand.Rand

var data []*Game = []*Game{
	{
		Name:      "The Sims 2",
		FandomURL: "https://sims.fandom.com/wiki/The_Sims_2",
		Messages: []string{
			"Chlorinating Car Pools",
			"Partitioning Social Network",
			"Prelaminating Drywall Inventory",
			"Blurring Reality Lines",
			"Reticulating 3-Dimensional Splines",
			"Preparing Captive Simulators",
			"Capacitating Genetic Modifiers",
			"Destabilizing Orbital Payloads",
			"Sequencing Cinematic Specifiers",
			"Branching Family Trees",
			"Manipulating Modal Memory",
		},
		Packs: []*Game{
			{
				Name:      "University",
				FandomURL: "https://sims.fandom.com/wiki/The_Sims_2:_University",
				Messages: []string{
					"Pressurizing Fruit Punch Barrel Hydraulics",
					"Testing Underworld Telecommunications",
					"Priming Mascot Mischief Coefficients",
					"Caffeinating Student Body",
					"Initializing Secret Societies",
					"Securing Online Grades Database",
					"Reticulating Graduated Splines",
					"Requisitioning Alumni Donations",
					"Pre-Inking Simoleon Plates",
					"Loading School Spirit Algorithm",
				},
			},
			{
				Name:      "Nightlife",
				FandomURL: "https://sims.fandom.com/wiki/The_Sims_2:_Nightlife",
				Messages: []string{
					"Shampooing Dirty Rugs",
					"Restocking Sim Inventories",
					"Compositing Vampiric Complexions",
					"Replacing Wheel Bearings",
					"Re-Re-Re-Re-Re-Reticulating Splines",
					"Loading 'Vroom' Sounds",
					"Turning On Turn-Ons",
					"Preparing a Tasty Grilled Cheese Sandwich",
					"Infuriating Furious Bits",
					"Flavorizing Side-Dishes",
				},
			},
			{
				Name:      "Open For Business",
				FandomURL: "https://sims.fandom.com/wiki/The_Sims_2:_Open_for_Business",
				Messages: []string{
					"Disgruntling Employees",
					"Managing Managers' Managers",
					"Configuring Lemony Squeezation",
					"Preparing Bacon for Homeward Transportation",
					"Reticulated Splines for Sale: §2000",
					"Mitigating Time-Stream Discontinuities",
					"Loading 'First Batch' Metrics",
					"Initializing Forth-Rallying Protocol",
					"Neutralizing Shuriken Oxidization",
					"Roof = Roof(1/3*pi*r^2*h)",
				},
			},
			{
				Name:      "Pets",
				FandomURL: "https://sims.fandom.com/wiki/The_Sims_2:_Pets",
				Messages: []string{
					"Rasterizing Rodent Residences",
					"Limiting Litterbox Loads",
					"Scheduling Copious Catnaps",
					"Calibrating Canine Customization",
					"Dumbing Down Doofuses",
					"Scolding Splines for Reticulating",
					"Distilling Doggie Dynamics",
					"Atomizing Atomic Particles",
					"Decrementing Feline Life-Count",
					"Dampening Stray Generators",
				},
			},
			{
				Name:      "Seasons",
				FandomURL: "https://sims.fandom.com/wiki/The_Sims_2:_Seasons",
				Messages: []string{
					"Gleefully Stacking Inventories",
					"De-chlorophyllizing Leaves",
					"Predicting Puddle Prevalence",
					"Calculating Snowball Trajectories",
					"Unexpectedly Reticulating Splines",
					"Assessing Loam Particle Sizes",
					"Timing Temperature Transference",
					"Individualizing Snowflakes",
					"Hydrating Harvestables",
					"Stocking Ponds",
				},
			},
			{
				Name:      "Bon Voyage",
				FandomURL: "https://sims.fandom.com/wiki/The_Sims_2:_Bon_Voyage",
				Messages: []string{
					"Readying Relaxation Receptors",
					"Predicting Pagoda Peaks",
					"Originating Ocean Currents",
					"Faceting Precious Gems",
					"Preparing Vacation Days",
					"Spawning Sights to See",
					"Reticulating Ninja Splines",
					"Analyzing Axe Trajectories",
					"Training Tour Guides",
					"Initializing Dastardly Schemes",
				},
			},
			{
				Name:      "Free Time",
				FandomURL: "https://sims.fandom.com/wiki/The_Sims_2:_FreeTime",
				Messages: []string{
					"Factoring Hobby Enthusiasm",
					"Calculating Lifetime Aspirations",
					"Predicating Predestined Paths",
					"Populating Yards with Bugs and Birds",
					"Writing Scrolling Startup String Text",
					"Reticulating Splines in the Zone",
					"Recruiting Snooty Food Judges",
					"Breaking Down Restorable Cars",
					"Threading Sewing Needles",
					"Lacing Football Cleats",
				},
			},
			{
				Name:      "Apartment Life",
				FandomURL: "https://sims.fandom.com/wiki/The_Sims_2:_Apartment_Life",
				Messages: []string{
					"Going Apartment Hunting",
					"Determining Rent Guidelines",
					"Preparing for Pops and Locks",
					"Generating Compatible Roommates",
					"Twisting Spiral Staircases",
					"Telling Splines to Reticulate More Quietly",
					"Making a Little Bit of Magic",
					"Rasterizing Reputation Algorithms",
					"Cluttering Closets",
					"Perfecting Playground Pieces",
				},
			},
			{
				Name:      "Family Fun Stuff",
				FandomURL: "https://sims.fandom.com/wiki/The_Sims_2:_Family_Fun_Stuff",
				Messages: []string{
					"Submerging Bedroom Furniture",
					"Desalinizing Snorkels",
					"Enhancing Crown Reflectivity",
					"Crenellating Crenellations",
					"Dragon-proofing Dressers",
					"Reticulating Underwater Splines",
					"Intensifying Hawaiian Prints",
					"Navigating Stormy Waters",
					"Pre-fluffing Pillows",
					"Factoring Fairy Frolicking Frequencies",
				},
			},
			{
				Name:      "Glamour Life Stuff",
				FandomURL: "https://sims.fandom.com/wiki/The_Sims_2:_Glamour_Life_Stuff",
				Messages: []string{
					"Modeling Marquetry",
					"Eschewing Everyday Aesthetics",
					"Cultivating Quality and Class",
					"Proscribing Plebeian Palates",
					"Falsifying Faux Finishes",
					"Invigorating Dull Habitations",
					"Abolishing Pedestrian Posturing",
					"Buffing Splines for Reticulation",
					"Appointing Appealing Appurtenances",
					"Simulating Sparkling Surfaces",
				},
			},
			{
				Name:      "Celebration! Stuff",
				FandomURL: "https://sims.fandom.com/wiki/The_Sims_2:_Celebration!_Stuff",
				Messages: []string{
					"Reverse-Engineering Party Scores",
					"Unfolding Foldy Chairs",
					"Rehearsing Dinner",
					"Crash-Proofing Parties",
					"Grooming Grooms",
					"Mingling",
					"De-inviting Don Lothario",
					"Borrowing Something Blue",
					"Happy 14th Birthday Reticulated Splines!",
					"Applying Lampshade Headwear",
				},
			},
			{
				Name:      "H&M Fasion Stuff",
				FandomURL: "https://sims.fandom.com/wiki/The_Sims_2:_H%26M_Fashion_Stuff",
				Messages: []string{
					"Stocking Clearance Racks",
					"Fiercely Reticulating Splines",
					"Fashioning Late Arrivals",
					"De-wrinkling Worry-Free Clothing",
					"Distressing Jeans",
					"Developing Delicious Designs",
					"Formulating Fitting Rooms",
					"Tailoring Trendy Threads",
					"Constructing Clothes Hangers",
					"Adjusting Acceptable Apparel",
				},
			},
			{
				Name:      "Teen Style Stuff",
				FandomURL: "https://sims.fandom.com/wiki/The_Sims_2:_Teen_Style_Stuff",
				Messages: []string{
					"Capturing Youthful Exuberance",
					"Analyzing Adolescent Angst",
					"Preparing Personal Spaces",
					"Making a Mess",
					"Like, Totally Reticulating Splines, Dude",
					"Generating Gothic Glamour",
					"Monitoring Moody Minors",
					"Sweetening Sixteens",
					"Teasing Teenage Hair-dos",
					"Building Boring Bedrooms? As If!",
				},
			},
			{
				Name:      "Kitchen & Bath Stuff",
				FandomURL: "https://sims.fandom.com/wiki/The_Sims_2:_Kitchen_%26_Bath_Interior_Design_Stuff",
				Messages: []string{
					"Taking Countertops for Granite",
					"Preparing Perfect Plumbing",
					"Baking Bread for Toasters",
					"Igniting Pilot Lights",
					"Putting Down Toilet Seats",
					"Remodeling Spline Reticulator",
					"Assembling Shower Stalls",
					"Examining Tiles from All Zooms and Angles",
					"Cooling Down Refrigerators",
					"Stocking Stylish Sinks",
				},
			},
			{
				Name:      "IKEA Home Stuff",
				FandomURL: "https://sims.fandom.com/wiki/The_Sims_2:_IKEA_Home_Stuff",
				Messages: []string{
					"Creating Handmade Lampshades",
					"Making Many Mini Wrenches",
					"Supplying Self-Serve Furniture Area",
					"Simmering Swedish Meatballs",
					"Building Bedroom Displays",
					"Stress-Testing POÄNG Chairs",
					"Some Spline Reticulating Required",
					"Upholstering Sofas and Loveseats",
					"Boxing BILLY Bookcases",
					"Spooling IKEA Awesomeness",
				},
			},
			{
				Name:      "Mansion & Garden Stuff",
				FandomURL: "https://sims.fandom.com/wiki/The_Sims_2:_Mansion_%26_Garden_Stuff",
				Messages: []string{
					"Growing Greener Gardens",
					"Making Manic Mansions",
					"Storing Solar Energy",
					"Over-Waxing Banisters",
					"Stopping To Smell The Flowers",
					"Extrapolating Empire Eigenvectors",
					"Ceiling Fan Rotation = dL/dT",
					"Increasing Water Plant Population",
					"Redistributing Resources",
					"Reticulating Splines One Last Time",
				},
			},
		},
	},
	{
		Name:      "The Sims 1",
		FandomURL: "https://sims.fandom.com/wiki/The_Sims",
		Messages: []string{
			"Extruding Mesh Terrain",
			"Balancing Domestic Coefficients",
			"Inverting Career Ladder",
			"Calculating Money Supply",
			"Normalizing Social Network",
			"Reticulating Splines",
			"Adjusting Emotional Weights",
			"Calibrating Personality Matrix",
			"Inserting Chaos Generator",
		},
		Packs: []*Game{
			{
				Name:      "Livin' Large",
				FandomURL: "https://sims.fandom.com/wiki/The_Sims:_Livin%27_Large",
				Messages: []string{
					"Concatenating Vertex Nodes",
					"Balancing Domestic Coefficients",
					"Inverting Career Ladder",
					"Mapping Influence Attributes",
					"Assigning Mimic Propagation",
					"Reticulating Splines",
					"Iterating Chaos Array",
					"Importing Personality Anchors",
					"Inserting Extension Algorithms",
				},
			},
			{
				Name:      "House Party",
				FandomURL: "https://sims.fandom.com/wiki/The_Sims:_House_Party",
				Messages: []string{
					"Concatenating Vertex Nodes",
					"Balancing Domestic Coefficients",
					"Re-Inverting Career Ladder",
					"Mapping Influence Attributes",
					"Aggregating Need Agents",
					"Reticulating Splines",
					"Interpreting Family Values",
					"Cabalizing NPC Controls",
					"Maximizing Social Network",
				},
			},
			{
				Name:      "Hot Date",
				FandomURL: "https://sims.fandom.com/wiki/The_Sims:_Hot_Date",
				Messages: []string{
					"Renewing Urban Combinatorics",
					"Redefining Family Values",
					"Calibrating Personality Matrix",
					"Generating Population Model",
					"Homogenizing Interest Anatomy",
					"Reticulating Splines",
					"Establishing Gift Registry",
					"Randomizing Inhabitant Characteristics",
					"Readjusting Emotional Weights",
				},
			},
			{
				Name:      "Vacation",
				FandomURL: "https://sims.fandom.com/wiki/The_Sims:_Vacation",
				Messages: []string{
					"Activating Hotel Staff",
					"Importing Entertainment Talent",
					"Updating Vacancy Request Hotline",
					"Downloading Weather Data",
					"Hyperactivating Children",
					"Still Reticulating Splines",
					"Updating Hotel Registry",
					"Calculating Exchange Rate",
					"Activating Deviance Threshold",
				},
			},
			{
				Name:      "Unleashed",
				FandomURL: "https://sims.fandom.com/wiki/The_Sims:_Unleashed",
				Messages: []string{
					"Adapting Behavioral Model",
					"Reconfiguring Genetic Algorithms",
					"Hybridizing Plant Material",
					"Reticulating Splines Again",
					"Unfolding Helix Packet",
					"Synthesizing Natural Selection",
					"Enabling Lot Commercialization",
					"Recomputing Mammal Matrix",
					"Augmenting Occupational Conduits",
					"Initializing Operant Construct",
				},
			},
			{
				Name:      "Superstar",
				FandomURL: "https://sims.fandom.com/wiki/The_Sims:_Superstar",
				Messages: []string{
					"Generating Schmoozing Algorithm",
					"Populating Empyreal Entities",
					"Configuring Studio Operations",
					"Reticulating Golden Splines",
					"Composing Melodic Euphony",
					"Spreading Rumors",
					"Polarizing Image Conduits",
					"Calibrating Fame Indicant",
					"Strengthening Award Foundations",
					"Abstracting Loading Procedures",
				},
			},
			{
				Name:      "Makin Magic",
				FandomURL: "https://sims.fandom.com/wiki/The_Sims:_Makin%27_Magic",
				Messages: []string{
					"Locating Misplaced Calculations",
					"Eliminating Would-be Chicanery",
					"Tabulating Spell Effectors",
					"Reticulating Unreticulated Splines",
					"Recycling Hex Decimals",
					"Binding Trace Enchantments",
					"Fabricating Imaginary Infrastructure",
					"Optimizing Baking Temperature",
					"Ensuring Transplanar Synergy",
					"Simulating Program Execution",
				},
			},
			{
				Name:      "Complete Collection",
				FandomURL: "https://sims.fandom.com/wiki/Compilations_of_The_Sims#The_Sims:_Complete_Collection",
				Messages: []string{
					"Reticulating Splines",
					"Interpreting Family Values",
					"Fabricating Imaginary Infrastructure",
					"Recomputing Mammal Matrix",
					"Activating Deviance Threshold",
					"Composing Melodic Euphony",
					"Homogenizing Interest Anatomy",
					"Normalizing Social Network",
					"Compiling Reticulated Splines",
					"Simulating Program Execution",
				},
			},
		},
	},
}

func init() {
	simsRand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func flattenGameMessages(gamesSlice []*Game) []string {
	var messages []string
	for _, game := range gamesSlice {
		if game.Messages != nil {
			messages = append(messages, game.Messages...)
		}
		if game.Packs != nil {
			messages = append(messages, flattenGameMessages(game.Packs)...)
		}
	}
	return messages
}

func gameNameSearch(name string, set []*Game) *Game {
	for _, game := range set {
		if fuzzy.MatchNormalizedFold(name, game.Name) {
			return game
		}
	}
	return nil
}

func GameByName(name string) *Game {
	return gameNameSearch(name, data)
}

func (g *Game) PackByName(name string) *Game {
	return gameNameSearch(name, g.Packs)
}

func GetRandomGame(gameSlice []*Game) *Game {
	return gameSlice[simsRand.Intn(len(gameSlice))]
}

func (g *Game) GetRandomPack() *Game {
	return g.Packs[simsRand.Intn(len(g.Packs))]
}

func (g *Game) GetRandomMessage() string {
	messages := flattenGameMessages([]*Game{g})
	return messages[simsRand.Intn(len(messages))]
}
