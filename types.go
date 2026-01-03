package mollie

import (
	"encoding/json"
	"fmt"
	"time"
)

// Amount represents a single price amount with a currency.
//
// See the [documentation].
//
// [documentation]: https://docs.mollie.com/reference/common-data-types#amount-object
type Amount struct {
	Currency string `json:"currency,omitempty" url:"currency,omitempty"`
	Value    string `json:"value,omitempty"    url:"value,omitempty"`
}

// Address represents a single address.
//
// See the [documentation].
//
// [documentation]: https://docs.mollie.com/reference/common-data-types#address-object
type Address struct {
	GivenName        string `json:"givenName,omitempty"        url:"givenName,omitempty"`
	FamilyName       string `json:"familyName,omitempty"       url:"familyName,omitempty"`
	StreetAndNumber  string `json:"streetAndNumber,omitempty"  url:"streetAndNumber,omitempty"`
	StreetAdditional string `json:"streetAdditional,omitempty" url:"streetAdditional,omitempty"`
	PostalCode       string `json:"postalCode,omitempty"       url:"postalCode,omitempty"`
	City             string `json:"city,omitempty"             url:"city,omitempty"`
	Region           string `json:"region,omitempty"           url:"region,omitempty"`
	Country          string `json:"country,omitempty"          url:"country,omitempty"`
}

// BusinessCategory specifies a single business category.
//
// See the [documentation].
//
// [documentation]: https://docs.mollie.com/reference/common-data-types#business-category
type BusinessCategory string

