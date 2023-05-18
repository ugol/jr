// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     NetDevice.avsc
 *     User.avsc
 *     campaign_finance.avsc
 *     clickstream.avsc
 *     clickstream_codes.avsc
 *     clickstream_users.avsc
 *     credit_cards.avsc
 *     device_information.avsc
 *     fleet_mgmt_description.avsc
 *     fleet_mgmt_location.avsc
 *     fleet_mgmt_sensors.avsc
 *     gaming_games.avsc
 *     gaming_player_activity.avsc
 *     gaming_players.avsc
 *     insurance_customer_activity.avsc
 *     insurance_customers.avsc
 *     insurance_offers.avsc
 *     inventory.avsc
 *     map_dumb_schema.avsc
 *     orders.avsc
 *     pageviews.avsc
 *     payroll_bonus.avsc
 *     payroll_employee.avsc
 *     payroll_employee_location.avsc
 *     pizza_orders.avsc
 *     pizza_orders_cancelled.avsc
 *     pizza_orders_completed.avsc
 *     product.avsc
 *     purchase.avsc
 *     ratings.avsc
 *     shoe_clickstream.avsc
 *     shoe_customers.avsc
 *     shoe_orders.avsc
 *     shoes_product.avsc
 *     siem_logs.avsc
 *     stockTrades.avsc
 *     stores.avsc
 *     syslog_logs.avsc
 *     transactions.avsc
 *     users.avsc
 */
package types

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Insurancecustomers struct {
	Customer_id int32 `json:"customer_id"`

	First_name string `json:"first_name"`

	Last_name string `json:"last_name"`

	Email string `json:"email"`

	Gender string `json:"gender"`

	Income int32 `json:"income"`

	Fico int32 `json:"fico"`

	Years_active int32 `json:"years_active"`
}

const InsurancecustomersAvroCRC64Fingerprint = "jۇYT\xcaĕ"

func NewInsurancecustomers() Insurancecustomers {
	r := Insurancecustomers{}
	return r
}