const (
	// BusinessCategoryPetShops specifies pet shops, pet food, and supplies.
	BusinessCategoryPetShops BusinessCategory = "PET_SHOPS"
	// BusinessCategoryVeterinaryServices specifies veterinary services.
	BusinessCategoryVeterinaryServices BusinessCategory = "VETERINARY_SERVICES"

	// BusinessCategoryACAndHeatingContractors specifies A/C and heating contractors.
	BusinessCategoryACAndHeatingContractors BusinessCategory = "AC_AND_HEATING_CONTRACTORS"
	// BusinessCategoryCarpentryContractors specifies carpentry contractors.
	BusinessCategoryCarpentryContractors BusinessCategory = "CARPENTRY_CONTRACTORS"
	// BusinessCategoryElectricalContractors specifies electrical contractors.
	BusinessCategoryElectricalContractors BusinessCategory = "ELECTRICAL_CONTRACTORS"
	// BusinessCategoryEquipmentToolsFurnitureRentalLeasing specifies equipment, tools or furniture rental/leasing.
	BusinessCategoryEquipmentToolsFurnitureRentalLeasing BusinessCategory = "EQUIPMENT_TOOLS_FURNITURE_RENTAL_LEASING"
	// BusinessCategoryGeneralContractors specifies general contractors.
	BusinessCategoryGeneralContractors BusinessCategory = "GENERAL_CONTRACTORS"
	// BusinessCategorySpecialTradeContractors specifies special trade contractors.
	BusinessCategorySpecialTradeContractors BusinessCategory = "SPECIAL_TRADE_CONTRACTORS"

	// BusinessCategoryCharityAndDonations specifies charity and donations.
	BusinessCategoryCharityAndDonations BusinessCategory = "CHARITY_AND_DONATIONS"
	// BusinessCategoryFundraisingCrowdfundingSocialService specifies fundraising, crowdfunding and social service organizations.
	BusinessCategoryFundraisingCrowdfundingSocialService BusinessCategory = "FUNDRAISING_CROWDFUNDING_SOCIAL_SERVICE"

	// BusinessCategoryApps specifies apps.
	BusinessCategoryApps BusinessCategory = "APPS"
	// BusinessCategoryBooksMediaMoviesMusic specifies books, media, movies, music.
	BusinessCategoryBooksMediaMoviesMusic BusinessCategory = "BOOKS_MEDIA_MOVIES_MUSIC"
	// BusinessCategoryGames specifies games.
	BusinessCategoryGames BusinessCategory = "GAMES"
	// BusinessCategorySoftwareAndSubscriptions specifies software and subscriptions.
	BusinessCategorySoftwareAndSubscriptions BusinessCategory = "SOFTWARE_AND_SUBSCRIPTIONS"

	// BusinessCategoryChildCareServices specifies child care services.
	BusinessCategoryChildCareServices BusinessCategory = "CHILD_CARE_SERVICES"
	// BusinessCategoryCollegesUniversities specifies colleges or universities.
	BusinessCategoryCollegesUniversities BusinessCategory = "COLLEGES_UNIVERSITIES"
	// BusinessCategoryElementarySecondarySchools specifies elementary or secondary schools.
	BusinessCategoryElementarySecondarySchools BusinessCategory = "ELEMENTARY_SECONDARY_SCHOOLS"
	// BusinessCategoryOtherEducationalServices specifies other educational services.
	BusinessCategoryOtherEducationalServices BusinessCategory = "OTHER_EDUCATIONAL_SERVICES"
	// BusinessCategoryVocationalSchoolsTradeSchools specifies vocational schools or trade schools.
	BusinessCategoryVocationalSchoolsTradeSchools BusinessCategory = "VOCATIONAL_SCHOOLS_TRADE_SCHOOLS"

	// BusinessCategoryAmusementParks specifies amusement parks, circuses, carnivals, and fortune tellers.
	BusinessCategoryAmusementParks BusinessCategory = "AMUSEMENT_PARKS"
	// BusinessCategoryEventTicketing specifies event ticketing.
	BusinessCategoryEventTicketing BusinessCategory = "EVENT_TICKETING"
	// BusinessCategoryGamingEstablishments specifies gaming establishments, incl. billiards, pool, bowling, arcades.
	BusinessCategoryGamingEstablishments BusinessCategory = "GAMING_ESTABLISHMENTS"
	// BusinessCategoryMovieTheatres specifies movie theatres.
	BusinessCategoryMovieTheatres BusinessCategory = "MOVIE_THEATRES"
	// BusinessCategoryMusiciansBandsOrchestras specifies musicians, bands, or orchestras.
	BusinessCategoryMusiciansBandsOrchestras BusinessCategory = "MUSICIANS_BANDS_ORCHESTRAS"
	// BusinessCategoryOnlineGambling specifies online gambling.
	BusinessCategoryOnlineGambling BusinessCategory = "ONLINE_GAMBLING"
	// BusinessCategoryOtherEntertainmentRecreation specifies other entertainment and recreation.
	BusinessCategoryOtherEntertainmentRecreation BusinessCategory = "OTHER_ENTERTAINMENT_RECREATION"
	// BusinessCategorySportingRecreationalCamps specifies sporting and recreational camps.
	BusinessCategorySportingRecreationalCamps BusinessCategory = "SPORTING_RECREATIONAL_CAMPS"
	// BusinessCategorySportsForecasting specifies sports forecasting or prediction services.
	BusinessCategorySportsForecasting BusinessCategory = "SPORTS_FORECASTING"

	// BusinessCategoryCreditCounsellingRepair specifies credit counselling or credit repair.
	BusinessCategoryCreditCounsellingRepair BusinessCategory = "CREDIT_COUNSELLING_REPAIR"
	// BusinessCategoryDigitalWallets specifies digital wallets.
	BusinessCategoryDigitalWallets BusinessCategory = "DIGITAL_WALLETS"
	// BusinessCategoryInvestmentServices specifies investment services.
	BusinessCategoryInvestmentServices BusinessCategory = "INVESTMENT_SERVICES"
	// BusinessCategoryMoneyServices specifies money services or transmission.
	BusinessCategoryMoneyServices BusinessCategory = "MONEY_SERVICES"
	// BusinessCategoryMortgagesInsurancesLoansFinancialAdvice specifies mortgages, insurances, loans and financial advice.
	BusinessCategoryMortgagesInsurancesLoansFinancialAdvice BusinessCategory = "MORTGAGES_INSURANCES_LOANS_FINANCIAL_ADVICE"
	// BusinessCategorySecurityBrokersDealers specifies security brokers or dealers.
	BusinessCategorySecurityBrokersDealers BusinessCategory = "SECURITY_BROKERS_DEALERS"
	// BusinessCategoryTrustOffices specifies trust offices.
	BusinessCategoryTrustOffices BusinessCategory = "TRUST_OFFICES"
	// BusinessCategoryVirtualCryptoCurrencies specifies virtual currencies and crypto currencies.
	BusinessCategoryVirtualCryptoCurrencies BusinessCategory = "VIRTUAL_CRYPTO_CURRENCIES"

	// BusinessCategoryCaterers specifies caterers (prepare and delivery).
	BusinessCategoryCaterers BusinessCategory = "CATERERS"
	// BusinessCategoryFastFoodRestaurants specifies fast food restaurants.
	BusinessCategoryFastFoodRestaurants BusinessCategory = "FAST_FOOD_RESTAURANTS"
	// BusinessCategoryFoodProductStores specifies grocery stores, supermarkets and food product stores.
	BusinessCategoryFoodProductStores BusinessCategory = "FOOD_PRODUCT_STORES"
	// BusinessCategoryRestaurantsNightlife specifies restaurants, nightlife & other on-premise consumption.
	BusinessCategoryRestaurantsNightlife BusinessCategory = "RESTAURANTS_NIGHTLIFE"

	// BusinessCategoryBoatRentalsLeasing specifies boat rentals and leasing.
	BusinessCategoryBoatRentalsLeasing BusinessCategory = "BOAT_RENTALS_LEASING"
	// BusinessCategoryCruiseLines specifies cruise lines.
	BusinessCategoryCruiseLines BusinessCategory = "CRUISE_LINES"
	// BusinessCategoryLodging specifies hotels, motels, resorts, inns and other lodging and hospitality.
	BusinessCategoryLodging BusinessCategory = "LODGING"
	// BusinessCategoryPropertyRentalsCamping specifies property rentals / camping.
	BusinessCategoryPropertyRentalsCamping BusinessCategory = "PROPERTY_RENTALS_CAMPING"

	// BusinessCategoryMarketplaces specifies marketplaces.
	BusinessCategoryMarketplaces BusinessCategory = "MARKETPLACES"

	// BusinessCategoryDentalEquipmentSupplies specifies dental, lab and/or ophthalmic equipment and supplies.
	BusinessCategoryDentalEquipmentSupplies BusinessCategory = "DENTAL_EQUIPMENT_SUPPLIES"
	// BusinessCategoryDentistsOrthodontists specifies dentists and orthodontists.
	BusinessCategoryDentistsOrthodontists BusinessCategory = "DENTISTS_ORTHODONTISTS"
	// BusinessCategoryMedicalServices specifies doctors, physicians and other medical services.
	BusinessCategoryMedicalServices BusinessCategory = "MEDICAL_SERVICES"
	// BusinessCategoryDrugPharmaciesPrescription specifies drug stores, pharmacies and prescription medicine.
	BusinessCategoryDrugPharmaciesPrescription BusinessCategory = "DRUG_PHARMACIES_PRESCRIPTION"
	// BusinessCategoryMedicalDevices specifies medical devices.
	BusinessCategoryMedicalDevices BusinessCategory = "MEDICAL_DEVICES"
	// BusinessCategoryMedicalOrganizations specifies medical organizations.
	BusinessCategoryMedicalOrganizations BusinessCategory = "MEDICAL_ORGANIZATIONS"
	// BusinessCategoryMentalHealthServices specifies mental health services.
	BusinessCategoryMentalHealthServices BusinessCategory = "MENTAL_HEALTH_SERVICES"
	// BusinessCategoryNursing specifies nursing or personal care facilities and assisted living.
	BusinessCategoryNursing BusinessCategory = "NURSING"
	// BusinessCategoryOpticiansEyeglasses specifies opticians and eyeglasses.
	BusinessCategoryOpticiansEyeglasses BusinessCategory = "OPTICIANS_EYEGLASSES"

	// BusinessCategorySocialAssociations specifies civic, fraternal, or social associations.
	BusinessCategorySocialAssociations BusinessCategory = "SOCIAL_ASSOCIATIONS"
	// BusinessCategoryMembershipFeeBasedSports specifies gyms, membership fee based sports.
	BusinessCategoryMembershipFeeBasedSports BusinessCategory = "MEMBERSHIP_FEE_BASED_SPORTS"
	// BusinessCategoryOtherMembershipOrganizations specifies other membership organizations.
	BusinessCategoryOtherMembershipOrganizations BusinessCategory = "OTHER_MEMBERSHIP_ORGANIZATIONS"

	// BusinessCategoryAdultContentServices specifies adult content or services.
	BusinessCategoryAdultContentServices BusinessCategory = "ADULT_CONTENT_SERVICES"
	// BusinessCategoryCounselingServices specifies counseling services.
	BusinessCategoryCounselingServices BusinessCategory = "COUNSELING_SERVICES"
	// BusinessCategoryDatingServices specifies dating services.
	BusinessCategoryDatingServices BusinessCategory = "DATING_SERVICES"
	// BusinessCategoryHealthBeautySpas specifies health and beauty spas.
	BusinessCategoryHealthBeautySpas BusinessCategory = "HEALTH_BEAUTY_SPAS"
	// BusinessCategoryLandscapingServices specifies landscaping services.
	BusinessCategoryLandscapingServices BusinessCategory = "LANDSCAPING_SERVICES"
	// BusinessCategoryLaundryDrycleaningServices specifies laundry or (dry)cleaning services.
	BusinessCategoryLaundryDrycleaningServices BusinessCategory = "LAUNDRY_DRYCLEANING_SERVICES"
	// BusinessCategoryMassageParlours specifies massage parlours.
	BusinessCategoryMassageParlours BusinessCategory = "MASSAGE_PARLOURS"
	// BusinessCategoryOtherPersonalServices specifies other personal services.
	BusinessCategoryOtherPersonalServices BusinessCategory = "OTHER_PERSONAL_SERVICES"
	// BusinessCategoryPhotographyStudios specifies photography studios.
	BusinessCategoryPhotographyStudios BusinessCategory = "PHOTOGRAPHY_STUDIOS"
	// BusinessCategorySalonsBarbers specifies salons or barbers.
	BusinessCategorySalonsBarbers BusinessCategory = "SALONS_BARBERS"

	// BusinessCategoryPoliticalParties specifies political parties.
	BusinessCategoryPoliticalParties BusinessCategory = "POLITICAL_PARTIES"

	// BusinessCategoryAccountingAuditingBookkeepingTaxPreparationServices specifies accounting, auditing, bookkeeping and tax preparation services.
	BusinessCategoryAccountingAuditingBookkeepingTaxPreparationServices BusinessCategory = "ACCOUNTING_AUDITING_BOOKKEEPING_TAX_PREPARATION_SERVICES"
	// BusinessCategoryAdvertisingServices specifies advertising services.
	BusinessCategoryAdvertisingServices BusinessCategory = "ADVERTISING_SERVICES"
	// BusinessCategoryCleaningMaintenanceJanitorialServices specifies cleaning and maintenance, janitorial services.
	BusinessCategoryCleaningMaintenanceJanitorialServices BusinessCategory = "CLEANING_MAINTENANCE_JANITORIAL_SERVICES"
	// BusinessCategoryComputerRepair specifies computer repair.
	BusinessCategoryComputerRepair BusinessCategory = "COMPUTER_REPAIR"
	// BusinessCategoryConsultancy specifies consultancy.
	BusinessCategoryConsultancy BusinessCategory = "CONSULTANCY"
	// BusinessCategorySecurityServices specifies detective/protective agencies, security services.
	BusinessCategorySecurityServices BusinessCategory = "SECURITY_SERVICES"
	// BusinessCategoryDirectMarketing specifies direct marketing.
	BusinessCategoryDirectMarketing BusinessCategory = "DIRECT_MARKETING"
	// BusinessCategoryFuneralServices specifies funeral services and crematories.
	BusinessCategoryFuneralServices BusinessCategory = "FUNERAL_SERVICES"
	// BusinessCategoryGovernmentServices specifies government services.
	BusinessCategoryGovernmentServices BusinessCategory = "GOVERNMENT_SERVICES"
	// BusinessCategoryHostingVPNServices specifies hosting and vpn services.
	BusinessCategoryHostingVPNServices BusinessCategory = "HOSTING_VPN_SERVICES"
	// BusinessCategoryIndustrialSuppliesNotElsewhereClassified specifies industrial supplies, not elsewhere classified.
	BusinessCategoryIndustrialSuppliesNotElsewhereClassified BusinessCategory = "INDUSTRIAL_SUPPLIES_NOT_ELSEWHERE_CLASSIFIED"
	// BusinessCategoryLegalServicesAttorneys specifies legal services and attorneys.
	BusinessCategoryLegalServicesAttorneys BusinessCategory = "LEGAL_SERVICES_ATTORNEYS"
	// BusinessCategoryMotionPicturesDistribution specifies motion picture / video tape production and/or distribution.
	BusinessCategoryMotionPicturesDistribution BusinessCategory = "MOTION_PICTURES_DISTRIBUTION"
	// BusinessCategoryOtherBusinessServices specifies other business services.
	BusinessCategoryOtherBusinessServices BusinessCategory = "OTHER_BUSINESS_SERVICES"
	// BusinessCategoryPrintingPublishing specifies printing and publishing.
	BusinessCategoryPrintingPublishing BusinessCategory = "PRINTING_PUBLISHING"
	// BusinessCategoryRealEstateAgents specifies real estate agents.
	BusinessCategoryRealEstateAgents BusinessCategory = "REAL_ESTATE_AGENTS"
	// BusinessCategorySanitationPolishingSpecialtyCleaning specifies sanitation, polishing and specialty cleaning.
	BusinessCategorySanitationPolishingSpecialtyCleaning BusinessCategory = "SANITATION_POLISHING_SPECIALTY_CLEANING"
	// BusinessCategoryOfficeSupplies specifies stationery / office supplies.
	BusinessCategoryOfficeSupplies BusinessCategory = "OFFICE_SUPPLIES"
	// BusinessCategoryTestingLaboratoriesNotMedical specifies testing laboratories (not medical).
	BusinessCategoryTestingLaboratoriesNotMedical BusinessCategory = "TESTING_LABORATORIES_NOT_MEDICAL"
	// BusinessCategoryTrainingAndCoaching specifies training and coaching.
	BusinessCategoryTrainingAndCoaching BusinessCategory = "TRAINING_AND_COACHING"
	// BusinessCategoryUtilities specifies utilities.
	BusinessCategoryUtilities BusinessCategory = "UTILITIES"

	// BusinessCategoryReligiousOrganizations specifies religious organizations.
	BusinessCategoryReligiousOrganizations BusinessCategory = "RELIGIOUS_ORGANIZATIONS"

	// BusinessCategoryClothingShoesAccesories specifies (sports) clothing, shoes and accessories.
	BusinessCategoryClothingShoesAccesories BusinessCategory = "CLOTHING_SHOES_ACCESSORIES"
	// BusinessCategoryCommercialArt specifies art dealers, galleries, (commercial) photography and graphics.
	BusinessCategoryCommercialArt BusinessCategory = "COMMERCIAL_ART"
	// BusinessCategoryBeautyProducts specifies beauty products.
	BusinessCategoryBeautyProducts BusinessCategory = "BEAUTY_PRODUCTS"
	// BusinessCategoryBooksPeriodicalsNewspapers specifies books, periodicals and newspapers.
	BusinessCategoryBooksPeriodicalsNewspapers BusinessCategory = "BOOKS_PERIODICALS_NEWSPAPERS"
	// BusinessCategoryHomeImprovement specifies building, home improvement and equipment.
	BusinessCategoryHomeImprovement BusinessCategory = "HOME_IMPROVEMENT"
	// BusinessCategoryGiftsShops specifies cards, gifts, novelty and souvenir shops.
	BusinessCategoryGiftsShops BusinessCategory = "GIFTS_SHOPS"
	// BusinessCategoryCBDMarijuanaProducts specifies cbd/marijuana (related) products.
	BusinessCategoryCBDMarijuanaProducts BusinessCategory = "CBD_MARIJUANA_PRODUCTS"
	// BusinessCategoryCoffeeShops specifies coffee shops / grow shops.
	BusinessCategoryCoffeeShops BusinessCategory = "COFFEE_SHOPS"
	// BusinessCategoryConvenienceStores specifies convenience stores, specialty markets, health food stores.
	BusinessCategoryConvenienceStores BusinessCategory = "CONVENIENCE_STORES"
	// BusinessCategoryGiftCards specifies credits, vouchers, gift cards (excl. sim cards) for non-financial institutions.
	BusinessCategoryGiftCards BusinessCategory = "GIFT_CARDS"
	// BusinessCategoryEroticToys specifies erotic toys.
	BusinessCategoryEroticToys BusinessCategory = "EROTIC_TOYS"
	// BusinessCategoryFlorists specifies florists, florist supplier.
	BusinessCategoryFlorists BusinessCategory = "FLORISTS"
	// BusinessCategoryFuelDealers specifies fuel dealers (i.e. oil, pertroleum).
	BusinessCategoryFuelDealers BusinessCategory = "FUEL_DEALERS"
	// BusinessCategoryFurnitureFurnishingsEquipmentStores specifies furniture, home furnishings and equipment stores.
	BusinessCategoryFurnitureFurnishingsEquipmentStores BusinessCategory = "FURNITURE_FURNISHINGS_EQUIPMENT_STORES"
	// BusinessCategoryGameToyHobbyShops specifies game, toy and hobby shops.
	BusinessCategoryGameToyHobbyShops BusinessCategory = "GAME_TOY_HOBBY_SHOPS"
	// BusinessCategoryOutdoorEquipment specifies garden and outdoor equipment.
	BusinessCategoryOutdoorEquipment BusinessCategory = "OUTDOOR_EQUIPMENT"
	// BusinessCategoryHomeElectronics specifies home electronics & (personal) computers.
	BusinessCategoryHomeElectronics BusinessCategory = "HOME_ELECTRONICS"
	// BusinessCategoryHouseholdApplianceStores specifies household appliance stores.
	BusinessCategoryHouseholdApplianceStores BusinessCategory = "HOUSEHOLD_APPLIANCE_STORES"
	// BusinessCategoryJewelryWatchClockAndSilverwareStoresUnder1000 specifies jewelry, watch, clock, and silverware stores (<1000 euro).
	BusinessCategoryJewelryWatchClockAndSilverwareStoresUnder1000 BusinessCategory = "JEWELRY_WATCH_CLOCK_AND_SILVERWARE_STORES_UNDER_1000"
	// BusinessCategoryMusicStores specifies music stores, instruments and records.
	BusinessCategoryMusicStores BusinessCategory = "MUSIC_STORES"
	// BusinessCategoryOtherMerchandise specifies other merchandise.
	BusinessCategoryOtherMerchandise BusinessCategory = "OTHER_MERCHANDISE"
	// BusinessCategoryLiquorStores specifies package stores–beer, wine, and liquor.
	BusinessCategoryLiquorStores BusinessCategory = "LIQUOR_STORES"
	// BusinessCategoryPaidTelevisionRadio specifies paid television or radio services (cable/satellite).
	BusinessCategoryPaidTelevisionRadio BusinessCategory = "PAID_TELEVISION_RADIO"
	// BusinessCategoryPreciousStonesMetalsJewelryOver1000 specifies precious stones, metals, watches and jewelry (>1000 euro).
	BusinessCategoryPreciousStonesMetalsJewelryOver1000 BusinessCategory = "PRECIOUS_STONES_METALS_JEWELRY_OVER_1000"
	// BusinessCategoryRepairShops specifies repair shops and related services, not elsewhere classified.
	BusinessCategoryRepairShops BusinessCategory = "REPAIR_SHOPS"
	// BusinessCategorySecondHandStores specifies second hand / used merchandise stores.
	BusinessCategorySecondHandStores BusinessCategory = "SECOND_HAND_STORES"
	// BusinessCategorySportingGoodsSpecialtyRetailShops specifies sporting goods stores, miscellaneous and specialty retail shops.
	BusinessCategorySportingGoodsSpecialtyRetailShops BusinessCategory = "SPORTING_GOODS_SPECIALTY_RETAIL_SHOPS"
	// BusinessCategorySupplementsStores specifies supplements, nutrition, vitamin stores.
	BusinessCategorySupplementsStores BusinessCategory = "SUPPLEMENTS_STORES"
	// BusinessCategoryTelecomEquipment specifies telecom equipment (i.e. chargers, phones).
	BusinessCategoryTelecomEquipment BusinessCategory = "TELECOM_EQUIPMENT"
	// BusinessCategoryTelecomServices specifies telecom services (incl. (anonymous) sim cards).
	BusinessCategoryTelecomServices BusinessCategory = "TELECOM_SERVICES"
	// BusinessCategoryTobaccoProducts specifies tobacco, cigars, e-cigarettes and related products.
	BusinessCategoryTobaccoProducts BusinessCategory = "TOBACCO_PRODUCTS"
	// BusinessCategoryTradersDiamonds specifies traders in diamonds.
	BusinessCategoryTradersDiamonds BusinessCategory = "TRADERS_DIAMONDS"
	// BusinessCategoryTradersGold specifies traders in gold.
	BusinessCategoryTradersGold BusinessCategory = "TRADERS_GOLD"
	// BusinessCategoryWeaponsAmmunition specifies weapons or ammunition.
	BusinessCategoryWeaponsAmmunition BusinessCategory = "WEAPONS_AMMUNITION"

	// BusinessCategoryCommuterTransportation specifies commuter transportation.
	BusinessCategoryCommuterTransportation BusinessCategory = "COMMUTER_TRANSPORTATION"
	// BusinessCategoryCourierServices specifies courier services and freight forwarders.
	BusinessCategoryCourierServices BusinessCategory = "COURIER_SERVICES"
	// BusinessCategoryOtherTransportationServices specifies other transportation services.
	BusinessCategoryOtherTransportationServices BusinessCategory = "OTHER_TRANSPORTATION_SERVICES"
	// BusinessCategoryRidesharing specifies taxis, limos and ridesharing.
	BusinessCategoryRidesharing BusinessCategory = "RIDESHARING"

	// BusinessCategoryTravelServices specifies travel agencies, tour operators and other travel services.
	BusinessCategoryTravelServices BusinessCategory = "TRAVEL_SERVICES"

	// BusinessCategoryAutomotivePartsAccessories specifies auto(motive) parts and accessories.
	BusinessCategoryAutomotivePartsAccessories BusinessCategory = "AUTOMOTIVE_PARTS_ACCESSORIES"
	// BusinessCategoryCarTruckCompanies specifies auto and truck sales and service dealers and leasing companies.
	BusinessCategoryCarTruckCompanies BusinessCategory = "CAR_TRUCK_COMPANIES"
	// BusinessCategoryAutomotiveServices specifies automotive services.
	BusinessCategoryAutomotiveServices BusinessCategory = "AUTOMOTIVE_SERVICES"
	// BusinessCategoryBicyclePartsShopsService specifies bicycle (parts) shops and service.
	BusinessCategoryBicyclePartsShopsService BusinessCategory = "BICYCLE_PARTS_SHOPS_SERVICE"
	// BusinessCategoryCarBoatCamperMobileHomeDealer specifies car, boat, camper, mobile home dealer.
	BusinessCategoryCarBoatCamperMobileHomeDealer BusinessCategory = "CAR_BOAT_CAMPER_MOBILE_HOME_DEALER"
	// BusinessCategoryCarRentals specifies car rentals.
	BusinessCategoryCarRentals BusinessCategory = "CAR_RENTALS"
	// BusinessCategoryMotorcyclePartsShopsAndDealers specifies motorcycle (parts) shops and dealers.
	BusinessCategoryMotorcyclePartsShopsAndDealers BusinessCategory = "MOTORCYCLE_PARTS_SHOPS_AND_DEALERS"
)

// Date specifies a single date.
//
// See the [documentation].
//
// [documentation]: https://docs.mollie.com/reference/common-data-types#date
type Date struct {
	time.Time
}

// MarshalJSON marshals a date in a YYYY-MM-DD format.
func (d Date) MarshalJSON() ([]byte, error) {
	dateJSON, err := json.Marshal(d.Format("2006-01-02"))

	if err != nil {
		return nil, fmt.Errorf("cannot marshal Date JSON: %s", err.Error())
	}

	return dateJSON, nil
}

// UnmarshalJSON unmarshals a date from a YYYY-MM-DD format.
func (d *Date) UnmarshalJSON(dateBytes []byte) error {
	date, err := time.Parse(`"2006-01-02"`, string(dateBytes))

	if err != nil {
		return fmt.Errorf("could not parse date: %s", err.Error())
	}

	d.Time = date

	return nil
}

// Datetime specifies a single datetime.
//
// See the [documentation].
//
// [documentation]: https://docs.mollie.com/reference/common-data-types#datetime
type Datetime struct {
	time.Time
}

// MarshalJSON marshals a datetime in an ISO8601 format.
func (d Datetime) MarshalJSON() ([]byte, error) {
	dateJSON, err := json.Marshal(time.RFC3339)

	if err != nil {
		return nil, fmt.Errorf("cannot marshal DateTime JSON: %s", err.Error())
	}

	return dateJSON, nil
}