func DeserializeInsurancecustomers(r io.Reader) (Insurancecustomers, error) {
	t := NewInsurancecustomers()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeInsurancecustomersFromSchema(r io.Reader, schema string) (Insurancecustomers, error) {
	t := NewInsurancecustomers()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeInsurancecustomers(r Insurancecustomers, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Customer_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.First_name, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Last_name, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Email, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Gender, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Income, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Fico, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Years_active, w)
	if err != nil {
		return err
	}
	return err
}

func (r Insurancecustomers) Serialize(w io.Writer) error {
	return writeInsurancecustomers(r, w)
}

func (r Insurancecustomers) Schema() string {
	return "{\"arg.properties\":{\"options\":[{\"customer_id\":415,\"email\":\"hagdahlbi@last.fm\",\"fico\":423,\"first_name\":\"Herold\",\"gender\":\"Male\",\"income\":444152,\"last_name\":\"Agdahl\",\"years_active\":1},{\"customer_id\":871,\"email\":\"npenningtono6@hao123.com\",\"fico\":767,\"first_name\":\"Nikkie\",\"gender\":\"Female\",\"income\":452642,\"last_name\":\"Pennington\",\"years_active\":36},{\"customer_id\":597,\"email\":\"jcolefordgk@unesco.org\",\"fico\":730,\"first_name\":\"Josi\",\"gender\":\"Female\",\"income\":430798,\"last_name\":\"Coleford\",\"years_active\":1},{\"customer_id\":210,\"email\":\"abrockway5t@wiley.com\",\"fico\":661,\"first_name\":\"Alis\",\"gender\":\"Female\",\"income\":210965,\"last_name\":\"Brockway\",\"years_active\":30},{\"customer_id\":833,\"email\":\"etzarn4@umich.edu\",\"fico\":771,\"first_name\":\"Eugine\",\"gender\":\"Female\",\"income\":390882,\"last_name\":\"Tzar\",\"years_active\":43},{\"customer_id\":371,\"email\":\"lskidmoreaa@moonfruit.com\",\"fico\":482,\"first_name\":\"Larina\",\"gender\":\"Female\",\"income\":261165,\"last_name\":\"Skidmore\",\"years_active\":41},{\"customer_id\":387,\"email\":\"rbecclesaq@cpanel.net\",\"fico\":821,\"first_name\":\"Raffarty\",\"gender\":\"Male\",\"income\":348862,\"last_name\":\"Beccles\",\"years_active\":43},{\"customer_id\":313,\"email\":\"wbreadon8o@altervista.org\",\"fico\":880,\"first_name\":\"Winn\",\"gender\":\"Male\",\"income\":489764,\"last_name\":\"Breadon\",\"years_active\":10},{\"customer_id\":882,\"email\":\"higooh@nhs.uk\",\"fico\":893,\"first_name\":\"Harold\",\"gender\":\"Male\",\"income\":55800,\"last_name\":\"Igo\",\"years_active\":31},{\"customer_id\":854,\"email\":\"bhenrichsennp@hp.com\",\"fico\":839,\"first_name\":\"Barbara\",\"gender\":\"Female\",\"income\":111818,\"last_name\":\"Henrichsen\",\"years_active\":2},{\"customer_id\":895,\"email\":\"nminciniou@blinklist.com\",\"fico\":811,\"first_name\":\"Nani\",\"gender\":\"Female\",\"income\":212199,\"last_name\":\"Mincini\",\"years_active\":36},{\"customer_id\":198,\"email\":\"rfowgies5h@meetup.com\",\"fico\":469,\"first_name\":\"Rudd\",\"gender\":\"Male\",\"income\":337648,\"last_name\":\"Fowgies\",\"years_active\":19},{\"customer_id\":468,\"email\":\"ckarolyicz@paginegialle.it\",\"fico\":686,\"first_name\":\"Cecily\",\"gender\":\"Female\",\"income\":287668,\"last_name\":\"Karolyi\",\"years_active\":39},{\"customer_id\":860,\"email\":\"jkendernv@ucsd.edu\",\"fico\":753,\"first_name\":\"Jessalin\",\"gender\":\"Female\",\"income\":91516,\"last_name\":\"Kender\",\"years_active\":44},{\"customer_id\":383,\"email\":\"hramageam@alibaba.com\",\"fico\":414,\"first_name\":\"Hollis\",\"gender\":\"Male\",\"income\":367829,\"last_name\":\"Ramage\",\"years_active\":43},{\"customer_id\":226,\"email\":\"ekerrod69@sciencedaily.com\",\"fico\":793,\"first_name\":\"Eleonore\",\"gender\":\"Female\",\"income\":260176,\"last_name\":\"Kerrod\",\"years_active\":41},{\"customer_id\":630,\"email\":\"samayahh@phpbb.com\",\"fico\":536,\"first_name\":\"Shaylynn\",\"gender\":\"Female\",\"income\":454307,\"last_name\":\"Amaya\",\"years_active\":25},{\"customer_id\":546,\"email\":\"hidiensf5@shop-pro.jp\",\"fico\":645,\"first_name\":\"Had\",\"gender\":\"Male\",\"income\":223522,\"last_name\":\"Idiens\",\"years_active\":44},{\"customer_id\":949,\"email\":\"fdumphriesqc@clickbank.net\",\"fico\":756,\"first_name\":\"Farrand\",\"gender\":\"Female\",\"income\":410888,\"last_name\":\"Dumphries\",\"years_active\":41},{\"customer_id\":276,\"email\":\"ccloughton7n@phoca.cz\",\"fico\":416,\"first_name\":\"Carena\",\"gender\":\"Female\",\"income\":495407,\"last_name\":\"Cloughton\",\"years_active\":20},{\"customer_id\":595,\"email\":\"gsteingi@imgur.com\",\"fico\":597,\"first_name\":\"Graham\",\"gender\":\"Male\",\"income\":346149,\"last_name\":\"Stein\",\"years_active\":24},{\"customer_id\":406,\"email\":\"cfarrinb9@yellowpages.com\",\"fico\":826,\"first_name\":\"Consolata\",\"gender\":\"Female\",\"income\":126493,\"last_name\":\"Farrin\",\"years_active\":11},{\"customer_id\":565,\"email\":\"tmollenefo@stanford.edu\",\"fico\":730,\"first_name\":\"Trev\",\"gender\":\"Male\",\"income\":76961,\"last_name\":\"Mollene\",\"years_active\":34},{\"customer_id\":129,\"email\":\"nsidary3k@netscape.com\",\"fico\":821,\"first_name\":\"Nicholle\",\"gender\":\"Female\",\"income\":82035,\"last_name\":\"Sidary\",\"years_active\":23},{\"customer_id\":973,\"email\":\"resposir0@histats.com\",\"fico\":409,\"first_name\":\"Robby\",\"gender\":\"Male\",\"income\":315754,\"last_name\":\"Esposi\",\"years_active\":5},{\"customer_id\":10,\"email\":\"ahardstaff9@ask.com\",\"fico\":842,\"first_name\":\"Antonietta\",\"gender\":\"Female\",\"income\":417732,\"last_name\":\"Hardstaff\",\"years_active\":9},{\"customer_id\":579,\"email\":\"esteereg2@domainmarket.com\",\"fico\":722,\"first_name\":\"Emyle\",\"gender\":\"Female\",\"income\":380771,\"last_name\":\"Steere\",\"years_active\":38},{\"customer_id\":167,\"email\":\"srenfield4m@mayoclinic.com\",\"fico\":453,\"first_name\":\"Storm\",\"gender\":\"Female\",\"income\":386862,\"last_name\":\"Renfield\",\"years_active\":37},{\"customer_id\":273,\"email\":\"phyder7k@cnn.com\",\"fico\":808,\"first_name\":\"Phillipp\",\"gender\":\"Male\",\"income\":223150,\"last_name\":\"Hyder\",\"years_active\":15},{\"customer_id\":548,\"email\":\"nklassf7@usgs.gov\",\"fico\":672,\"first_name\":\"Noelle\",\"gender\":\"Female\",\"income\":50831,\"last_name\":\"Klass\",\"years_active\":39},{\"customer_id\":175,\"email\":\"gjannasch4u@umich.edu\",\"fico\":791,\"first_name\":\"Garek\",\"gender\":\"Male\",\"income\":58157,\"last_name\":\"Jannasch\",\"years_active\":44},{\"customer_id\":957,\"email\":\"rreapeqk@g.co\",\"fico\":835,\"first_name\":\"Rena\",\"gender\":\"Female\",\"income\":372940,\"last_name\":\"Reape\",\"years_active\":35},{\"customer_id\":45,\"email\":\"kfydoe18@npr.org\",\"fico\":465,\"first_name\":\"Kelcey\",\"gender\":\"Female\",\"income\":245023,\"last_name\":\"Fydoe\",\"years_active\":44},{\"customer_id\":582,\"email\":\"wmerigeaug5@friendfeed.com\",\"fico\":660,\"first_name\":\"Wolfy\",\"gender\":\"Male\",\"income\":366902,\"last_name\":\"Merigeau\",\"years_active\":18},{\"customer_id\":66,\"email\":\"chughes1t@google.it\",\"fico\":709,\"first_name\":\"Charita\",\"gender\":\"Female\",\"income\":455085,\"last_name\":\"Hughes\",\"years_active\":41},{\"customer_id\":472,\"email\":\"dmarcoolynd3@toplist.cz\",\"fico\":485,\"first_name\":\"Davie\",\"gender\":\"Male\",\"income\":369945,\"last_name\":\"Marcoolyn\",\"years_active\":30},{\"customer_id\":578,\"email\":\"ssallansg1@gnu.org\",\"fico\":435,\"first_name\":\"Suzann\",\"gender\":\"Female\",\"income\":137396,\"last_name\":\"Sallans\",\"years_active\":35},{\"customer_id\":825,\"email\":\"bshopcottmw@edublogs.org\",\"fico\":876,\"first_name\":\"Brooks\",\"gender\":\"Male\",\"income\":362624,\"last_name\":\"Shopcott\",\"years_active\":33},{\"customer_id\":572,\"email\":\"gblethinfv@reddit.com\",\"fico\":798,\"first_name\":\"Grant\",\"gender\":\"Male\",\"income\":154295,\"last_name\":\"Blethin\",\"years_active\":43},{\"customer_id\":322,\"email\":\"imallen8x@networksolutions.com\",\"fico\":714,\"first_name\":\"Idaline\",\"gender\":\"Female\",\"income\":74954,\"last_name\":\"Mallen\",\"years_active\":35},{\"customer_id\":864,\"email\":\"gcaddynz@pcworld.com\",\"fico\":633,\"first_name\":\"Gill\",\"gender\":\"Male\",\"income\":298967,\"last_name\":\"Caddy\",\"years_active\":42},{\"customer_id\":272,\"email\":\"mepton7j@oaic.gov.au\",\"fico\":551,\"first_name\":\"Mariele\",\"gender\":\"Female\",\"income\":482891,\"last_name\":\"Epton\",\"years_active\":45},{\"customer_id\":910,\"email\":\"gfarmanp9@seesaa.net\",\"fico\":599,\"first_name\":\"Gerry\",\"gender\":\"Male\",\"income\":350830,\"last_name\":\"Farman\",\"years_active\":19},{\"customer_id\":507,\"email\":\"cwillavizee2@networksolutions.com\",\"fico\":518,\"first_name\":\"Cirillo\",\"gender\":\"Male\",\"income\":87378,\"last_name\":\"Willavize\",\"years_active\":30},{\"customer_id\":942,\"email\":\"wakredq5@twitter.com\",\"fico\":808,\"first_name\":\"Wendye\",\"gender\":\"Female\",\"income\":311450,\"last_name\":\"Akred\",\"years_active\":17},{\"customer_id\":368,\"email\":\"dedensora7@wired.com\",\"fico\":440,\"first_name\":\"Delmor\",\"gender\":\"Male\",\"income\":225683,\"last_name\":\"Edensor\",\"years_active\":11},{\"customer_id\":643,\"email\":\"rbiggenhu@yellowpages.com\",\"fico\":722,\"first_name\":\"Rori\",\"gender\":\"Female\",\"income\":119504,\"last_name\":\"Biggen\",\"years_active\":32},{\"customer_id\":421,\"email\":\"rdronbo@hibu.com\",\"fico\":596,\"first_name\":\"Reiko\",\"gender\":\"Female\",\"income\":339053,\"last_name\":\"Dron\",\"years_active\":39},{\"customer_id\":92,\"email\":\"wtollett2j@amazon.de\",\"fico\":803,\"first_name\":\"Willey\",\"gender\":\"Male\",\"income\":278712,\"last_name\":\"Tollett\",\"years_active\":19},{\"customer_id\":433,\"email\":\"tlittleyc0@t-online.de\",\"fico\":759,\"first_name\":\"Tadio\",\"gender\":\"Male\",\"income\":294165,\"last_name\":\"Littley\",\"years_active\":6},{\"customer_id\":75,\"email\":\"csherewood22@loc.gov\",\"fico\":815,\"first_name\":\"Claiborn\",\"gender\":\"Male\",\"income\":128802,\"last_name\":\"Sherewood\",\"years_active\":32},{\"customer_id\":305,\"email\":\"tkrikorian8g@wordpress.com\",\"fico\":719,\"first_name\":\"Tierney\",\"gender\":\"Female\",\"income\":175243,\"last_name\":\"Krikorian\",\"years_active\":3},{\"customer_id\":173,\"email\":\"gmallya4s@mozilla.com\",\"fico\":552,\"first_name\":\"Gwendolin\",\"gender\":\"Female\",\"income\":435535,\"last_name\":\"Mallya\",\"years_active\":16},{\"customer_id\":863,\"email\":\"hbrunkerny@go.com\",\"fico\":799,\"first_name\":\"Hali\",\"gender\":\"Female\",\"income\":450239,\"last_name\":\"Brunker\",\"years_active\":37},{\"customer_id\":315,\"email\":\"elitherland8q@icq.com\",\"fico\":786,\"first_name\":\"Eba\",\"gender\":\"Female\",\"income\":369614,\"last_name\":\"Litherland\",\"years_active\":9},{\"customer_id\":344,\"email\":\"drheam9j@ustream.tv\",\"fico\":884,\"first_name\":\"Darby\",\"gender\":\"Female\",\"income\":105164,\"last_name\":\"Rheam\",\"years_active\":45},{\"customer_id\":526,\"email\":\"hocleryel@plala.or.jp\",\"fico\":759,\"first_name\":\"Hadria\",\"gender\":\"Female\",\"income\":404047,\"last_name\":\"O'Clery\",\"years_active\":12},{\"customer_id\":264,\"email\":\"gandri7b@stumbleupon.com\",\"fico\":615,\"first_name\":\"Gwenni\",\"gender\":\"Female\",\"income\":164274,\"last_name\":\"Andri\",\"years_active\":28},{\"customer_id\":225,\"email\":\"carmell68@tinyurl.com\",\"fico\":621,\"first_name\":\"Carmelita\",\"gender\":\"Female\",\"income\":355234,\"last_name\":\"Armell\",\"years_active\":28},{\"customer_id\":40,\"email\":\"fludl13@chron.com\",\"fico\":451,\"first_name\":\"Flemming\",\"gender\":\"Male\",\"income\":50188,\"last_name\":\"Ludl\",\"years_active\":4},{\"customer_id\":866,\"email\":\"cikringillo1@newsvine.com\",\"fico\":500,\"first_name\":\"Concettina\",\"gender\":\"Female\",\"income\":380505,\"last_name\":\"Ikringill\",\"years_active\":31},{\"customer_id\":473,\"email\":\"lmapplesd4@icq.com\",\"fico\":850,\"first_name\":\"Legra\",\"gender\":\"Female\",\"income\":109376,\"last_name\":\"Mapples\",\"years_active\":13},{\"customer_id\":662,\"email\":\"vthoresbyid@wordpress.org\",\"fico\":418,\"first_name\":\"Viviana\",\"gender\":\"Female\",\"income\":253060,\"last_name\":\"Thoresby\",\"years_active\":41},{\"customer_id\":611,\"email\":\"dgerrietzgy@pcworld.com\",\"fico\":769,\"first_name\":\"Dita\",\"gender\":\"Female\",\"income\":323327,\"last_name\":\"Gerrietz\",\"years_active\":24},{\"customer_id\":401,\"email\":\"slatheeb4@foxnews.com\",\"fico\":848,\"first_name\":\"Sapphire\",\"gender\":\"Female\",\"income\":477165,\"last_name\":\"Lathee\",\"years_active\":9},{\"customer_id\":993,\"email\":\"bcornessrk@craigslist.org\",\"fico\":474,\"first_name\":\"Bengt\",\"gender\":\"Male\",\"income\":394992,\"last_name\":\"Corness\",\"years_active\":14},{\"customer_id\":158,\"email\":\"mteffrey4d@amazon.com\",\"fico\":826,\"first_name\":\"Morton\",\"gender\":\"Male\",\"income\":185934,\"last_name\":\"Teffrey\",\"years_active\":19},{\"customer_id\":997,\"email\":\"lcrocroftro@marriott.com\",\"fico\":833,\"first_name\":\"Levi\",\"gender\":\"Male\",\"income\":364541,\"last_name\":\"Crocroft\",\"years_active\":15},{\"customer_id\":43,\"email\":\"kflancinbaum16@w3.org\",\"fico\":713,\"first_name\":\"Kele\",\"gender\":\"Male\",\"income\":468948,\"last_name\":\"Flancinbaum\",\"years_active\":42},{\"customer_id\":36,\"email\":\"skitleez@goo.gl\",\"fico\":569,\"first_name\":\"Shaylynn\",\"gender\":\"Female\",\"income\":128959,\"last_name\":\"Kitlee\",\"years_active\":21},{\"customer_id\":843,\"email\":\"mspringtorpene@netlog.com\",\"fico\":545,\"first_name\":\"Maia\",\"gender\":\"Female\",\"income\":204079,\"last_name\":\"Springtorpe\",\"years_active\":1},{\"customer_id\":814,\"email\":\"ncosterml@newyorker.com\",\"fico\":404,\"first_name\":\"Nolie\",\"gender\":\"Female\",\"income\":260330,\"last_name\":\"Coster\",\"years_active\":31},{\"customer_id\":965,\"email\":\"mmorgonqs@blogspot.com\",\"fico\":433,\"first_name\":\"Marty\",\"gender\":\"Male\",\"income\":424541,\"last_name\":\"Morgon\",\"years_active\":8},{\"customer_id\":967,\"email\":\"ctinniswoodqu@macromedia.com\",\"fico\":606,\"first_name\":\"Ceciley\",\"gender\":\"Female\",\"income\":368036,\"last_name\":\"Tinniswood\",\"years_active\":39},{\"customer_id\":470,\"email\":\"ddeportd1@xrea.com\",\"fico\":699,\"first_name\":\"Davie\",\"gender\":\"Male\",\"income\":57483,\"last_name\":\"Deport\",\"years_active\":22},{\"customer_id\":55,\"email\":\"jborgars1i@wp.com\",\"fico\":837,\"first_name\":\"Jae\",\"gender\":\"Male\",\"income\":388179,\"last_name\":\"Borgars\",\"years_active\":15},{\"customer_id\":615,\"email\":\"mbrewoodh2@acquirethisname.com\",\"fico\":890,\"first_name\":\"Melvin\",\"gender\":\"Male\",\"income\":456571,\"last_name\":\"Brewood\",\"years_active\":26},{\"customer_id\":30,\"email\":\"sbent@spiegel.de\",\"fico\":603,\"first_name\":\"Staford\",\"gender\":\"Male\",\"income\":362731,\"last_name\":\"Ben\",\"years_active\":4},{\"customer_id\":894,\"email\":\"gfallowfieldot@51.la\",\"fico\":422,\"first_name\":\"Garreth\",\"gender\":\"Male\",\"income\":310827,\"last_name\":\"Fallowfield\",\"years_active\":40},{\"customer_id\":333,\"email\":\"kdarnbrough98@parallels.com\",\"fico\":580,\"first_name\":\"Krystalle\",\"gender\":\"Female\",\"income\":276156,\"last_name\":\"Darnbrough\",\"years_active\":36},{\"customer_id\":763,\"email\":\"mmealel6@omniture.com\",\"fico\":723,\"first_name\":\"Major\",\"gender\":\"Male\",\"income\":309416,\"last_name\":\"Meale\",\"years_active\":38},{\"customer_id\":929,\"email\":\"blaingps@scientificamerican.com\",\"fico\":769,\"first_name\":\"Bennie\",\"gender\":\"Female\",\"income\":450578,\"last_name\":\"Laing\",\"years_active\":8},{\"customer_id\":994,\"email\":\"rlinebargerrl@si.edu\",\"fico\":890,\"first_name\":\"Raeann\",\"gender\":\"Female\",\"income\":484790,\"last_name\":\"Linebarger\",\"years_active\":33},{\"customer_id\":771,\"email\":\"esouthgatele@g.co\",\"fico\":761,\"first_name\":\"Emory\",\"gender\":\"Male\",\"income\":161393,\"last_name\":\"Southgate\",\"years_active\":14},{\"customer_id\":601,\"email\":\"mchethamgo@google.it\",\"fico\":598,\"first_name\":\"Madella\",\"gender\":\"Female\",\"income\":191473,\"last_name\":\"Chetham\",\"years_active\":29},{\"customer_id\":634,\"email\":\"ekrollhl@sun.com\",\"fico\":672,\"first_name\":\"Elita\",\"gender\":\"Female\",\"income\":318111,\"last_name\":\"Kroll\",\"years_active\":10},{\"customer_id\":955,\"email\":\"hewbankeqi@globo.com\",\"fico\":563,\"first_name\":\"Hazel\",\"gender\":\"Female\",\"income\":392441,\"last_name\":\"Ewbanke\",\"years_active\":11},{\"customer_id\":310,\"email\":\"lcrafts8l@about.me\",\"fico\":445,\"first_name\":\"Lodovico\",\"gender\":\"Male\",\"income\":125866,\"last_name\":\"Crafts\",\"years_active\":21},{\"customer_id\":463,\"email\":\"cpiggfordcu@google.com.au\",\"fico\":835,\"first_name\":\"Corilla\",\"gender\":\"Female\",\"income\":194180,\"last_name\":\"Piggford\",\"years_active\":25},{\"customer_id\":887,\"email\":\"bbutteryom@soundcloud.com\",\"fico\":536,\"first_name\":\"Brennan\",\"gender\":\"Male\",\"income\":67549,\"last_name\":\"Buttery\",\"years_active\":40},{\"customer_id\":770,\"email\":\"nhallgathld@ehow.com\",\"fico\":659,\"first_name\":\"Nicol\",\"gender\":\"Female\",\"income\":77799,\"last_name\":\"Hallgath\",\"years_active\":24},{\"customer_id\":386,\"email\":\"rdemschkeap@webnode.com\",\"fico\":641,\"first_name\":\"Rhianna\",\"gender\":\"Female\",\"income\":434414,\"last_name\":\"Demschke\",\"years_active\":11},{\"customer_id\":188,\"email\":\"bwhitley57@cbsnews.com\",\"fico\":683,\"first_name\":\"Bellina\",\"gender\":\"Female\",\"income\":427043,\"last_name\":\"Whitley\",\"years_active\":3},{\"customer_id\":62,\"email\":\"jchadd1p@adobe.com\",\"fico\":684,\"first_name\":\"Joleen\",\"gender\":\"Female\",\"income\":161541,\"last_name\":\"Chadd\",\"years_active\":17},{\"customer_id\":41,\"email\":\"wbleddon14@yellowpages.com\",\"fico\":842,\"first_name\":\"Willie\",\"gender\":\"Female\",\"income\":401028,\"last_name\":\"Bleddon\",\"years_active\":5},{\"customer_id\":465,\"email\":\"gthomlinsoncw@about.com\",\"fico\":650,\"first_name\":\"Griffith\",\"gender\":\"Male\",\"income\":240613,\"last_name\":\"Thomlinson\",\"years_active\":34},{\"customer_id\":932,\"email\":\"kgrabbampv@ft.com\",\"fico\":793,\"first_name\":\"Kev\",\"gender\":\"Male\",\"income\":442943,\"last_name\":\"Grabbam\",\"years_active\":43},{\"customer_id\":832,\"email\":\"lskucen3@about.me\",\"fico\":450,\"first_name\":\"Libbie\",\"gender\":\"Female\",\"income\":414333,\"last_name\":\"Skuce\",\"years_active\":37},{\"customer_id\":819,\"email\":\"dstallworthmq@qq.com\",\"fico\":705,\"first_name\":\"Derward\",\"gender\":\"Male\",\"income\":150254,\"last_name\":\"Stallworth\",\"years_active\":41},{\"customer_id\":103,\"email\":\"rhellmer2u@cbsnews.com\",\"fico\":513,\"first_name\":\"Russ\",\"gender\":\"Male\",\"income\":437338,\"last_name\":\"Hellmer\",\"years_active\":28},{\"customer_id\":377,\"email\":\"nmccartag@domainmarket.com\",\"fico\":680,\"first_name\":\"Nicholle\",\"gender\":\"Female\",\"income\":106811,\"last_name\":\"McCart\",\"years_active\":29},{\"customer_id\":216,\"email\":\"gdawson5z@ted.com\",\"fico\":885,\"first_name\":\"Gwen\",\"gender\":\"Female\",\"income\":235140,\"last_name\":\"Dawson\",\"years_active\":1},{\"customer_id\":958,\"email\":\"amacdermottql@github.io\",\"fico\":715,\"first_name\":\"Anthea\",\"gender\":\"Female\",\"income\":248104,\"last_name\":\"MacDermott\",\"years_active\":3},{\"customer_id\":39,\"email\":\"gmendes12@spotify.com\",\"fico\":501,\"first_name\":\"Gates\",\"gender\":\"Female\",\"income\":94187,\"last_name\":\"Mendes\",\"years_active\":32},{\"customer_id\":699,\"email\":\"placerje@discovery.com\",\"fico\":630,\"first_name\":\"Parsifal\",\"gender\":\"Male\",\"income\":168013,\"last_name\":\"Lacer\",\"years_active\":18},{\"customer_id\":557,\"email\":\"abrumhamfg@huffingtonpost.com\",\"fico\":450,\"first_name\":\"Alan\",\"gender\":\"Male\",\"income\":165570,\"last_name\":\"Brumham\",\"years_active\":25},{\"customer_id\":195,\"email\":\"chobell5e@mapquest.com\",\"fico\":656,\"first_name\":\"Celle\",\"gender\":\"Female\",\"income\":158933,\"last_name\":\"Hobell\",\"years_active\":3},{\"customer_id\":966,\"email\":\"acarbettqt@list-manage.com\",\"fico\":694,\"first_name\":\"Anabella\",\"gender\":\"Female\",\"income\":179331,\"last_name\":\"Carbett\",\"years_active\":9},{\"customer_id\":382,\"email\":\"tdeaconsonal@wordpress.org\",\"fico\":446,\"first_name\":\"Tracee\",\"gender\":\"Female\",\"income\":358295,\"last_name\":\"Deaconson\",\"years_active\":35},{\"customer_id\":796,\"email\":\"cvanesm3@clickbank.net\",\"fico\":802,\"first_name\":\"Conn\",\"gender\":\"Male\",\"income\":322834,\"last_name\":\"Vanes\",\"years_active\":13},{\"customer_id\":223,\"email\":\"lnorcutt66@theglobeandmail.com\",\"fico\":505,\"first_name\":\"Livia\",\"gender\":\"Female\",\"income\":312518,\"last_name\":\"Norcutt\",\"years_active\":34},{\"customer_id\":58,\"email\":\"ahornung1l@census.gov\",\"fico\":482,\"first_name\":\"Aurilia\",\"gender\":\"Female\",\"income\":137743,\"last_name\":\"Hornung\",\"years_active\":3},{\"customer_id\":232,\"email\":\"gkornes6f@skyrock.com\",\"fico\":853,\"first_name\":\"Gorden\",\"gender\":\"Male\",\"income\":205095,\"last_name\":\"Kornes\",\"years_active\":14},{\"customer_id\":486,\"email\":\"dchitteydh@globo.com\",\"fico\":501,\"first_name\":\"Delinda\",\"gender\":\"Female\",\"income\":76898,\"last_name\":\"Chittey\",\"years_active\":18},{\"customer_id\":462,\"email\":\"dchurchillct@zdnet.com\",\"fico\":506,\"first_name\":\"Dalton\",\"gender\":\"Male\",\"income\":190931,\"last_name\":\"Churchill\",\"years_active\":15},{\"customer_id\":441,\"email\":\"ecourtinc8@parallels.com\",\"fico\":706,\"first_name\":\"Em\",\"gender\":\"Female\",\"income\":383650,\"last_name\":\"Courtin\",\"years_active\":37},{\"customer_id\":731,\"email\":\"afieldska@imdb.com\",\"fico\":708,\"first_name\":\"Annetta\",\"gender\":\"Female\",\"income\":72063,\"last_name\":\"Fields\",\"years_active\":8},{\"customer_id\":388,\"email\":\"gsteddallar@google.com.au\",\"fico\":794,\"first_name\":\"Gates\",\"gender\":\"Female\",\"income\":212122,\"last_name\":\"Steddall\",\"years_active\":29},{\"customer_id\":395,\"email\":\"egaltoneay@fastcompany.com\",\"fico\":744,\"first_name\":\"Ellerey\",\"gender\":\"Male\",\"income\":302819,\"last_name\":\"Galtone\",\"years_active\":40},{\"customer_id\":606,\"email\":\"bburletongt@washington.edu\",\"fico\":523,\"first_name\":\"Bastien\",\"gender\":\"Male\",\"income\":306724,\"last_name\":\"Burleton\",\"years_active\":36},{\"customer_id\":307,\"email\":\"afish8i@stumbleupon.com\",\"fico\":742,\"first_name\":\"Alfi\",\"gender\":\"Female\",\"income\":115555,\"last_name\":\"Fish\",\"years_active\":31},{\"customer_id\":137,\"email\":\"aneilands3s@wordpress.com\",\"fico\":755,\"first_name\":\"Amaleta\",\"gender\":\"Female\",\"income\":417631,\"last_name\":\"Neilands\",\"years_active\":11},{\"customer_id\":56,\"email\":\"mskyram1j@sciencedaily.com\",\"fico\":871,\"first_name\":\"Maiga\",\"gender\":\"Female\",\"income\":451031,\"last_name\":\"Skyram\",\"years_active\":33},{\"customer_id\":719,\"email\":\"bcrasswelljy@php.net\",\"fico\":615,\"first_name\":\"Barbra\",\"gender\":\"Female\",\"income\":208742,\"last_name\":\"Crasswell\",\"years_active\":20},{\"customer_id\":104,\"email\":\"ahinchshaw2v@state.tx.us\",\"fico\":476,\"first_name\":\"Anestassia\",\"gender\":\"Female\",\"income\":204589,\"last_name\":\"Hinchshaw\",\"years_active\":16},{\"customer_id\":493,\"email\":\"jforlongedo@usatoday.com\",\"fico\":827,\"first_name\":\"Joan\",\"gender\":\"Female\",\"income\":365601,\"last_name\":\"Forlonge\",\"years_active\":42},{\"customer_id\":114,\"email\":\"jpelchat35@exblog.jp\",\"fico\":447,\"first_name\":\"Jeddy\",\"gender\":\"Male\",\"income\":97390,\"last_name\":\"Pelchat\",\"years_active\":26},{\"customer_id\":392,\"email\":\"jfilippyevav@pinterest.com\",\"fico\":741,\"first_name\":\"Jermayne\",\"gender\":\"Male\",\"income\":211833,\"last_name\":\"Filippyev\",\"years_active\":25},{\"customer_id\":343,\"email\":\"pkadwallider9i@toplist.cz\",\"fico\":889,\"first_name\":\"Pepillo\",\"gender\":\"Male\",\"income\":323547,\"last_name\":\"Kadwallider\",\"years_active\":27},{\"customer_id\":640,\"email\":\"sjolleyhr@netvibes.com\",\"fico\":789,\"first_name\":\"Smith\",\"gender\":\"Male\",\"income\":442136,\"last_name\":\"Jolley\",\"years_active\":24},{\"customer_id\":102,\"email\":\"dpeckitt2t@symantec.com\",\"fico\":505,\"first_name\":\"David\",\"gender\":\"Male\",\"income\":267979,\"last_name\":\"Peckitt\",\"years_active\":22},{\"customer_id\":268,\"email\":\"elusted7f@paginegialle.it\",\"fico\":828,\"first_name\":\"Emmott\",\"gender\":\"Male\",\"income\":109105,\"last_name\":\"Lusted\",\"years_active\":34},{\"customer_id\":294,\"email\":\"ksteers85@yahoo.com\",\"fico\":839,\"first_name\":\"Kellsie\",\"gender\":\"Female\",\"income\":432978,\"last_name\":\"Steers\",\"years_active\":24},{\"customer_id\":169,\"email\":\"rmiall4o@wikipedia.org\",\"fico\":526,\"first_name\":\"Ricoriki\",\"gender\":\"Male\",\"income\":240458,\"last_name\":\"Miall\",\"years_active\":43},{\"customer_id\":57,\"email\":\"ktuminelli1k@wufoo.com\",\"fico\":578,\"first_name\":\"Karalee\",\"gender\":\"Female\",\"income\":102314,\"last_name\":\"Tuminelli\",\"years_active\":25},{\"customer_id\":35,\"email\":\"wrigdeny@simplemachines.org\",\"fico\":609,\"first_name\":\"Willie\",\"gender\":\"Male\",\"income\":348820,\"last_name\":\"Rigden\",\"years_active\":12},{\"customer_id\":956,\"email\":\"cdenormanvilleqj@jimdo.com\",\"fico\":619,\"first_name\":\"Chico\",\"gender\":\"Male\",\"income\":108575,\"last_name\":\"De Normanville\",\"years_active\":36},{\"customer_id\":49,\"email\":\"wshawcross1c@miitbeian.gov.cn\",\"fico\":468,\"first_name\":\"Winifred\",\"gender\":\"Female\",\"income\":219323,\"last_name\":\"Shawcross\",\"years_active\":2},{\"customer_id\":752,\"email\":\"kchungkv@europa.eu\",\"fico\":755,\"first_name\":\"Konstantin\",\"gender\":\"Male\",\"income\":154355,\"last_name\":\"Chung\",\"years_active\":39},{\"customer_id\":522,\"email\":\"doldmeadoweh@tamu.edu\",\"fico\":868,\"first_name\":\"Didi\",\"gender\":\"Female\",\"income\":83238,\"last_name\":\"Oldmeadow\",\"years_active\":10},{\"customer_id\":845,\"email\":\"eranscombng@ucla.edu\",\"fico\":814,\"first_name\":\"Elisabet\",\"gender\":\"Female\",\"income\":305152,\"last_name\":\"Ranscomb\",\"years_active\":30},{\"customer_id\":980,\"email\":\"kchasemorer7@accuweather.com\",\"fico\":859,\"first_name\":\"Kali\",\"gender\":\"Female\",\"income\":274697,\"last_name\":\"Chasemore\",\"years_active\":1},{\"customer_id\":204,\"email\":\"lgainforth5n@qq.com\",\"fico\":544,\"first_name\":\"Lily\",\"gender\":\"Female\",\"income\":268598,\"last_name\":\"Gainforth\",\"years_active\":44},{\"customer_id\":581,\"email\":\"rmcanenyg4@msu.edu\",\"fico\":813,\"first_name\":\"Rea\",\"gender\":\"Female\",\"income\":369682,\"last_name\":\"McAneny\",\"years_active\":34},{\"customer_id\":886,\"email\":\"hfullickol@nps.gov\",\"fico\":846,\"first_name\":\"Hadleigh\",\"gender\":\"Male\",\"income\":219324,\"last_name\":\"Fullick\",\"years_active\":43},{\"customer_id\":694,\"email\":\"lfreckinghamj9@nydailynews.com\",\"fico\":472,\"first_name\":\"Lindsay\",\"gender\":\"Male\",\"income\":121780,\"last_name\":\"Freckingham\",\"years_active\":20},{\"customer_id\":229,\"email\":\"hsygroves6c@chicagotribune.com\",\"fico\":702,\"first_name\":\"Hattie\",\"gender\":\"Female\",\"income\":254510,\"last_name\":\"Sygroves\",\"years_active\":14},{\"customer_id\":753,\"email\":\"wbeatsonkw@europa.eu\",\"fico\":660,\"first_name\":\"Werner\",\"gender\":\"Male\",\"income\":249245,\"last_name\":\"Beatson\",\"years_active\":2},{\"customer_id\":190,\"email\":\"fivons59@symantec.com\",\"fico\":554,\"first_name\":\"Frederica\",\"gender\":\"Female\",\"income\":436162,\"last_name\":\"Ivons\",\"years_active\":10},{\"customer_id\":206,\"email\":\"bmcneely5p@huffingtonpost.com\",\"fico\":675,\"first_name\":\"Brandtr\",\"gender\":\"Male\",\"income\":414909,\"last_name\":\"McNeely\",\"years_active\":29}]},\"fields\":[{\"name\":\"customer_id\",\"type\":\"int\"},{\"name\":\"first_name\",\"type\":\"string\"},{\"name\":\"last_name\",\"type\":\"string\"},{\"name\":\"email\",\"type\":\"string\"},{\"name\":\"gender\",\"type\":\"string\"},{\"name\":\"income\",\"type\":\"int\"},{\"name\":\"fico\",\"type\":\"int\"},{\"name\":\"years_active\",\"type\":\"int\"}],\"name\":\"insurance.insurancecustomers\",\"type\":\"record\"}"
}

func (r Insurancecustomers) SchemaName() string {
	return "insurance.insurancecustomers"
}

func (_ Insurancecustomers) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Insurancecustomers) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Insurancecustomers) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Insurancecustomers) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Insurancecustomers) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Insurancecustomers) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Insurancecustomers) SetString(v string)   { panic("Unsupported operation") }
func (_ Insurancecustomers) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Insurancecustomers) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Customer_id}

		return w

	case 1:
		w := types.String{Target: &r.First_name}

		return w

	case 2:
		w := types.String{Target: &r.Last_name}

		return w

	case 3:
		w := types.String{Target: &r.Email}

		return w

	case 4:
		w := types.String{Target: &r.Gender}

		return w

	case 5:
		w := types.Int{Target: &r.Income}

		return w

	case 6:
		w := types.Int{Target: &r.Fico}

		return w

	case 7:
		w := types.Int{Target: &r.Years_active}

		return w

	}
	panic("Unknown field index")
}

func (r *Insurancecustomers) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Insurancecustomers) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Insurancecustomers) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Insurancecustomers) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Insurancecustomers) HintSize(int)                     { panic("Unsupported operation") }
func (_ Insurancecustomers) Finalize()                        {}

func (_ Insurancecustomers) AvroCRC64Fingerprint() []byte {
	return []byte(InsurancecustomersAvroCRC64Fingerprint)
}

func (r Insurancecustomers) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["customer_id"], err = json.Marshal(r.Customer_id)
	if err != nil {
		return nil, err
	}
	output["first_name"], err = json.Marshal(r.First_name)
	if err != nil {
		return nil, err
	}
	output["last_name"], err = json.Marshal(r.Last_name)
	if err != nil {
		return nil, err
	}
	output["email"], err = json.Marshal(r.Email)
	if err != nil {
		return nil, err
	}
	output["gender"], err = json.Marshal(r.Gender)
	if err != nil {
		return nil, err
	}
	output["income"], err = json.Marshal(r.Income)
	if err != nil {
		return nil, err
	}
	output["fico"], err = json.Marshal(r.Fico)
	if err != nil {
		return nil, err
	}
	output["years_active"], err = json.Marshal(r.Years_active)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Insurancecustomers) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["customer_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Customer_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for customer_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["first_name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.First_name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for first_name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["last_name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Last_name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for last_name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["email"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Email); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for email")
	}
	val = func() json.RawMessage {
		if v, ok := fields["gender"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Gender); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for gender")
	}
	val = func() json.RawMessage {
		if v, ok := fields["income"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Income); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for income")
	}
	val = func() json.RawMessage {
		if v, ok := fields["fico"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Fico); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for fico")
	}
	val = func() json.RawMessage {
		if v, ok := fields["years_active"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Years_active); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for years_active")
	}
	return nil
}