// Locale specifies a single locale.
//
// See the [documentation].
//
// [documentation]: https://docs.mollie.com/reference/common-data-types#locale
type Locale string

const (
	// LocaleEnglishUnitedStates specifies the English (United States) locale.
	LocaleEnglishUnitedStates Locale = "en_US"
	// LocaleEnglishUnitedKingdom specifies the English (United Kingdom) locale.
	LocaleEnglishUnitedKingdom Locale = "en_GB"
	// LocaleDutchNetherlands specifies the Dutch (Netherlands) locale.
	LocaleDutchNetherlands Locale = "nl_NL"
	// LocaleDutchBelgium specifies the Dutch (Belgium) locale.
	LocaleDutchBelgium Locale = "nl_BE"
	// LocaleFrenchFrance specifies the French (France) locale.
	LocaleFrenchFrance Locale = "fr_FR"
	// LocaleFrenchBelgium specifies the French (Belgium) locale.
	LocaleFrenchBelgium Locale = "fr_BE"
	// LocaleGermanGermany specifies the German (Germany) locale.
	LocaleGermanGermany Locale = "de_DE"
	// LocaleGermanAustria specifies the German (Austria) locale.
	LocaleGermanAustria Locale = "de_AT"
	// LocaleGermanSwitzerland specifies the German (Switzerland) locale.
	LocaleGermanSwitzerland Locale = "de_CH"
	// LocaleSpanishSpain specifies the Spanish (Spain) locale.
	LocaleSpanishSpain Locale = "es_ES"
	// LocaleCatalanSpain specifies the Catalan (Spain) locale.
	LocaleCatalanSpain Locale = "ca_ES"
	// LocalePortuguesePortugal specifies the Portuguese (Portugal) locale.
	LocalePortuguesePortugal Locale = "pt_PT"
	// LocaleItalianItaly specifies the Italian (Italy) locale.
	LocaleItalianItaly Locale = "it_IT"
	// LocaleNorwegianNorway specifies the Norwegian (Norway) locale.
	LocaleNorwegianNorway Locale = "nb_NO"
	// LocaleSwedishSweden specifies the Swedish (Sweden) locale.
	LocaleSwedishSweden Locale = "sv_SE"
	// LocaleFinnishFinland specifies the Finnish (Finland) locale.
	LocaleFinnishFinland Locale = "fi_FI"
	// LocaleDanishDenmark specifies the Danish (Denmark) locale.
	LocaleDanishDenmark Locale = "da_DK"
	// LocaleIcelandicIceland specifies the Icelandic (Iceland) locale.
	LocaleIcelandicIceland Locale = "is_IS"
	// LocaleHungarianHungary specifies the Hungarian (Hungary) locale.
	LocaleHungarianHungary Locale = "hu_HU"
	// LocalePolishPoland specifies the Polish (Poland) locale.
	LocalePolishPoland Locale = "pl_PL"
	// LocaleLatvianLatvia specifies the Latvian (Latvia) locale.
	LocaleLatvianLatvia Locale = "lv_LV"
	// LocaleLithuanianLithuania specifies the Lithuanian (Lithuania) locale.
	LocaleLithuanianLithuania Locale = "lt_LT"
)

// PhoneNumber specifies a single phone number.
//
// See the [documentation].
//
// [documentation]: https://docs.mollie.com/reference/common-data-types#phone-number
type PhoneNumber string

// QRCode specifies a single QR code.
//
// See the [documentation].
//
// [documentation]: https://docs.mollie.com/reference/common-data-types#qr-code-object
type QRCode struct {
	Height int    `json:"height,omitempty" url:"height,omitempty"`
	Width  int    `json:"width,omitempty"  url:"width,omitempty"`
	Src    string `json:"src,omitempty"    url:"src,omitempty"`
}

// URL specifies a single URL.
//
// See the [documentation].
//
// [documentation]: https://docs.mollie.com/reference/common-data-types#url-object
type URL struct {
	Href string `json:"href,omitempty" url:"href,omitempty"`
	Type string `json:"type,omitempty" url:"type,omitempty"`
}

// Mode specifies the API mode.
type Mode string

const (
	// ModeTest specifies the test mode.
	ModeTest = "test"
	// ModeLive specifies the live mode.
	ModeLive = "live"
)

// Link specifies a single link.
type Link struct {
	Href string `json:"href" url:"href"`
	Type string `json:"type" url:"type"`
}

// Payment specifies a single payment instance.
type Payment struct {
	Resource     string            `json:"resource" url:"resource"`
	ID           string            `json:"id" url:"id"`
	Mode         string            `json:"mode" url:"mode"`
	CreatedAt    time.Time         `json:"createdAt" url:"createdAt"`
	Amount       Amount            `json:"amount" url:"amount"`
	Description  string            `json:"description" url:"description"`
	Method       string            `json:"method" url:"method"`
	Metadata     map[string]string `json:"metadata" url:"metadata"`
	Status       string            `json:"status" url:"status"`
	IsCancelable bool              `json:"isCancelable" url:"isCancelable"`
	ExpiresAt    time.Time         `json:"expiresAt" url:"expiresAt"`
	MandateID    string            `json:"mandateId" url:"mandateId"`
	CustomerID   string            `json:"customerId" url:"customerId"`
	ProfileID    string            `json:"profileId" url:"profileId"`
	SequenceType string            `json:"sequenceType" url:"sequenceType"`
	RedirectURL  string            `json:"redirectUrl" url:"redirectUrl"`
	Links        map[string]Link   `json:"_links" url:"_links"`
}

// Payments specifies a payments response.
type Payments struct {
	Count    int `json:"count" url:"count"`
	Embedded struct {
		Payments []Payment `json:"payments" url:"payments"`
	} `json:"_embedded" url:"_embedded"`
}

// Customer specifies a single customer instance.
type Customer struct {
	Resource  string            `json:"resource" url:"resource"`
	ID        string            `json:"id" url:"id"`
	Mode      string            `json:"mode" url:"mode"`
	Name      string            `json:"name" url:"name"`
	Email     string            `email:"name" url:"email"`
	Locale    Locale            `email:"locale" url:"locale"`
	Metadata  map[string]string `json:"metadata" url:"metadata"`
	CreatedAt time.Time         `json:"createdAt" url:"createdAt"`
	Links     map[string]Link   `json:"_links" url:"_links"`
}

// Customers specifies a customers response.
type Customers struct {
	Count    int `json:"count" url:"count"`
	Embedded struct {
		Customers []Customer `json:"customers" url:"customers"`
	} `json:"_embedded" url:"_embedded"`
}

// Subscription specifies a single subscription instance.
type Subscription struct {
	Resource        string            `json:"resource" url:"resource"`
	ID              string            `json:"id" url:"id"`
	Mode            string            `json:"mode" url:"mode"`
	Amount          Amount            `json:"amount" url:"amount"`
	Times           int               `json:"times" url:"times"`
	TimesRemaining  int               `json:"timesRemaining" url:"timesRemaining"`
	StartDate       time.Time         `json:"startDate" url:"startDate"`
	NextPaymentDate time.Time         `json:"nextPaymentDate" url:"nextPaymentDate"`
	Description     string            `json:"description" url:"description"`
	Metadata        map[string]string `json:"metadata" url:"metadata"`
	Method          string            `json:"method" url:"method"`
	Status          string            `json:"status" url:"status"`
	WebhookURL      string            `json:"webhookUrl" url:"webhookUrl"`
	CustomerID      string            `json:"customerId" url:"customerId"`
	MandateID       string            `json:"mandateId" url:"mandateId"`
	CreatedAt       time.Time         `json:"createdAt" url:"createdAt"`
	Links           map[string]Link   `json:"_links" url:"_links"`
}

// Mandate specifies a single mandate instance.
type Mandate struct {
	Resource string `json:"resource" url:"resource"`
	ID       string `json:"id" url:"id"`
	Mode     string `json:"mode" url:"mode"`
	Status   string `json:"status" url:"status"`
	Method   string `json:"method" url:"method"`
	Details  struct {
		ConsumerName    string `json:"consumerName" url:"consumerName"`
		ConsumerAccount string `json:"consumerAccount" url:"consumerAccount"`
		ConsumerBIC     string `json:"consumerBic" url:"consumerBic"`
	} `json:"details" url:"details"`
	MandateReference string          `json:"mandateReference" url:"mandateReference"`
	SignatureDate    Date            `json:"signatureDate" url:"signatureDate"`
	CustomerID       string          `json:"customerId" url:"customerId"`
	CreatedAt        time.Time       `json:"createdAt" url:"createdAt"`
	Links            map[string]Link `json:"_links" url:"_links"`
}
